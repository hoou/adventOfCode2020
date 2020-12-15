package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testInput := "0,3,6"
	input := "1,17,0,10,18,11,6"

	fmt.Println(getLastNumber(testInput, 2020))
	fmt.Println(getLastNumber(input, 2020))
	fmt.Println(getLastNumber(input, 30000000))
}

func getLastNumber(input string, index int) int {
	numbers := getNumbers(input)
	history := initHistory(numbers)
	lastNumber := numbers[len(numbers)-1]

	for step := len(numbers) + 1; step != index+1; step++ {
		history, lastNumber = updateHistory(history, lastNumber, step)
	}

	return lastNumber
}

func updateHistory(history map[int][]int, lastNumber, step int) (map[int][]int, int) {
	nextNumber := -1
	if len(history[lastNumber]) == 1 {
		nextNumber = 0
	} else {
		nextNumber = history[lastNumber][len(history[lastNumber])-1] - history[lastNumber][len(history[lastNumber])-2]
	}

	before, existing := history[nextNumber]
	if !existing {
		before = []int{}
	}
	before = append(before, step)
	history[nextNumber] = before

	return history, nextNumber
}

func initHistory(numbers []int) map[int][]int {
	history := make(map[int][]int)
	for i, number := range numbers {
		history[number] = []int{i + 1}
	}
	return history
}

func getNumbers(input string) []int {
	numbersRaw := strings.Split(input, ",")
	var numbers []int
	for _, numberRaw := range numbersRaw {
		number, _ := strconv.Atoi(numberRaw)
		numbers = append(numbers, number)
	}
	return numbers
}
