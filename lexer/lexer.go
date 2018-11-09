package lexer

import (
	"bellamy/token"
	"bellamy/utils"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // read in first byte to initialize
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.FromMultiChar(token.EQ, []byte{ch, l.ch})
		} else {
			t = token.FromChar(token.ASSIGN, l.ch)
		}
	case ';':
		t = token.FromChar(token.SEMICOLON, l.ch)
	case ':':
		t = token.FromChar(token.COLON, l.ch)
	case ',':
		t = token.FromChar(token.COMMA, l.ch)
	case '(':
		t = token.FromChar(token.LPAREN, l.ch)
	case ')':
		t = token.FromChar(token.RPAREN, l.ch)
	case '{':
		t = token.FromChar(token.LBRACE, l.ch)
	case '}':
		t = token.FromChar(token.RBRACE, l.ch)
	case '[':
		t = token.FromChar(token.LBRACKET, l.ch)
	case ']':
		t = token.FromChar(token.RBRACKET, l.ch)

	case '+':
		t = token.FromChar(token.PLUS, l.ch)
	case '-':
		t = token.FromChar(token.MINUS, l.ch)
	case '<':
		t = token.FromChar(token.LT, l.ch)
	case '>':
		t = token.FromChar(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.FromMultiChar(token.NE, []byte{ch, l.ch})
		} else {
			t = token.FromChar(token.BANG, l.ch)
		}
	case '*':
		t = token.FromChar(token.ASTERISK, l.ch)
	case '/':
		t = token.FromChar(token.SLASH, l.ch)
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	case 0:
		t = token.FromChar(token.EOF, '0')
	default:
		if utils.IsLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if utils.IsDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = token.FromChar(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return t
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readString() string {
	b := []byte{}
	l.readChar()
	for l.ch != '"' && l.ch != '0' {
		b = append(b, l.ch)
		l.readChar()
	}
	return string(b)
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for utils.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for utils.IsWhitespace(l.ch) {
		l.readChar() // advance the token ahead
	}
}
