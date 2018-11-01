package main

import (
	"Bellamy/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
