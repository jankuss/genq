package parser

type GenqClassDeclaration struct {
	Name                  string
	Annotation            GenqAnnotation
	HasPrivateConstructor bool
	Constructor           GenqConstructor
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
	ParamList    GenqFormalParameterList
	IsFunction   bool
}

type GenqAnnotation struct {
	Identifier GenqIdentifier
	Arguments  GenqArgumentList
}

type GenqNamedExpression struct {
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

type GenqConstructor struct {
	ParamList GenqFormalParameterList
	IsConst   bool
}

type GenqPositionalFormalParameter struct {
	ParamType GenqNamedType
	Name      string
}

type GenqFormalNamedParameter struct {
	Required   bool
	Name       string
	ParamType  GenqNamedType
	Annotation GenqAnnotation
}

type GenqFormalParameterList struct {
	PositionalParams []GenqPositionalFormalParameter
	NamedParams      []GenqFormalNamedParameter
}

type GenqArgumentList struct {
	PositionalArgs []GenqValue
	NamedArgs      []GenqNamedExpression
}
