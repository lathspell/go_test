package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Identifiable interface {
	GetId() int
}

type Foo struct {
	Id int
}

type Bar struct {
	Id int
}

func (it Foo) GetId() int {
	return it.Id
}

func isIdentifiable(i interface{}) bool {
	_, isIdentifiable := i.(Identifiable)
	return isIdentifiable
}

/* After declaring the `GetId()` method for the Foo type, the compiler recognizes that Foo now implements the Identifiable interface. */
func TestIfInterface(t *testing.T) {
	f := Foo{123}
	t.Logf("f: %+v", f)
	assert.True(t, isIdentifiable(f))
	assert.Equal(t, 123, f.GetId())

	b := Bar{ 123 }
	t.Logf("b: %+v", b)
	assert.False(t, isIdentifiable(b))
}
