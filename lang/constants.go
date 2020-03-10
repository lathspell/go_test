package lang

import "fmt"

// Constants are also written using Camelcase!
const ABitOfPi = 3.14

// A list of constants can be grouped just like dependencies
const (
	Category1           = 201
	Category2           = 456
	Category3           = 35
	packagePrivateConst = 3
)

// Often type aliases are used to make clear what a parameter value should be
type AsciiChar int

// Iota tells the compiler to increment following constants by 1
const (
	AsciiA AsciiChar = 65 + iota
	AsciiB
	AsciiC
	_ // skipping
	_ // skipping
	AsciiE
)

type Direction int

const (
	DirRight Direction = 1 << iota // using iota for expressions
	DirLeft
	DirBackward
	DirForward
)

func GetAsciiCharAsString(a AsciiChar) string {
	return fmt.Sprintf("%c", a)
}
