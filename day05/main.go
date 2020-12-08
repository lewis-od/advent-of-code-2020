package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

const (
	UPPER = "U"
	LOWER = "L"
)

func main() {
	boardingPassCodes := common.ReadAndSanitiseRows("input.txt")
	fmt.Println("Max seat ID:", part1(boardingPassCodes))
	if mySeat, err := part2(boardingPassCodes); err == nil {
		fmt.Println("My seat ID:", mySeat)
	} else {
		fmt.Println(err)
	}
}

func part1(boardingPassCodes []string) int {
	maxSeatID := 0
	for _, code := range boardingPassCodes {
		seatID := calcSeatID(code)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func part2(boardingPassCodes []string) (int, error) {
	presentSeatIDs := make([]int, len(boardingPassCodes))
	for i, seatCode := range boardingPassCodes {
		presentSeatIDs[i] = calcSeatID(seatCode)
	}
	sort.Ints(presentSeatIDs)

	for i := 0; i < len(presentSeatIDs)-1; i++ {
		current := presentSeatIDs[i]
		next := presentSeatIDs[i+1]
		if next == current+2 {
			return current + 1, nil
		}
	}

	return 0, common.SolutionNotFoundError("Can't find my seat :(")
}

func standardiseCode(code string) string {
	standardised := strings.ReplaceAll(code, "B", UPPER)
	standardised = strings.ReplaceAll(standardised, "F", LOWER)
	standardised = strings.ReplaceAll(standardised, "R", UPPER)
	standardised = strings.ReplaceAll(standardised, "L", LOWER)
	return standardised
}

func calcSeatID(code string) int {
	standardised := standardiseCode(code)

	rowCode := standardised[:7]
	rowNum := calcRowNumber(rowCode)

	columnCode := standardised[7:]
	columnNum := calcColumnNumber(columnCode)

	return 8*rowNum + columnNum
}

func calcRowNumber(code string) int {
	return calcSeatPosition(127, code)
}

func calcColumnNumber(code string) int {
	return calcSeatPosition(7, code)
}

func calcSeatPosition(max int, code string) int {
	minPosition := 0
	maxPosition := max
	for _, character := range code {
		positionIndicator := string(character)
		step := (maxPosition - minPosition + 1) / 2
		if positionIndicator == UPPER {
			minPosition += step
		} else if positionIndicator == LOWER {
			maxPosition -= step
		} else {
			fmt.Println("Invalid character", string(positionIndicator))
		}
	}
	if minPosition != maxPosition {
		log.Fatal("Ambiguous row")
	}
	return minPosition
}
