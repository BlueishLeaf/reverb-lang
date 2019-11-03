package main

import (
	"fmt"
	"github.com/BlueishLeaf/reverb-lang/repl"
	"os"
)

func main() {
	fmt.Printf("Welcome to the Reverb REPL.\n")
	repl.Start(os.Stdin, os.Stdout)
}