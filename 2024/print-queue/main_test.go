package main

import (
	"strings"
	"testing"
)

type input struct {
	Updates     []int
	Constraints string
}

func TestIsValidUpdate(t *testing.T) {
	cases := []struct {
		Description string
		Input       input
		Want        bool
	}{
		{Description: "Valid 1", Input: input{
			Updates:     []int{75, 47, 61, 53, 29},
			Constraints: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"},
			Want: true},
		{Description: "Valid 2", Input: input{
			Updates:     []int{97, 61, 53, 29, 13},
			Constraints: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"},
			Want: true},
		{Description: "Valid 3", Input: input{
			Updates:     []int{75, 29, 13},
			Constraints: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"},
			Want: true},
		{Description: "Invalid 1", Input: input{
			Updates:     []int{75, 97, 47, 61, 53},
			Constraints: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"},
			Want: false},
		{Description: "Invalid 2", Input: input{
			Updates:     []int{61, 13, 29},
			Constraints: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"},
			Want: false},
		{Description: "Invalid 3", Input: input{
			Updates:     []int{97, 13, 75, 29, 47},
			Constraints: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"},
			Want: false},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			constraints := buildConstraints(strings.Split(c.Input.Constraints, "\n"))
			valid := isValidUpdate(constraints, c.Input.Updates)
			if valid != c.Want {
				t.Errorf("isValidUpdate= %v, want %v", valid, c.Want)
			}
		})
	}
}
