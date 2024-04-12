package parser

func genqAnnotationParser(annotation GenqAnnotation, p *Parser, listener ParserListener) *ParsingError {
	genqClass, err := p.parseGenqClass(annotation)
	if err != nil {
		return err
	}

	// When the class is successfully parsed, we call the listener with the parsed class.
	listener.OnGenqClass(genqClass)
	return nil
}
