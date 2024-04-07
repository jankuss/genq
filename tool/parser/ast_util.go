package parser

func (t GenqTypeReference) IsCollectionType() bool {
	return t.Name == "List" || t.Name == "Set" || t.Name == "Map"
}

func ReadAnnotationNamedParameter(annotation GenqAnnotation, name string) *GenqAnnotationParameter {
	for _, param := range annotation.NamedParams {
		if param.Name == name {
			return &param
		}
	}

	return nil
}
