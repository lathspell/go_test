package lang

// Guidelines for using a pointer vs values: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// There are no classes but "struct" types that contains attributes
type Animal struct {
	Legs   int
	Origin *AnimalOrigin // pointer to separate struct
}

type AnimalOrigin struct {
	Continent string
}

// Methods are not defined inside the struct but as separate functions.
// Methods have a "receiver" specified after the "func" keyword.
func (it Animal) hasFourLegs() bool {
	return it.Legs == 4
}

func (it Animal) setLegsByValue(n int) {
	it.Legs = n
}

func (it *Animal) setLegsByReference(n int) {
	it.Legs = n
}

func TestMethods(t *testing.T) {
	f := Animal{Legs: 4}
	assert.True(t, f.hasFourLegs())
	assert.Nil(t, f.Origin)

	g := Animal{
		Legs:   2,
		Origin: &AnimalOrigin{"Europe"},
	}
	assert.Equal(t, "Europe", g.Origin.Continent)
}

func TestCallingByRefVsByValue(t *testing.T) {
	a := Animal{Legs: 2}
	a.setLegsByValue(9)
	assert.Equal(t, 2, a.Legs) // no change

	b := Animal{Legs: 2}
	b.setLegsByReference(9)
	assert.Equal(t, 9, b.Legs) // change
}
