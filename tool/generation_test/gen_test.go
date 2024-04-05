package generationtest

import (
	"testing"
)

func TestClassWithSingleParam(t *testing.T) {
	testGenOutput(t, "class_with_single_param")
}

func TestClassWithSingleOptionalParam(t *testing.T) {
	testGenOutput(t, "class_with_single_optional_param")
}

func TestClassWithMultipleParam(t *testing.T) {
	testGenOutput(t, "class_with_multiple_param")
}

func TestMultipleClassesWithMultipleParam(t *testing.T) {
	testGenOutput(t, "multiple_classes_with_multiple_param")
}

func TestAdditionalMembers(t *testing.T) {
	testGenOutput(t, "additional_members")
}

// Enums are not supported yet
// func TestGenerics(t *testing.T) {
//  testGenOutput(t, "generics");
// }

func TestListFields(t *testing.T) {
	testGenOutput(t, "list_fields")
}

func TestEnum(t *testing.T) {
	testGenOutput(t, "enum")
}

func TestCollectionTypes(t *testing.T) {
	testGenOutput(t, "collection_types")
}

func TestFunctionTypes(t *testing.T) {
	testGenOutput(t, "function_types")
}

func TestEmpty(t *testing.T) {
	testGenOutput(t, "empty")
}

func TestParameterAnnotations(t *testing.T) {
	testGenOutput(t, "parameter_annotations")
}

func TestNotRequiredParam(t *testing.T) {
	testGenOutput(t, "not_required_param")
}

func TestJsonBasic(t *testing.T) {
	testGenOutput(t, "json_basic")
}

func TestJsonEnum(t *testing.T) {
	testGenOutput(t, "json_enum")
}
