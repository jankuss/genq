package parser

func genqJsonEnumAnnotationParser(annotation GenqAnnotation, p *Parser, listener ParserListener) *ParsingError {
	v, err := p.parseGenqJsonEnum()
	if err != nil {
		return err
	}

	listener.OnGenqJsonEnum(v)
	return nil
}

func (p *Parser) parseGenqJsonEnum() (GenqJsonEnum, *ParsingError) {
	_, err := p.eat(TOKEN_ENUM)
	if err != nil {
		return GenqJsonEnum{}, err
	}

	name, err := p.eat(TOKEN_IDENTIFIER)
	if err != nil {
		return GenqJsonEnum{}, err
	}

	_, err = p.eat(TOKEN_CURLY_START)
	if err != nil {
		return GenqJsonEnum{}, err
	}

	enumValues := []GenqJsonEnumValue{}
	for p.lookahead != TOKEN_CURLY_END {
		annotation := GenqAnnotation{}
		if p.lookahead == TOKEN_ANNOTATION {
			annotation, err = p.parseAnnotation()
			if err != nil {
				return GenqJsonEnum{}, err
			}
		}

		enumValue, err := p.eat(TOKEN_IDENTIFIER)
		if err != nil {
			return GenqJsonEnum{}, err
		}

		enumValues = append(enumValues, GenqJsonEnumValue{
			Annotation: annotation,
			Name:       enumValue,
		})

		if p.lookahead == TOKEN_COMMA {
			_, err = p.eat(TOKEN_COMMA)
			if err != nil {
				return GenqJsonEnum{}, err
			}
		}
	}

	return GenqJsonEnum{
		Name:       name,
		EnumValues: enumValues,
	}, nil
}
