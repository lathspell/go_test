package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeSequenceSearch(t *testing.T) {
	ch := make(chan int)
	go Walk(NewTree(1), ch)

	var actual []int
	for x := range ch {
		actual = append(actual, x)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, actual)
}

func TestBinaryTreeEquality(t *testing.T) {
	assert.True(t, Same(NewTree(1), NewTree(1)))
	assert.False(t, Same(NewTree(1), NewTree(2)))
}
