package main

import (
	"os"

	"github.com/sapslaj/pl-challenges/go/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
