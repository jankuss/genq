//

package parser

import (
	"fmt"
)

type Parser struct {
	lexer          *Lexer
	lookahead      TokenType
	lookaheadValue string

	// The current scope is used to keep track of the current level of nesting.
	// When a "{" is encountered, the scope is increased by one. When a "}" is
	// encountered, the scope is decreased by one.
	currentScope      int
	restorationPoints []restorationPoint
}

type ParseResult struct {
	Errors []*ParsingError
}

// Parse the input. The listener will be called when nodes are parsed.
// If parsing errors are encountered, they will be returned in the ParseResult.
func (p *Parser) Parse(listener ParserListener) ParseResult {
	token, value := p.lexer.nextToken()
	p.lookahead = token
	p.lookaheadValue = value

	errors := []*ParsingError{}

	for p.lexer.hasMoreTokens() {
		// The way genq works is by looking for specific annotations and only parsing
		// the nodes that are annotated. This is a simple way to avoid parsing the entire
		// file and only focus on the parts that are relevant.
		// Thus: We just skip tokens until we encounter an annotation
		if p.lookahead != TOKEN_ANNOTATION {
			p.eatUntil(TOKEN_ANNOTATION)
			continue
		}

		// We only parse annotations at the top level. If we are inside a scope, we skip
		if p.currentScope > 0 {
			p.eat(p.lookahead)
			continue
		}

		// We found an annotation at the top level. Now lets parse it into a node.
		n, err := p.parseAnnotation()
		if err != nil {
			// If a parsing error is encountered while parsing the annotation, we do not
			// want to stop parsing. Instead, we just ignore this annotation and continue
			// parsing the rest of the file.
			if p.lookahead == TOKEN_ANNOTATION {
				p.eat(TOKEN_ANNOTATION)
			}
			continue
		}

		if n.Identifier.Name == "genq" || n.Identifier.Name == "Genq" {
			err := genqAnnotationParser(n, p, listener)
			if err != nil {
				errors = append(errors, err)
				continue
			}
		}

		if n.Identifier.Name == "GenqJsonEnum" {
			err := genqJsonEnumAnnotationParser(n, p, listener)
			if err != nil {
				errors = append(errors, err)
				continue
			}
		}

	}

	return ParseResult{
		Errors: errors,
	}
}

// ParserListener is an interface for listening to nodes parsed by the parser.
// When a node is parsed, the parser will call the appropriate method on the listener.
type ParserListener interface {
	OnGenqClass(genqClass GenqClassDeclaration)
	OnGenqJsonEnum(genqJsonEnum GenqJsonEnum)
}

// NewParser creates a new parser with the given string as input.
func NewParser(str string) *Parser {
	return &Parser{
		lexer: newLexer(str),
	}
}

// Marks the restoration point for the parser. Used in combination with "restore".
// In most cases, it is enough to look at the lookahead to determine the path.
// Sometimes though, it is necessary to parse optimistically and check wether or not
// an error occurs.
func (p *Parser) markRestorationPoint() {
	// To support multiple restoration points, we keep a stack of restoration points.
	p.restorationPoints = append(p.restorationPoints, restorationPoint{
		lookaheadValue: p.lookaheadValue,
		lookahead:      p.lookahead,
		cursor:         p.lexer.cursor,
	})
}

// Restores the parser to the restoration point.
func (p *Parser) restore() {
	restorationPoint := p.restorationPoints[len(p.restorationPoints)-1]

	// Remove the last restoration point
	p.restorationPoints = p.restorationPoints[:len(p.restorationPoints)-1]

	p.lookahead = restorationPoint.lookahead
	p.lookaheadValue = restorationPoint.lookaheadValue
	p.lexer.cursor = restorationPoint.cursor
}

func (p *Parser) dontRestore() {
	p.restorationPoints = p.restorationPoints[:len(p.restorationPoints)-1]
}

type restorationPoint struct {
	lookahead      TokenType
	lookaheadValue string
	cursor         int
}

func (p *Parser) parseArgumentList() (GenqArgumentList, *ParsingError) {
	positional := []GenqValue{}
	params := []GenqNamedExpression{}
	if p.lookahead == TOKEN_PAREN_START {
		// Annotation is an invocation
		_, err := p.eat(TOKEN_PAREN_START)
		if err != nil {
			return GenqArgumentList{}, err
		}

		for p.lookahead != TOKEN_PAREN_END {
			p.markRestorationPoint()
			v, err := p.parseNamedAssignment()
			if err == nil {
				p.dontRestore()

				params = append(params, *v)
			} else {
				p.restore()

				v, err := p.parseRawCode(TOKEN_COMMA, TOKEN_PAREN_END, TOKEN_CURLY_END)
				if err != nil {
					return GenqArgumentList{}, err
				}

				positional = append(positional, v)
			}

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)
				if err != nil {
					return GenqArgumentList{}, err
				}
			}
		}

		_, err = p.eat(TOKEN_PAREN_END)
		if err != nil {
			return GenqArgumentList{}, err
		}
	}

	return GenqArgumentList{
		NamedArgs:      params,
		PositionalArgs: positional,
	}, nil
}

// Our little cheat code for "parsing" expressions. Normally, we are not really interested in the
// AST of user defined values, as we mostly just want to pass them through to the output.
// This function can be used to parse it as a string. Feels a bit hacky, but works.
//
// It works as follows:
// - We set the lexer into a special mode, where whitespaces are not skipped.
// - We keep track of the encountered parentheses and curly braces to keep track of the current scope.
// - We build a string by concatenating the values of the tokens we encounter.
// - If we encounter a token that is in the stopTokens list, we stop parsing.
func (p *Parser) parseRawCode(stopTokens ...TokenType) (GenqValue, *ParsingError) {
	p.lexer.setMode(MODE_DONT_SKIP_COMMENT)
	defer p.lexer.setMode(MODE_DEFAULT)

	code := ""

	// We keep track of the encountered parentheses and curly braces to keep track of the current scope.
	// When the stack is empty, and we encounter a stop token, we stop parsing.
	tokenStack := []TokenType{}

	for {
		if !p.lexer.hasMoreTokens() {
			break
		}

		cur := p.lookahead
		if cur == TOKEN_PAREN_START {
			tokenStack = append(tokenStack, TOKEN_PAREN_START)
		}

		if cur == TOKEN_CURLY_START {
			tokenStack = append(tokenStack, TOKEN_CURLY_START)
		}

		if len(tokenStack) > 0 && cur == TOKEN_PAREN_END {
			if tokenStack[len(tokenStack)-1] == TOKEN_PAREN_START {
				tokenStack = tokenStack[:len(tokenStack)-1]
			} else {
				err := fmt.Errorf("Unexpected token: %s (`%s`).", cur, p.lookaheadValue)
				return GenqValue{}, p.produceError(err)
			}
		}

		if len(tokenStack) > 0 && cur == TOKEN_CURLY_END {
			if tokenStack[len(tokenStack)-1] == TOKEN_CURLY_START {
				tokenStack = tokenStack[:len(tokenStack)-1]
			} else {
				err := fmt.Errorf("Unexpected token: %s (`%s`).", cur, p.lookaheadValue)
				return GenqValue{}, p.produceError(err)
			}
		}

		if contains(stopTokens, cur) {
			break
		} else {
			value, err := p.eat(cur)
			if err != nil {
				return GenqValue{}, err
			}
			code = code + value
		}
	}

	return GenqValue{
		RawValue: code,
	}, nil
}

func (p *Parser) parseInsideParen() (string, *ParsingError) {
	p.lexer.setMode(MODE_ONLY_PAREN)
	defer p.lexer.setMode(MODE_DEFAULT)

	parenCounter := 0
	code := ""

	for {
		if p.lookahead == TOKEN_PAREN_START {
			parenCounter++
		}

		if p.lookahead == TOKEN_PAREN_END {
			parenCounter--
		}

		value, err := p.eat(p.lookahead)
		if err != nil {
			return "", err
		}

		code = code + value

		if parenCounter == 0 {
			break
		}
	}

	return code[1 : len(code)-1], nil
}

func (p *Parser) parseAnnotation() (GenqAnnotation, *ParsingError) {
	_, err := p.eat(TOKEN_ANNOTATION)
	if err != nil {
		return GenqAnnotation{}, err
	}

	annotationIdentifier, err := p.parseGenqIdentifier()
	if err != nil {
		return GenqAnnotation{}, err
	}

	var argumentList GenqArgumentList = GenqArgumentList{}
	if p.lookahead == TOKEN_PAREN_START {
		// Annotation is an invocation
		argumentList, err = p.parseArgumentList()
	}

	return GenqAnnotation{
		Identifier: *annotationIdentifier,
		Arguments:  argumentList,
	}, nil
}

func (p *Parser) parseNamedAssignment() (*GenqNamedExpression, *ParsingError) {
	name, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return nil, err
	}

	_, err = p.eat(TOKEN_COLON)
	if err != nil {
		return nil, err
	}

	value, err := p.parseRawCode(TOKEN_COMMA, TOKEN_PAREN_END, TOKEN_CURLY_END)
	if err != nil {
		return nil, err
	}

	return &GenqNamedExpression{
		Name:  name,
		Value: value,
	}, nil
}

func (p *Parser) parseGenqIdentifier() (*GenqIdentifier, *ParsingError) {
	v, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return &GenqIdentifier{}, err
	}

	top := &GenqIdentifier{
		Name: v,
	}
	cur := top

	for p.lookahead == TOKEN_DOT {
		_, err := p.eat(TOKEN_DOT)
		if err != nil {
			return &GenqIdentifier{}, err
		}

		v, err := p.eat(TOKEN_IDENTIFIER)
		if err != nil {
			return &GenqIdentifier{}, err
		}

		cur.Next = &GenqIdentifier{
			Name: v,
		}

		cur = cur.Next
	}

	return top, nil
}

// Due to dart syntax, it is possible for a type reference to be the return type of a function.
// This function acts as a utility, to check wether or not we just parsed a return type.
// See: https://github.com/dart-lang/language/issues/2972 for more info
func (p *Parser) isReturnTypeOfFunction() bool {
	p.markRestorationPoint()
	defer p.restore()

	identifier, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return false
	}

	_, err = p.eat(TOKEN_PAREN_START)
	if err != nil {
		return false
	}

	if identifier != "Function" {
		return false
	}

	return true
}

func (p *Parser) parseTypeReference() (GenqNamedType, *ParsingError) {
	typeIdentifier, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqNamedType{}, err
	}

	genericTypes := []GenqNamedType{}
	if p.lookahead == TOKEN_GENERIC_START {
		_, err := p.eat(TOKEN_GENERIC_START)
		if err != nil {
			return GenqNamedType{}, nil
		}

		for {
			if p.lookahead == TOKEN_GENERIC_END {
				_, err := p.eat(TOKEN_GENERIC_END)
				if err != nil {
					return GenqNamedType{}, err
				}

				break
			}

			typeRef, err := p.parseTypeReference()
			if err != nil {
				return GenqNamedType{}, err
			}

			genericTypes = append(genericTypes, typeRef)

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)

				if err != nil {
					return GenqNamedType{}, err
				}
			}
		}
	}

	optional := false
	if p.lookahead == TOKEN_OPTIONAL {
		_, err := p.eat(TOKEN_OPTIONAL)
		if err != nil {
			return GenqNamedType{}, err
		}

		optional = true
	}

	typeRef := GenqNamedType{
		Name:         typeIdentifier,
		Optional:     optional,
		GenericTypes: genericTypes,
		IsFunction:   false,
	}

	for p.isReturnTypeOfFunction() {
		p.eat(TOKEN_IDENTIFIER)
		paramList, err := p.parseFormalParameterList()

		if err != nil {
			return GenqNamedType{}, err
		}

		prev := typeRef
		typeRef = GenqNamedType{
			ReturnType: &prev,
			IsFunction: true,
			ParamList:  paramList,
		}
	}

	return typeRef, nil
}

func (p *Parser) parseFormalParameterList() (GenqFormalParameterList, *ParsingError) {
	_, err := p.eat(TOKEN_PAREN_START)
	if err != nil {
		return GenqFormalParameterList{}, err
	}

	positionalParams := []GenqPositionalFormalParameter{}
	namedParams := []GenqFormalNamedParameter{}

	for p.lookahead != TOKEN_PAREN_END && p.lookahead != TOKEN_CURLY_START {
		if p.lookahead == TOKEN_IDENTIFIER {
			typeRef, err := p.parseTypeReference()
			if err != nil {
				return GenqFormalParameterList{}, err
			}

			name, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqFormalParameterList{}, err
			}

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)
				if err != nil {
					return GenqFormalParameterList{}, err
				}
			}

			positionalParams = append(positionalParams, GenqPositionalFormalParameter{
				ParamType: typeRef,
				Name:      name,
			})
		}
	}

	if p.lookahead == TOKEN_CURLY_START {
		_, err := p.eat(TOKEN_CURLY_START)
		if err != nil {
			return GenqFormalParameterList{}, err
		}

		for p.lookahead != TOKEN_CURLY_END {
			namedParam, err := p.parseNamedParam()
			if err != nil {
				return GenqFormalParameterList{}, err
			}

			namedParams = append(namedParams, namedParam)

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)
				if err != nil {
					return GenqFormalParameterList{}, err
				}
			}
		}

		_, err = p.eat(TOKEN_CURLY_END)
		if err != nil {
			return GenqFormalParameterList{}, err
		}
	}

	_, err = p.eat(TOKEN_PAREN_END)
	if err != nil {
		return GenqFormalParameterList{}, err
	}

	return GenqFormalParameterList{
		PositionalParams: positionalParams,
		NamedParams:      namedParams,
	}, nil
}

func (p *Parser) parseGenqClass(annotation GenqAnnotation) (GenqClassDeclaration, *ParsingError) {
	_, err := p.eat(TOKEN_CLASS)
	if err != nil {
		return GenqClassDeclaration{}, err
	}

	classIdentifier, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqClassDeclaration{}, err
	}

	err = p.eatUntil(TOKEN_CURLY_START)
	if err != nil {
		return GenqClassDeclaration{}, err
	}

	_, err = p.eat(TOKEN_CURLY_START)
	if err != nil {
		return GenqClassDeclaration{}, err
	}

	var genqConstructorSignature GenqConstructor
	hasPrivateConstructor := false

	for {
		if !p.lexer.hasMoreTokens() {
			break
		}

		isConst := false
		if p.lookahead == TOKEN_CONST && p.currentScope == 1 {
			_, err := p.eat(TOKEN_CONST)
			if err != nil {
				return GenqClassDeclaration{}, err
			}
			isConst = true
		}

		if p.lookahead == TOKEN_IDENTIFIER {
			cn, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqClassDeclaration{}, err
			}

			if cn == classIdentifier && p.lookahead == TOKEN_DOT {
				_, err := p.eat(TOKEN_DOT)
				if err != nil {
					return GenqClassDeclaration{}, err
				}

				constructorName, err := p.eat(TOKEN_IDENTIFIER)
				if err != nil {
					return GenqClassDeclaration{}, err
				}

				if constructorName == "_" {
					// This is a private constructor
					err := p.eatUntil(TOKEN_SEMICOLON)
					if err != nil {
						return GenqClassDeclaration{}, err
					}

					_, err = p.eat(TOKEN_SEMICOLON)
					if err != nil {
						return GenqClassDeclaration{}, err
					}

					hasPrivateConstructor = true
				}
			}
		} else if p.lookahead == TOKEN_CURLY_END {
			_, err := p.eat(TOKEN_CURLY_END)
			if err != nil {
				return GenqClassDeclaration{}, err
			}

			if p.currentScope == 0 {
				break
			}
		} else if p.lookahead == TOKEN_FACTORY {
			_, err := p.eat(TOKEN_FACTORY)
			if err != nil {
				return GenqClassDeclaration{}, err
			}

			cn, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqClassDeclaration{}, err
			}

			if p.lookahead == TOKEN_PAREN_START && cn == classIdentifier {
				// Factory matches class name
				// This is the chosen constructor
				genqConstructorSignature, err = p.parseGenqConstructor(isConst)
				if err != nil {
					return GenqClassDeclaration{}, err
				}

				err := p.eatUntil(TOKEN_SEMICOLON)
				if err != nil {
					return GenqClassDeclaration{}, err
				}

				_, err = p.eat(TOKEN_SEMICOLON)
				if err != nil {
					return GenqClassDeclaration{}, err
				}
			}
		} else {
			_, err := p.eat(p.lookahead)
			if err != nil {
				return GenqClassDeclaration{}, err
			}
		}
	}

	return GenqClassDeclaration{
		Name:                  classIdentifier,
		Constructor:           genqConstructorSignature,
		HasPrivateConstructor: hasPrivateConstructor,
		Annotation:            annotation,
	}, nil
}

func (p *Parser) parseGenqConstructor(isConst bool) (GenqConstructor, *ParsingError) {
	formalParams, err := p.parseFormalParameterList()
	if err != nil {
		return GenqConstructor{}, err
	}

	return GenqConstructor{
		ParamList: formalParams,
		IsConst:   isConst,
	}, nil
}

func (p *Parser) parseNamedParam() (GenqFormalNamedParameter, *ParsingError) {
	required := false
	annotation := GenqAnnotation{}

	if p.lookahead == TOKEN_ANNOTATION {
		ann, err := p.parseAnnotation()
		if err != nil {
			p.eatUntil(TOKEN_COMMA, TOKEN_PAREN_END)
			p.eat(p.lookahead)
		}

		annotation = ann
	}

	if p.lookahead == TOKEN_REQUIRED {
		_, err := p.eat(TOKEN_REQUIRED)
		if err != nil {
			return GenqFormalNamedParameter{}, nil
		}

		required = true
	}

	paramType, err := p.parseTypeReference()
	if err != nil {
		return GenqFormalNamedParameter{}, err
	}

	paramName, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqFormalNamedParameter{}, nil
	}

	return GenqFormalNamedParameter{
		ParamType:  paramType,
		Name:       paramName,
		Required:   required,
		Annotation: annotation,
	}, nil
}

// Eats tokens until one of the given tokens is encountered
func (p *Parser) eatUntil(tokens ...TokenType) *ParsingError {
	for !contains(tokens, p.lookahead) {
		if !p.lexer.hasMoreTokens() {
			return nil
		}

		_, err := p.eat(p.lookahead)
		if err != nil {
			return err
		}
	}

	return nil
}

// Moves the parser forward by one token
func (p *Parser) progress() {
	if p.lookahead == TOKEN_CURLY_START {
		p.currentScope++
	}

	if p.lookahead == TOKEN_CURLY_END {
		p.currentScope--
	}

	token, v := p.lexer.nextToken()
	p.lookahead = token
	p.lookaheadValue = v
}

// Eats the token currently in the lookahead. If the token does not match the
// expected token, a ParsingError is returned.
func (p *Parser) eat(token TokenType) (string, *ParsingError) {
	currentToken := p.lookahead
	currentTokenValue := p.lookaheadValue

	if currentToken != token {
		err := fmt.Errorf("Unexpected token: %s (`%s`). Expected %s.", currentToken, currentTokenValue, token)
		return "", p.produceError(err)
	}

	p.progress()
	return currentTokenValue, nil
}

func (p *Parser) produceError(err error) *ParsingError {
	return &ParsingError{
		Err: err,
		Pos: p.lexer.cursor - len(p.lookaheadValue),
	}
}

type ParsingError struct {
	Pos      int
	Err      error
	internal bool
}

func contains(s []TokenType, e TokenType) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
