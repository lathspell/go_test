package lang

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPrintWithQ(t *testing.T) {
	a := [...]int{1, 2, 3}

	assert.Equal(t, 3, len(a))
	assert.Equal(t, "['\\x01' '\\x02' '\\x03']", fmt.Sprintf("%q", a))
	assert.Equal(t, "[3]int", reflect.TypeOf(a).String())
}

func TestArraySlices(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, 3, len(s))
	assert.Equal(t, "[]int", reflect.TypeOf(s).String())

	// Array Slices are "n..n-1" so "1:1" translates to elements "1..0"
	assert.Equal(t, 0, len(s[1:1]))
	assert.Equal(t, []int{2}, s[1:2])

	// Create Slice using make() instead of literal
	s2 := make([]string, 3)
	s2[1] = "b"
	assert.Equal(t, []string{"", "b", ""}, s2)

	// Grow slice, create a bigger underlying array, if necessary
	s2 = append(s2, "c")
	assert.Equal(t, []string{"", "b", "", "c"}, s2)
}

func TestEmptySlices(t *testing.T) {
	var a []int // a points to NIL
	assert.Equal(t, 0, len(a))

	b := []int{} // b points to an empty array slice
	assert.Equal(t, 0, len(b))

	// both can be used to append values
	a = append(a, 3)
	b = append(b, 3)
	assert.Equal(t, a, b)
}
