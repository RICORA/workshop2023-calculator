package evaluator

import (
	"fmt"
	"strconv"
)

type TokenType int

const (
	NUMERIC TokenType = iota
	ADD
	SUB
	MUL
	DIV
	LPARENTHESIS
	RPARENTHESIS
)

type Token struct {
	kind  TokenType
	value int
}

func Lex(formula string) []Token {
	tokens := make([]Token, 0)
	runes := []rune(formula)
	for i := 0; i < len(runes); {
		switch runes[i] {
		case ' ':
			i += 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			add, str := lexNumeric(runes[i:])
			value, err := strconv.Atoi(str)
			if err != nil {
				errStr := fmt.Sprintf("invalid numeric error: %s", str)
				panic(errStr)
			}
			token := Token{NUMERIC, value}
			tokens = append(tokens, token)
			i += add
		case '+':
			token := Token{ADD, 0}
			tokens = append(tokens, token)
			i += 1
		case '-':
			token := Token{SUB, 0}
			tokens = append(tokens, token)
			i += 1
		case '*':
			token := Token{MUL, 0}
			tokens = append(tokens, token)
			i += 1
		case '/':
			token := Token{DIV, 0}
			tokens = append(tokens, token)
			i += 1
		case '(':
			token := Token{LPARENTHESIS, 0}
			tokens = append(tokens, token)
			i += 1
		case ')':
			token := Token{RPARENTHESIS, 0}
			tokens = append(tokens, token)
			i += 1
		default:
			errStr := fmt.Sprintf("invalid character error: %c", runes[i])
			panic(errStr)
		}
	}
	return tokens
}

func lexNumeric(runes []rune) (int, string) {
	for i, v := range runes {
		switch v {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			return i, string(runes[:i])
		}
	}
	return len(runes), string(runes)
}
