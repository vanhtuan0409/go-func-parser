package parser

import (
	"bufio"
	"strings"
)

type Param struct {
	Name string
	Type string
}

type FuncSignature struct {
	Name       string
	ReturnType string
	ParamList  []*Param
}

func Parse(s string) (*FuncSignature, error) {
	reader := bufio.NewReader(strings.NewReader(s))
	l := newLexer(reader)

	t := l.scanWithoutWS()
	if t.typ != funcLit {
		return nil, newInvalidSyntax(t, "function")
	}

	fSignature := new(FuncSignature)

	t = l.scanWithoutWS()
	if t.typ != typeLit {
		return nil, newInvalidSyntax(t, "data type")
	}
	fSignature.ReturnType = t.lexeme

	t = l.scanWithoutWS()
	if t.typ != identLit {
		return nil, newInvalidSyntax(t, "function name")
	}
	fSignature.Name = t.lexeme

	t = l.scanWithoutWS()
	if t.typ != openBracket {
		return nil, newInvalidSyntax(t, "(")
	}

	paramList := []*Param{}
	for {
		// handle empty param list
		t := l.scanWithoutWS()
		if t.typ == closeBracket {
			break
		}

		// handle param list
		param := new(Param)
		if t.typ != typeLit {
			return nil, newInvalidSyntax(t, "data type")
		}
		param.Type = t.lexeme

		t = l.scanWithoutWS()
		if t.typ != identLit {
			return nil, newInvalidSyntax(t, "param name")
		}
		param.Name = t.lexeme

		t = l.scanWithoutWS()
		if t.typ != closeBracket && t.typ != comma {
			return nil, newInvalidSyntax(t, ") or ,")
		}
		paramList = append(paramList, param)
		if t.typ == closeBracket {
			break
		}
	}

	fSignature.ParamList = paramList
	return fSignature, nil
}
