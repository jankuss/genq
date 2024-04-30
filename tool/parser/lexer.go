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

	var bestMatchIndex int = -1
	var bestMatchStr string

	remaining := l.input[l.cursor:]

	for i, mapping := range TOKEN_MAPPINGS {
		if mapping.regex.MatchString(remaining) {
			checkStr := mapping.regex.FindString(remaining)

			if bestMatchIndex == -1 || len(TOKEN_MAPPINGS[bestMatchIndex].regex.FindString(remaining)) < len(checkStr) {
				bestMatchIndex = i
				bestMatchStr = checkStr
			}
		}
	}

	if bestMatchIndex != -1 {
		bestMatch := TOKEN_MAPPINGS[bestMatchIndex]
		l.cursor += len(bestMatchStr)

		if bestMatch.token == TOKEN_SKIP || bestMatch.token == TOKEN_SINGLE_LINE_COMMENT {
			return l.nextToken()
		}

		return bestMatch.token, bestMatchStr
	}

	unknownToken := l.input[l.cursor : l.cursor+1]
	l.cursor++
	return UNRECOGNIZED_TOKEN, unknownToken
}
