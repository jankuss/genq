package templates

import (
	. "genq/parser"
)

// map of all converters per type
var converters = map[string]jsonConverter{
	"String":   noopConverter{},
	"int":      intConverter{},
	"double":   doubleConverter{},
	"bool":     noopConverter{},
	"num":      noopConverter{},
	"dynamic":  noopConverter{},
	"Object":   noopConverter{},
	"DateTime": dateTimeConverter{},
	"BigInt":   bigIntConverter{},
	"Uri":      uriConverter{},
	"List":     mappableConverter{mappableType: "List", methodFromJson: "toList", methodToJson: "toList"},
	"Set":      mappableConverter{mappableType: "Set", methodFromJson: "toSet", methodToJson: "toList"},
	"Map":      mapConverter{},
}

type jsonConverter interface {
	ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string
	FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool)
}

type doubleConverter struct {
}

func (d doubleConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName
}

func (d doubleConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return "(" + valueName + " as num).toDouble()", false
}

type intConverter struct {
}

func (d intConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName
}

func (d intConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return "(" + valueName + " as num).toInt()", false
}

type noopConverter struct {
}

func (d noopConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName
}

func (d noopConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return valueName, true
}

type mappableConverter struct {
	mappableType   string
	methodFromJson string
	methodToJson   string
}

func (d mappableConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName + ".map((e) => " + typeToJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[0], "e", false) + ")." + d.methodToJson + "()"
}

func (d mappableConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return d.mappableType + ".of(" + valueName + ").map((e) => " + typeFromJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[0], "e") + ")." + d.methodFromJson + "()", false
}

type dateTimeConverter struct {
}

func (d dateTimeConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName + ".toIso8601String()"
}

func (d dateTimeConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return "DateTime.parse(" + valueName + ")", false
}

type bigIntConverter struct {
}

func (d bigIntConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName + ".toString()"
}

func (d bigIntConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return "BigInt.parse(" + valueName + ")", false
}

type uriConverter struct {
}

func (d uriConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName + ".toString()"
}

func (d uriConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return "Uri.parse(" + valueName + ")", false
}

type mapConverter struct{}

func (d mapConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	keyType := typeToJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[0], "key", false)
	valueType := typeToJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[1], "value", false)

	return "Map.of(" + valueName + ").map((key, value) => MapEntry(" + keyType + ", " + valueType + "))"
}

func (d mapConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	keyType := typeFromJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[0], "key")
	valueType := typeFromJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[1], "value")

	return "Map.of(" + valueName + ").map((key, value) => MapEntry(" + keyType + ", " + valueType + "))", false
}
