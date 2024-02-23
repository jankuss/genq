package templates

import . "genq/parser"

func Template(str []string, params GenqClass) []string {
  str = templateMixin(str, params)
	str = append(str, "")
  str = templateConstructor(str, params)
  str = append(str, "")
  str = templateCopyWith(str, params)

	return str
}
