package main

import (
	"fmt"
	"os"
	"regexp"
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

	sum := resolve(string(buf))

	fmt.Printf("Total sum is : %d\n", sum)
}

/*
Find all "mul(X,Y)" occurrences in input.txt then return
mul(1,2) + mul(1,3)... evaluating mul(x,y) as x*y
If you find a do(), enable all next mul until you find a don't()
If you find a don't(), disable all next mul until you find a do()


Full problem here : https://adventofcode.com/2024/day/3
**/

func resolve(buf string) int {

	var (
		do  bool = true
		sum int
	)

	for i, char := range buf {
		// reading do/don't
		if char == 'd' {
			// max chars based on don't
			for j := i; j <= i+7; j++ {
				if string(buf[j]) == ")" {
					result := isValidCommand(buf[i : j+1])
					if result == 1 {
						fmt.Printf("Found a do at index %d, enabling..\n", i)
						do = true
					}
					if result == -1 {
						fmt.Printf("Found a don't at index %d, disabling..\n", i)
						do = false
					}
				}
			}
		}

		// Reading mul
		if char == 'm' {
			// Max chars based on mul(999,999)
			for j := i; j <= i+11; j++ {
				if string(buf[j]) == ")" {
					left, right, err := isValidMul(buf[i : j+1])
					if err == nil {
						if do {
							fmt.Printf("Found mul : %s\n", buf[i:j+1])
							sum += left * right
						} else {
							fmt.Printf("Found disabled: %s\n", buf[i:j+1])
						}
					}
				}
			}

		}
	}
	return sum
}

// isVaildMul checks whether the string is of type mul(x,y). If valid, returns x, y and no error, otherwise 0, 0 and error
func isValidMul(s string) (int, int, error) {

	var leftString, rightString string

	fmt.Printf("Testing %s\n", s)

	pattern, err := regexp.Compile("^mul\\([0-9]{1,3},[0-9]{1,3}\\)$")

	if err != nil {
		return 0, 0, err
	}

	matches := pattern.MatchString(s)

	if !matches {
		return 0, 0, err
	}

	split := strings.Split(s, ",")

	// find left number
	// Max digits are 3
	i := len(split[0]) - 1
	for j := 0; j < 3; j++ {
		if _, err = strconv.Atoi(string(split[0][i])); err == nil {
			leftString = string(split[0][i]) + leftString
			i--
		}
	}

	// find right number
	// Max digits are 3
	i = 0
	for j := 0; j < 3; j++ {
		if _, err = strconv.Atoi(string(split[1][i])); err == nil {
			rightString += string(split[1][i])
			i++
		}
	}

	// find right number

	left, err := strconv.Atoi(leftString)
	right, err := strconv.Atoi(rightString)

	return left, right, nil
}

// isValidCommand returns 1 if command is a Do, -1 if command is a don't 0 otherwise
func isValidCommand(s string) int {

	patternDo, err := regexp.Compile("^do\\(\\)$")
	patternDont, err := regexp.Compile("^don't\\(\\)$")

	if err != nil {
		return 0
	}

	if patternDo.MatchString(s) {
		return 1
	}

	if patternDont.MatchString(s) {
		return -1
	}

	return 0
}
