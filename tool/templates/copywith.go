package templates

import "fmt"
import . "genq/parser"

func templateCopyWith(str []string, name string, valueType string, constructor GenqConstructor) []string {
	str = append(str, fmt.Sprintf("abstract class $%sCopyWith {", name))

	if len(constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(2, fmt.Sprintf("%s call({", name)))
		for _, param := range constructor.ParamList.NamedParams {
			str = append(str, indent(4, fmt.Sprintf("%s %s,", param.ParamType.String(), param.Name)))
		}
		str = append(str, indent(2, fmt.Sprintf("});")))
	} else {
		str = append(str, indent(2, fmt.Sprintf("%s call();", name)))
	}

	str = append(str, fmt.Sprintf("}"))
	str = append(str, "")

	str = append(str, fmt.Sprintf("class _$%sCopyWithImpl implements $%sCopyWith {", name, name))
	str = append(str, indent(2, fmt.Sprintf("final %s value;", valueType)))
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("_$%sCopyWithImpl(this.value);", name)))
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	if len(constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(2, fmt.Sprintf("%s call({", name)))
		for _, param := range constructor.ParamList.NamedParams {
			str = append(str, indent(4, fmt.Sprintf("Object? %s = genq,", param.Name)))
		}
		str = append(str, indent(2, fmt.Sprintf("}) {")))

		str = append(str, indent(4, fmt.Sprintf("return %s(", name)))
		for _, param := range constructor.ParamList.NamedParams {
			str = append(str, indent(6, fmt.Sprintf("%s: %s == genq ? value.%s : %s as %s,", param.Name, param.Name, param.Name, param.Name, param.ParamType.String())))
		}
		str = append(str, indent(4, fmt.Sprintf(");")))

		str = append(str, indent(2, fmt.Sprintf("}")))
	} else {
		str = append(str, indent(2, fmt.Sprintf("%s call() {", name)))
		str = append(str, indent(4, fmt.Sprintf("return %s();", name)))
		str = append(str, indent(2, fmt.Sprintf("}")))
	}
	str = append(str, fmt.Sprintf("}"))
	return str
}
