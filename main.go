package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/hunterwilkins2/monkey/compiler"
	"github.com/hunterwilkins2/monkey/evaluator"
	"github.com/hunterwilkins2/monkey/lexer"
	"github.com/hunterwilkins2/monkey/object"
	"github.com/hunterwilkins2/monkey/parser"
	"github.com/hunterwilkins2/monkey/repl"
	"github.com/hunterwilkins2/monkey/vm"
)

func main() {
	filename := flag.String("filename", "", "file to parse")
	useInterpreter := flag.Bool("interpt", false, "Use interpreter instead of compiler")
	flag.Parse()

	if *filename != "" {
		f, err := os.Open(*filename)
		if err != nil {
			fmt.Printf("could not open file: %s\n", *filename)
		}
		input, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("could not read %s: %v\n", *filename, err)
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

		if *useInterpreter {
			env := object.NewEnvironment()
			evaluated := evaluator.Eval(program, env)
			fmt.Println(evaluated.Inspect())
		} else {
			comp := compiler.New()
			err := comp.Compile(program)
			if err != nil {
				fmt.Printf("could not compile program:\n %s\n", err)
				os.Exit(1)
			}

			machine := vm.New(comp.ByteCode())
			err = machine.Run()
			if err != nil {
				fmt.Printf("error running virtual machine:\n %s\n", err)
				os.Exit(1)
			}

			stackTop := machine.StackTop()
			fmt.Println(stackTop.Inspect())
		}
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		if *useInterpreter {
			repl.StartInterperter(os.Stdin, os.Stdout)
		} else {
			repl.StartCompiler(os.Stdin, os.Stdout)
		}
	}
}
