package main

import (
	"fmt"

	"uk.co.lewis-od.aoc2020/common"
)

const TREE string = "#"

type SkiMap struct {
	data       []string
	frameWidth int
}

func (m SkiMap) get(x, y int) string {
	row := m.data[y]
	return string(row[x%m.frameWidth])
}

func (m SkiMap) traverse(xStep, yStep int) int {
	x, y := 0, 0
	numTrees := 0
	for y < len(m.data) {
		mapLocation := m.get(x, y)
		if mapLocation == TREE {
			numTrees++
		}
		x += xStep
		y += yStep
	}
	return numTrees
}

func main() {
	fileRows := common.ReadAndSanitise("input.txt")
	skiMap := SkiMap{
		data:       fileRows,
		frameWidth: len(fileRows[0]),
	}

	fmt.Println("Part 1:")
	numTrees := skiMap.traverse(3, 1)
	fmt.Printf("%d trees\n", numTrees)

	fmt.Println("Part 2:")
	treeProduct := part2(skiMap)
	fmt.Println(treeProduct)
}

func part2(skiMap SkiMap) int {
	trajectories := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	numTrees := 1
	for _, trajectory := range trajectories {
		numTrees *= skiMap.traverse(trajectory[0], trajectory[1])
	}
	return numTrees
}
