package gosleep

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSleepConcurrently(t *testing.T) {
	var wg sync.WaitGroup

	done := make(chan bool)
	wg.Add(1)

	SleepConcurrently(1*time.Second, done, &wg)

	wg.Wait()

	select {
	case <-done:
		t.Log("SleepConcurrently completed")
	case <-time.After(time.Second + (time.Duration(rand.Intn(300)+200) * time.Millisecond)):
		t.Error("SleepConcurrently did not complete")
	}

}
