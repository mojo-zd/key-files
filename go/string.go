package tool

import (
	"strings"
	"unicode"
)

func ToUnderLineLower(value string) string {
	r := ""
	for i, v := range value {
		if unicode.IsUpper(v) && i != 0 {
			r += "_"
		}
		r += string(v)
	}
	return strings.ToLower(r)
}

