package main

import (
	"brunoroux/advent-2022/common"
	"fmt"
	"sort"
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
	var elfSums []int
	for _, elf := range data {
		sum := sumArrayValues(elf)
		elfSums = append(elfSums, sum)
	}
	sort.Ints(elfSums)

	return sumArrayValues(elfSums[len(elfSums)-3:])
}

func sumArrayValues(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}

	return sum
}
