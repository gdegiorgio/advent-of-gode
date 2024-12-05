package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("cannot get working direcory - %s", err))
	}

	file, err := os.Open(fmt.Sprintf("%s/input.txt", workingDirectory))

	if err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	updates := []string{}
	inputConstraints := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			inputConstraints = append(inputConstraints, line)
		} else if strings.Contains(line, ",") {
			updates = append(updates, line)
		}

	}

	constraints := buildConstraints(inputConstraints)
	totalValid := count(constraints, updates)
	totalInvalid := countInvalid(constraints, updates)

	fmt.Printf("Total valid updates : %d\n", totalValid)
	fmt.Printf("Total invalid updates : %d\n", totalInvalid)

}

func countInvalid(constraints map[int][]int, updates []string) int {
	c := 0
	for _, update := range updates {
		var test []int
		for _, item := range strings.Split(update, ",") {
			val, _ := strconv.Atoi(item)
			test = append(test, val)
		}
		if !isValidUpdate(constraints, test) {
			for !isValidUpdate(constraints, test) {
				validate(constraints, test)
			}
			c += test[len(test)/2]
		}
	}
	return c
}

func count(constraints map[int][]int, updates []string) int {
	c := 0
	for _, update := range updates {
		var test []int
		for _, item := range strings.Split(update, ",") {
			val, _ := strconv.Atoi(item)
			test = append(test, val)
		}
		if isValidUpdate(constraints, test) {
			c += test[len(test)/2]
		}
	}

	return c
}

// [24 56 78 32 45]
// [56 78 32 45]
// [78 32 45]
// [32 45]
// [45]

func validate(constraints map[int][]int, update []int) {
	if len(update) == 1 {
		return
	}

	current := update[0]
	next := update[1]

	chain := constraints[current]

	mustSwap := true
	for _, item := range chain {
		if item == next {
			mustSwap = false
		}
	}

	if mustSwap {
		tmp := update[0]
		update[0] = update[1]
		update[1] = tmp
	}

	validate(constraints, update[1:])

}

func isValidUpdate(constraints map[int][]int, update []int) bool {

	//fmt.Printf("Testing : %v\n", update)
	if len(update) == 1 {
		return true
	}
	current := update[0]
	next := update[1]

	chain := constraints[current]

	for _, item := range chain {
		if item == next {
			return isValidUpdate(constraints, update[1:])
		}
	}
	return false
}

func buildConstraints(input []string) map[int][]int {
	var c map[int][]int = make(map[int][]int)
	for _, row := range input {
		split := strings.Split(row, "|")
		key, _ := strconv.Atoi(strings.TrimSpace(split[0]))
		val, _ := strconv.Atoi(strings.TrimSpace(split[1]))

		if _, ok := c[key]; ok {
			c[key] = append(c[key], val)
		} else {
			c[key] = []int{val}
		}

	}
	return c
}
