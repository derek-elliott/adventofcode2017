package handlers

import (
	"fmt"
	"strings"
)

// DaySix solves the sixth day of Advent of Code 2017
func DaySix(data []int) int {
	test := make([]int, len(data))
	copy(test, data)
	results := make(map[string]bool)
	steps := 0
	for results[intListToString(test)] == false {
		results[intListToString(test)] = true
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
	return steps
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
