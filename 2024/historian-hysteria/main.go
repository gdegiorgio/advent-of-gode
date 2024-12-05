package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("cannot get working direcory - %s", err))
	}

	buf, err := os.ReadFile(fmt.Sprintf("%s/input.txt", workingDirectory))

	if err != nil {
		panic(fmt.Errorf("Error reading input file: %w", err))
	}

	a, b, err := transformInput(buf)

	distance, similarity := resolve(a, b)

	fmt.Printf("Distance is %.0f\n Similarity is %d\n", distance, similarity)

}

/*

Given two unordered integer array, find the total distance and similarity among them

Pair up the smallest number in the left list with the smallest number in the right list, then the red-nosed-reports-smallest left number with the red-nosed-reports-smallest right number, and so on.
Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances.

For example, if you pair up a 3 from the left list with a 7 from the right list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.

To find the total distance between the left list and the right list, add up the distances between all the pairs you found.
Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.

Full problem here : https://adventofcode.com/2024/day/1

**/

func resolve(a []int, b []int) (float64, int) {

	var (
		distance    float64
		similarity  int
		occurrences map[int]int = make(map[int]int, len(b))
	)

	slices.Sort(a)
	slices.Sort(b)

	for _, item := range b {
		occurrences[item]++
	}

	for i, item := range a {
		similarity += item * occurrences[item]
		distance += math.Abs(float64(a[i] - b[i]))
	}

	return distance, similarity
}

func transformInput(buf []byte) ([]int, []int, error) {

	lines := strings.Split(string(buf), "\n")

	a := make([]int, len(lines))
	b := make([]int, len(lines))

	for i, line := range lines {
		values := strings.Split(strings.TrimSpace(line), "   ")
		aValue, err := strconv.Atoi(values[0])
		bValue, err := strconv.Atoi(values[1])

		if err != nil {
			fmt.Errorf("%w", err)
		}

		a[i] = aValue
		b[i] = bValue
	}
	return a, b, nil
}
