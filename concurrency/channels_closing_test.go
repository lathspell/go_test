package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func produceValues(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // range waits forever without this!
}

// "range ch" will read values until the channel is closed by the sender
func TestChannelClosing(t *testing.T) {
	ch := make(chan int)

	go produceValues(ch)

	sum := 0
	for i := range ch {
		sum += i
	}
	assert.Equal(t, 6, sum)
}
