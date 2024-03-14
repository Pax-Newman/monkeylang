package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		lex := lexer.New(line)

		for tok := lex.Next(); tok.Type != token.EOF; tok = lex.Next() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
