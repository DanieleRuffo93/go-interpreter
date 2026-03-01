package main

import (
	"fmt"
	"github.com/DanieleRuffo93/go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Moneky Programming Language!\n", user.Username)
	fmt.Printf("Type in some commands to see how it will be tokenized from the Lexer\n")
	repl.Start(os.Stdin, os.Stdout)
}
