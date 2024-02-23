package templates

import "fmt"
import . "genq/parser"

func templateCopyWith(str []string, params GenqClass) []string {
	str = append(str, fmt.Sprintf("abstract class $%sCopyWith {", params.Name))

	if len(params.Constructor.Params) > 0 {
		str = append(str, indent(2, fmt.Sprintf("%s call({", params.Name)))
		for _, param := range params.Constructor.Params {
			str = append(str, indent(4, fmt.Sprintf("%s %s,", param.ParamType.String(), param.Name)))
		}
		str = append(str, indent(2, fmt.Sprintf("});")))
	} else {
		str = append(str, indent(2, fmt.Sprintf("%s call();", params.Name)))
	}

	str = append(str, fmt.Sprintf("}"))
	str = append(str, "")

	str = append(str, fmt.Sprintf("class _$%sCopyWithImpl implements $%sCopyWith {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("final _$%s value;", params.Name)))
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("_$%sCopyWithImpl(this.value);", params.Name)))
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	if len(params.Constructor.Params) > 0 {
		str = append(str, indent(2, fmt.Sprintf("%s call({", params.Name)))
		for _, param := range params.Constructor.Params {
			str = append(str, indent(4, fmt.Sprintf("Object? %s = genq,", param.Name)))
		}
		str = append(str, indent(2, fmt.Sprintf("}) {")))

		str = append(str, indent(4, fmt.Sprintf("return %s(", params.Name)))
		for _, param := range params.Constructor.Params {
			str = append(str, indent(6, fmt.Sprintf("%s: %s == genq ? value.%s : %s as %s,", param.Name, param.Name, param.Name, param.Name, param.ParamType.String())))
		}
		str = append(str, indent(4, fmt.Sprintf(");")))

		str = append(str, indent(2, fmt.Sprintf("}")))
	} else {
		str = append(str, indent(2, fmt.Sprintf("%s call() {", params.Name)))
    str = append(str, indent(4, fmt.Sprintf("return %s();", params.Name)))
		str = append(str, indent(2, fmt.Sprintf("}")))
	}
	str = append(str, fmt.Sprintf("}"))
	return str
}
