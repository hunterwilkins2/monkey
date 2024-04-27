package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/hunterwilkins2/monkey/evaluator"
	"github.com/hunterwilkins2/monkey/lexer"
	"github.com/hunterwilkins2/monkey/object"
	"github.com/hunterwilkins2/monkey/parser"
	"github.com/hunterwilkins2/monkey/repl"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("cannot read more than one file")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("could not open file: %s\n", os.Args[1])
		}
		input, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("could not read %s: %v\n", os.Args[1], err)
		}
		l := lexer.New(string(input))
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			fmt.Println("could not parse program:")
			for _, msg := range p.Errors() {
				fmt.Println("\t" + msg)
			}
			os.Exit(1)
		}

		env := object.NewEnvironment()
		evaluated := evaluator.Eval(program, env)
		fmt.Println(evaluated.Inspect())
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)

	}
}
