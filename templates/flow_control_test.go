package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	template := "Hello {{if .name}}{{.name}}{{else}}unknown{{end}}!"
	assert.Equal(t, "Hello Tim!", renderTemplate(template, map[string]string{"name": "Tim"}))
	assert.Equal(t, "Hello unknown!", renderTemplate(template, nil))
}

func TestRange(t *testing.T) {
	// the first dot refers the passed values as an int array has no members
	// the second dot refers to the loop variable as its only an int with no inner member attributes
	template := "{{range .}}calling {{.}} {{end}}"
	assert.Equal(t, "calling 1 calling 2 calling 3 ", renderTemplate(template, []int{1, 2, 3}))
}

func TestWith(t *testing.T) {
	template := "{{with .}}{{.Name}} calls {{range .Contacts}}{{.Name}}, {{end}}{{end}}"
	person := Person{Name: "Tim", Contacts: []Person{{Name: "Tom"}}}
	assert.Equal(t, "Tim calls Tom, ", renderTemplate(template, person))
}
