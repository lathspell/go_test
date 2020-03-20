package templates

type Person struct {
	Name     string // works only with public attributes
	Contacts []Person
}
