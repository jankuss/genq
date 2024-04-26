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
	var hasPrimaryConstructor bool
	for _, constructor := range classDecl.Constructors {
		if constructor.Name == "" {
			primaryConstructor = constructor
			hasPrimaryConstructor = true
			break
		}
	}

	if hasPrimaryConstructor {
		str = templateMixin(str, classDecl.Name, primaryConstructor)
		str = append(str, "")
		str = templateConstructor(str, classDecl, primaryConstructor)
		str = append(str, "")
		str = templateCopyWith(str, classDecl.Name, "_$"+classDecl.Name, primaryConstructor)

		if shouldGenerateJson {
			str = append(str, "")
			str = templateFromJson(str, classDecl, primaryConstructor)
			str = append(str, "")
			str = templateToJson(str, classDecl, primaryConstructor)
		}

		str = append(str, "")
	} else {
		str = append(str, "mixin _$"+classDecl.Name+" {")
		str = append(str, indent(2, "dynamic get copyWith => throw UnimplementedError();"))
		str = append(str, "}")
		str = append(str, "")

		for _, constructor := range classDecl.Constructors {
			if constructor.Name != "" {
				str = templateSubConstructor(str, classDecl, constructor)
				str = append(str, "")
				str = templateCopyWith(str, constructor.RedirectTo, constructor.RedirectTo, constructor)

				if shouldGenerateJson {
					str = templateFromJson(str, classDecl, constructor)
					str = append(str, "")
					str = templateToJson(str, classDecl, constructor)
          str = append(str, "")
				}
			}
		}
	}

	return str
}
