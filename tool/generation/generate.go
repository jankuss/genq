package generation

import (
	"fmt"
	. "genq/parser"
	. "genq/templates"
	"path/filepath"
	"strings"
	"time"
)

// Given a string and a position, return the line number, the position in the
// line and the string of the line
func GetContextForPosition(input string, pos int) ContextForPosition {
	  line := 0
	posInLine := 0
	lineString := ""

	for i, value := range input {
		if value == '\n' {
			lineString = ""
			posInLine = 0
			line++
		} else {
			lineString = lineString + string(value)
		}

		if pos == i {
			break
		}

		posInLine++
	}

	for _, value := range input[pos+1:] {
		if value == '\n' {
			break
		}

		lineString = lineString + string(value)
	}

	return ContextForPosition{
		Line:       line,
		PosInLine:  posInLine,
		LineString: lineString,
	}
}

type ContextForPosition struct {
	Line       int
	PosInLine  int
	LineString string
}

type GenerateResultError struct {
	Error   *ParsingError
	Context ContextForPosition
}

type GenerateResult struct {
	OutputFile  string
	Noop        bool
	Duration    time.Duration
	Errors      []GenerateResultError
	WriteString []string
	GenCount    int
}
type ReadContent func(path string) (string, error)

func Generate(path string, readContent ReadContent) (GenerateResult, error) {
	start := time.Now()

	ext := filepath.Ext(path)
	if ext != ".dart" {
		return GenerateResult{}, fmt.Errorf("File %s is not a Dart file", path)
	}

	if strings.HasSuffix(path, ".genq.dart") {
		return GenerateResult{Noop: true}, nil
	}

	extLength := len(ext)

	bytes, err := readContent(path)
	if err != nil {
		return GenerateResult{}, err
	}

	str := string(bytes)

	generatorContext := &generatorContext{}
	listener := &generatorParserListener{
		generatorContext: generatorContext,
		genCount:         0,
	}

	parser := NewParser(str)
	parseResult := parser.Parse(listener)
	generateErrors := []GenerateResultError{}

	if len(parseResult.Errors) > 0 {
		for _, value := range parseResult.Errors {
			generateErrors = append(generateErrors, GenerateResultError{
				Context: GetContextForPosition(str, value.Pos),
				Error:   value,
			})
		}
	}

	if len(generatorContext.outputLines) > 0 {
		generatedPath := path[:len(path)-extLength] + ".genq" + ext
		partOf := fmt.Sprintf("part of '%s';", filepath.Base(path))

		allLines := append([]string{partOf, ""}, generatorContext.outputLines...)

		if err != nil {
			return GenerateResult{}, err
		}

		return GenerateResult{OutputFile: generatedPath, Duration: time.Since(start), Errors: generateErrors, WriteString: allLines, GenCount: listener.genCount}, nil
	}

	return GenerateResult{Noop: true, Errors: generateErrors}, nil
}

type generatorParserListener struct {
	generatorContext *generatorContext
	genCount         int
}

func (l *generatorParserListener) OnGenqClass(genqClass GenqClass) {
	if l.genCount > 0 {
		l.generatorContext.addOutput([]string{""})
	}

	l.generatorContext.addOutput(Template([]string{}, genqClass))

	// JSON Generation is not production ready yet
	// if genqClass.HasJsonConstructor {
	//  l.generatorContext.addOutput([]string{""})
	//  l.generatorContext.addOutput(TemplateFromJson([]string{}, genqClass))
	// }

	l.genCount++
}

type generatorContext struct {
	count       int
	outputLines []string
}

func (ctx *generatorContext) addOutput(lines []string) {
	for _, line := range lines {
		ctx.outputLines = append(ctx.outputLines, line)
	}
}
