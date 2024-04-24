package templates

import (
	. "genq/parser"
)

func Template(str []string, classDecl GenqClassDeclaration) []string {
	shouldGenerateJson := false
	for _, param := range classDecl.Annotation.Arguments.NamedArgs {
		if param.Name == "json" {
			shouldGenerateJson = param.Value.BooleanValue
		}
	}

	var primaryConstructor GenqConstructor
	for _, constructor := range classDecl.Constructors {
		if constructor.Name == "" {
			primaryConstructor = constructor
			break
		}
	}

	str = templateMixin(str, classDecl.Name, primaryConstructor)
	str = append(str, "")
	str = templateConstructor(str, classDecl, primaryConstructor)
	str = append(str, "")
	str = templateCopyWith(str, classDecl.Name, primaryConstructor)

	if shouldGenerateJson {
		str = append(str, "")
		str = templateFromJson(str, classDecl, primaryConstructor)
		str = append(str, "")
		str = templateToJson(str, classDecl, primaryConstructor)
	}

	// Create empty mixin
	str = append(str, "")
	for _, constructor := range classDecl.Constructors {
		if constructor.Name != "" {
			str = templateMixin(str, constructor.RedirectTo, constructor)
			str = append(str, "")
			str = templateSubConstructor(str, classDecl, constructor)
			str = append(str, "")
			str = templateCopyWith(str, constructor.RedirectTo, constructor)
		}

	}

	return str
}
