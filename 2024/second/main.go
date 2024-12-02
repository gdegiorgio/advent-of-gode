package main

import (
	"fmt"
	"math"
	"os"
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
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	levelsMatrix := bufferToLevels(buf)
	safe := resolve(levelsMatrix)

	fmt.Printf("Safe levels : %d\n", safe)

}

/*

Given an integer matrix return the number of "safe" array

An array of integer is "safe" if is sorted (asc or desc) and distances among elements is not greater than 3

2nd Part : An unsafe array is considered safe if removing an item meets the above property

Distance(i,j) = abs(i-j)




Full problem here : https://adventofcode.com/2024/day/2

**/

func resolve(levelsMatrix [][]int) int {
	var safe int = 0
	for _, levels := range levelsMatrix {
		safe += isSafe(levels, true)
	}
	return safe
}

// isSafe returns 1 if level is safe, 0 otherwise
func isSafe(levels []int, itemRemoval bool) int {

	// Guess the sorting
	var sortingAsc bool = levels[1] >= levels[0]
	var unsafe bool = false

	for i := 0; i < len(levels)-1; i++ {

		// If number sorting is asc but numbers are decreasing then is unsafe

		if sortingAsc && levels[i+1] < levels[i] {
			unsafe = true
		}

		// If number sorting is desc but numbers are increasing then is unsafe
		if !sortingAsc && levels[i+1] > levels[i] {
			unsafe = true
		}

		// If distance is more than 3 or less than 1
		if (math.Abs(float64(levels[i]-levels[i+1])) > 3) || (math.Abs(float64(levels[i]-levels[i+1])) < 1) {
			unsafe = true
		}

		// If I can try to remove one of the two items
		if unsafe && itemRemoval {
			return isSafe(removeItem(levels, i), false) + isSafe(removeItem(levels, i+1), false)
		}
	}

	// Still unsafe after itemRemoval or safe
	if unsafe {
		return 0
	}
	return 1
}

func bufferToLevels(buf []byte) [][]int {

	lines := strings.Split(string(buf), "\n")
	matrix := make([][]int, len(lines))

	for i, line := range lines {
		items := strings.Split(strings.TrimSpace(line), " ")
		var levelArray []int = make([]int, len(items))
		for j, item := range items {
			levelArray[j], _ = strconv.Atoi(item)
		}
		matrix[i] = levelArray
	}
	return matrix
}

func removeItem(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}
