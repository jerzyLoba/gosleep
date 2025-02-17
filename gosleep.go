package gosleep

import (
	"sync"
	"time"
)

func SleepConcurrently(d time.Duration, done chan<- bool, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		defer close(done)
		time.Sleep(d)
		done <- true
	}()
}
