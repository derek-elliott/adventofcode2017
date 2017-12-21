package handlers

import (
	"fmt"
	"strings"
)

// Visited holds whether a block has been visited and which step it was visited
type Visited struct {
	Visited     bool
	StepVisited int
}

// DaySix solves the sixth day of Advent of Code 2017
func DaySix(data []int) (int, int) {
	test := make([]int, len(data))
	copy(test, data)
	results := make(map[string]Visited)
	steps := 0
	for results[intListToString(test)].Visited == false {
		visited := &Visited{
			true,
			steps,
		}
		results[intListToString(test)] = *visited
		index, value := firstMax(test)
		test[index] = 0
		for value != 0 {
			index++
			if index >= len(test) {
				index = 0
			}
			test[index]++
			value--
		}
		steps++
	}
	return steps, steps - results[intListToString(test)].StepVisited
}

func firstMax(data []int) (index, max int) {
	max = data[0]
	index = 0
	for i, val := range data {
		if val == max {
			continue
		}
		if val > max {
			max = val
			index = i
		}
	}
	return index, max
}

func intListToString(list []int) string {
	return fmt.Sprint(strings.Trim(strings.Replace(fmt.Sprint(list), " ", "", -1), "[]"))
}
