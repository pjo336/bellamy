package main

import (
	"bellamy/repl"
	"os"
)

func main() {
	repl.StartEvalRepl(os.Stdin, os.Stdout)
	//repl.StartParseRepl(os.Stdin, os.Stdout)
}
