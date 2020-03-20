package main

import (
	"github.com/lathspell/go_test/cli/complex-app/cmd"
)

/* `go build  &&  CMPLX_ENV=beta ./complex-app` */
func main() {
	cmd.CliHandling()
}
