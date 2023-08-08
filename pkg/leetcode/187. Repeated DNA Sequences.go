package leetcode

func findRepeatedDnaSequences(s string) []string {
	sequenceMap := make(map[string]Sequence)
	result := make([]string, 0)
	for i := 10; i <= len(s); i++ {
		if _, exist := sequenceMap[s[i-10:i]]; !exist {
			sequenceMap[s[i-10:i]] = Sequence{
				sequence: s[i-10 : i],
				count:    1,
			}
		} else {
			tem := sequenceMap[s[i-10:i]]
			tem.count = sequenceMap[s[i-10:i]].count + 1
			sequenceMap[s[i-10:i]] = tem
			if tem.count == 2 {
				result = append(result, s[i-10:i])
			}
		}
	}
	return result
}

type Sequence struct {
	sequence string
	count    int
}
