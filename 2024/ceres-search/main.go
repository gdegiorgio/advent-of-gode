package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {

	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("cannot get working direcory - %s", err))
	}

	file, err := os.Open(fmt.Sprintf("%s/_.txt", workingDirectory))

	if err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	var m [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		m = append(m, row)
	}

	fmt.Printf("Matrix is %dx%d\n", len(m), len(m))

	count := resolvePartOne(m)
	masCount := resolvePartTwo(m)

	fmt.Printf("XMAS count is : %d\n", count)
	fmt.Printf("MAS count is : %d\n", masCount)

}

/*

Given a string matrix, find the number of "XMAS" occurrences

Full problem here : https://adventofcode.com/2024/day/4
**/

func resolvePartTwo(m [][]string) int {

	crosses := []string{}
	mas := 0

	for i, row := range m {
		if i <= len(row)-3 {
			for j, cell := range row {
				if (cell == "M" || cell == "S") && (j <= len(row)-3) {
					column := cell + m[i+1][j+1] + m[i+2][j+2]
					if isValidMAS(column) {
						crosses = append(crosses, fmt.Sprintf("%d:%d", i+1, j+1))
					}
				}
			}
		}
	}

	for i := 0; i <= len(m)-3; i++ {
		for j := len(m) - 1; j >= 2; j-- {
			if m[i][j] == "M" || m[i][j] == "S" {
				column := m[i][j] + m[i+1][j-1] + m[i+2][j-2]
				if isValidMAS(column) {
					if slices.Contains(crosses, fmt.Sprintf("%d:%d", i+1, j-1)) {
						mas++
					}
				}
			}
		}
	}

	return mas
}

func resolvePartOne(m [][]string) int {

	// Search for XMAS horizontally

	horizontal := 0

	for _, row := range m {
		for j, cell := range row {
			if (cell == "X" || cell == "S") && j <= len(row)-4 {
				if isValidXMAS(strings.Join(row[j:j+4], "")) {
					horizontal++
				}
			}
		}
	}
	fmt.Printf("Horizontal XMAS count is : %d\n", horizontal)

	vertical := 0

	for i, row := range m {
		for j, cell := range row {
			if (cell == "X" || cell == "S") && i <= len(row)-4 {
				column := cell + m[i+1][j] + m[i+2][j] + m[i+3][j]
				if isValidXMAS(column) {
					vertical++
				}
			}
		}
	}

	fmt.Printf("Vertical XMAS count is : %d\n", vertical)

	diagonal := 0

	for i, row := range m {
		if i <= len(row)-4 {
			for j, cell := range row {
				if (cell == "X" || cell == "S") && (j <= len(row)-4) {
					column := cell + m[i+1][j+1] + m[i+2][j+2] + m[i+3][j+3]
					if isValidXMAS(column) {
						diagonal++
					}
				}
			}
		}
	}

	fmt.Printf("Diagonal XMAS count is : %d\n", diagonal)

	inverseDiagonal := 0

	for i := 0; i <= len(m)-4; i++ {
		for j := len(m) - 1; j >= 3; j-- {
			if m[i][j] == "X" || m[i][j] == "S" {
				column := m[i][j] + m[i+1][j-1] + m[i+2][j-2] + m[i+3][j-3]
				if isValidXMAS(column) {
					inverseDiagonal++
				}
			}
		}
	}

	fmt.Printf("Inverse XMAS count is : %d\n", inverseDiagonal)

	return horizontal + vertical + diagonal + inverseDiagonal
}

func isValidXMAS(s string) bool {

	pattern, err := regexp.Compile("^(XMAS)|(SAMX)$")

	if err != nil {
		return false
	}

	matches := pattern.MatchString(strings.TrimSpace(s))

	if !matches {
		return false
	}
	return true
}

func isValidMAS(s string) bool {

	pattern, err := regexp.Compile("^(MAS)|(SAM)$")

	if err != nil {
		return false
	}

	matches := pattern.MatchString(strings.TrimSpace(s))

	if !matches {
		return false
	}
	return true
}
