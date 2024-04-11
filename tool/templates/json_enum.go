package templates

import (
	"fmt"
	. "genq/parser"
)

func TemplateJsonEnum(str []string, classDecl GenqJsonEnum) []string {
	str = TemplateJsonEnumFromJson(str, classDecl)
	str = append(str, "")
	str = TemplateJsonEnumToJson(str, classDecl)

	return str
}

func TemplateJsonEnumFromJson(str []string, classDecl GenqJsonEnum) []string {
	str = append(str, fmt.Sprintf("%s $%sFromJson(Object json, [%s? unknownEnumValue]) {", classDecl.Name, classDecl.Name, classDecl.Name))
	str = append(str, indent(2, fmt.Sprintf("switch (json) {")))
	for _, value := range classDecl.EnumValues {
		jsonValue := "\"" + value.Name + "\""
		if len(value.Annotation.Arguments.PositionalArgs) > 0 {
			jsonValue = value.Annotation.Arguments.PositionalArgs[0].RawValue
		}

		str = append(str, indent(4, fmt.Sprintf("case %s:", jsonValue)))
		str = append(str, indent(6, fmt.Sprintf("return %s.%s;", classDecl.Name, value.Name)))
	}
	str = append(str, indent(4, fmt.Sprintf("default:")))
	str = append(str, indent(6, fmt.Sprintf("if (unknownEnumValue != null) {")))
	str = append(str, indent(8, fmt.Sprintf("return unknownEnumValue;")))
	str = append(str, indent(6, fmt.Sprintf("} else {")))
	str = append(str, indent(8, fmt.Sprintf("throw UnsupportedError('The value $json is not a valid value for enum %s');", classDecl.Name)))
	str = append(str, indent(6, fmt.Sprintf("}")))
	str = append(str, indent(2, fmt.Sprintf("}")))
	str = append(str, fmt.Sprintf("}"))
	return str
}

func TemplateJsonEnumToJson(str []string, classDecl GenqJsonEnum) []string {
	str = append(str, fmt.Sprintf("Object $%sToJson(%s value) {", classDecl.Name, classDecl.Name))
	str = append(str, indent(2, fmt.Sprintf("switch (value) {")))
	for _, value := range classDecl.EnumValues {
		str = append(str, indent(4, fmt.Sprintf("case %s.%s:", classDecl.Name, value.Name)))
		jsonValue := "\"" + value.Name + "\""
		if len(value.Annotation.Arguments.PositionalArgs) > 0 {
			jsonValue = value.Annotation.Arguments.PositionalArgs[0].RawValue
		}
		str = append(str, indent(6, fmt.Sprintf("return %s;", jsonValue)))
	}
	str = append(str, indent(4, fmt.Sprintf("default:")))
	str = append(str, indent(6, fmt.Sprintf("throw UnsupportedError('Could not map $value to a JSON value');")))
	str = append(str, indent(2, fmt.Sprintf("}")))
	str = append(str, fmt.Sprintf("}"))
	return str
}
