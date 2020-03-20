package templates

import (
	"bytes"
	"text/template"
)

func renderTemplate(source string, values interface{}) string {
	parsed, err := template.New("test").Parse(source)
	if err != nil {
		panic(err.Error())
	}

	buf := new(bytes.Buffer)
	parsed.Execute(buf, values)
	return buf.String()
}
