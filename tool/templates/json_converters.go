package templates

import (
	. "genq/parser"
)

// map of all converters per type
var converters = map[string]jsonConverter{
  "String":  noopConverter{},
  "int":     noopConverter{},
  "double":  noopConverter{},
  "bool":    noopConverter{},
  "num":     noopConverter{},
  "List":    mappableConverter{mappableType: "List"},
  "Set":     mappableConverter{mappableType: "Set"},
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
}

func (d mappableConverter) ToJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) string {
		return valueName + ".map((e) => " + typeToJson(GenqAnnotation{}, typeRef.GenericTypes[0], "e") + ").toList()"
}

func (d mappableConverter) FromJson(annotation GenqAnnotation, typeRef GenqNamedType, valueName string) (string, bool) {
		return d.mappableType + ".of(" + valueName + ").map((e) => " + typeFromJsonNullable(GenqAnnotation{}, typeRef.GenericTypes[0], "e") + ").toList()", false
}
