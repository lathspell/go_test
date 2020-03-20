package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionCall(t *testing.T) {
	template := `0x{{ printf "%02x" . }}`
	assert.Equal(t, "0x41", renderTemplate(template, 65))
}

func TestPipeline(t *testing.T) {
	template := `0x{{ . | printf "%02x" }}`
	assert.Equal(t, "0x41", renderTemplate(template, 65))
}

func TestLongPipeline(t *testing.T) {
	template := `{{ . | printf "%02x" | printf "0x%s" }}`
	assert.Equal(t, "0x41", renderTemplate(template, 65))
}

func TestChainingMixedWithPipeline(t *testing.T) {
	template := `{{ printf "0x%s" ( . | printf "%02x" ) }}`
	assert.Equal(t, "0x41", renderTemplate(template, 65))
}
