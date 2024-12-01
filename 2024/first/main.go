package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
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

	a, b, err := readInput(buf)

	fmt.Printf("Distance is %.0f\n", resolve(a, b))
}

func resolve(a []int, b []int) float64 {

	var distance float64 = 0

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(group *sync.WaitGroup) {
		slices.Sort(a)
		defer group.Done()
	}(&wg)

	go func(group *sync.WaitGroup) {
		slices.Sort(b)
		defer group.Done()
	}(&wg)

	wg.Wait()

	for i, _ := range a {
		distance += math.Abs(float64(a[i] - b[i]))
	}

	return distance
}

func readInput(buf []byte) ([]int, []int, error) {

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
