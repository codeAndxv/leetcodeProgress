package leetcode

import "strings"

func isAcronym(words []string, s string) bool {
	var builder strings.Builder

	for i := 0; i < len(words); i++ {
		builder.WriteByte(words[i][0])
	}
	return builder.String() == s
}
