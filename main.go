package main

import (
	"fmt"
	"golang.org/x/exp/ebnf"
	"strings"
)

func main() {
	// Define the grammar for the arithmetic expressions DSL
	grammar := `
		expr     = term { "+" term } .
		term     = factor { "*" factor } .
		factor   = number | "(" expr ")" .
		number   = digit { digit } .
		digit    = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" .
	`

	// Parse the grammar
	parser, err := ebnf.Parse("arithmetic", grammar)
	if err != nil {
		fmt.Println("Error parsing grammar:", err)
		return
	}

	// Input expression to parse
	expression := "2 + (3 * 4)"

	// Parse the expression using the generated parser
	ast, err := ebnf.ParseAST(parser, strings.NewReader(expression))
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	// Print the abstract syntax tree (AST)
	fmt.Println("AST:", ast)
}
