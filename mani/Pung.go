package mani

import (
	"regexp"
	"strings"
)

func Pung(str string) string {
	pung := regexp.MustCompile(`\s+([,.:;!?]+)`)
	pun := regexp.MustCompile(`([,.:;!?]+)(\w+)`)

	str = pung.ReplaceAllString(str, "$1")
	str1 := pun.ReplaceAllString(str, "$1 $2")
	return str1
}

func SingleQ(str string) string {
	str = strings.ReplaceAll(str, "' ", " '")
	str = strings.ReplaceAll(str, " '", "'")

	return strings.TrimSpace(str)
}
