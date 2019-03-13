package newyearchaos

import "testing"

var testCases = []struct {
	Given    []int32
	Expected string
}{
	{
		Given:    []int32{1, 2, 3, 4},
		Expected: "0",
	},
	{
		Given:    []int32{1, 3, 2, 4},
		Expected: "1",
	},
	{
		Given:    []int32{1, 3, 4, 2},
		Expected: "2",
	},
	{
		Given:    []int32{2, 1, 5, 3, 4},
		Expected: "3",
	},
	{
		Given:    []int32{2, 5, 1, 3, 4},
		Expected: "Too chaotic",
	},
	{
		Given:    []int32{5, 1, 2, 3, 7, 8, 6, 4},
		Expected: "Too chaotic",
	},
	{
		Given:    []int32{1, 2, 5, 3, 7, 8, 6, 4},
		Expected: "7",
	},
	{
		Given:    []int32{1, 2, 5, 3, 4, 7, 8, 6},
		Expected: "4",
	},
}

func TestMinimumBribes(t *testing.T) {
	for _, c := range testCases {
		res := minimumBribes(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}
