# Monkey Lang

Go implementation of the Monkey Programming language from the books _Writing a Interpreter in Go_ and _Writing a Compiler in Go_. 

## Requirments

* go 1.22

## Usage

### Build Binary
```sh
$ go build -o bin/monkey main.go
```


### Running the REPL

```sh
$ bin/monkey [-interpt]
```

### Parse a file

```sh
$ bin/monkey [-interpt] -filename file_name
```

## Example 

The [example.monkey](https://github.com/hunterwilkins2/monkey/blob/master/example.monkey) is a sample program written in monkey lang.  
It includes functions to `map`, `filter`, `reduce` and, `search` arrays and map operations!

Run this file with:

```sh
$ bin/monkey -filename example.monkey
```

## Benchmark

A benchmark between the interpreter and bytecode compiler shows a 4x in performance of the bytecode compiler vs the interpreter.

### Run benchmark

```sh
$ go build -o bin/fibonacci ./benchmark
$ ./bin/fibonacci
starting interpreter `fibonacci(35);`...
starting compiler `fibonacci(35);`...

engine=interpreter, result=9227465, duration=8.709372729s
engine=compiler, result=9227465, duration=2.431848157s
```

### System Specs

```
CPU: 13th Gen Intel i7-13700K (24) @ 5.300GHz 
GPU: AMD ATI Radeon RX 6700/6700 XT/6750 XT / 6800M/6850M XT 
Memory: 3847MiB / 31863MiB 
```

## Features

* C-like syntax
* Interger, boolean, and string primitive data types
* Array and hashtables
* Variable bindings
* If-else conditions
* First class higher order functions and closures
* Built-ins

### Types

| Type      | Examples         |
|-----------|------------------|
| `int` | `0`, `-42`, `99`     |
| `boolean` | `true`, `false`  |
| `strings` | `"Hello, World!` |
| `array`   | `[]`, `[1, 2 ,3]`, `[true, "", fn(x) { x }]` |
| `hashes`  | `{}`, `{"key": "value"}`, `{true: "yup!", 1: fn(x) { x + 1}}` |

### Integer Arithmetis
```
>> let result = (5 + 10 * 2 + 15 / 3) * 2 + -10;
>> result;
55
```

### Let and Return Statements
```
>> let one = 1;
1
>> return 42;
42
```
### If-Else Conditionals

```
>> if(1 < 2) {
    return true;
} else {
    return "Borked."
}
true
```
### Functions and Recursion
```
>> let factorial = fn(x) {
    if (x < 1) {
        return 1;
    }
    return factorial(x - 1) * x;
};
>> factorial(5);
125
```

### Build-ins

```
>> len("abc")
3
>> len([1, 2])
2

>> push([1], 2)
[1, 2]

>> first([1, 2, 3])
1
 
>> last([1, 2, 3])
3

>> rest([1, 2, 3])
[2, 3]
```