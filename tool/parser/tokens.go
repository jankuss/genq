package parser

import (
  "regexp"
)

type TokenType string

const TOKEN_SKIP TokenType = "SKIP"

var REGEXP_SKIP = regexp.MustCompile("^[\\s+\\n]")

const TOKEN_SINGLE_LINE_COMMENT TokenType = "SINGLE_LINE_COMMENT"

var REGEXP_SINGLE_LINE_COMMENT = regexp.MustCompile("^\\/\\/.*")

const TOKEN_SEMICOLON TokenType = "SEMICOLON"

var REGEXP_SEMICOLON = regexp.MustCompile("^;")

const TOKEN_COLON TokenType = "COLON"

var REGEXP_COLON = regexp.MustCompile("^:")

const TOKEN_ANNOTATION TokenType = "ANNOTATION"

var REGEXP_ANNOTATION = regexp.MustCompile("^@")

const TOKEN_CLASS TokenType = "CLASS"

var REGEXP_CLASS = regexp.MustCompile("^class")

const TOKEN_CURLY_START TokenType = "CURLY_START"

var REGEXP_CURLY_START = regexp.MustCompile("^{")

const TOKEN_CURLY_END TokenType = "CURLY_END"

var REGEXP_CURLY_END = regexp.MustCompile("^}")

const TOKEN_CONST TokenType = "CONST"

var REGEXP_CONST = regexp.MustCompile("^const")

const TOKEN_FACTORY TokenType = "FACTORY"

var REGEXP_FACTORY = regexp.MustCompile("^factory")

const TOKEN_SINGLE_STRING TokenType = "SINGLE_STRING"

var REGEXP_SINGLE_STRING = regexp.MustCompile("^'[^']*'")

const TOKEN_DOUBLE_STRING TokenType = "DOUBLE_STRING"

var REGEXP_DOUBLE_STRING = regexp.MustCompile("^\"[^\"]*\"")

const TOKEN_BOOLEAN_TRUE TokenType = "BOOLEAN_TRUE"

var REGEXP_BOOLEAN_TRUE = regexp.MustCompile("^true")

const TOKEN_BOOLEAN_FALSE TokenType = "BOOLEAN_FALSE"

var REGEXP_BOOLEAN_FALSE = regexp.MustCompile("^false")

const TOKEN_FINAL TokenType = "FINAL"

var REGEXP_FINAL = regexp.MustCompile("^final")

const TOKEN_VAR TokenType = "VAR"

var REGEXP_VAR = regexp.MustCompile("^var")

const TOKEN_PAREN_START TokenType = "PAREN_START"

var REGEXP_PAREN_START = regexp.MustCompile("^\\(")

const TOKEN_PAREN_END TokenType = "PAREN_END"

var REGEXP_PAREN_END = regexp.MustCompile("^\\)")

const TOKEN_REQUIRED TokenType = "REQUIRED"

var REGEXP_REQUIRED = regexp.MustCompile("^required")

const TOKEN_COMMA TokenType = "COMMA"

var REGEXP_COMMA = regexp.MustCompile("^,")

const TOKEN_DOT TokenType = "DOT"

var REGEXP_DOT = regexp.MustCompile("^\\.")

const TOKEN_GENERIC_START TokenType = "GENERIC_START"

var REGEXP_GENERIC_START = regexp.MustCompile("^<")

const TOKEN_GENERIC_END TokenType = "GENERIC_END"

var REGEXP_GENERIC_END = regexp.MustCompile("^>")

const TOKEN_OPTIONAL TokenType = "OPTIONAL"

var REGEXP_OPTIONAL = regexp.MustCompile("^\\?")

const TOKEN_IDENTIFIER TokenType = "IDENTIFIER"

var REGEXP_IDENTIFIER = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*")

var TOKEN_MAPPINGS = []TokenMapping{
  {REGEXP_SKIP, TOKEN_SKIP},
  {REGEXP_SINGLE_LINE_COMMENT, TOKEN_SINGLE_LINE_COMMENT},
  {REGEXP_SEMICOLON, TOKEN_SEMICOLON},
  {REGEXP_COLON, TOKEN_COLON},
  {REGEXP_ANNOTATION, TOKEN_ANNOTATION},
  {REGEXP_CLASS, TOKEN_CLASS},
  {REGEXP_CURLY_START, TOKEN_CURLY_START},
  {REGEXP_CURLY_END, TOKEN_CURLY_END},
  {REGEXP_CONST, TOKEN_CONST},
  {REGEXP_FACTORY, TOKEN_FACTORY},
  {REGEXP_SINGLE_STRING, TOKEN_SINGLE_STRING},
  {REGEXP_DOUBLE_STRING, TOKEN_DOUBLE_STRING},
  {REGEXP_BOOLEAN_TRUE, TOKEN_BOOLEAN_TRUE},
  {REGEXP_BOOLEAN_FALSE, TOKEN_BOOLEAN_FALSE},
  {REGEXP_FINAL, TOKEN_FINAL},
  {REGEXP_VAR, TOKEN_VAR},
  {REGEXP_PAREN_START, TOKEN_PAREN_START},
  {REGEXP_PAREN_END, TOKEN_PAREN_END},
  {REGEXP_REQUIRED, TOKEN_REQUIRED},
  {REGEXP_COMMA, TOKEN_COMMA},
  {REGEXP_DOT, TOKEN_DOT},
  {REGEXP_GENERIC_START, TOKEN_GENERIC_START},
  {REGEXP_GENERIC_END, TOKEN_GENERIC_END},
  {REGEXP_OPTIONAL, TOKEN_OPTIONAL},
  {REGEXP_IDENTIFIER, TOKEN_IDENTIFIER},
}

