package templates

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("test").Parse("Hello {{.name}}!")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	values := map[string]string{"name": "Tim"}

	tmpl.Execute(buf, values)

	assert.Equal(t, "Hello Tim!", buf.String())
}
