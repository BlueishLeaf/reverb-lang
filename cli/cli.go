package cli

import (
	"bufio"
	"fmt"
	"github.com/BlueishLeaf/reverb-lang/evaluator"
	"github.com/BlueishLeaf/reverb-lang/lexer"
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/parser"
	"github.com/BlueishLeaf/reverb-lang/repl"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func Start() {
	app := &cli.App{
		Name:     "Reverb CLi",
		Usage:    "Provides tools for use with the Reverb programming language.",
		Version:  "1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Killian Kelly",
				Email: "killiank360@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "reverb",
				Usage: "Compiles or interprets a reverb file.",
				Action: func(context *cli.Context) error {
					numArgs := context.Args().Len()
					var fileName string
					var flag string
					if numArgs == 0 {
						// Open REPL
						fmt.Printf("Welcome to the Reverb REPL. Press Ctrl + D to exit.\n")
						repl.Start(os.Stdin, os.Stdout)
					} else if numArgs == 1 {
						// Interpret Reverb fileName
						fileName = context.Args().Get(0)
						file, err := os.Open(fileName)
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
					} else if numArgs == 2 {
						// Compile Reverb fileName to mp3
						fileName = context.Args().Get(0)
						flag = context.Args().Get(1)
						fmt.Println(fileName)
						fmt.Println(flag)
					} else {
						fmt.Println("Invalid number of arguments. Try again.")
					}
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
