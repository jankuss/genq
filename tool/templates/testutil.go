package templates

import (
	"genq/parser"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func getNormalized(str string) []string {
	return strings.Split(strings.Trim(str, "\n "), "\n")
}

func getFirstGenqClass(str string) parser.GenqClassDeclaration {
	p := parser.NewParser(str)
	listener := &testParserListener{}
	p.Parse(listener)

	return listener.genqClass
}

func compare(t *testing.T, actual []string, expect []string) {
	if !cmp.Equal(actual, expect) {
		t.Fatalf(cmp.Diff(expect, actual))
	}
}

type testParserListener struct {
	genqClass parser.GenqClassDeclaration
}

func (l *testParserListener) OnGenqClass(genqClass parser.GenqClassDeclaration) {
	l.genqClass = genqClass
}

func (l *testParserListener) OnGenqJsonEnum(genqClass parser.GenqJsonEnum) {
}
