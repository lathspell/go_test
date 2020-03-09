package lang

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Custom error class
type CalculationError struct {
	ts time.Time
	msg string
}

// Make custom error class compliant a valid "error"
func (c CalculationError) Error() string {
	return fmt.Sprintf("%s: %s", c.ts.Format(time.RFC3339), c.msg)
}

// The signature just declares the function as returning an "error"
// while in fact it does return our custom error type
func calc() error {
	return CalculationError{
		ts:  time.Now(),
		msg: "Oops!",
	}
}

func TestCustomError(t *testing.T ) {
	e := calc()
	t.Logf("Custom error: %s", e.Error())
	assert.True(t, strings.HasSuffix(e.Error(), "Oops!"))
}