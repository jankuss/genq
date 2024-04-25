package parser

const NO_MORE_TOKENS = "NO_MORE_TOKENS"
const UNRECOGNIZED_TOKEN = "UNRECOGNIZED_TOKEN"

const MODE_DEFAULT = 0;
const MODE_ONLY_PAREN = 1;


var tokensForMode map[int][]TokenMapping = map[int][]TokenMapping{
  MODE_DEFAULT: TOKEN_MAPPINGS,
  MODE_ONLY_PAREN: TOKEN_MAPPINGS,
}

type Lexer struct {
	input            string
	cursor           int
	restorationPoint int
  mode              int
}

func newLexer(str string) *Lexer {
	return &Lexer{
		input:  str,
		cursor: 0,
	}
}

func (l *Lexer) setMode(mode int) {
  l.mode = mode;
}

func (l *Lexer) hasMoreTokens() bool {
	return l.cursor < len(l.input)
}

func (l *Lexer) nextToken() (TokenType, string) {
	if !l.hasMoreTokens() {
		return NO_MORE_TOKENS, ""
	}

	for _, mapping := range tokensForMode[l.mode] {
		if mapping.regex.MatchString(l.input[l.cursor:]) {
			str := mapping.regex.FindString(l.input[l.cursor:])
			l.cursor += len(str)

			if l.mode == MODE_DEFAULT && (mapping.token == TOKEN_SKIP || mapping.token == TOKEN_SINGLE_LINE_COMMENT) {
				return l.nextToken()
			}

			return mapping.token, str
		}
	}

	unknownToken := l.input[l.cursor : l.cursor+1]
	l.cursor++
	return UNRECOGNIZED_TOKEN, unknownToken
}

