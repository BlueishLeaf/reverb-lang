package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/BlueishLeaf/reverb-lang/evaluator"
	"github.com/BlueishLeaf/reverb-lang/lexer"
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/parser"
	"github.com/BlueishLeaf/reverb-lang/repl"
	"github.com/urfave/cli/v2"
)

func main() {
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
				Name:  "run",
				Usage: "Interprets a specified reverb file",
				Action: func(ctx *cli.Context) error {
					numArgs := ctx.Args().Len()
					var fileName string
					if numArgs == 1 {
						fileName = ctx.Args().Get(0)
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
					} else {
						fmt.Println("Invalid number of arguments. Try again.")
					}
					return nil
				},
			},
			{
				Name:  "compile",
				Usage: "Compiles a specified reverb file to MP3 format",
				Action: func(ctx *cli.Context) error {
					// TODO: This...
					return nil
				},
			},
			{
				Name:  "repl",
				Usage: "Initiates the reverb REPL environment",
				Action: func(ctx *cli.Context) error {
					fmt.Printf("Welcome to the Reverb REPL. Press Ctrl + D to exit.\n")
					repl.Start(os.Stdin, os.Stdout)
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
