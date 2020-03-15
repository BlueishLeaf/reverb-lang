package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/BlueishLeaf/reverb-lang/evaluator"
	"github.com/BlueishLeaf/reverb-lang/lexer"
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/parser"
	"github.com/BlueishLeaf/reverb-lang/repl"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		code := strings.Join(lines, "\n")
		l := lexer.New(code)
		p := parser.New(l)
		program := p.ParseProgram()
		environment := object.NewEnvironment()
		evaluated := evaluator.Eval(program, environment)
		if evaluated != nil {
			_, _ = io.WriteString(os.Stdout, evaluated.Inspect())
			_, _ = io.WriteString(os.Stdout, "\n")
		}
	} else {
		fmt.Printf("Welcome to the Reverb REPL. Press Ctrl + C to exit.\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
