package main

import (
	"brunoroux/advent-2022/common"
	"fmt"
	"strconv"
)

func main() {
	data := formatInput()
	result := processData(data)
	fmt.Println(result)
}

func formatInput() [][]int {
	input := common.ReadInput()

	var data [][]int
	var current []int

	for _, calories := range input {
		if calories != "" {
			caloriesInt, _ := strconv.Atoi(calories)
			current = append(current, caloriesInt)
		} else {
			data = append(data, current)
			current = nil
		}
	}

	return data
}

func processData(data [][]int) int {
	max := 0
	for _, elf := range data {
		sum := 0
		for _, calories := range elf {
			sum += calories
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
