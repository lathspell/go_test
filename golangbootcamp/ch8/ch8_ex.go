package ch8

import "math/rand"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk tree and send values to channel
func Walk(t *Tree, ch chan int) {
	recWalk(t, ch)
	close(ch) // to make "range" happy
}

// Recursive walk from left to right
func recWalk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	// left side
	recWalk(t.Left, ch)
	// send current value
	ch <- t.Value
	// right side
	recWalk(t.Right, ch)
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			// not same size (one channel is closed, the other one not)
			return false
		case !ok1:
			// both channels are empty (ok2 must be equal to ok1 or the first check would have failed)
			// the trees must hold the same value as all previous values passed the third check
			return true
		case x1 != x2:
			// the last values that were just read from each channel are different
			return false
		default:
			// keep iterating
		}
	}
}

func NewTree(k int) *Tree {
	var t *Tree // Tree{} would create a tree with value "0"
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{
			Left:  nil,
			Value: v,
			Right: nil,
		}
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}

	return t
}
