package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"uk.co.lewis-od.aoc2020/common"
)

func main() {
	input := common.ReadAndSanitiseRows("example2.txt")
	adapterRatings := parseNumbers(input)

	adapterChain, _ := findChain(adapterRatings, make([]int, 0, len(adapterRatings)), 0)
	differences := append(adapterChain, 0)
	laptopRating := differences[0] + 3
	differences = append([]int{laptopRating}, differences...)
	ones, threes := countDifferences(differences)
	fmt.Println("Part 1:", ones*threes)

	numChains, _ := countArrangements(adapterRatings, 0, laptopRating)
	fmt.Println(numChains)
}

func parseNumbers(input []string) []int {
	numbers := make([]int, len(input))
	for i, inputText := range input {
		number, err := strconv.Atoi(inputText)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = number
	}
	return numbers
}

func findChain(remainingAdapters, usedAdapters []int, currentOutput int) ([]int, error) {
	options := findOptions(remainingAdapters, currentOutput)
	sort.Ints(options)
	if len(options) == 0 {
		if len(remainingAdapters) != 0 {
			return []int{}, common.SolutionNotFoundError("No solution here")
		} else {
			return []int{}, nil
		}
	}

	for _, option := range options {
		remaining := removeElement(remainingAdapters, option)
		used, err := findChain(remaining, append(usedAdapters, option), option)
		if err != nil {
			continue
		}
		return append(used, option), nil
	}

	return []int{}, common.SolutionNotFoundError("No solution found")
}

func countArrangements(remainingAdapters []int, currentOutput, targetRating int) (int, error) {
	if targetRating-currentOutput <= 3 && targetRating-currentOutput > 0 {
		return 1, nil
	}
	options := findOptions(remainingAdapters, currentOutput)
	sort.Ints(options)
	if len(options) == 0 || currentOutput > targetRating {
		return 0, common.SolutionNotFoundError("Reached end of this path")
	}
	numChains := 1
	for _, option := range options {
		numberFound, err := countArrangements(removeElement(remainingAdapters, option), option, targetRating)
		if err == nil {
			numChains += numberFound
		}
	}

	return numChains, nil
}

func removeElement(slice []int, element int) []int {
	i := indexOf(slice, element)
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	newSlice[i] = newSlice[len(newSlice)-1]
	newSlice = newSlice[:len(newSlice)-1]
	return newSlice
}

func indexOf(slice []int, element int) int {
	for i, x := range slice {
		if x == element {
			return i
		}
	}
	return -1
}

func findOptions(adapters []int, output int) []int {
	options := make([]int, 0)
	for _, adapter := range adapters {
		delta := adapter - output
		if delta <= 3 && delta > 0 {
			options = append(options, adapter)
		}
	}
	return options
}

func countDifferences(adapterChain []int) (int, int) {
	ones := 0
	threes := 0
	for i := 0; i < len(adapterChain)-1; i++ {
		diff := adapterChain[i] - adapterChain[i+1]
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
	}
	return ones, threes
}
