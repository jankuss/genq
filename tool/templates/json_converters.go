package templates

import (
	. "genq/parser"
)

// map of all converters per type
var converters = map[string]jsonConverter{
	"String":   noopConverter{},
	"int":      noopConverter{},
	"double":   noopConverter{},
	"bool":     noopConverter{},
	"num":      noopConverter{},
	"dynamic":  noopConverter{},
	"Object":   noopConverter{},
	"DateTime": dateTimeConverter{},
	"BigInt":   bigIntConverter{},
  "Uri":      uriConverter{},
	"List":     mappableConverter{mappableType: "List", method: "toList"},
	"Set":      mappableConverter{mappableType: "Set", method: "toSet"},
}

type jsonConverter interface {
	ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string
	FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool)
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
	mappableType string
	method       string
}

func (d mappableConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
	return valueName + ".map((e) => " + typeToJson(GenqAnnotation{}, typeRef.GenericTypes[0], "e") + ")." + d.method + "()"
}

func (d mappableConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
	return d.mappableType + ".of(" + valueName + ").map((e) => " + typeFromJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[0], "e") + ")." + d.method + "()", false
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
