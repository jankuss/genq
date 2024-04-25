package templates

import . "genq/parser"

func Template(str []string, classDecl GenqClassDeclaration) []string {
	str = templateMixin(str, classDecl)
	str = append(str, "")
	str = templateConstructor(str, classDecl)
	str = append(str, "")
	str = templateCopyWith(str, classDecl)

	shouldGenerateJson := false
	for _, param := range classDecl.Annotation.Arguments.NamedArgs {
		if param.Name == "json" {
			shouldGenerateJson = param.Value.RawValue == "true"
		}
	}

	if shouldGenerateJson {
		str = append(str, "")
		str = templateFromJson(str, classDecl)
		str = append(str, "")
		str = templateToJson(str, classDecl)
	}

	return str
}
