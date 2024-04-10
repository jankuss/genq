package parser

import "fmt"

func (t GenqNamedType) String() string {
	if t.IsFunction {
		return fmt.Sprintf("%s Function%s", t.ReturnType.String(), t.ParamList.String())
	}

	str := t.Name
	if len(t.GenericTypes) > 0 {
		str += "<"
		for i, gt := range t.GenericTypes {
			str += gt.String()
			if i < len(t.GenericTypes)-1 {
				str += ", "
			}
		}
		str += ">"
	}
	if t.Optional {
		str += "?"
	}
	return str
}

func (t GenqParamList) String() string {
	str := "("
	for i, p := range t.PositionalParams {
		str += p.ParamType.String()
		str += " " + p.Name
		if i < len(t.PositionalParams)-1 || len(t.NamedParams) > 0 {
			str += ", "
		}
	}

	if len(t.NamedParams) > 0 {
		str += "{"

		for i, p := range t.NamedParams {
			str += p.ParamType.String()
			str += " " + p.Name
			if i < len(t.NamedParams)-1 {
				str += ", "
			}
		}

		str += "}"
	}
	str += ")"
	return str
}

func (t GenqNamedParam) String() string {
	str := ""
	if t.Required {
		str += "required "
	}
	str += t.ParamType.String()
	str += " " + t.Name
	return str
}

func (t *GenqIdentifier) String() string {
	cur := t
	str := cur.Name
	for cur != nil {
		if cur.Next != nil {
			str += "." + cur.Next.Name
		}
		cur = cur.Next
	}
	return str
}
