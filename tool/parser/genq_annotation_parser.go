package parser

func genqAnnotationParser(p *Parser, listener ParserListener) *ParsingError {
	// Finally, we found a class annotated with "genq". Lets parse it.
	genqClass, err := p.parseGenqClass()
	if err != nil {
		return err
	}

	// When the class is successfully parsed, we call the listener with the parsed class.
	listener.OnGenqClass(genqClass)
	return nil
}
