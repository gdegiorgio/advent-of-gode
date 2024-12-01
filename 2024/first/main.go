package main

import (
	"math"
	"slices"
	"sync"
)

func main() {
	resolve([]int{1, 2}, []int{2, 3})
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
