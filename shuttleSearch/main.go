package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := `1003681
23,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,x,x,x,x,x,x,x,x,409,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29`

	//input := `939
	//7,13,x,x,59,x,31,19`

	arrival, buses := parseInput(input)
	result := part1(buses, arrival)
	fmt.Println(result)

	//inputTest := `
	//x,7,x,5,x,x,8`

	input2 := `
17,x,13,19`

	//fmt.Println(part2(parseBusTimes(inputTest)))
	fmt.Println(part2(parseBusTimes(input2)))
	fmt.Println(part2(parseBusTimes(`
67,7,59,61`)))
	fmt.Println(part2(parseBusTimes(input)))
}

func parseInput(input string) (int, []int) {
	inputParts := strings.Split(input, "\n")
	arrival, _ := strconv.Atoi(inputParts[0])
	var buses []int
	for _, busRaw := range strings.Split(inputParts[1], ",") {
		if busRaw == "x" {
			continue
		}
		bus, _ := strconv.Atoi(busRaw)
		buses = append(buses, bus)
	}
	return arrival, buses
}

func part1(buses []int, arrival int) int {
	sort.Ints(buses)
	if len(buses) > 0 {
		min := buses[len(buses)-1]
		closestBusNumber := buses[len(buses)-1]
		for _, bus := range buses {
			newMin := bus - arrival%bus
			if newMin < min {
				min = newMin
				closestBusNumber = bus
			}
		}
		return min * closestBusNumber
	}
	return -1
}

func parseBusTimes(input string) []BusTime {
	var result []BusTime

	inputParts := strings.Split(input, "\n")
	for i, part := range strings.Split(inputParts[1], ",") {
		if part == "x" {
			continue
		}
		id, _ := strconv.Atoi(part)
		result = append(result, BusTime{id, i})
	}

	return result
}

type BusTime struct {
	id    int
	place int
}

func part2(times []BusTime) int {
	result := 0

	var products []int

	varN := getN(times)

	for _, time := range times {
		product := 0
		if time.place != 0 {
			remainder := time.id - time.place
			ni := varN / time.id
			xi := getXi(ni, time.id)
			product = remainder * ni * xi
		}

		products = append(products, product)
	}

	for _, product := range products {
		result += product
	}
	result %= varN

	return result
}

func getXi(ni int, place int) int {
	i := 0
	result := 0
	for result != 1 {
		i++
		result = (i * ni) % place
	}
	return i
}

func getN(times []BusTime) int {
	result := 1
	for _, time := range times {
		result *= time.id
	}
	return result
}
