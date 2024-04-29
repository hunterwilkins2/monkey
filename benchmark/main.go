package main

import (
	"fmt"
	"time"

	"github.com/hunterwilkins2/monkey/compiler"
	"github.com/hunterwilkins2/monkey/evaluator"
	"github.com/hunterwilkins2/monkey/lexer"
	"github.com/hunterwilkins2/monkey/object"
	"github.com/hunterwilkins2/monkey/parser"
	"github.com/hunterwilkins2/monkey/vm"
)

var input = `
let fibonacci = fn(x) {
	if (x == 0) {
		0
	} else {
		if (x == 1) {
			return 1;
		} else {
			fibonacci(x - 1) + fibonacci(x - 2);
		}
	}
};
fibonacci(35);
`

func main() {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	fmt.Println("starting interpreter `fibonacci(35);`...")
	env := object.NewEnvironment()
	start := time.Now()
	evalResult := evaluator.Eval(program, env)
	evalDuration := time.Since(start)

	fmt.Println("starting compiler `fibonacci(35);`...")
	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compiler error: %s", err)
		return
	}

	machine := vm.New(comp.ByteCode())
	start = time.Now()

	err = machine.Run()
	if err != nil {
		fmt.Printf("vm error: %s", err)
		return
	}
	vmDuration := time.Since(start)
	vmResult := machine.LastPoppedStackElem()

	fmt.Println()
	fmt.Printf("engine=interpreter, result=%s, duration=%s\n", evalResult.Inspect(), evalDuration)
	fmt.Printf("engine=compiler, result=%s, duration=%s\n", vmResult.Inspect(), vmDuration)
}
