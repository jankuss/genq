package parser

import "regexp"

type TokenMapping struct {
	regex *regexp.Regexp
	token TokenType
}

