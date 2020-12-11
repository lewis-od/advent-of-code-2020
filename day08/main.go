package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

type GamesConsole struct {
	instructions       []Instruction
	accumulator        int
	instructionCounter int
	counterHistory     []int
}

func (gc *GamesConsole) execute() bool {
	if gc.instructionCounter == len(gc.instructions) {
		return true
	} else if common.ArrayContains(gc.counterHistory, gc.instructionCounter) {
		return false
	}

	currentInstruction := gc.instructions[gc.instructionCounter]
	op := currentInstruction.operation

	counterDelta := 1
	if op == "jmp" {
		counterDelta = currentInstruction.argument
	} else if op == "acc" {
		gc.accumulator += currentInstruction.argument
	} else if op != "nop" {
		log.Fatal("Encountered unknown operation", op)
	}

	gc.counterHistory = append(gc.counterHistory, gc.instructionCounter)
	gc.instructionCounter += counterDelta
	return gc.execute()
}

type Instruction struct {
	operation string
	argument  int
}

func main() {
	inputRows := common.ReadAndSanitiseRows("input.txt")

	instructions := make([]Instruction, len(inputRows))
	for i, inputRow := range inputRows {
		instructions[i] = parseInstruction(inputRow)
	}

	fmt.Println("Part 1:")
	fmt.Println("acc =", part1(instructions))

	fmt.Println()
	fmt.Println("Part 2:")
	accumulator, err := part2(instructions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("acc =", accumulator)
}

func parseInstruction(text string) Instruction {
	parts := strings.Split(text, " ")
	argument, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Error converting", parts[1], "to int")
	}
	return Instruction{
		operation: parts[0],
		argument:  argument,
	}
}

func part1(instructions []Instruction) int {
	accumulator, _ := runInstructions(instructions)
	return accumulator
}

func part2(instructions []Instruction) (int, error) {
	for index, instruction := range instructions {
		if instruction.operation == "jmp" {
			flippedInstruction := instruction
			flippedInstruction.operation = "nop"

			modifiedInstructions := make([]Instruction, len(instructions))
			copy(modifiedInstructions, instructions)
			modifiedInstructions[index] = flippedInstruction

			accumulator, didTerminate := runInstructions(modifiedInstructions)
			if didTerminate {
				return accumulator, nil
			}
		}
	}

	return 0, common.SolutionNotFoundError("Unable to find solution")
}

func runInstructions(instructions []Instruction) (int, bool) {
	console := GamesConsole{
		instructions:       instructions,
		accumulator:        0,
		instructionCounter: 0,
		counterHistory:     make([]int, 0),
	}
	didTerminate := console.execute()
	return console.accumulator, didTerminate
}
