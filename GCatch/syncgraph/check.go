package syncgraph

import (
	"fmt"
	"time"

	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
)

type blockingPos struct {
	pathId  int
	pNodeId int
}

// newStopChannel creates a channel that is stopped after a certain time period.
func newStopChannel() chan struct{} {
	stop := make(chan struct{})

	go func() {
		// If no timeout has been established, do not instrument an abort.
		if config.MAX_GCATCH_FRAGMENT_ANALYSIS_TIME == 0 {
			return
		}

		<-time.After(time.Duration(config.MAX_GCATCH_FRAGMENT_ANALYSIS_TIME) * time.Second)
		close(stop)
	}()

	return stop
}

// iterateOrAbort takes an abort channel, a list of items, and a function that operates over individual items
// and their index. It iterates over every element of the list, and either executes the function or aborts
// if the stop channel has been closed. It returns true if stopped prematurely, or false if it succesfully
// iterated over the entire list.
func iterateOrAbort[T any](stop chan struct{}, ts []T, f func(int, T) bool) bool {
	for i, t := range ts {
		select {
		case <-stop:
			return true
		default:
			if f(i, t) {
				return true
			}
		}
	}

	return false
}

// iterateOrAbort takes an abort channel and a function that is repeatedly executed until the guard returned
// is false aborts if the stop channel has been closed. It returns true if stopped prematurely, or false if
// it all iterations succeed before the timer expires.
func loopUntilTimeout(stop chan struct{}, f func() bool) bool {
	for f() {
		select {
		case <-stop:
			return true
		default:
		}
	}

	return false
}

func (g SyncGraph) CheckWithZ3() bool {
	var countBlockPoint int64 = 0

	bugFound := false

	// Create stop channel
	stop := newStopChannel()

	// Main loop: for each pathCombination
	iterateOrAbort(stop, g.PathCombinations, func(_ int, pathComb *pathCombination) bool {
		// If the pathCombination satisfies: the MainGoroutine's path is nil, skip this pathCombination

		// Store path's nodes into a slice of PNode. Make sure that each PNode is unique, though PNode.Node may appear multiple times
		goroutines := []*Goroutine{}
		paths := []*PPath{}

		iterateOrAbort(stop, pathComb.go_paths, func(_ int, goPath *tupleGoroutinePath) bool {
			goroutines = append(goroutines, goPath.goroutine)

			vecNewPNode := []*PNode{}
			for i, oldNode := range goPath.path.Path {
				newNode := &PNode{
					Path:     goPath.path,
					Index:    i,
					Node:     oldNode,
					Blocked:  false,
					Executed: false,
				}
				vecNewPNode = append(vecNewPNode, newNode)
			}

			newPath := &PPath{
				Path:      vecNewPNode,
				localPath: goPath.path,
			}
			paths = append(paths, newPath)
			return false
		})

		if config.BoolChSafety {
			// Check if the program has double close
			vecClose := []*instinfo.ChClose{}
			iterateOrAbort(stop, paths, func(_ int, pPath *PPath) bool {
				iterateOrAbort(stop, pPath.Path, func(i int, pNode *PNode) bool {
					if syncNode, ok := pNode.Node.(SyncOp); ok {
						if g.Task.IsPrimATarget(syncNode.Primitive()) {
							if op, ok := syncNode.(*ChanOp); ok {
								if chClose, ok := op.Op.(*instinfo.ChClose); ok {
									vecClose = append(vecClose, chClose)
								}
							}
						}
					}
					return false
				})
				return false
			})

			for _, aClose := range vecClose {
				for _, bClose := range vecClose {
					if aClose == bClose {
						continue
					}
					if aClose.Parent == bClose.Parent {
						config.BugIndex++
						fmt.Print("----------Bug[")
						fmt.Print(config.BugIndex)
						fmt.Print("]----------\n\tType: Channel Safety \tReason: Double close.\n")
						fmt.Println("Location of closes:")
						output.PrintIISrc(aClose.Inst)
						output.PrintIISrc(bClose.Inst)
						bugFound = true
						return true
					}
				}
			}
		}

		// List all blocking op of target channel on any path
		pathId2AllBlockPos := make(map[int][]blockingPos)

		const emptyPNodeId = -2
		iterateOrAbort(stop, paths, func(i int, pPath *PPath) bool {
			iterateOrAbort(stop, pPath.Path, func(j int, pNode *PNode) bool {
				if syncNode, ok := pNode.Node.(SyncOp); ok {
					if g.Task.IsPrimATarget(syncNode.Primitive()) {
						if canSyncOpTriggerGl(syncNode) {
							newBlockPos := blockingPos{
								pathId:  i,
								pNodeId: j,
							}
							pathId2AllBlockPos[i] = append(pathId2AllBlockPos[i], newBlockPos)
						}
					}
				}
				return false
			})

			emptyBlockPos := blockingPos{
				pathId:  i,
				pNodeId: emptyPNodeId,
			}
			pathId2AllBlockPos[i] = append(pathId2AllBlockPos[i], emptyBlockPos)
			return false
		})

		allBlockPosComb := []map[int]blockingPos{}

		vecIndex := []int{}
		for _, _ = range pathId2AllBlockPos {
			vecIndex = append(vecIndex, 0)
		}
		loopUntilTimeout(stop, func() bool {
			newComb := make(map[int]blockingPos)
			boolCanSync := false
			for pathId, indice := range vecIndex {
				blockPos := pathId2AllBlockPos[pathId][indice]
				if blockPos.pNodeId != emptyPNodeId {
					// check if blockPos can sync with a previous blockPos
					thisSyncNode, ok := paths[pathId].Path[blockPos.pNodeId].Node.(SyncOp)
					if !ok {
						fmt.Println("Panic when enumerate blockPos combination: Node is not SyncOp")
						panic(1)
					}
					for otherPathId, otherBlockPos := range newComb {
						if otherBlockPos.pNodeId == emptyPNodeId || otherPathId == pathId {
							continue
						}
						otherPath := paths[otherPathId]
						otherSyncNode, ok2 := otherPath.Path[otherBlockPos.pNodeId].Node.(SyncOp)
						if !ok2 {
							fmt.Println("Panic when enumerate blockPos combination: Node is not SyncOp")
							panic(1)
						}
						if thisSyncNode.Primitive() != otherSyncNode.Primitive() {
							continue
						}
						if canSync(thisSyncNode, otherSyncNode) {
							boolCanSync = true
							break
						}
					}
				}
				newComb[pathId] = blockPos
			}
			if boolCanSync == false {
				boolAllEmpty := true
				for _, blockPos := range newComb {
					if blockPos.pNodeId != emptyPNodeId {
						boolAllEmpty = false
						break
					}
				}
				if boolAllEmpty == false {
					allBlockPosComb = append(allBlockPosComb, newComb)
				}
			}

			nextPathId := -1
			for pathId, indice_ := range vecIndex {
				if indice_ >= len(pathId2AllBlockPos[pathId])-1 {
					continue
				} else {
					nextPathId = pathId
					break
				}
			}

			if nextPathId == -1 {
				return false
			}

			vecIndex[nextPathId] += 1

			for pathId, _ := range vecIndex {
				if pathId == nextPathId {
					break
				}
				vecIndex[pathId] = 0
			}

			return true
		})

		// For every blocking op of target channel on any path
		return iterateOrAbort(stop, allBlockPosComb, func(i int, blockPosComb map[int]blockingPos) bool {
			for _, blockPos := range blockPosComb {
				if blockPos.pNodeId != emptyPNodeId {
					inst := paths[blockPos.pathId].Path[blockPos.pNodeId].Node.Instruction()
					str := output.StringIISrc(inst)
					if _, printed := PrintedBlockPosStr[str]; printed {
						return true
					}
				}
			}
			// Make some paths block and other paths exit
			for j, path := range paths {
				blockPos, exist := blockPosComb[j]
				if exist && blockPos.pNodeId != emptyPNodeId {
					path.SetBlockAt(blockPos.pNodeId)
				} else {
					path.SetAllReached()
				}
			}

			// See if Sync-rule is satisfied. Sync-rule: the number of ops of one prim must match, except the blocking one
			flagSyncRuleSatisfied := true

			mapPrim2OpNodes := make(map[interface{}][]*PNode) // A map from target primitive to its reached ops in these paths
			for i := 0; i < len(goroutines); i++ {
				path := paths[i]
				for _, pNode := range path.Path {
					syncNode, ok := pNode.Node.(SyncOp)
					if !ok {
						continue
					}
					if pNode.Executed == false {
						continue
					}
					prim := syncNode.Primitive()
					if g.Task.IsPrimATarget(prim) == false { // Only consider prim that is in Task.Target
						continue
					}
					mapPrim2OpNodes[prim] = append(mapPrim2OpNodes[prim], pNode)
				}
			}

			for p, vecOpNodes := range mapPrim2OpNodes {
				nodes := []Node{}
				for _, pNode := range vecOpNodes {
					nodes = append(nodes, pNode.Node)
				}

				switch prim := p.(type) {
				case *instinfo.Channel:
					flagSyncRuleSatisfied = checkChOpsLegal(prim, nodes)
				case *instinfo.Locker:
					// Do we really need a rule for Locker?
					// let's add it anyway
					flagSyncRuleSatisfied = checkLockerOpsLegal(prim, nodes)
				}
				if flagSyncRuleSatisfied == false {
					break
				}
			}

			if flagSyncRuleSatisfied == false {
				return false
			}

			vecBlockingPos := []blockingPos{}
			for _, blockPos := range blockPosComb {
				if blockPos.pNodeId == emptyPNodeId {
					continue
				}
				vecBlockingPos = append(vecBlockingPos, blockPos)
			}
			var foundBug bool

			if config.BoolChSafety {
				config.BoolChSafety = false
				// first run, check only blocking of ch or lock
				z3Sys := NewZ3ForGl()
				z3Sat1 := z3Sys.Z3Main(paths, vecBlockingPos)

				config.BoolChSafety = true
				// second run, also check ch safety bugs
				z3Sys2 := NewZ3ForGl()
				z3Sat2 := z3Sys2.Z3Main(paths, vecBlockingPos)

				foundBug = z3Sat1 || z3Sat2
			} else {
				z3Sys := NewZ3ForGl()
				z3Sat := z3Sys.Z3Main(paths, vecBlockingPos)

				foundBug = z3Sat
			}

			// Report a bug
			if foundBug {
				//z3Sys.PrintAssert()
				config.BugIndexMu.Lock()
				config.BugIndex++
				fmt.Print("----------Bug[")
				fmt.Print(config.BugIndex)
				config.BugIndexMu.Unlock()
				if config.BoolChSafety {
					fmt.Print("]----------\n\tType: BMOC/Channel Safety \tReason: One or multiple channel operation is blocked/panic.\n")
					fmt.Println("-----Blocking/Panic at:")
				} else {
					fmt.Print("]----------\n\tType: BMOC \tReason: One or multiple channel operation is blocked.\n")
					fmt.Println("-----Blocking at:")
				}
				for _, blockPos := range blockPosComb {
					if blockPos.pNodeId != emptyPNodeId {
						inst := paths[blockPos.pathId].Path[blockPos.pNodeId].Node.Instruction()
						str := output.StringIISrc(inst)
						fmt.Print(str)
						PrintedBlockPosStr[str] = struct{}{}
					}
				}

				for _, blockPos := range blockPosComb {
					if blockPos.pNodeId != emptyPNodeId {
						fmt.Println("-----Blocking/Panic Path NO.", blockPos.pathId)
						paths[blockPos.pathId].PrintPPath()
					} else {
						fmt.Println("-----Path NO.", blockPos.pathId, "\tEntry func at:", goroutines[blockPos.pathId].EntryFn.String())
						paths[blockPos.pathId].PrintPPath()
					}
				}

				fmt.Println()

				bugFound = true
				return true
			}

			countBlockPoint++
			return false
		})
	})

	//fmt.Println("=========Total path sets:",countBlockPoint)
	//output.Wait_for_input()
	return bugFound
}

var PrintedBlockPosStr map[string]struct{} = make(map[string]struct{})
