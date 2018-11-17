package main

import (
	"bellamy/repl"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		repl.StartEvalRepl(os.Stdin, os.Stdout)
	} else {
		arg := os.Args[1]
		switch arg[0] {
		case 'p':
			repl.StartParseRepl(os.Stdin, os.Stdout)
		case 'l':
			repl.StartLexRepl(os.Stdin, os.Stdout)
		default:
			repl.StartEvalRepl(os.Stdin, os.Stdout)
		}
	}
}
