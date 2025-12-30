package stdfuncs

import (
	"fmt"
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"cat":     cat,
		"replace": replace,
	}
}

// [replace] replaces all occurrences of `old` with `new` in the source string `src`.
func replace(old, new, src string) string {
	return strings.Replace(src, old, new, -1)
}

// [cat] concatenates all non-nil arguments into a single space-separated string.
func cat(v ...interface{}) string {
	v = removeNilElements(v)
	r := strings.TrimSpace(strings.Repeat("%v ", len(v)))
	return fmt.Sprintf(r, v...)
}

func removeNilElements(v []interface{}) []interface{} {
	newSlice := make([]interface{}, 0, len(v))
	for _, i := range v {
		if i != nil {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}
