package newyearchaos

import (
	"fmt"
)

func minimumBribes(q []int32) string {
	var cnt int32
	for i := len(q) - 1; i >= 0; i-- {
		if q[i]-int32(i+1) > 2 {
			return "Too chaotic"
		}
		for j := max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				cnt++
			}
		}
	}
	return fmt.Sprintf("%v", cnt)
}

func max(x, y int32) int {
	if x > y {
		return int(x)
	}
	return int(y)
}
