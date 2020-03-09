package lang

import (
	"fmt"
	"testing"
	"time"
)

// The "time.After()" method creates a channel that yields a value after a certain time
func waitReallyLong(ch chan int) {
	for {
		select {

		case x := <-ch:
			fmt.Printf("Received %d\n", x)

		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Timeout!!!\n")
			return
		}
	}
}

func TestChannelTimeout(t *testing.T) {
	ch := make(chan int)
	go waitReallyLong(ch)

	ch <- 1
	ch <- 2
	time.Sleep(1 * time.Second)
}
