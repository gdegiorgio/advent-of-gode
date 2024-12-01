package first

import "testing"

type Input struct {
	a []int
	b []int
}

func TestResolve(t *testing.T) {
	cases := []struct {
		Description string
		Input       Input
		Expected    []int
	}{
		{
			Description: "Example Input",
			Input:       Input{a: []int{3, 4, 2, 1, 3, 3}, b: []int{4, 3, 5, 3, 9, 3}},
			Expected:    []int{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {

		})
	}
}
