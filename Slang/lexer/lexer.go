package lexer

import (
	"regexp"
	"strings"
)

func Lexer(code string) [][]string {
	var LexerdCode [][]string
	raw := strings.Split(strings.ReplaceAll(code, "/r/n", ""), ";")
	for _, l := range raw {
		re := regexp.MustCompile(`"[^"]*"|\S+|(<\+|<-|<\*|</|<)`)

		matches := re.FindAllString(l, -1)

		for i, match := range matches {
			if strings.HasPrefix(match, `"`) && strings.HasSuffix(match, `"`) {
				matches[i] = match[1 : len(match)-1]
			}
		}

		if len(matches) != 0 {
			LexerdCode = append(LexerdCode, matches)
		}
	}
	return LexerdCode
}
