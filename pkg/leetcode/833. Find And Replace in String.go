package leetcode

import "strings"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	sourceStr := []byte(s)
	var builder strings.Builder
	j := 0
	sortIndices(indices, sources, targets)
	for i := 0; i < len(indices); i++ {
		for j < indices[i] {
			builder.WriteByte(s[j])
			j++
		}
		match := true
		for k := 0; k < len(sources[i]); k++ {
			if j+k >= len(sourceStr) || sources[i][k] != sourceStr[j+k] {
				match = false
				break
			}
		}
		if match {
			builder.WriteString(targets[i])
			j = j + len(sources[i])
		}
	}
	for ; j < len(sourceStr); j++ {
		builder.WriteByte(s[j])
	}
	return builder.String()
}

func sortIndices(indices []int, sources []string, targets []string) {
	for i := 0; i < len(indices); i++ {
		for j := i + 1; j < len(indices); j++ {
			if indices[i] > indices[j] {
				indices[i], indices[j] = indices[j], indices[i]
				sources[i], sources[j] = sources[j], sources[i]
				targets[i], targets[j] = targets[j], targets[i]
			}
		}
	}
}
