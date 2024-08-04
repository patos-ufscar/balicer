package common

import (
	"regexp"
	"strings"
)

func ParseHostConfig(hostStr string) regexp.Regexp {
	str := ExtractRegExpFromHostStr(hostStr)

	return *regexp.MustCompile(str)
}

func ExtractRegExpFromHostStr(hostStr string) string {
	if strings.HasPrefix(hostStr, "r`") && strings.HasSuffix(hostStr, "`") {
		return hostStr[1+1 : len(hostStr)-1]
	}

	switch hostStr {
	case "*":
		return `.+`
	default:
		s := strings.ReplaceAll(hostStr, ".", "\\.")
		s = "^" + s + "$"
		return s
	}
}
