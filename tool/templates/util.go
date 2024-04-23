package templates

import "strings"

func indent(num int, str string) string {
	return strings.Repeat(" ", num) + str
}

