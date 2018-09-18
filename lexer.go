package parser

import (
	"bufio"
)

type lexer struct {
	r *bufio.Reader
}

func newLexer(r *bufio.Reader) *lexer {
	return &lexer{r: r}
}

func (l *lexer) scanWithoutWS() *token {
	t := l.scan()
	if t.typ == ws {
		return l.scanWithoutWS()
	}
	return t
}

func (l *lexer) scan() *token {
	r, _, err := l.r.ReadRune()
	if err != nil {
		return &token{typ: eof}
	}
	if r == ' ' || r == '\t' || r == '\n' {
		l.r.UnreadRune()
		return l.scaneWs()
	}
	if r == '(' {
		return &token{typ: openBracket, lexeme: string(r)}
	}
	if r == ')' {
		return &token{typ: closeBracket, lexeme: string(r)}
	}
	if r == ',' {
		return &token{typ: comma, lexeme: string(r)}
	}

	if isAlphabet(r) {
		l.r.UnreadRune()
		return l.scanIdent()
	}

	return &token{typ: illegal, lexeme: string(r)}
}

func (l *lexer) scaneWs() *token {
	lexeme := ""
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			l.r.UnreadRune()
			break
		}

		if r != ' ' && r != '\t' && r != '\n' {
			l.r.UnreadRune()
			break
		}

		lexeme += string(r)
	}

	return &token{
		typ:    ws,
		lexeme: lexeme,
	}
}

func (l *lexer) scanIdent() *token {
	lexeme := ""
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			l.r.UnreadRune()
			break
		}

		if !isAlphabet(r) && !isNumber(r) {
			l.r.UnreadRune()
			break
		}

		lexeme += string(r)
	}

	if typ := keywords[lexeme]; typ != 0 {
		return &token{typ: typ, lexeme: lexeme}
	}

	if typ := dataType[lexeme]; typ != 0 {
		return &token{typ: typ, lexeme: lexeme}
	}

	return &token{
		typ:    identLit,
		lexeme: lexeme,
	}
}

func isAlphabet(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
