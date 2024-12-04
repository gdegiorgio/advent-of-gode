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
	fmt.Printf("Unsafe levels : %d\n", len(levelsMatrix)-safe)

}

/*

Given an integer matrix return the number of "safe" array

An array of integer is "safe" if is sorted (asc or desc) and distances among elements is not greater than 3 or less than 1

An unsafe array is considered safe if removing an item meets the above property

Distance(i,j) = abs(i-j)




Full problem here : https://adventofcode.com/2024/day/2

**/

func resolve(levelsMatrix [][]int) int {
	var safe int = 0
	for _, levels := range levelsMatrix {
		val := isSafe(levels, true)
		safe += val
	}
	return safe
}

// isSafe returns 1 if level is safe, 0 otherwise
func isSafe(levels []int, itemRemoval bool) int {

	var safe bool = true

	for i := 0; i < len(levels)-1; i++ {

		if levels[i] > levels[i+1] {
			if !isSortedDesc(levels) {
				safe = false
			}
		}

		if levels[i] < levels[i+1] {
			if !isSortedAsc(levels) {
				safe = false
			}
		}

		// If distance is more than 3 or less than 1
		if (math.Abs(float64(levels[i]-levels[i+1])) > 3) || (math.Abs(float64(levels[i]-levels[i+1])) < 1) {
			safe = false
		}

		// Can I remove one item?
		if !safe && itemRemoval {
			left := isSafe(removeItem(levels, i), false)
			right := isSafe(removeItem(levels, i+1), false)
			if left+right > 0 {
				return 1
			}

		}
	}
	if safe {
		return 1
	}
	return 0
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
	fmt.Printf("Removing item %d from %v\n", i, a)
	res := make([]int, 0, len(a)-1)
	res = append(res, a[:i]...)
	res = append(res, a[i+1:]...)
	fmt.Printf("Removed : %v from %v\n", res, a)
	return res
}

func isSortedAsc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func isSortedDesc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			return false
		}
	}
	return true
}
