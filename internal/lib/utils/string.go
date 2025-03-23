package utils

import "strings"

func GetSubstringIfSymbolExists(str string, symbol string) string {
	index := strings.Index(str, symbol)
	if index == -1 {
		return ""
	}
	return str[:index]
}
