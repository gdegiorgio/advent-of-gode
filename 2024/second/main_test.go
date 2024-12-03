package main

import "testing"

func TestResolve(t *testing.T) {
	cases := []struct {
		Description string
		Input       [][]int
		Safe        int
	}{
		{
			Description: "Example Input",
			Input: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			Safe: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			safe := resolve(c.Input)
			if safe != c.Safe {
				t.Errorf("resolve(%v) = %d ; want %d", c.Input, safe, c.Safe)
			}
		})
	}
}

func TestIsSafe(t *testing.T) {
	cases := []struct {
		Description string
		Input       []int
		Safe        int
	}{
		{
			Description: "Example Safe No Removing",
			Input:       []int{1, 3, 6, 7, 9},
			Safe:        1,
		},
		{
			Description: "Example Safe One Removal",
			Input:       []int{8, 6, 4, 4, 1},
			Safe:        1,
		},
		{
			Description: "Example Unsafe 1 ",
			Input:       []int{8, 6, 4, 4, 4, 1},
			Safe:        0,
		},
		{
			Description: "Example Unsafe 2 ",
			Input:       []int{57, 53, 52, 49, 48, 47, 43, 39},
			Safe:        0,
		},
		{
			Description: "Example Unsafe 3 ",
			Input:       []int{57, 53, 52, 49, 48, 47, 43, 39},
			Safe:        0,
		},
		{
			Description: "Example Unsafe 4",
			Input:       []int{81, 85, 88, 91, 93, 96, 93, 96},
			Safe:        0,
		},
		{
			Description: "Example Unsafe 5",
			Input:       []int{37, 33, 32, 29, 29, 23},
			Safe:        0,
		},
		{
			Description: "Example Unsafe 6 ",
			Input:       []int{79, 86, 88, 91, 98},
			Safe:        0,
		},
		{
			Description: "Example Unsafe 7",
			Input:       []int{61, 67, 67, 70, 72, 73, 78},
			Safe:        0, // Assuming index 1 (67) gets removed
		},
		{
			Description: "Example Unsafe 8",
			Input:       []int{55, 51, 48, 43, 45},
			Safe:        0, // Assuming index 2 (48) gets removed
		},
		{
			Description: "Example Unsafe 9",
			Input:       []int{65, 66, 65, 62, 64, 62, 55},
			Safe:        0, // Assuming index 4 (64) gets removed
		},
		{
			Description: "Example Unsafe 10",
			Input:       []int{57, 53, 51, 50, 53, 55},
			Safe:        0, // Assuming index 3 (50) gets removed
		},
		{
			Description: "Example Unsafe 11",
			Input:       []int{75, 69, 66, 69, 68, 66, 62},
			Safe:        0, // Assuming index 2 (66) gets removed
		},
		{
			Description: "Example Unsafe 12",
			Input:       []int{45, 47, 48, 48, 49, 51, 53, 53},
			Safe:        0, // Assuming index 2 (48) gets removed
		},
		{
			Description: "Example Unsafe 13",
			Input:       []int{57, 53, 52, 49, 48, 47, 43, 39},
			Safe:        0, // Assuming index 0 (57) gets removed
		},
		{
			Description: "Edge Case Remove Last",
			Input:       []int{8, 6, 4, 2, 7},
			Safe:        1,
		},
		{
			Description: "Edge Case Remove First",
			Input:       []int{7, 8, 6, 4, 2},
			Safe:        1,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			safe := isSafe(c.Input, true)
			if safe != c.Safe {
				t.Errorf("isSafe(%v) = %d ; want %d", c.Input, safe, c.Safe)
			}
		})
	}
}

func TestRemoveItem(t *testing.T) {
	arr := []int{1, 2, 4}
	expected := []int{1, 4}
	actual := removeItem(arr, 1)
	if len(actual) != len(arr)-1 {
		t.Errorf("removeItem(%v, 1) = %v ; want %v", arr, actual, expected)
	}

}
