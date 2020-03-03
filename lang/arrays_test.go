package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPrintWithQ(t *testing.T) {
	a := [...]int{1, 2, 3}

	assert.Equal(t, 3, len(a))
	assert.Equal(t, "['\\x01' '\\x02' '\\x03']", fmt.Sprintf("%q", a))
}
