package config

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type PkgPath struct {
	StrPath string
	VecChildrenPath [] string
	NumLock int
	NumSend int
}

type PathStat struct {
	StrPath string
	NumLock int
	NumSend int
}

func ContainsAnyStr(str string, strVec [] string) bool {
	flag := false
	for _, sub := range strVec {
		if strings.Contains(str, sub) {
			flag = true
			break
		}
	}
	return flag
}

func IsPathIncluded(strPath string) bool {
	if strings.Contains(strPath, StrRelativePath) &&
		!ContainsAnyStr(strPath, VecExcludePaths) {
		return true
	}

	return false
}



func ListAllPkgPaths() []string {

	strRoot := StrAbsolutePath + StrRelativePath
	var vecPkgPaths [] string

	err := filepath.Walk(strRoot, func(path string, info os.FileInfo, err error) error {
		if ContainsAnyStr(path, VecExcludePaths) || info.IsDir() == false {
			return nil
		}

		fInfos, err := ioutil.ReadDir(path)
		if err != nil {
			return nil
		}

		flag := false
		for _, finfo := range fInfos {
			if strings.HasSuffix(finfo.Name(),".go") {
				flag = true
				break
			}
		}
		if !flag {
			return nil
		}

		vecPkgPaths = append(vecPkgPaths, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return vecPkgPaths
}

func CountOccurrenceFile( strFileName string, strQuery string) int {

	file, err := os.Open(strFileName)
	defer file.Close()

	if err != nil {
		return 0
	}

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), strQuery) {
			total ++
		}
	}


	return total
}

func LastButOneIndex(str string, strQuery string) int {
	str = str[:strings.LastIndex(str,strQuery)]
	return strings.LastIndex(str, strQuery)
}


func CountOccurrence(strRootPath string, strQuery string ) int {
	total := 0

	err := filepath.Walk(strRootPath, func(path string, finfo os.FileInfo, err error) error {
		if !finfo.IsDir() && LastButOneIndex(path,"/") == strings.LastIndex(strRootPath,"/") {
			total += CountOccurrenceFile(path, strQuery)
		}
		return nil
	})

	if err != nil {
		return 0
	}

	return total
}

func CountOccurrenceRecursive(strRootPath string, strQuery string) int {

	total := 0

	err := filepath.Walk(strRootPath, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			total += CountOccurrenceFile(path, strQuery)
		}
		return nil
	})


	if err != nil {
		return 0
	}

	return total
}

func ListWorthyPaths() (wPaths [] PkgPath) {

	wPaths = *new([] PkgPath)
	vecPathStats := *new([]PathStat)

	vecPkgPaths := ListAllPkgPaths()

	for _, path := range vecPkgPaths {
		numLock := CountOccurrence( path, ".Lock()")
		numSend := CountOccurrence(path, "<-")

		//fmt.Println(path, numLock, numSend)
		pathStat := PathStat{
			path,
			numLock,
			numSend,
		}
		vecPathStats = append(vecPathStats,pathStat)
	}

	sort.Slice(vecPathStats, func(i, j int) bool {
		return ( 1 * vecPathStats[i].NumLock + 1 * vecPathStats[i].NumSend) > ( 1 * vecPathStats[j].NumLock + 1 * vecPathStats[j].NumSend)
	})

	VecPathStats = vecPathStats

outer:
	for _, pathStat := range vecPathStats {

		index := 0

		for index < len(wPaths) {
			if strings.HasPrefix(pathStat.StrPath, wPaths[index].StrPath) { //this path is a child of an existing path in worthy_paths
				wPaths[index].VecChildrenPath = append(wPaths[index].VecChildrenPath, pathStat.StrPath)
				continue outer
			}

			if strings.HasPrefix(wPaths[index].StrPath, pathStat.StrPath) { //this path is a parent of an existing path in worthy_paths
				break
			}

			index += 1
		}

		newPkgPath := PkgPath{
			StrPath:         pathStat.StrPath,
			VecChildrenPath: []string{},
			NumLock:         0,
			NumSend:         0,
		}

		if index < len(wPaths) {
			newPkgPath.VecChildrenPath = append(newPkgPath.VecChildrenPath, wPaths[index].StrPath)
			for _, str := range wPaths[index].VecChildrenPath {
				newPkgPath.VecChildrenPath = append(newPkgPath.VecChildrenPath, str)
			}
		}

		wPaths = append(wPaths, newPkgPath)

	}

	for index,_ := range wPaths {
		wPaths[index].NumLock= CountOccurrenceRecursive(wPaths[index].StrPath, ".Lock()")
		wPaths[index].NumSend = CountOccurrenceRecursive(wPaths[index].StrPath, "<-")
		wPaths[index].StrPath = strings.ReplaceAll(wPaths[index].StrPath, StrAbsolutePath,"")

		for j,_ := range wPaths[index].VecChildrenPath {
			wPaths[index].VecChildrenPath[j] = strings.ReplaceAll(wPaths[index].VecChildrenPath[j], StrAbsolutePath,"")
		}
	}

	sort.Slice(wPaths, func(i, j int) bool {
		return ( 1 * wPaths[i].NumLock + 1 * wPaths[i].NumSend) > ( 1 * wPaths[j].NumLock + 1 * wPaths[j].NumSend)
	})

	return
}

