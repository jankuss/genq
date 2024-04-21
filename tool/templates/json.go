package templates

import (
	"fmt"
	. "genq/parser"
	"strings"
)

func typeFromJsonNullable(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	str, requiresCast := typeFromJson(annotation, typeRef, valueName)

	defaultValueAnnotation := ReadAnnotationNamedParameter(annotation, "defaultValue")
	defaultValue := "null"
	if defaultValueAnnotation != nil {
		if defaultValueAnnotation.Value.RawValue != "" {
			defaultValue = defaultValueAnnotation.Value.RawValue
		} else {
			defaultValue = defaultValueAnnotation.Value.Reference.String()
		}
	}

	expr := fmt.Sprintf("%s", str)
	if requiresCast {
		expr = fmt.Sprintf("%s as %s", str, typeRef.String())
	}

	if typeRef.Optional {
		return fmt.Sprintf("%s == null ? %s : %s", valueName, defaultValue, expr)
	} else {
		if defaultValueAnnotation != nil {
			return fmt.Sprintf("%s == null ? %s : %s", valueName, defaultValue, expr)
		}

		return expr
	}
}

func typeFromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	customFromJson := ReadAnnotationNamedParameter(annotation, "fromJson")
	if customFromJson != nil && customFromJson.Value.Reference != nil {
		return customFromJson.Value.Reference.String() + "(" + valueName + ")", false
	}

	converter, ok := converters[typeRef.Name]
	if ok {
		return converter.FromJson(annotation, typeRef, valueName)
	}

	// For every other type, we call the generated ${Type}FromJson method.
	params := []string{}
	params = append(params, valueName)

	unknownEnumValue := ReadAnnotationNamedParameter(annotation, "unknownEnumValue")
	if unknownEnumValue != nil && unknownEnumValue.Value.Reference != nil {
		params = append(params, unknownEnumValue.Value.Reference.String())
	}

	return "$" + typeRef.Name + "FromJson(" + strings.Join(params, ", ") + ")", false
}

func typeToJsonNullable(annotation GenqAnnotation, typeRef GenqNamedType, valueName string, requiresNonNullCast bool) string {
	nonNullCast := ""
	if requiresNonNullCast {
		nonNullCast = "!"
	}

	if typeRef.Optional {
		return valueName + " == null ? null : " + typeToJson(annotation, typeRef, valueName+nonNullCast)
	} else {
		return typeToJson(annotation, typeRef, valueName)
	}
}

func typeToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	customToJson := ReadAnnotationNamedParameter(annotation, "toJson")
	if customToJson != nil && customToJson.Value.Reference != nil {
		return customToJson.Value.Reference.String() + "(" + valueName + ")"
	}

	converter, ok := converters[typeRef.Name]
	if ok {
		return converter.ToJson(annotation, typeRef, valueName)
	}

	// For every other type, we call the generated ${Type}ToJson method.
	return "$" + typeRef.Name + "ToJson(" + valueName + ")"
}

func templateFromJson(str []string, classDecl GenqClassDeclaration) []string {
	str = append(str, fmt.Sprintf("%s $%sFromJson(Map<String, dynamic> json) {", classDecl.Name, classDecl.Name))
	str = append(str, indent(2, fmt.Sprintf("return %s(", classDecl.Name)))
	for _, param := range classDecl.Constructor.ParamList.NamedParams {
		jsonKey := param.Name
		if param.Annotation.Identifier.Name == "JsonKey" {
			for _, annotationParam := range param.Annotation.Arguments.NamedArgs {
				if annotationParam.Name == "name" {
					jsonKey = annotationParam.Value.StringValue
				}
			}
		}

		convName := typeFromJsonNullable(param.Annotation, param.ParamType, "json['"+jsonKey+"']")
		str = append(str, indent(4, fmt.Sprintf("%s: %s,", param.Name, convName)))
	}
	str = append(str, indent(2, fmt.Sprintf(");")))
	str = append(str, fmt.Sprintf("}"))

	return str
}

func templateToJson(str []string, params GenqClassDeclaration) []string {
	str = append(str, fmt.Sprintf("Map<String, dynamic> $%sToJson(%s obj) {", params.Name, params.Name))
	str = append(str, indent(2, fmt.Sprintf("return {")))
	for _, param := range params.Constructor.ParamList.NamedParams {
		jsonKey := param.Name
		if param.Annotation.Identifier.Name == "JsonKey" {
			for _, annotationParam := range param.Annotation.Arguments.NamedArgs {
				if annotationParam.Name == "name" {
					jsonKey = annotationParam.Value.StringValue
				}
			}
		}

		str = append(str, indent(4, fmt.Sprintf("'%s': %s,", jsonKey, typeToJsonNullable(param.Annotation, param.ParamType, "obj."+param.Name, true))))
	}
	str = append(str, indent(2, fmt.Sprintf("};")))
	str = append(str, fmt.Sprintf("}"))

	return str
}
