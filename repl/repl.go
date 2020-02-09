package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/BlueishLeaf/reverb-lang/evaluator"
	"github.com/BlueishLeaf/reverb-lang/lexer"
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/parser"
)

// Prompt represents the symbol to prompt the user to enter code
const Prompt = ">>"

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}

// Start initiates the Reverb REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	environment := object.NewEnvironment()
	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, environment)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}
