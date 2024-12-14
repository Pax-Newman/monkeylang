package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
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

		// for tok := lex.Next(); tok.Type != token.EOF; tok = lex.Next() {
		// 	io.WriteString(out, fmt.Sprintf("%+v\n", tok))
		// }

		prs := parser.New(lex)

		program := prs.Parse()
		if len(prs.Errors()) != 0 {
			printParserErrors(out, prs.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
