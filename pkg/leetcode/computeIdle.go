package leetcode

import (
	"fmt"
	"time"
)

func computeIdleTime(a int, b int) time.Duration {
	idleRate := float32(a) / float32(b) * 100
	fmt.Println(idleRate)
	durationInSeconds := int64(-1.7*idleRate + 180)
	return time.Duration(durationInSeconds) * time.Second
}
