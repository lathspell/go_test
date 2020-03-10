package lang

import (
	format "fmt"  // Import package "fmt" using alias name "format"
	. "math"      // import into current namespace, see function Sqrt() below
	_ "math/rand" // suppress errors/warnings about unused package
	"testing"     // regular import

	"github.com/stretchr/testify/assert"
)

func TestImports(t *testing.T) {
	assert.Equal(t, "42", format.Sprintf("%d", 42))
	assert.Equal(t, 3.0, Sqrt(9))
}
