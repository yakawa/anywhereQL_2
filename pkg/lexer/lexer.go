package lexer

import (
	"strings"

	"github.com/yakawa1128/anywhereQL/pkg/token"
)

var basicIdentifierCharacter = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '_'}
var initialIdentifierCharacter = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '_'}
var lineEndMark = [...]byte{'\r', '\n', 0}
var numericalCharacter = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '.', 'e', '+', 'E'}
var whiteSpaceCharacter = [...]byte{' ', '\t', '\r', '\n'}

type Lexer struct {
	input        string
	position     uint
	readPosition uint
	ch           byte
}

func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}

func (l *Lexer) Tokenize() []token.Token {
	t := make([]token.Token, 0)

	for {
		tok := l.NextToken()
		if tok.Type != token.EOF {
			t = append(t, tok)
		} else {
			break
		}
	}
	return t
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	l.skipWhiteSpace()

	switch l.ch {
	case '"':
		tok.Type = token.IDENTIFIER
		tok.Literal = l.readQuotedIdentifier('"')
	case '%':
		tok = newToken(token.PERCENT, l.ch)
	case '&':
		tok = newToken(token.AMPERSAND, l.ch)
	case '\'':
		tok.Type = token.IDENTIFIER
		tok.Literal = l.readQuotedIdentifier('\'')
	case '(':
		tok = newToken(token.LEFTPAREN, l.ch)
	case ')':
		tok = newToken(token.RIGHTPAREN, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '+':
		tok = newToken(token.PLUSSIGN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '-':
		if l.peekChar() == '-' {
			tok.Literal = l.readComment()
			tok.Type = token.COMMENT
			return tok
		} else {
			tok = newToken(token.MINUSSIGN, l.ch)
		}
	case '.':
		if isNumberCharacter(l.peekChar()) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok = newToken(token.PERIOD, l.ch)
		}
	case '/':
		tok = newToken(token.SOLIDAS, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '<':
		tok = newToken(token.LESSTHANOPERATOR, l.ch)
	case '=':
		tok = newToken(token.EQUALSOPERATOR, l.ch)
	case '>':
		tok = newToken(token.GREATERTHANOPERATOR, l.ch)
	case '?':
		tok = newToken(token.QUESTIONMARK, l.ch)
	case '[':
		tok = newToken(token.LEFTBRACKET, l.ch)
	case ']':
		tok = newToken(token.RIGHTBRACKET, l.ch)
	case '^':
		tok = newToken(token.CIRCUMFLEX, l.ch)
	case '_':
		if isIdentifierCharacter(l.peekChar()) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupKeyWord(tok.Literal)
			return tok
		} else {
			tok = newToken(token.UNDERSCORE, l.ch)
		}
	case '|':
		tok = newToken(token.VERTICALBAR, l.ch)
	case '{':
		tok = newToken(token.LEFTBRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHTBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		return tok
	default:
		if isIdentifierCharacter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupKeyWord(tok.Literal)
			return tok
		} else if isNumberCharacter(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok.Literal = ""
			tok.Type = token.ILLIEGAL
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= uint(len(l.input)) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= uint(len(l.input)) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isIdentifierCharacter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readQuotedIdentifier(q byte) string {
	pos := l.position + 1
	l.readChar()
	for l.ch != q {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readComment() string {
	isContinueMinus := true
	for isContinueMinus {
		l.readChar()
		if l.ch != '-' {
			isContinueMinus = false
		}
	}
	pos := l.position + 1
	for !isLineEnd(l.ch) {
		l.readChar()
	}

	return strings.TrimSpace(l.input[pos:l.position])
}

func (l *Lexer) readNumber() string {
	isAppearDot := false
	isExpandMark := false
	isExpandSign := false

	pos := l.position
	for isNumberCharacter(l.ch) {
		if l.ch == '.' {
			if isAppearDot == true {
				break
			}
			isAppearDot = true
		}
		if l.ch == 'e' || l.ch == 'E' {
			if isExpandMark == true {
				break
			}
			isExpandMark = true
		}
		if l.ch == '+' || l.ch == '-' {
			if isExpandMark == false {
				break
			} else {
				if isExpandSign == true {
					break
				}
				isExpandSign = true
			}
		}
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(ch),
	}
}

func isIdentifierCharacter(ch byte) bool {
	for _, v := range basicIdentifierCharacter {
		if v == ch {
			return true
		}
	}
	return false
}

func isLineEnd(ch byte) bool {
	for _, v := range lineEndMark {
		if v == ch {
			return true
		}
	}
	return false
}

func isNumberCharacter(ch byte) bool {
	for _, v := range numericalCharacter {
		if v == ch {
			return true
		}
	}
	return false
}

func isWhiteSpace(ch byte) bool {
	for _, v := range whiteSpaceCharacter {
		if v == ch {
			return true
		}
	}
	return false
}
