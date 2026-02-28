> 🚧 Work in progress

A tree-walking interpreter implemented in Go, built for learning purposes.

## Goal

This project exists to deeply understand how interpreters work by implementing one from scratch. Rather than relying on parser generators or external tools, everything is built by hand: lexer, parser, AST, and evaluator.

The project follows [**Writing An Interpreter In Go**](https://interpreterbook.com/) by Thorsten Ball, a highly recommended resource in the Go and programming languages community. The book guides the reader through building a complete interpreter for the Monkey programming language, step by step.

## What this project covers

Starting from zero, the interpreter is built incrementally:

- **Lexer** — tokenizes raw source code into meaningful symbols
- **Parser** — builds an Abstract Syntax Tree (AST) using a Pratt parser (top-down operator precedence)
- **AST** — the in-memory representation of the program structure
- **Evaluator** — walks the AST and executes the program

## The language

The interpreter targets **Monkey**, a small language designed specifically for this book. It supports variables, functions, closures, conditionals, and more.

**Monkey** has the following features:
- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

Example:

```
let add = fn(x, y) {
  x + y;
};

let result = add(5, 10);
```

I am following the book chapter by chapter. The implementation may diverge from the book where it makes sense to explore alternative approaches or extend the language.

## Reference

- 📖 [Writing An Interpreter In Go — Thorsten Ball](https://interpreterbook.com/)

## License

MIT
