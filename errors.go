package parser

import (
	"errors"
	"fmt"
)

func newInvalidSyntax(t *token, expected string) error {
	msg := fmt.Sprintf("ERR: Invalid syntax. Unexpected token %s. Expected %s", t.lexeme, expected)
	return errors.New(msg)
}
