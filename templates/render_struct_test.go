package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	values := Person{Name: "Tim"}
	actual := renderTemplate("Hello {{.Name}}!", values)
	assert.Equal(t, "Hello Tim!", actual)
}
