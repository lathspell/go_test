package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mail struct {
	From string
	To   string
	Body string
}

func TestBlocks(t *testing.T) {
	// FROM and TO are simple templates
	// MAIL is a composite template
	// the last line is the one that actually renders anything
	// the "-" are used to strip whitespaces at both ends of the rendered text
	template := `
{{- define "FROM" -}} From: {{.From}} {{- end -}}
{{- define "TO" -}} To: {{.To}} {{- end -}}
{{- define "MAIL" -}}
{{ template "FROM" . }}
{{ template "TO" . }}

{{ .Body }}
{{- end -}}
{{ template "MAIL" . }}
`

	mail := mail{"Tim", "Tom", "blah"}
	expected := "From: Tim\nTo: Tom\n\nblah\n"

	assert.Equal(t, expected, renderTemplate(template, mail))
}
