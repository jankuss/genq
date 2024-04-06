package templates

import . "genq/parser"

func Template(str []string, params GenqClass) []string {
	str = templateMixin(str, params)
	str = append(str, "")
	str = templateConstructor(str, params)
	str = append(str, "")
	str = templateCopyWith(str, params)

	shouldGenerateJson := false
	for _, param := range params.Annotation.NamedParams {
		if param.Name == "json" {
			shouldGenerateJson = param.Value.BooleanValue
		}
	}

	if shouldGenerateJson {
		str = append(str, "")
		str = templateFromJson(str, params)
		str = append(str, "")
		str = templateToJson(str, params)
	}

	return str
}
