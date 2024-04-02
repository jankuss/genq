package parser

type GenqClass struct {
	Name                  string
	HasPrivateConstructor bool
	Constructor           GenqConstructorSignature
	FromJsonConstructor   GenqFromJsonConstructor
	HasJsonConstructor    bool
}

type GenqTypeReference struct {
	Name         string
	Optional     bool
	GenericTypes []GenqTypeReference
	FunctionType *FunctionType
	ReturnType   *GenqTypeReference
	ParamList    GenqParamList
	IsFunction   bool
}

type GenqAnnotation struct {
	Name   string
	Params []GenqAnnotationParameter
}

type GenqAnnotationParameter struct {
	Name  string
	Value GenqValue
}

type GenqValue struct {
	BooleanValue bool
	StringValue  string
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
