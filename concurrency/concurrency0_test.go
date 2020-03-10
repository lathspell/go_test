package concurrency

import (
	"testing"
	"time"
)

var counter0 = 1

// WRONG: Does not use locking when accessing counter0!
func blubber0(t *testing.T, id string) {
	for i:=0; i<10; i++ {
		t.Logf("%s blubbers about %d", id, counter0)
		counter0++
		time.Sleep(50 * time.Millisecond)
	}
}

// WRONG: Does not keep track if all go routines have finished!
func TestConcurrency0(t *testing.T) {
	go blubber0(t, "A")
	go blubber0(t, "B")
	time.Sleep(1 * time.Second)
}
