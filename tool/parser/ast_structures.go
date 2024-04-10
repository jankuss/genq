package parser

type GenqClass struct {
	Name                  string
	Annotation            GenqAnnotation
	HasPrivateConstructor bool
	Constructor           GenqConstructorSignature
}

type GenqJsonEnum struct {
	Name       string
	EnumValues []GenqJsonEnumValue
}

type GenqJsonEnumValue struct {
	Annotation GenqAnnotation
	Name       string
}

type GenqNamedType struct {
	Name         string
	Optional     bool
	GenericTypes []GenqNamedType
	ReturnType   *GenqNamedType
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
	Reference    *GenqIdentifier
}

type GenqIdentifier struct {
	Name string
	Next *GenqIdentifier
}

type GenqConstructorSignature struct {
	Params []GenqNamedParam
}

type GenqNamedParam struct {
	Required   bool
	ParamType  GenqNamedType
	Name       string
	Annotation GenqAnnotation
}

type GenqJsonKeyAnnotation struct {
	Name string
}

type GenqPositionalParam struct {
	ParamType GenqNamedType
	Name      string
}

type GenqParamList struct {
	PositionalParams []GenqPositionalParam
	NamedParams      []GenqNamedParam
}
