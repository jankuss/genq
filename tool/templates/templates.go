package templates

import . "genq/parser"

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

	str = templateMixin(str, classDecl, primaryConstructor)

	for _, constructor := range classDecl.Constructors {
		str = append(str, "")
		str = templateConstructor(str, classDecl, constructor)
		str = append(str, "")
		str = templateCopyWith(str, classDecl, constructor)

		if shouldGenerateJson {
			str = append(str, "")
			str = templateFromJson(str, classDecl, constructor)
			str = append(str, "")
			str = templateToJson(str, classDecl, constructor)
		}
	}

	return str
}
