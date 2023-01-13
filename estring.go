package estring

import (
	"regexp"
	"strings"
)

// UcFirst 字符串首字母大写
func UcFirst(s string) string {
	if len(s) == 0 || s[0] < 'a' || s[0] > 'z' {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// LcFirst 字符串首字母小写
func LcFirst(s string) string {
	if len(s) == 0 || s[0] < 'A' || s[0] > 'Z' {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func Studly(s string) string {
	s = regexp.MustCompile("[-._ ]").ReplaceAllString(s, " ")
	splits := strings.Split(s, " ")

	for i, split := range splits {
		splits[i] = UcFirst(split)
	}

	return strings.Join(splits, "")
}

func Camel(s string) string {
	return LcFirst(Studly(s))
}

func PadLeft(s string, length int, pad string) string {
	sLen := len(s)
	if length < 1 || sLen >= length {
		return s
	}

	return strings.Repeat(pad, length-sLen) + s
}

func PadRight(s string, length int, pad string) string {
	sLen := len(s)
	if length < 1 || sLen >= length {
		return s
	}

	return s + strings.Repeat(pad, length-sLen)
}
