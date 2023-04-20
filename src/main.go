package main

import (
	"bufio"
	"fmt"
	"github.com/RICORA/workshop/calculator/evaluator"
	"os"
)

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
