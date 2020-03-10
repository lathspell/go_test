package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelsWithBuffers(t *testing.T) {
	ch := make(chan int, 2 /* buffer size */)

	assert.Equal(t, 2, cap(ch)) // channel capacity

	ch <- 1
	ch <- 2
	// ch <- 3 would produce "fatal error: deadlock"
}
