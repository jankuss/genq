package templates

import (
	"fmt"
	. "genq/parser"
)

func extendsOrImplements(classDecl GenqClassDeclaration) string {
	if classDecl.HasPrivateConstructor {
		return "extends"
	} else {
		return "implements"
	}
}

func templateMixin(str []string, name string, constructor GenqConstructor) []string {
	str = append(str, fmt.Sprintf("mixin _$%s {", name))
	for _, param := range constructor.ParamList.NamedParams {
		str = append(str, indent(2, fmt.Sprintf("%s get %s => throw UnimplementedError();", param.ParamType.String(), param.Name)))
	}
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("$%sCopyWith get copyWith => throw UnimplementedError();", name)))
	str = append(str, fmt.Sprintf("}"))

	return str
}

func templateSubConstructor(str []string, classDecl GenqClassDeclaration, constructor GenqConstructor) []string {
	str = append(str, fmt.Sprintf("class %s implements %s {", constructor.RedirectTo, classDecl.Name))

	for _, param := range constructor.ParamList.NamedParams {
		str = append(str, indent(2, fmt.Sprintf("@override")))
		str = append(str, indent(2, fmt.Sprintf("final %s %s;", param.ParamType.String(), param.Name)))
		str = append(str, "")
	}

	if len(constructor.ParamList.NamedParams) > 0 {
		var constructorDecl string
		if constructor.IsConst {
			constructorDecl = fmt.Sprintf("const %s({", constructor.RedirectTo)
		} else {
			constructorDecl = fmt.Sprintf("%s({", constructor.RedirectTo)
		}

		str = append(str, indent(2, constructorDecl))
		for _, param := range constructor.ParamList.NamedParams {
			if param.Required {
				str = append(str, indent(4, fmt.Sprintf("required this.%s,", param.Name)))
			} else {
				str = append(str, indent(4, fmt.Sprintf("this.%s,", param.Name)))
			}
		}
		if classDecl.HasPrivateConstructor {
			str = append(str, indent(2, fmt.Sprintf("}) : super._();")))
		} else {
			str = append(str, indent(2, fmt.Sprintf("});")))
		}
	} else {
		if classDecl.HasPrivateConstructor {
			if constructor.IsConst {
				str = append(str, indent(2, fmt.Sprintf("const %s() : super._();", constructor.RedirectTo)))
			} else {
				str = append(str, indent(2, fmt.Sprintf("%s() : super._();", constructor.RedirectTo)))
			}
		} else {
			if constructor.IsConst {
				str = append(str, indent(2, fmt.Sprintf("const %s();", constructor.RedirectTo)))
			} else {
				str = append(str, indent(2, fmt.Sprintf("%s();", constructor.RedirectTo)))
			}
		}
	}

	// copyWith
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("$%sCopyWith get copyWith => _$%sCopyWithImpl(this);", constructor.RedirectTo, constructor.RedirectTo)))
	str = append(str, "")

	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("String toString() {")))

	toStringParams := ""
	for i, param := range constructor.ParamList.NamedParams {
		if i == len(constructor.ParamList.NamedParams)-1 {
			toStringParams += fmt.Sprintf("%s: $%s", param.Name, param.Name)
		} else {
			toStringParams += fmt.Sprintf("%s: $%s, ", param.Name, param.Name)
		}
	}

	str = append(str, indent(4, fmt.Sprintf("return \"%s(%s)\";", constructor.RedirectTo, toStringParams)))
	str = append(str, indent(2, fmt.Sprintf("}")))
	str = append(str, "")

	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("bool operator ==(Object other) {")))
	str = templateEqualityBody(str, constructor.RedirectTo, constructor)
	str = append(str, indent(2, fmt.Sprintf("}")))

	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("int get hashCode {")))

	if len(constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(4, fmt.Sprintf("return Object.hash(")))
		str = append(str, indent(6, "runtimeType,"))
		for _, param := range constructor.ParamList.NamedParams {
			str = append(str, indent(6, fmt.Sprintf("%s,", param.Name)))
		}
		str = append(str, indent(4, fmt.Sprintf(");")))
	} else {
		str = append(str, indent(4, fmt.Sprintf("return runtimeType.hashCode;")))
	}

	str = append(str, indent(2, fmt.Sprintf("}")))

	str = append(str, "}")

	return str
}

func templateConstructor(str []string, classDecl GenqClassDeclaration, constructor GenqConstructor) []string {
	str = append(str, fmt.Sprintf("class _%s %s %s {", classDecl.Name, extendsOrImplements(classDecl), classDecl.Name))

	for _, param := range constructor.ParamList.NamedParams {
		str = append(str, indent(2, fmt.Sprintf("@override")))
		str = append(str, indent(2, fmt.Sprintf("final %s %s;", param.ParamType.String(), param.Name)))
		str = append(str, "")
	}

	if len(constructor.ParamList.NamedParams) > 0 {
		var constructorDecl string
		if constructor.IsConst {
			constructorDecl = fmt.Sprintf("const _%s({", classDecl.Name)
		} else {
			constructorDecl = fmt.Sprintf("_%s({", classDecl.Name)
		}

		str = append(str, indent(2, constructorDecl))
		for _, param := range constructor.ParamList.NamedParams {
			if param.Required {
				str = append(str, indent(4, fmt.Sprintf("required this.%s,", param.Name)))
			} else {
				str = append(str, indent(4, fmt.Sprintf("this.%s,", param.Name)))
			}
		}
		if classDecl.HasPrivateConstructor {
			str = append(str, indent(2, fmt.Sprintf("}) : super._();")))
		} else {
			str = append(str, indent(2, fmt.Sprintf("});")))
		}
	} else {
		if classDecl.HasPrivateConstructor {
			if constructor.IsConst {
				str = append(str, indent(2, fmt.Sprintf("const _%s() : super._();", classDecl.Name)))
			} else {
				str = append(str, indent(2, fmt.Sprintf("_%s() : super._();", classDecl.Name)))
			}
		} else {
			if constructor.IsConst {
				str = append(str, indent(2, fmt.Sprintf("const _%s();", classDecl.Name)))
			} else {
				str = append(str, indent(2, fmt.Sprintf("_%s();", classDecl.Name)))
			}
		}
	}

	// copyWith
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("$%sCopyWith get copyWith => _$%sCopyWithImpl(this);", classDecl.Name, classDecl.Name)))
	str = append(str, "")

	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("String toString() {")))

	toStringParams := ""
	for i, param := range constructor.ParamList.NamedParams {
		if i == len(constructor.ParamList.NamedParams)-1 {
			toStringParams += fmt.Sprintf("%s: $%s", param.Name, param.Name)
		} else {
			toStringParams += fmt.Sprintf("%s: $%s, ", param.Name, param.Name)
		}
	}

	str = append(str, indent(4, fmt.Sprintf("return \"%s(%s)\";", classDecl.Name, toStringParams)))
	str = append(str, indent(2, fmt.Sprintf("}")))
	str = append(str, "")

	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("bool operator ==(Object other) {")))
	str = templateEqualityBody(str, classDecl.Name, constructor)
	str = append(str, indent(2, fmt.Sprintf("}")))

	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("int get hashCode {")))

	if len(constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(4, fmt.Sprintf("return Object.hash(")))
		str = append(str, indent(6, "runtimeType,"))
		for _, param := range constructor.ParamList.NamedParams {
			str = append(str, indent(6, fmt.Sprintf("%s,", param.Name)))
		}
		str = append(str, indent(4, fmt.Sprintf(");")))
	} else {
		str = append(str, indent(4, fmt.Sprintf("return runtimeType.hashCode;")))
	}

	str = append(str, indent(2, fmt.Sprintf("}")))

	str = append(str, "}")

	return str
}

func templateEqualityBody(str []string, name string, constructor GenqConstructor) []string {
	str = append(str, indent(4, fmt.Sprintf("if (identical(this, other)) return true;")))
	str = append(str, indent(4, fmt.Sprintf("if (other is! %s) return false;", name)))

	for _, param := range constructor.ParamList.NamedParams {
		if param.ParamType.IsCollectionType() {
			str = append(str, indent(4, fmt.Sprintf("if (!const DeepCollectionEquality().equals(other.%s, %s)) return false;", param.Name, param.Name)))
		} else {
			str = append(str, indent(4, fmt.Sprintf("if (!identical(other.%s, %s) && other.%s != %s) return false;", param.Name, param.Name, param.Name, param.Name)))
		}
	}

	str = append(str, indent(4, fmt.Sprintf("return true;")))

	return str
}
