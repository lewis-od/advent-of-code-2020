package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

type PasswordPolicy struct {
	firstNumber  int
	secondNumber int
	character    string
}

type PasswordValidator func(firstNumber, secondNumber int, character, password string) bool

func (p PasswordPolicy) validate(password string, validator PasswordValidator) bool {
	return validator(p.firstNumber, p.secondNumber, p.character, password)
}

type DatabaseEntry struct {
	policy   PasswordPolicy
	password string
}

func (entry DatabaseEntry) validate(validator PasswordValidator) bool {
	return entry.policy.validate(entry.password, validator)
}

func main() {
	fileRows := common.ReadAndSanitiseRows("input.txt")

	oldValidPasswords := 0
	newValidPasswords := 0
	for _, row := range fileRows {
		entry := parseEntry(row)
		if entry.validate(oldValidator) {
			oldValidPasswords++
		}
		if entry.validate(newValidator) {
			newValidPasswords++
		}
	}

	fmt.Println("Part 1:")
	fmt.Println(oldValidPasswords, "valid passwords")
	fmt.Println("Part 2:")
	fmt.Println(newValidPasswords, "valid passwords")
}

func parseEntry(entryText string) DatabaseEntry {
	parts := strings.SplitN(entryText, " ", 3)
	policy := parsePolicy(parts[0], parts[1])
	return DatabaseEntry{
		policy:   policy,
		password: parts[2],
	}
}

func parsePolicy(rangeText, characterText string) PasswordPolicy {
	numbers := strings.SplitN(rangeText, "-", 2)
	min, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Fatal("Error parsing integer: ", numbers[0])
	}
	max, err := strconv.Atoi(numbers[1])
	if err != nil {
		log.Fatal("Error parsing integer: ", numbers[1])
	}
	character := string(characterText[0])
	return PasswordPolicy{
		firstNumber:  min,
		secondNumber: max,
		character:    character,
	}
}

func oldValidator(firstNumber, secondNumber int, character, password string) bool {
	occurrences := 0
	for _, passwordCharacter := range password {
		if string(passwordCharacter) == character {
			occurrences++
		}
	}
	return (occurrences >= firstNumber) && (occurrences <= secondNumber)
}

func newValidator(firstNumber, secondNumber int, character, password string) bool {
	firstPosition := string(password[firstNumber-1])
	firstPositionValid := firstPosition == character

	secondPosition := string(password[secondNumber-1])
	secondPositionValid := secondPosition == character

	return (firstPositionValid && !secondPositionValid) || (!firstPositionValid && secondPositionValid)
}
