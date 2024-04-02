package parser

const NO_MORE_TOKENS = "NO_MORE_TOKENS"
const UNRECOGNIZED_TOKEN = "UNRECOGNIZED_TOKEN"

type Lexer struct {
	input            string
	cursor           int
	restorationPoint int
}

func newLexer(str string) *Lexer {
	return &Lexer{
		input:  str,
		cursor: 0,
	}
}

func (l *Lexer) hasMoreTokens() bool {
	return l.cursor < len(l.input)
}

func (l *Lexer) nextToken() (TokenType, string) {
	if !l.hasMoreTokens() {
		return NO_MORE_TOKENS, ""
	}

	for _, mapping := range TOKEN_MAPPINGS {
		if mapping.regex.MatchString(l.input[l.cursor:]) {
			str := mapping.regex.FindString(l.input[l.cursor:])
			l.cursor += len(str)

			if mapping.token == TOKEN_SKIP || mapping.token == TOKEN_SINGLE_LINE_COMMENT {
				return l.nextToken()
			}

			return mapping.token, str
		}
	}

	unknownToken := l.input[l.cursor : l.cursor+1]
	l.cursor++
	return UNRECOGNIZED_TOKEN, unknownToken
}
