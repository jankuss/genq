package templates

import (
	"fmt"
	. "genq/parser"
)

func TemplateJsonEnum(str []string, params GenqJsonEnum) []string {
	str = TemplateJsonEnumFromJson(str, params)
	str = append(str, "")
	str = TemplateJsonEnumToJson(str, params)

	return str
}

func TemplateJsonEnumFromJson(str []string, params GenqJsonEnum) []string {
	str = append(str, fmt.Sprintf("%s $%sFromJson(Object json) {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("switch (json) {")))
	for _, value := range params.EnumValues {
		jsonValue := "\"" + value.Name + "\""
		if len(value.Annotation.Params) > 0 {
			jsonValue = value.Annotation.Params[0].RawValue
		}

		str = append(str, indent(4, fmt.Sprintf("case %s:", jsonValue)))
		str = append(str, indent(6, fmt.Sprintf("return %s.%s;", params.Name, value.Name)))
	}
	str = append(str, indent(4, fmt.Sprintf("default:")))
	str = append(str, indent(6, fmt.Sprintf("throw Exception(\"Unknown %s value: $json\");", params.Name)))
	str = append(str, indent(2, fmt.Sprintf("}")))
	str = append(str, fmt.Sprintf("}"))
	return str
}

func TemplateJsonEnumToJson(str []string, params GenqJsonEnum) []string {
	str = append(str, fmt.Sprintf("Object $%sToJson(%s value) {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("switch (value) {")))
	for _, value := range params.EnumValues {
		str = append(str, indent(4, fmt.Sprintf("case %s.%s:", params.Name, value.Name)))
		jsonValue := "\"" + value.Name + "\""
		if len(value.Annotation.Params) > 0 {
			jsonValue = value.Annotation.Params[0].RawValue
		}
		str = append(str, indent(6, fmt.Sprintf("return %s;", jsonValue)))
	}
	str = append(str, indent(4, fmt.Sprintf("default:")))
	str = append(str, indent(6, fmt.Sprintf("throw Exception(\"Unknown %s value: $value\");", params.Name)))
	str = append(str, indent(2, fmt.Sprintf("}")))
	str = append(str, fmt.Sprintf("}"))
	return str
}
