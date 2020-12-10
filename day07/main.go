package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

type Rule struct {
	parentBag       string
	childQuantities []ChildQuantity
}

type ChildQuantity struct {
	child    string
	quantity int
}

type BagGraph struct {
	bagIds          map[string]int
	bagNames        map[int]string
	adjacencyMatrix [][]int
}

func (g *BagGraph) addEdge(from, to string, weight int) {
	fromId := g.bagIds[from]
	toId := g.bagIds[to]
	g.adjacencyMatrix[toId][fromId] = weight
}

func newGraph(bags []string) BagGraph {
	numNodes := len(bags)
	adjacencyMatrix := make([][]int, numNodes)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, numNodes)
	}

	bagIds := make(map[string]int, numNodes)
	bagNames := make(map[int]string, numNodes)
	for i, bag := range bags {
		bagIds[bag] = i
		bagNames[i] = bag
	}

	return BagGraph{
		bagIds:          bagIds,
		bagNames:        bagNames,
		adjacencyMatrix: adjacencyMatrix,
	}
}

func main() {
	inputText := common.ReadAndSanitiseRows("input.txt")
	rules := make([]Rule, len(inputText))

	bagNameSet := make(map[string]int, len(inputText))
	for i, ruleText := range inputText {
		rule := parseRule(ruleText)
		rules[i] = rule
		bagNameSet[rule.parentBag] = 1
		for _, childRule := range rule.childQuantities {
			bagNameSet[childRule.child] = childRule.quantity
		}
	}

	bagNames := make([]string, len(bagNameSet))
	i := 0
	for bagName := range bagNameSet {
		bagNames[i] = bagName
		i++
	}

	graph := newGraph(bagNames)
	for _, rule := range rules {
		for _, childQuantity := range rule.childQuantities {
			graph.addEdge(rule.parentBag, childQuantity.child, childQuantity.quantity)
		}
	}

	fmt.Println(graph.bagIds)
	for _, row := range graph.adjacencyMatrix {
		fmt.Println(row)
	}
}

func parseRule(ruleText string) Rule {
	ruleParts := strings.Split(ruleText, " bags contain ")
	parentBag := ruleParts[0]

	childBagsText := ruleParts[1]
	if childBagsText == "no other bags." {
		return Rule{
			parentBag:       parentBag,
			childQuantities: make([]ChildQuantity, 0),
		}
	}

	childQuantityParts := strings.Split(childBagsText, ", ")
	childQuantities := make([]ChildQuantity, len(childQuantityParts))
	for i, childQuantityText := range childQuantityParts {
		childQuantities[i] = parseChildRule(childQuantityText)
	}

	return Rule{
		parentBag:       parentBag,
		childQuantities: childQuantities,
	}
}

var childRegexp = regexp.MustCompile(`([0-9]) ([a-z ]*) bags?`)

func parseChildRule(childBagText string) ChildQuantity {
	matches := childRegexp.FindStringSubmatch(childBagText)
	quantity, _ := strconv.Atoi(matches[1])
	return ChildQuantity{
		child:    matches[2],
		quantity: quantity,
	}
}
