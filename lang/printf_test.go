package lang

import (
	"fmt"
	"testing"

	"github.com/golang/go/src/pkg/regexp"
	"github.com/stretchr/testify/assert"
)

func TestPrintf(t *testing.T) {
	ai := [...]int {1,2,3}
	aSlice := ai[0:1]
	m := map[string]int { "A": 65 }

	// %v - value in default format
	assert.Equal(t, "map[A:65]", fmt.Sprintf("%v", m))

	// %#v - value in Go-syntax
	assert.Equal(t, "map[string]int{\"A\":65}", fmt.Sprintf("%#v", m))

	// %q - quoted values
	assert.Equal(t, "\"abc\"", fmt.Sprintf("%q", "abc"))
	assert.Equal(t, "['\\x01' '\\x02' '\\x03']", fmt.Sprintf("%q", ai))

	// %T - type
	assert.Equal(t, "func(*testing.T)", fmt.Sprintf("%T", TestPrintf))
	assert.Equal(t, "int", fmt.Sprintf("%T", 42))
	assert.Equal(t, "[3]int", fmt.Sprintf("%T", ai))

	// %b - binary
	assert.Equal(t, "0100", fmt.Sprintf("%04b", 4))

	// %x - hex
	assert.Equal(t, "2d", fmt.Sprintf("%x", 45))

	// %c - Unicode code point
	assert.Equal(t, "A", fmt.Sprintf("%c", 65))

	// %p - pointer
	addr := fmt.Sprintf("%p", aSlice)
	wasMatch,_ := regexp.MatchString("^0[a-z0-9]+$", addr)
	assert.True(t, wasMatch)
}
