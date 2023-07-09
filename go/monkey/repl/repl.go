package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sapslaj/pl-challenges/go/monkey/lexer"
	"github.com/sapslaj/pl-challenges/go/monkey/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("\n>> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
