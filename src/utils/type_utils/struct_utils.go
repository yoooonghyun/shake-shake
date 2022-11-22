package type_utils

import (
	"reflect"
	"strings"
)

func GetStructNameSnake(data interface{}) string {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return toSnakeCase(t.Name())
}

func toSnakeCase(name string) string {
	var result strings.Builder
	for i, char := range name {
		if i > 0 && (char >= 'A' && char <= 'Z') {
			result.WriteByte('_')
		}
		result.WriteRune(char)
	}
	return strings.ToLower(result.String())
}
