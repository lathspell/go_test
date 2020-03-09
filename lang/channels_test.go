package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func TestChanneledSum(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) // first half
	go sum(a[len(a)/2:], c) // second half

	// receive two values (in random order)
	x, y := <-c, <-c

	assert.Equal(t, 12, x+y)
}
