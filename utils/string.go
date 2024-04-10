package utils

import (
	"bytes"
	"strings"
	"unicode"
)

func GetRandomStringFromList(strings ...string) string {
	if len(strings) <= 0 {
		return ""
	}
	return strings[GetRandomInt(0, len(strings)-1)]
}

func GetRandomString(chars string, length int) string {
	l := len(chars) - 1
	var b bytes.Buffer
	for i := 0; i < length; i++ {
		b.WriteByte(chars[GetRandomInt(0, l)])
	}
	return b.String()
}

func RemoveNonPrintable(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
}
