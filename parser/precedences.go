package parser

import "bellamy/token"

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myFn(x)
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NE:       EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.PERIOD:   CALL,
	token.LBRACKET: INDEX,
}
