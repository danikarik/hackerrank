package minswap

import (
	"fmt"
	"math"
	"sort"
)

func minimumSwaps(arr []int32) int32 {
	n := len(arr)
	var cnt int32
	for i := 0; i < n; i++ {
		if arr[i] == int32(i+1) {
			continue
		}
		idx := findMin(arr, i)
		arr[i], arr[idx] = arr[idx], arr[i]
		cnt++
	}
	return cnt
}

func findMin(arr []int32, start int) int {
	min := int32(math.MaxInt32)
	idx := -1
	for i := start; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			idx = i
		}
	}
	return idx
}

type pair struct {
	first  int32
	second int
}

func (p pair) String() string {
	return fmt.Sprintf("(%v %v)", p.first, p.second)
}

func minimumSwapsWithCycle(arr []int32) int32 {
	n := len(arr)
	arrPos := make([]pair, n)

	for i := 0; i < n; i++ {
		arrPos[i].first = arr[i]
		arrPos[i].second = i
	}
	sort.SliceStable(arrPos, func(i, j int) bool {
		return arrPos[i].first < arrPos[j].first
	})

	seen := make([]bool, n)
	var ans int32

	for i := 0; i < n; i++ {
		if seen[i] || arrPos[i].second == i {
			continue
		}
		cycleSize := 0
		j := i
		for !seen[j] {
			seen[j] = true
			j = arrPos[j].second
			cycleSize++
		}
		if cycleSize > 0 {
			ans += int32(cycleSize - 1)
		}
	}

	return ans
}
