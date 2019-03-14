package minswap

import "testing"

var testCases = []struct {
	Given    []int32
	Expected int32
}{
	{
		Given:    []int32{4, 3, 1, 2},
		Expected: 3,
	},
	{
		Given:    []int32{2, 3, 4, 1, 5},
		Expected: 3,
	},
	{
		Given:    []int32{7, 1, 3, 2, 4, 5, 6},
		Expected: 5,
	},
}

var testCasesWithCycle = []struct {
	Given    []int32
	Expected int32
}{
	{
		Given:    []int32{4, 3, 1, 2},
		Expected: 3,
	},
	{
		Given:    []int32{2, 3, 4, 1, 5},
		Expected: 3,
	},
	{
		Given:    []int32{7, 1, 3, 2, 4, 5, 6},
		Expected: 5,
	},
}

func TestMinimumSwaps(t *testing.T) {
	for _, c := range testCases {
		res := minimumSwaps(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

func TestMinimumSwapsWithCycle(t *testing.T) {
	for _, c := range testCasesWithCycle {
		res := minimumSwapsWithCycle(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var (
	result          int32
	resultWithCycle int32
)

func BenchmarkMinimumSwaps(b *testing.B) {
	var r int32
	for i := 0; i < b.N; i++ {
		r = minimumSwaps(testCases[0].Given)
	}
	result = r
}

func BenchmarkMinimumSwapsWithCycle(b *testing.B) {
	var r int32
	for i := 0; i < b.N; i++ {
		r = minimumSwapsWithCycle(testCases[0].Given)
	}
	result = r
}
