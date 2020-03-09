package ch7

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	y, e := MySqrt(9.0)
	assert.Equal(t, 3.0, y)
	assert.Nil(t, e)
}

func TestMySqrtBad(t *testing.T) {
	y, e := MySqrt(-3.0)
	t.Logf("Error: %s", e.Error())
	assert.Equal(t, 0.0, y)
	assert.NotNil(t, e)
	assert.True(t, strings.Contains(e.Error(), "negative number"))
}
