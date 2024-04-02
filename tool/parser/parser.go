//

package parser

import "fmt"

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

		// If the annotation is not a "genq" annotation
		if n.Name != "genq" {
			continue
		}

		// As of now, we are only interested in classes annotated with "genq" and.
		// Skip if the next token is not "class"
		if p.lookahead != TOKEN_CLASS {
			continue
		}

		// Finally, we found a class annotated with "genq". Lets parse it.
		genqClass, err := p.parseGenqClass()
		if err != nil {
			// When there was an error parsing the class, we add it to the list of Errors
			// and continue parsing the rest of the file.
			errors = append(errors, err)
			continue
		}

		// When the class is successfully parsed, we call the listener with the parsed class.
		listener.OnGenqClass(genqClass)
	}

	return ParseResult{
		Errors: errors,
	}
}

// ParserListener is an interface for listening to nodes parsed by the parser.
// When a node is parsed, the parser will call the appropriate method on the listener.
type ParserListener interface {
	OnGenqClass(genqClass GenqClass)
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

type restorationPoint struct {
	lookahead      TokenType
	lookaheadValue string
	cursor         int
}

func (p *Parser) parseAnnotation() (GenqAnnotation, *ParsingError) {
	p.eat(TOKEN_ANNOTATION)
	annotationName, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqAnnotation{}, err
	}

	params := []GenqAnnotationParameter{}

	if p.lookahead == TOKEN_PAREN_START {
		// Annotation is an invocation
		p.eat(TOKEN_PAREN_START)

		for p.lookahead != TOKEN_PAREN_END {
			name, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqAnnotation{}, err
			}

			_, err = p.eat(TOKEN_COLON)
			if err != nil {
				return GenqAnnotation{}, err
			}

			value, err := p.parseValue()
			if err != nil {
				return GenqAnnotation{}, err
			}

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)
				if err != nil {
					return GenqAnnotation{}, err
				}
			}

			params = append(params, GenqAnnotationParameter{
				Name:  name,
				Value: value,
			})
		}
		p.eat(TOKEN_PAREN_END)
	}

	return GenqAnnotation{
		Name:   annotationName,
		Params: params,
	}, nil
}

func (p *Parser) parseValue() (GenqValue, *ParsingError) {
	if p.lookahead == TOKEN_SINGLE_STRING || p.lookahead == TOKEN_DOUBLE_STRING {
		v, err := p.eat(p.lookahead)
		if err != nil {
			return GenqValue{}, err
		}

		return GenqValue{
			StringValue: v,
		}, nil
	}

	if p.lookahead == TOKEN_BOOLEAN_TRUE || p.lookahead == TOKEN_BOOLEAN_FALSE {
		v, err := p.eat(p.lookahead)
		if err != nil {
			return GenqValue{}, nil
		}

		return GenqValue{
			BooleanValue: v == "true",
		}, nil
	}

	return GenqValue{}, nil
}

type FunctionType struct {
	ReturnType GenqTypeReference
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

func (p *Parser) parseTypeReference() (GenqTypeReference, *ParsingError) {
	typeIdentifier, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqTypeReference{}, err
	}

	genericTypes := []GenqTypeReference{}
	if p.lookahead == TOKEN_GENERIC_START {
		_, err := p.eat(TOKEN_GENERIC_START)
		if err != nil {
			return GenqTypeReference{}, nil
		}

		for {
			if p.lookahead == TOKEN_GENERIC_END {
				_, err := p.eat(TOKEN_GENERIC_END)
				if err != nil {
					return GenqTypeReference{}, err
				}

				break
			}

			typeRef, err := p.parseTypeReference()
			if err != nil {
				return GenqTypeReference{}, err
			}

			genericTypes = append(genericTypes, typeRef)

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)

				if err != nil {
					return GenqTypeReference{}, err
				}
			}
		}
	}

	optional := false
	if p.lookahead == TOKEN_OPTIONAL {
		_, err := p.eat(TOKEN_OPTIONAL)
		if err != nil {
			return GenqTypeReference{}, err
		}

		optional = true
	}

	typeRef := GenqTypeReference{
		Name:         typeIdentifier,
		Optional:     optional,
		GenericTypes: genericTypes,
		IsFunction:   false,
	}

	for p.isReturnTypeOfFunction() {
		p.eat(TOKEN_IDENTIFIER)
		paramList, err := p.parseParamList()

		if err != nil {
			return GenqTypeReference{}, err
		}

		prev := typeRef
		typeRef = GenqTypeReference{
			ReturnType: &prev,
			IsFunction: true,
			ParamList:  paramList,
		}
	}

	return typeRef, nil
}

func (p *Parser) parseParamList() (GenqParamList, *ParsingError) {
	p.eat(TOKEN_PAREN_START)

	positionalParams := []GenqPositionalParam{}
	namedParams := []GenqNamedParam{}

	for p.lookahead != TOKEN_PAREN_END && p.lookahead != TOKEN_CURLY_START {
		if p.lookahead == TOKEN_IDENTIFIER {
			typeRef, err := p.parseTypeReference()
			if err != nil {
				return GenqParamList{}, err
			}

			name, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqParamList{}, err
			}

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)
				if err != nil {
					return GenqParamList{}, err
				}
			}

			positionalParams = append(positionalParams, GenqPositionalParam{
				ParamType: typeRef,
				Name:      name,
			})
		}
	}

	if p.lookahead == TOKEN_CURLY_START {
		_, err := p.eat(TOKEN_CURLY_START)
		if err != nil {
			return GenqParamList{}, err
		}

		for p.lookahead != TOKEN_CURLY_END {
			namedParam, err := p.parseNamedParam()
			if err != nil {
				return GenqParamList{}, err
			}

			namedParams = append(namedParams, namedParam)

			if p.lookahead == TOKEN_COMMA {
				_, err := p.eat(TOKEN_COMMA)
				if err != nil {
					return GenqParamList{}, err
				}
			}
		}

		_, err = p.eat(TOKEN_CURLY_END)
		if err != nil {
			return GenqParamList{}, err
		}
	}

	p.eat(TOKEN_PAREN_END)

	return GenqParamList{
		PositionalParams: positionalParams,
		NamedParams:      namedParams,
	}, nil
}

func (p *Parser) parseGenqClass() (GenqClass, *ParsingError) {
	_, err := p.eat(TOKEN_CLASS)
	if err != nil {
		return GenqClass{}, err
	}

	classIdentifier, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqClass{}, err
	}

	err = p.eatUntil(TOKEN_CURLY_START)
	if err != nil {
		return GenqClass{}, err
	}

	_, err = p.eat(TOKEN_CURLY_START)
	if err != nil {
		return GenqClass{}, err
	}

	var genqConstructorSignature GenqConstructorSignature
	var jsonConstructorSignature GenqFromJsonConstructor
	var hasJsonConstructor bool
	hasPrivateConstructor := false

	for {
		if !p.lexer.hasMoreTokens() {
			break
		}

		if p.lookahead == TOKEN_CONST && p.currentScope == 1 {
			_, err := p.eat(TOKEN_CONST)
			if err != nil {
				return GenqClass{}, err
			}
		}

		if p.lookahead == TOKEN_IDENTIFIER {
			cn, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqClass{}, err
			}

			if cn == classIdentifier && p.lookahead == TOKEN_DOT {
				_, err := p.eat(TOKEN_DOT)
				if err != nil {
					return GenqClass{}, err
				}

				constructorName, err := p.eat(TOKEN_IDENTIFIER)
				if err != nil {
					return GenqClass{}, err
				}

				if constructorName == "_" {
					// This is a private constructor
					err := p.eatUntil(TOKEN_SEMICOLON)
					if err != nil {
						return GenqClass{}, err
					}

					_, err = p.eat(TOKEN_SEMICOLON)
					if err != nil {
						return GenqClass{}, err
					}

					hasPrivateConstructor = true
				}
			}
		} else if p.lookahead == TOKEN_CURLY_END {
			_, err := p.eat(TOKEN_CURLY_END)
			if err != nil {
				return GenqClass{}, err
			}

			if p.currentScope == 0 {
				break
			}
		} else if p.lookahead == TOKEN_FACTORY {
			_, err := p.eat(TOKEN_FACTORY)
			if err != nil {
				return GenqClass{}, err
			}

			cn, err := p.eat(TOKEN_IDENTIFIER)
			if err != nil {
				return GenqClass{}, err
			}

			if p.lookahead == TOKEN_PAREN_START && cn == classIdentifier {
				// Factory matches class name
				// This is the chosen constructor
				genqConstructorSignature, err = p.parseGenqConstructorSignature()
				if err != nil {
					return GenqClass{}, err
				}

				err := p.eatUntil(TOKEN_SEMICOLON)
				if err != nil {
					return GenqClass{}, err
				}

				_, err = p.eat(TOKEN_SEMICOLON)
				if err != nil {
					return GenqClass{}, err
				}
			} else if p.lookahead == TOKEN_DOT {
				_, err := p.eat(TOKEN_DOT)
				if err != nil {
					return GenqClass{}, err
				}

				constructorName, err := p.eat(TOKEN_IDENTIFIER)
				if err != nil {
					return GenqClass{}, nil
				}

				if constructorName == "fromJson" {
					if p.lookahead == TOKEN_PAREN_START {
						_, err := p.eat(TOKEN_PAREN_START)
						if err != nil {
							return GenqClass{}, err
						}

						paramType, err := p.parseTypeReference()
						if err != nil {
							return GenqClass{}, err
						}

						identifier, err := p.eat(TOKEN_IDENTIFIER)
						if err != nil {
							return GenqClass{}, err
						}

						_, err = p.eat(TOKEN_PAREN_END)
						if err != nil {
							return GenqClass{}, err
						}

						jsonConstructorSignature = GenqFromJsonConstructor{
							ParamType:  paramType,
							Identifier: identifier,
						}
						hasJsonConstructor = true
					}
				}

				err = p.eatUntil(TOKEN_SEMICOLON)
				if err != nil {
					return GenqClass{}, err
				}

				_, err = p.eat(TOKEN_SEMICOLON)
				if err != nil {
					return GenqClass{}, err
				}
			}
		} else {
			_, err := p.eat(p.lookahead)
			if err != nil {
				return GenqClass{}, err
			}
		}
	}

	return GenqClass{
		Name:                  classIdentifier,
		Constructor:           genqConstructorSignature,
		FromJsonConstructor:   jsonConstructorSignature,
		HasPrivateConstructor: hasPrivateConstructor,
		HasJsonConstructor:    hasJsonConstructor,
	}, nil
}

func (p *Parser) parseGenqConstructorSignature() (GenqConstructorSignature, *ParsingError) {
	_, err := p.eat(TOKEN_PAREN_START)
	if err != nil {
		return GenqConstructorSignature{}, err
	}

	if p.lookahead == TOKEN_PAREN_END {
		_, err := p.eat(TOKEN_PAREN_END)
		if err != nil {
			return GenqConstructorSignature{}, err
		}

		return GenqConstructorSignature{
			Params: []GenqNamedParam{},
		}, nil
	}

	_, err = p.eat(TOKEN_CURLY_START)
	if err != nil {
		return GenqConstructorSignature{}, err
	}

	params := []GenqNamedParam{}
	for p.lookahead != TOKEN_CURLY_END {
		namedParam, err := p.parseNamedParam()
		if err != nil {
			return GenqConstructorSignature{}, err
		}

		params = append(params, namedParam)

		if p.lookahead == TOKEN_COMMA {
			_, err := p.eat(TOKEN_COMMA)
			if err != nil {
				return GenqConstructorSignature{}, err
			}
		}
	}

	_, err = p.eat(TOKEN_CURLY_END)
	if err != nil {
		return GenqConstructorSignature{}, err
	}

	return GenqConstructorSignature{
		Params: params,
	}, nil
}

func (p *Parser) parseNamedParam() (GenqNamedParam, *ParsingError) {
	required := false
	annotation := GenqAnnotation{}

	if p.lookahead == TOKEN_ANNOTATION {
		_, err := p.parseAnnotation()
		if err != nil {
			p.eatUntil(TOKEN_COMMA, TOKEN_PAREN_END)
			p.eat(p.lookahead)
		}
	}

	if p.lookahead == TOKEN_REQUIRED {
		_, err := p.eat(TOKEN_REQUIRED)
		if err != nil {
			return GenqNamedParam{}, nil
		}

		required = true
	}

	paramType, err := p.parseTypeReference()
	if err != nil {
		return GenqNamedParam{}, err
	}

	paramName, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqNamedParam{}, nil
	}

	return GenqNamedParam{
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
		return "", &ParsingError{
			Err: fmt.Errorf("Unexpected token: %s (`%s`). Expected %s.", currentToken, currentTokenValue, token),
			Pos: p.lexer.cursor - len(p.lookaheadValue),
		}
	}

	p.progress()
	return currentTokenValue, nil
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
