package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// First, create an alias for "string" as it does not belong to our package
type MyString string

// Second, create the extension function on our alias
func (it MyString) MyCount() int {
	return len(it)
}

// Error: "Cannot define method on non-local type"
// func (it string) MyCount() int { ... }

func TestTypeAliasing(t *testing.T) {
	// using our alias directly
	myStr := MyString("test")
	assert.Equal(t, 4, myStr.MyCount())

	// using casting
	origString := "Hello"
	aliasedString := MyString(origString)
	assert.Equal(t, 5, aliasedString.MyCount())
}
