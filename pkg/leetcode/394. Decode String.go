package leetcode

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	decodeStr, _ := decodeStringBase(s, 0)
	return decodeStr
}

func decodeStringBase(s string, index int) (string, int) {
	result := ""
	numStr := ""
	for i := index; i < len(s); i++ {
		if s[i] == '[' {
			num, _ := strconv.Atoi(numStr)
			numStr = ""
			childrenStr, nextIndex := decodeStringBase(s, i+1)
			repeatedString := strings.Repeat(childrenStr, num)
			result = result + repeatedString
			i = nextIndex
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			numStr += string(s[index])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			result = result + string(s[i])
		}
		if s[i] == ']' {
			return result, i + 1
		}
	}
	return result, index
}

/*
// Stack 结构体
type MyStack struct {
	data []interface{}
}

// Push 将元素压入栈顶
func (s *MyStack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Pop 从栈顶弹出元素
func (s *MyStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, nil
}

// IsEmpty 检查栈是否为空
func (s *MyStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size 返回栈中元素的个数
func (s *MyStack) Size() int {
	return len(s.data)
}
*/
