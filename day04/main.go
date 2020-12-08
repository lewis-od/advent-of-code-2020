package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

type Passport map[string]string

type FieldValidator func(fieldValue string) bool

func (p *Passport) validate(requiredFields []string, validators map[string]FieldValidator) bool {
	for _, field := range requiredFields {
		fieldValue := (*p)[field]
		if fieldValue == "" {
			return false
		}
		validator := validators[field]
		if !validator(fieldValue) {
			return false
		}
	}
	return true
}

func main() {
	fileContents := common.ReadFileContents("input.txt")
	passports := parsePassports(fileContents)

	fmt.Println("Part 1")
	validPassports := part1(passports)
	fmt.Println(validPassports, "valid passports")

	fmt.Println("Part 2")
	validPassports = part2(passports)
	fmt.Println(validPassports, "valid passports")
}

func parsePassports(fileContents string) []Passport {
	passportRows := strings.Split(fileContents, "\n\n")

	passports := make([]Passport, len(passportRows))
	for i, passportText := range passportRows {
		passports[i] = parsePassport(passportText)
	}

	return passports
}

func parsePassport(passportText string) Passport {
	singleRow := strings.ReplaceAll(passportText, "\n", " ")
	trimmed := strings.TrimSpace(singleRow)
	tokens := strings.Split(trimmed, " ")
	passport := make(map[string]string)
	for _, token := range tokens {
		parts := strings.Split(token, ":")
		passport[parts[0]] = parts[1]
	}
	return Passport(passport)
}

func part1(passports []Passport) int {
	validators := map[string]FieldValidator{
		"byr": isPresentValidator,
		"iyr": isPresentValidator,
		"eyr": isPresentValidator,
		"hgt": isPresentValidator,
		"hcl": isPresentValidator,
		"ecl": isPresentValidator,
		"pid": isPresentValidator,
	}
	return countValidPassports(passports, validators)
}

func part2(passports []Passport) int {
	validators := map[string]FieldValidator{
		"byr": rangeValidator(1920, 2002),
		"iyr": rangeValidator(2010, 2020),
		"eyr": rangeValidator(2020, 2030),
		"hgt": heightValidator,
		"hcl": regexValidator(`^#[0-9a-f]{6}$`),
		"ecl": regexValidator(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
		"pid": regexValidator(`^[0-9]{9}$`),
	}
	return countValidPassports(passports, validators)
}

func countValidPassports(passports []Passport, validators map[string]FieldValidator) int {
	validPassports := 0
	for _, passport := range passports {
		if passport.validate(requiredFields, validators) {
			validPassports++
		}
	}
	return validPassports
}

func isPresentValidator(fieldValue string) bool {
	return fieldValue != ""
}

func rangeValidator(min, max int) FieldValidator {
	return func(fieldValue string) bool {
		intValue, err := strconv.Atoi(fieldValue)
		if err != nil {
			return false
		}
		return intValue >= min && intValue <= max
	}
}

func heightValidator(fieldValue string) bool {
	cmValidator := rangeValidator(150, 193)
	inchValidator := rangeValidator(59, 76)

	if strings.Contains(fieldValue, "cm") {
		return cmValidator(strings.ReplaceAll(fieldValue, "cm", ""))
	} else if strings.Contains(fieldValue, "in") {
		return inchValidator(strings.ReplaceAll(fieldValue, "in", ""))
	}

	return false
}

func regexValidator(regex string) FieldValidator {
	compiledRegex := regexp.MustCompile(regex)
	return func(fieldValue string) bool {
		return compiledRegex.MatchString(fieldValue)
	}
}
