package main

import "testing"

func TestResolve(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Want        int
	}{
		{Description: "Example", Input: "mul(542,794):{from()how() ^,mul(511,377)>:why():)", Want: 622995},
	}
	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			result := resolve(c.Input)
			if result != c.Want {
				t.Errorf("Resolve(%q) = %d, want %d", c.Input, result, c.Want)
			}
		})
	}
}

func TestIsValidMul(t *testing.T) {
	type Mul struct {
		left  int
		right int
	}

	cases := []struct {
		Description string
		Input       string
		Want        Mul
		err         bool
	}{
		{Description: "Valid String 1", Input: "mul(3,4)", Want: Mul{left: 3, right: 4}, err: false},
		{Description: "Valid String 2", Input: "mul(32,42)", Want: Mul{left: 32, right: 42}, err: false},
		{Description: "Valid String 3", Input: "mul(326,426)", Want: Mul{left: 326, right: 426}, err: false},
		{Description: "Empty String", Input: "", Want: Mul{left: 0, right: 0}, err: true},
		{Description: "Invalid Numbers", Input: "mul(1000,1000)", Want: Mul{left: 0, right: 0}, err: true},
		{Description: "Space Between Numbers", Input: "mul(3 , 4)", Want: Mul{left: 0, right: 0}, err: true},
		{Description: "Special Characters 1", Input: "mul(3!4)", Want: Mul{left: 0, right: 0}, err: true},
		{Description: "Special Characters 2", Input: "mul(6,9!", Want: Mul{left: 0, right: 0}, err: true},
		{Description: "Special Characters 3", Input: "mul?(12,34)", Want: Mul{left: 0, right: 0}},
		{Description: "Invalid Parenthesis 1", Input: "mul{3,4}", Want: Mul{left: 0, right: 0}, err: true},
		{Description: "Invalid Parenthesis 2", Input: "mul[3,4]", Want: Mul{left: 0, right: 0}, err: true},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			left, right, err := isValidMul(c.Input)

			if err != nil && !c.err {
				t.Errorf("isValidMul(%q) returned unexpected error: %q", c.Input, err)
			}

			if left != c.Want.left || right != c.Want.right {
				t.Errorf("isValidMul(%q) = (%d,%d), want (%d,%d)", c.Input, left, right, c.Want.left, c.Want.right)
			}
		})
	}
}

func TestIsValidCommand(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Want        int
	}{
		{Description: "Valid Do Command", Input: "do()", Want: 1},
		{Description: "Valid Don't Command", Input: "don't()", Want: -1},
		{Description: "Invalid Do Command", Input: "do_()", Want: 0},
		{Description: "Invalid Don't Command", Input: "dont()", Want: 0},
		{Description: "Special Characters 1", Input: "d0n't()", Want: 0},
		{Description: "Special Characters 2", Input: "d0()", Want: 0},
		{Description: "Special Characters 3", Input: "don!t()", Want: 0},
		{Description: "Special Characters 4", Input: "do(!)", Want: 0},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			if result := isValidCommand(c.Input); result != c.Want {
				t.Errorf("isValidCommand(%q) = %d, want %d", c.Input, result, c.Want)
			}
		})
	}
}
