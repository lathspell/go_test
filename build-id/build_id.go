package main

import "fmt"

// compile passing -ldflags "-X main.build=<build sha1>"
var build string

func main() {
	fmt.Printf("Version: %s\n", build)
}