package templates

import . "genq/parser"

func Template(str []string, params GenqClass) []string {
	str = templateMixin(str, params)
	str = append(str, "")
	str = templateConstructor(str, params)
	str = append(str, "")
	str = templateCopyWith(str, params)

	if params.HasJsonConstructor {
		str = append(str, "")
		str = templateFromJson(str, params)
		str = append(str, "")
		str = templateToJson(str, params)
	}

	return str
}
