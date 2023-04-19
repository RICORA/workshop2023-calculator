package evaluator

import "fmt"

func Eval(tokens []Token) int {
	res, next := expr(tokens)
	if len(next) > 0 {
		errStr := fmt.Sprintf("syntax error: unnecessary token found, the type is %s", next[0].kind)
		panic(errStr)
	}
	return res
}

func eat(tokens []Token) []Token {
	if len(tokens) > 0 {
		return tokens[1:]
	} else {
		panic("syntax error: tried to eat nil tokens")
	}
}

func expr(tokens []Token) (int, []Token) {
	left, next := term(tokens)
	return exprOpt(left, next)
}

func exprOpt(left int, tokens []Token) (int, []Token) {
	if len(tokens) == 0 {
		return left, tokens
	}
	switch tokens[0].kind {
	case ADD:
		right, next := term(eat(tokens))
		return exprOpt(left+right, next)
	case SUB:
		right, next := term(eat(tokens))
		return exprOpt(left-right, next)
	default:
		return left, tokens
	}
}

func term(tokens []Token) (int, []Token) {
	left, next := factor(tokens)
	return termOpt(left, next)
}

func termOpt(left int, tokens []Token) (int, []Token) {
	if len(tokens) == 0 {
		return left, tokens
	}
	switch tokens[0].kind {
	case MUL:
		right, next := factor(eat(tokens))
		return termOpt(left*right, next)
	case DIV:
		right, next := factor(eat(tokens))
		return termOpt(left/right, next)
	default:
		return left, tokens
	}
}

func factor(tokens []Token) (int, []Token) {
	if len(tokens) == 0 {
		panic("syntax error: numeric or ( token expected but no token found")
	}
	switch tokens[0].kind {
	case NUMERIC:
		return tokens[0].value, eat(tokens)
	case LPARENTHESIS:
		res, next := expr(eat(tokens))
		if len(next) > 0 && next[0].kind == RPARENTHESIS {
			return res, eat(next)
		} else {
			panic("syntax error: ) not found")
		}
	default:
		errStr := fmt.Sprintf("syntax error: numeric token or ( expected but found the type is ", tokens[0].kind)
		panic(errStr)
	}
}
