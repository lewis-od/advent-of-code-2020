package main

import (
	"fmt"
	"strconv"

	c "uk.co.lewis-od.aoc2020/common"
)

type result struct {
	numbers []int
	answer  int
}

func (r result) Print() {
	fmt.Printf("sum(%v) = 2020\n", r.numbers)
	fmt.Printf("prod(%v) = %d\n", r.numbers, r.answer)
}

func main() {
	numbers := parseNumberList("input.txt")

	fmt.Println("Part 1")
	result, err := part1(numbers)
	if err != nil {
		fmt.Println(err)
	} else {
		result.Print()
	}
	fmt.Println()

	fmt.Println("Part 2")
	result, err = part2(numbers)
	if err != nil {
		fmt.Println(err)
	} else {
		result.Print()
	}
}

func parseNumberList(filePath string) []int {
	fileRows := c.ReadAndSanitise(filePath)
	numbers := make([]int, len(fileRows))
	for index, numString := range fileRows {
		number, err := strconv.Atoi(numString)
		if err != nil {
			continue
		}
		numbers[index] = number
	}
	return numbers
}

func part1(numbers []int) (result, error) {
	for index, a := range numbers {
		for _, b := range numbers[index:] {
			if a+b == 2020 {
				return result{
					numbers: []int{a, b},
					answer:  a * b,
				}, nil
			}
		}
	}
	return result{}, c.SolutionNotFoundError("Solution not found")
}

func part2(numbers []int) (result, error) {
	for indexA, a := range numbers {
		for indexB, b := range numbers[indexA:] {
			for _, c := range numbers[indexB:] {
				if a+b+c == 2020 {
					return result{
						numbers: []int{a, b, c},
						answer:  a * b * c,
					}, nil
				}
			}
		}
	}
	return result{}, c.SolutionNotFoundError("Solution not found")
}
