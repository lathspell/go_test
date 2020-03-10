package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, 3.14, ABitOfPi)
	assert.Equal(t, 456, Category2)
	assert.Equal(t, 3, packagePrivateConst)
	assert.Equal(t, Direction(10), DirLeft | DirForward)
}

func TestGetAsciiCharAsString(t *testing.T) {
	// The compiler enforces type checking but in the end AsciiChar is just an int
	assert.Equal(t, 67, int(AsciiC))
	assert.Equal(t, AsciiChar(67), AsciiC)
	assert.True(t, 67 == AsciiC)

	assert.Equal(t, "B", GetAsciiCharAsString(66))
	assert.Equal(t, "B", GetAsciiCharAsString(AsciiChar(66)))
	assert.Equal(t, "B", GetAsciiCharAsString(AsciiB))
}
