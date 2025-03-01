package templates

import "fmt"
import . "genq/parser"

func templateCopyWith(str []string, classDecl GenqClassDeclaration) []string {
	str = append(str, fmt.Sprintf("abstract class $%sCopyWith {", classDecl.Name))

	if len(classDecl.Constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(2, fmt.Sprintf("%s call({", classDecl.Name)))
		for _, param := range classDecl.Constructor.ParamList.NamedParams {
			str = append(str, indent(4, fmt.Sprintf("%s %s,", param.ParamType.String(), param.Name)))
		}
		str = append(str, indent(2, fmt.Sprintf("});")))
	} else {
		str = append(str, indent(2, fmt.Sprintf("%s call();", classDecl.Name)))
	}

	str = append(str, fmt.Sprintf("}"))
	str = append(str, "")

	str = append(str, fmt.Sprintf("class _$%sCopyWithImpl implements $%sCopyWith {", classDecl.Name, classDecl.Name))
	if len(classDecl.Constructor.ParamList.NamedParams)+len(classDecl.Constructor.ParamList.PositionalParams) == 0 {
		str = append(str, indent(2, "// ignore: unused_field"))
	}
	str = append(str, indent(2, fmt.Sprintf("final _$%s __value;", classDecl.Name)))
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("_$%sCopyWithImpl(this.__value);", classDecl.Name)))
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	if len(classDecl.Constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(2, fmt.Sprintf("%s call({", classDecl.Name)))
		for _, param := range classDecl.Constructor.ParamList.NamedParams {
			str = append(str, indent(4, fmt.Sprintf("Object? %s = genq,", param.Name)))
		}
		str = append(str, indent(2, fmt.Sprintf("}) {")))

		str = append(str, indent(4, fmt.Sprintf("return %s(", classDecl.Name)))
		for _, param := range classDecl.Constructor.ParamList.NamedParams {
			if param.ParamType.String() == "Object?" {
				str = append(str, indent(6, fmt.Sprintf("%s: %s == genq ? __value.%s : %s,", param.Name, param.Name, param.Name, param.Name)))
			} else {
				str = append(str, indent(6, fmt.Sprintf("%s: %s == genq ? __value.%s : %s as %s,", param.Name, param.Name, param.Name, param.Name, param.ParamType.String())))
			}
		}
		str = append(str, indent(4, fmt.Sprintf(");")))

		str = append(str, indent(2, fmt.Sprintf("}")))
	} else {
		str = append(str, indent(2, fmt.Sprintf("%s call() {", classDecl.Name)))
		str = append(str, indent(4, fmt.Sprintf("return %s();", classDecl.Name)))
		str = append(str, indent(2, fmt.Sprintf("}")))
	}
	str = append(str, fmt.Sprintf("}"))
	return str
}
