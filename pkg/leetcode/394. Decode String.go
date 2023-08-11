package leetcode

func decodeString(s string) string {
	decode := ""
	source := []byte(s)

}

type ByteStack struct {
	tem []byte
}

func (this *ByteStack) push(b byte) {
	this.tem = append(this.tem, b)
}

func (this *ByteStack) pop() byte {
	if len(this.tem) == 0 {
		return nil
	}
}
