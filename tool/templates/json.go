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

	// For every other type, we call the generated ${Type}FromJson method.
	return "$" + typeRef.Name + "FromJson(" + valueName + ")"
}

func typeToJsonNullable(typeRef GenqTypeReference, valueName string) string {
	if typeRef.Optional {
		return valueName + " == null ? null : " + typeToJson(typeRef, valueName+"!")
	} else {
		return typeToJson(typeRef, valueName)
	}
}

func typeToJson(typeRef GenqTypeReference, valueName string) string {
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
		return valueName + ".map((e) => " + typeToJson(typeRef.GenericTypes[0], "e") + ").toList()"
	}

	if typeRef.Name == "Set" {
		return valueName + ".map((e) => " + typeToJson(typeRef.GenericTypes[0], "e") + ").toSet()"
	}

	// For every other type, we call the generated ${Type}ToJson method.
	return "$" + typeRef.Name + "ToJson(" + valueName + ")"
}

func templateFromJson(str []string, params GenqClass) []string {
	str = append(str, fmt.Sprintf("%s $%sFromJson(Map<String, dynamic> json) {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("return %s(", params.Name)))
	for _, param := range params.Constructor.Params {
		jsonKey := param.Name
		if param.Annotation.Name == "JsonKey" {
			for _, annotationParam := range param.Annotation.NamedParams {
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

func templateToJson(str []string, params GenqClass) []string {
	str = append(str, fmt.Sprintf("Map<String, dynamic> $%sToJson(%s obj) {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("return {")))
	for _, param := range params.Constructor.Params {
		jsonKey := param.Name
		if param.Annotation.Name == "JsonKey" {
			for _, annotationParam := range param.Annotation.NamedParams {
				if annotationParam.Name == "name" {
					jsonKey = annotationParam.Value.StringValue
				}
			}
		}

		str = append(str, indent(4, fmt.Sprintf("'%s': %s,", jsonKey, typeToJsonNullable(param.ParamType, "obj."+param.Name))))
	}
	str = append(str, indent(2, fmt.Sprintf("};")))
	str = append(str, fmt.Sprintf("}"))

	return str
}
