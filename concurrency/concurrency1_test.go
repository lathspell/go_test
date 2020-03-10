package concurrency

import (
	"sync"
	"testing"
	"time"
)

var counter1 = 1
var counter1Mutex = sync.Mutex{}

// Uses Mutex to access "counter1"
func blubber1(t *testing.T, id string) {
	for i:=0; i<10; i++ {
		counter1Mutex.Lock()
		t.Logf("%s blubbers about %d", id, counter1)
		counter1++
		counter1Mutex.Unlock()

		time.Sleep(50 * time.Millisecond)
	}
}

// WRONG: does not keep track if all go routines ended
func TestConcurrency1(t *testing.T) {
	go blubber1(t, "A")
	go blubber1(t, "B")
	time.Sleep(1 * time.Second)
}
