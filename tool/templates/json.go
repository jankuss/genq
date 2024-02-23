package templates

import (
	"fmt"
	. "genq/parser"
)


func typeFromJsonNullable(typeRef GenqTypeReference, valueName string) string {
	str := typeFromJson(typeRef, valueName)

	if typeRef.Optional {
		return valueName + " == null ? null : (" + str + " as " + typeRef.String() + ")"
	} else {
		return str + " as " + typeRef.String()
	}
}

func typeFromJson(typeRef GenqTypeReference, valueName string) string {
	if typeRef.Name == "String" {
		return valueName
	}

	if typeRef.Name == "int" {
		return valueName
	}

	if typeRef.Name == "double" {
		return valueName
	}

	if typeRef.Name == "bool" {
		return valueName
	}

	if typeRef.Name == "num" {
		return valueName
	}

	if typeRef.Name == "List" {
		return "List.of(" + valueName + ").map((e) => " + typeFromJsonNullable(typeRef.GenericTypes[0], "e") + ").toList()"
	}

	if typeRef.Name == "Set" {
		return "Set.of(" + valueName + ").map((e) => " + typeFromJsonNullable(typeRef.GenericTypes[0], "e") + ").toSet()"
	}

	// By default, we assume it's a class with a static member/factory method called fromJson
	return typeRef.Name + ".fromJson(" + valueName + ")"
}

func templateFromJson(str []string, params GenqClass) []string {
	str = append(str, fmt.Sprintf("%s _$%sFromJson(Map<String, dynamic> json) {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("return %s(", params.Name)))
	for _, param := range params.Constructor.Params {

		jsonKey := param.Name
		if param.Annotation.Name == "JsonKey" {
			for _, annotationParam := range param.Annotation.Params {
				if annotationParam.Name == "name" {
					jsonKey = annotationParam.Value.StringValue
				}
			}
		}

		convName := typeFromJsonNullable(param.ParamType, "json['"+jsonKey+"']")
		str = append(str, indent(4, fmt.Sprintf("%s: %s,", param.Name, convName)))
	}
	str = append(str, indent(2, fmt.Sprintf(");")))
	str = append(str, fmt.Sprintf("}"))

	return str
}
