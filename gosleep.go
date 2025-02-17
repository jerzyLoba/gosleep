package gosleep

import (
	"sync"
	"time"
)

func SleepConcurrently(d time.Duration, done chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		defer close(done)
		time.Sleep(d)
		done <- true
	}()
}
