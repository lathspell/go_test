package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Using "-" at the side where whitespaces should be trimmed to create a more readable template
func TestTrimming(t *testing.T) {
	template := "Hello {{if .name -}} {{.name}} {{- else -}} unknown {{- end}}!"
	assert.Equal(t, "Hello Tim!", renderTemplate(template, map[string]string{"name": "Tim"}))
	assert.Equal(t, "Hello unknown!", renderTemplate(template, nil))
}
