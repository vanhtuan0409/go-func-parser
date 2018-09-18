package parser

type tokenType int

type token struct {
	typ    tokenType
	lexeme string
}

const (
	illegal tokenType = iota
	eof
	ws

	openBracket
	closeBracket
	comma

	funcLit
	identLit
	typeLit
)

var (
	dataType = map[string]tokenType{
		"int":    typeLit,
		"float":  typeLit,
		"string": typeLit,
		"bool":   typeLit,
		"void":   typeLit,
	}

	keywords = map[string]tokenType{
		"function": funcLit,
	}
)
