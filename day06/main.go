package main

import (
	"fmt"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

func main() {
	fileContents := common.ReadFileContents("input.txt")
	responsesByGroup := strings.Split(fileContents, "\n\n")
	fmt.Println("Part 1:", part1(responsesByGroup))
	fmt.Println("Part 2:", part2(responsesByGroup))
}

func part1(responsesByGroup []string) int {
	anyoneAnsweredYes := 0
	for _, group := range responsesByGroup {
		tallies := tallyResponsesForGroup(group)
		anyoneAnsweredYes += len(tallies)
	}
	return anyoneAnsweredYes
}

func part2(responsesByGroup []string) int {
	everyoneAnsweredYes := 0
	for _, responses := range responsesByGroup {
		tallies := tallyResponsesForGroup(responses)
		groupSize := len(strings.Split(responses, "\n"))
		for _, numResponses := range tallies {
			if numResponses == groupSize {
				everyoneAnsweredYes++
			}
		}
	}
	return everyoneAnsweredYes
}

func tallyResponsesForGroup(groupResponses string) map[string]int {
	responsesTallies := make(map[string]int)

	responsesByPerson := strings.Split(groupResponses, "\n")
	for _, personResponses := range responsesByPerson {
		for _, questionAnswered := range personResponses {
			questionLetter := string(questionAnswered)
			responsesTallies[questionLetter] = responsesTallies[questionLetter] + 1
		}
	}

	return responsesTallies
}
