package main

import (
	"fmt"
	"os"
	"monkey/repl"
)

func main() {
	fmt.Printf("matthew's programming language (Jan. 17, 2026)\n")
	fmt.Printf("Type 'help' for more information about the language.\n")
	repl.Start(os.Stdin, os.Stdout)
}
