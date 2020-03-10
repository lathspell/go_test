package concurrency

import (
	"sync"
	"testing"
	"time"
)

var counter2 = 1
var counter2Mutex sync.Mutex

// Uses Mutex to access "counter2"
func blubber2(t *testing.T, id string) {
	for i:=0; i<10; i++ {
		counter2Mutex.Lock()
		t.Logf("%s: %s blubbers about %d", time.Now().Format(time.RFC3339Nano), id, counter2)
		counter2++
		counter2Mutex.Unlock()
		time.Sleep(50 * time.Millisecond)
	}
}

// Uses WaitGroup to keep track of its go routines
func TestConcurrency2(t *testing.T) {
	wg := sync.WaitGroup {}
	wg.Add(2)

	go func() {
		blubber2(t, "A")
		wg.Done()
	}()

	go func() {
		blubber2(t, "B")
		wg.Done()
	}()

	wg.Wait()
}
