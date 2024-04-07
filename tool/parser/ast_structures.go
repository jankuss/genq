package parser

type GenqClass struct {
	Name                  string
	HasPrivateConstructor bool
	Constructor           GenqConstructorSignature
	Annotation            GenqAnnotation
}

type GenqJsonEnum struct {
	Name       string
	EnumValues []GenqJsonEnumValue
}

type GenqJsonEnumValue struct {
	Annotation GenqAnnotation
	Name       string
}

type GenqTypeReference struct {
	Name         string
	Optional     bool
	GenericTypes []GenqTypeReference
	ReturnType   *GenqTypeReference
	ParamList    GenqParamList
	IsFunction   bool
}

type GenqAnnotation struct {
	Name        string
	Params      []GenqValue
	NamedParams []GenqAnnotationParameter
}

type GenqAnnotationParameter struct {
	Name  string
	Value GenqValue
}

type GenqValue struct {
	RawValue     string
	BooleanValue bool
	StringValue  string
	IntValue     int
	Reference    *GenqReference
}

type GenqReference struct {
	Name string
	Next *GenqReference
}

type GenqFromJsonConstructor struct {
	ParamType  GenqTypeReference
	Identifier string
}

type GenqConstructorSignature struct {
	Params []GenqNamedParam
}

type GenqNamedParam struct {
	Required   bool
	ParamType  GenqTypeReference
	Name       string
	Annotation GenqAnnotation
}

type GenqJsonKeyAnnotation struct {
	Name string
}

type GenqPositionalParam struct {
	ParamType GenqTypeReference
	Name      string
}

type GenqParamList struct {
	PositionalParams []GenqPositionalParam
	NamedParams      []GenqNamedParam
}
