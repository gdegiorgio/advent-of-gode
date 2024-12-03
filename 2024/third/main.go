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

// Find all "mul(X,Y)" occurrences in input.txt then return
// mul(1,2) + mul(1,3)...

func resolve(buf string) int {
	sum := 0
	for i, char := range buf {
		if char == 'm' {
			// Max chars based on mul(999,999)
			for j := i; j <= i+11; j++ {
				if string(buf[j]) == ")" {
					left, right, err := isValidMul(buf[i : j+1])
					if err != nil {
						i = j
					} else {
						fmt.Printf("Found mul : %s", buf[i:j+1])
						sum += left * right
					}
				}
			}

		}
	}
	return sum
}

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

	// find left number
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
