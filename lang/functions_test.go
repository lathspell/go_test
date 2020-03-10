package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousFunction(t *testing.T) {
	adder := func (a,b int) int {
		return a+b
	}

	assert.Equal(t, 5, adder(2,3))
}

func TestImmediatelyInvokedFunctionExpression(t *testing.T) {
	sum23 := func (a,b int) int {
		return a+b
	}(2,3)

	assert.Equal(t, 5, sum23)
}
