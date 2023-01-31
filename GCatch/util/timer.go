package util

import (
	"time"

	"github.com/system-pclub/GCatch/GCatch/config"
)

type Stopper chan struct{}

// NewStopper creates a channel that is stopped after a certain time period.
func NewStopper() Stopper {
	stop := make(Stopper)

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

// IterateUntilTimeout takes an abort channel, a list of items, and a function that operates over individual items
// and their index. It iterates over every element of the list, and either executes the function or aborts
// if the stop channel has been closed. It returns true if stopped prematurely, or false if it succesfully
// iterated over the entire list.
func IterateUntilTimeout[T any](stop Stopper, ts []T, f func(int, T) bool) bool {
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

// LoopUntilTimeout takes a stopper and a function that is repeatedly executed until the guard returned
// is false, or aborts if when the stopper has been closed. It returns true if stopped prematurely, or false if
// it all iterations succeed before the timer expires.
func (stop Stopper) LoopUntilTimeout(f func() bool) bool {
	for f() {
		select {
		case <-stop:
			return true
		default:
		}
	}

	return false
}
