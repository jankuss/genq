package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// / This script geneates a tokens.go file from a token_map.json file. This is just
// / a helper script to avoid having to manually write the tokens.go file.
func main() {
	file, err := os.ReadFile("./token_map.json")
	if err != nil {
		panic(err)
	}

	tokenMap := []TokenMap{}

	err = json.Unmarshal(file, &tokenMap)
	if err != nil {
		panic(err)
	}

	str := "package parser\n\nimport (\n  \"regexp\"\n)\n\n"

	str += "type TokenType string\n\n"

	for _, token := range tokenMap {
		str += fmt.Sprintf("const TOKEN_%s TokenType = \"%s\"\n\n", token.Name, token.Name)
		str += fmt.Sprintf("var REGEXP_%s = regexp.MustCompile(\"%s\")\n\n", token.Name, token.Regexp)
	}

	str += fmt.Sprintf("var TOKEN_MAPPINGS = []TokenMapping{\n")
	for _, token := range tokenMap {
		str += fmt.Sprintf("  {REGEXP_%s, TOKEN_%s},\n", token.Name, token.Name)
	}
	str += fmt.Sprintf("}\n\n")

	f, err := os.Create("./parser/tokens.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(str)
	if err != nil {
		panic(err)
	}
}

type TokenMap struct {
	Name   string `json:"type"`
	Regexp string `json:"regexp"`
}
