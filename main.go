package main

import (
	"fmt"
	"os"

	"github.com/iashyam/bclang/repl"
)

func main() {
	fmt.Println("Welcome to bclang 0.0.1")
	fmt.Println("enter exit to well... exit")

	repl.Start(os.Stdin, os.Stdout)
}
