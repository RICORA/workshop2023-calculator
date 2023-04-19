package main

import (
	"bufio"
	"fmt"
	"github.com/RICORA/workshop/calculator/evaluator"
	"os"
)

func eat(runes []rune) (rune, []rune) {
	switch {
	case len(runes) > 2:
		return runes[0], runes[1:]
	case len(runes) == 1:
		return runes[0], nil
	default:
		panic("syntax error")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	resScan := scanner.Scan()
	if !resScan {
		panic("failed to read a formula.")
	}
	formula := scanner.Text()
	tokens := evaluator.Lex(formula)
	res := evaluator.Eval(tokens)
	fmt.Println(res)
}
