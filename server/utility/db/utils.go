package db

import (
	"strings"
	"unicode"
)

// 判断字符是否为字母、数字或下划线
func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// 获取一行文本中的第一个完整单词
func getFirstWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !isWordChar(r) })
	if len(fields) > 0 {
		return fields[0]
	}
	return ""
}

// 获取一行文本中的最后一个完整单词
func getLastWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !isWordChar(r) })
	if len(fields) > 0 {
		return fields[len(fields)-1]
	}
	return ""
}