package main

import (
	"fmt"

	"uk.co.lewis-od.aoc2020/common"
)

type SeatingArea struct {
	Data          [][]byte
	width, height int
}

func FromInput(input []string) SeatingArea {
	width := len(input[0])
	height := len(input)

	data := make([][]byte, height)
	for i, row := range input {
		data[i] = []byte(row)
	}

	return SeatingArea{
		Data:   data,
		width:  width,
		height: height,
	}
}

func (sa *SeatingArea) Get(x, y int) string {
	if sa.isOutOfBounds(x, y) {
		return "."
	}
	return string(sa.Data[y][x])
}

func (sa *SeatingArea) isOutOfBounds(x, y int) bool {
	return (x < 0) || (y < 0) || (x >= sa.width) || (y >= sa.height)
}

func (sa *SeatingArea) Tick() bool {
	newData := make([][]byte, len(sa.Data))
	didChange := false
	for y, row := range sa.Data {
		newData[y] = make([]byte, len(row))
		copy(newData[y], row)

		for x, col := range row {
			seat := string(col)
			if seat == "." {
				continue
			}

			numOccupied := sa.countAdjacent(x, y, "#")
			if seat == "L" {
				if numOccupied == 0 {
					newData[y][x] = '#'
					didChange = true
				}
			} else if seat == "#" {
				if numOccupied >= 4 {
					newData[y][x] = 'L'
					didChange = true
				}
			}
		}
	}
	sa.Data = newData

	return didChange
}

func (sa *SeatingArea) countAdjacent(x, y int, targets ...string) int {
	numTargets := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == dy && dx == 0 {
				continue
			}
			for _, target := range targets {
				if sa.Get(x+dx, y+dy) == target {
					numTargets++
					continue
				}
			}
		}
	}
	return numTargets
}

func (sa *SeatingArea) NumOccupied() int {
	numOccupied := 0
	for _, row := range sa.Data {
		for _, character := range row {
			seat := string(character)
			if seat == "#" {
				numOccupied++
			}
		}
	}
	return numOccupied
}

func (sa *SeatingArea) Print() {
	for y := 0; y < sa.height; y++ {
		for x := 0; x < sa.width; x++ {
			fmt.Print(sa.Get(x, y))
		}
		fmt.Println()
	}
}

func main() {
	input := common.ReadAndSanitiseRows("input.txt")

	seatingArea := FromInput(input)
	for seatingArea.Tick() {
	}

	fmt.Println("Part 1:")
	fmt.Println(seatingArea.NumOccupied())

	// for round := 1; round <= 7; round++ {
	// 	fmt.Println("Round", round)
	// 	seatingArea.Tick()
	// 	seatingArea.Print()
	// 	fmt.Println()
	// }
}
