package main

import (
	"fmt"
	"log"
	"strconv"

	"uk.co.lewis-od.aoc2020/common"
)

func main() {
	input := common.ReadAndSanitiseRows("input.txt")
	numbers := parseNumbers(input)

	fmt.Println("Part 1:")
	answer, err := part1(numbers, 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)

	fmt.Println("Part 2:")
	answer, err = part2(numbers, answer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
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

func part1(numbers []int, preambleLength int) (int, error) {
	for i := preambleLength; i < (len(numbers) - preambleLength); i++ {
		preamble := numbers[i-preambleLength : i]
		currentNumber := numbers[i]

		if !isSumOfNumbersFrom(preamble, currentNumber) {
			return currentNumber, nil
		}
	}
	return 0, common.SolutionNotFoundError("Didn't find answer")
}

func isSumOfNumbersFrom(preamble []int, target int) bool {
	for j, a := range preamble {
		for _, b := range preamble[j:] {
			if a+b == target {
				return true
			}
		}
	}
	return false
}

func part2(numbers []int, target int) (int, error) {
	for start := range numbers {
		end := start + 1

		rangeSum := sum(numbers[start:end])
		for rangeSum <= target {
			rangeSum = sum(numbers[start:end])
			if rangeSum == target {
				smallest := min(numbers[start:end])
				largest := max(numbers[start:end])
				return smallest + largest, nil
			}
			end++
		}
	}
	return 0, common.SolutionNotFoundError("No solution found")
}

func sum(numbers []int) int {
	acc := 0
	for _, number := range numbers {
		acc += number
	}
	return acc
}

func min(numbers []int) int {
	smallest := numbers[0]
	for _, number := range numbers {
		if number < smallest {
			smallest = number
		}
	}
	return smallest
}

func max(numbers []int) int {
	biggest := numbers[0]
	for _, number := range numbers {
		if number > biggest {
			biggest = number
		}
	}
	return biggest
}
