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

func templateMixin(str []string, classDecl GenqClassDeclaration) []string {
	str = append(str, fmt.Sprintf("mixin _$%s {", classDecl.Name))
	for _, param := range classDecl.Constructor.ParamList.NamedParams {
		str = append(str, indent(2, fmt.Sprintf("%s get %s => throw UnimplementedError();", param.ParamType.String(), param.Name)))
	}
	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("$%sCopyWith get copyWith => throw UnimplementedError();", classDecl.Name)))
	str = append(str, fmt.Sprintf("}"))

	return str
}

func templateConstructor(str []string, classDecl GenqClassDeclaration) []string {
	str = append(str, fmt.Sprintf("class _%s %s %s {", classDecl.Name, extendsOrImplements(classDecl), classDecl.Name))

	for _, param := range classDecl.Constructor.ParamList.NamedParams {
		str = append(str, indent(2, fmt.Sprintf("@override")))
		str = append(str, indent(2, fmt.Sprintf("final %s %s;", param.ParamType.String(), param.Name)))
		str = append(str, "")
	}

	if len(classDecl.Constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(2, fmt.Sprintf("_%s({", classDecl.Name)))
		for _, param := range classDecl.Constructor.ParamList.NamedParams {
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
			str = append(str, indent(2, fmt.Sprintf("_%s() : super._();", classDecl.Name)))
		} else {
			str = append(str, indent(2, fmt.Sprintf("_%s();", classDecl.Name)))
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
	for i, param := range classDecl.Constructor.ParamList.NamedParams {
		if i == len(classDecl.Constructor.ParamList.NamedParams)-1 {
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
	str = templateEqualityBody(str, classDecl)
	str = append(str, indent(2, fmt.Sprintf("}")))

	str = append(str, "")
	str = append(str, indent(2, fmt.Sprintf("@override")))
	str = append(str, indent(2, fmt.Sprintf("int get hashCode {")))

	if len(classDecl.Constructor.ParamList.NamedParams) > 0 {
		str = append(str, indent(4, fmt.Sprintf("return Object.hash(")))
		str = append(str, indent(6, "runtimeType,"))
		for _, param := range classDecl.Constructor.ParamList.NamedParams {
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

func templateEqualityBody(str []string, classDecl GenqClassDeclaration) []string {
	str = append(str, indent(4, fmt.Sprintf("if (identical(this, other)) return true;")))
	str = append(str, indent(4, fmt.Sprintf("if (other is! %s) return false;", classDecl.Name)))

	for _, param := range classDecl.Constructor.ParamList.NamedParams {
		if param.ParamType.IsCollectionType() {
			str = append(str, indent(4, fmt.Sprintf("if (!const DeepCollectionEquality().equals(other.%s, %s)) return false;", param.Name, param.Name)))
		} else {
			str = append(str, indent(4, fmt.Sprintf("if (!identical(other.%s, %s) && other.%s != %s) return false;", param.Name, param.Name, param.Name, param.Name)))
		}
	}

	str = append(str, indent(4, fmt.Sprintf("return true;")))

	return str
}
