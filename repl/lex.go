package repl

import (
	"bellamy/lexer"
	"bellamy/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "--> "

func StartLexRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			out.Write([]byte(fmt.Sprintf("%+v\n", t)))

		}
	}
}
