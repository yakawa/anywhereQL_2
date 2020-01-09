// Copyrights (C) 2019-2020 KAWASAKI Yasukazu

package token

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) Debug() string {
	return fmt.Sprintf("Token{Type: %s, Literal: %s}\n", t.Type, t.Literal)
}

const (
	ILLIEGAL = "ILLIEGAL"
	EOF      = "EOF"

	// Data Type
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	COMMENT    = "COMMENT"

	// Characters
	DOUBLEQUOTE         = "\""
	PERCENT             = "%"
	AMPERSAND           = "&"
	QUOTE               = "'"
	LEFTPAREN           = "("
	RIGHTPAREN          = ")"
	ASTERISK            = "*"
	PLUSSIGN            = "+"
	COMMA               = ","
	MINUSSIGN           = "-"
	PERIOD              = "."
	SOLIDAS             = "/"
	COLON               = ":"
	SEMICOLON           = ";"
	LESSTHANOPERATOR    = "<"
	EQUALSOPERATOR      = "="
	GREATERTHANOPERATOR = ">"
	QUESTIONMARK        = "?"
	LEFTBRACKET         = "["
	RIGHTBRACKET        = "]"
	CIRCUMFLEX          = "^"
	UNDERSCORE          = "_"
	VERTICALBAR         = "|"
	LEFTBRACE           = "{"
	RIGHTBRACE          = "}"

	// SQL Key Words
	SELECT = "SELECT"
)

var reservedKeyWord = map[string]TokenType{
	"SELECT": SELECT,
}

func LookupKeyWord(w string) TokenType {
	if t, ok := reservedKeyWord[w]; ok {
		return t
	}
	return IDENTIFIER
}
