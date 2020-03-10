package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForLoops(t *testing.T) {
	counter := 0

	for i := 0; i < 42; i++ {
		counter++
	}

	assert.Equal(t, 42, counter)
}

func TestForLoops2(t *testing.T) {
	a := []byte("abcd")
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] // multi-assignment makes temporary variables unnecessary
	}
	assert.Equal(t, "dcba", string(a))
}

func TestWhileLoops(t *testing.T) {
	counter := 0

	for counter < 42 {
		counter++
	}

	assert.Equal(t, 42, counter)
}

func TestEndlessLoop(t *testing.T) {
	counter := 0

	for {
		counter++

		if counter == 42 {
			break
		}
	}

	assert.Equal(t, 42, counter)
}

func TestSwitch(t *testing.T) {
	n := 4
	switch n {
	case 1, 2:
		// Multiple cases in one line
		t.Fail()
	case 4:
		// OK
	case 5 + 6:
		// Expressions as case values
		t.Fail()
	case 9:
		// Fallthrough must be explicitly stated
		fallthrough
	case 10:
		t.Fail() // for 9 and 10!
	default:
		t.Fail()
	}
}

func helpTestDeferClose(s *[]byte) {
	(*s)[0] = '*'
}

func helpTestDefer() (out []byte) {
	defer helpTestDeferClose(&out) // will be called just after "return"
	return []byte("_abc")
}

// Example for "defer" (uses byte slice as strings are somewhat special)
func TestDefer(t *testing.T) {
	assert.Equal(t, "*abc", string(helpTestDefer()))
}
