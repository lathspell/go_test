package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * In Go interfaces are used to declare methods that can work on different structs
 * as long as they offer certain methods.
 * The structs do not have to declare their conformance to an interface explicitly,
 * as long as the proper methods are declare, the compiler is happy.
 */

// Interface for entities with some kind of name
type Person interface {
	getName() string
}

// Generic method that accepts anything that has a name
func greet(p Person) string {
	return "Hello " + p.getName()
}

// Employee

type Employee struct {
	fullName string
}

func (e Employee) getName() string {
	return e.fullName
}

// Customer

type Customer struct {
	companyName string
}

func (c Customer) getName() string {
	return c.companyName
}

func TestImplicitInterfaces(t *testing.T) {
	e := Employee{fullName:"Tim Taler"}
	assert.Equal(t, "Hello Tim Taler", greet(e))

	c := Customer{companyName:"ACME"}
	assert.Equal(t, "Hello ACME", greet(c))
}
