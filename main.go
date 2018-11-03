package main

import (
	"bellamy/repl"
	"os"
)

func main() {
	repl.StartParseRepl(os.Stdin, os.Stdout)
}
