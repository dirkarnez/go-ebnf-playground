package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/ebnf"
)

func main() {

	grammar := `
		Program = Start .
		Start = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" .
	`
	// Parse the grammar
	parser, err := ebnf.Parse("", strings.NewReader(grammar))
	if err != nil {
		fmt.Println("Error parsing grammar:", err)
		return
	}
	fmt.Println("parser:", parser)
	// Input expression to parse
	expression := `Program`

	// Parse the expression using the generated parser
	err = ebnf.Verify(parser, expression)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}
}
