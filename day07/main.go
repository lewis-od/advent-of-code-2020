package main

import (
	"fmt"
	"regexp"
	"strings"

	"uk.co.lewis-od.aoc2020/common"
)

type Rule struct {
	parentBag    string
	childrenBags []string
}

type Node string

type DirectedGraph struct {
	adjacentNodes map[Node][]Node
}

func (g *DirectedGraph) addNode(node Node, numChildren int) {
	g.adjacentNodes[node] = make([]Node, 0, numChildren)
}

func (g *DirectedGraph) addDependency(from, to Node) {
	g.adjacentNodes[from] = append(g.adjacentNodes[from], to)
}

func main() {
	rules := common.ReadAndSanitiseRows("input.txt")
	bagGraph := DirectedGraph{
		adjacentNodes: make(map[Node][]Node),
	}

	for _, ruleText := range rules {
		rule := parseRule(ruleText)
		bagGraph.addNode(Node(rule.parentBag), len(rule.childrenBags))
		for _, childBag := range rule.childrenBags {
			bagGraph.addDependency(Node(rule.parentBag), Node(childBag))
		}
	}

	for bagColour, canContain := range bagGraph.adjacentNodes {
		fmt.Println(bagColour, ":", canContain)
	}
}

func parseRule(ruleText string) Rule {
	ruleParts := strings.Split(ruleText, " bags contain ")
	parentBag := ruleParts[0]

	childBagsText := ruleParts[1]
	if childBagsText == "no other bags." {
		return Rule{
			parentBag:    parentBag,
			childrenBags: make([]string, 0),
		}
	}

	childBagParts := strings.Split(childBagsText, ", ")
	children := make([]string, len(childBagParts))
	for i, childBagText := range childBagParts {
		children[i] = parseChildBag(childBagText)
	}

	return Rule{
		parentBag:    parentBag,
		childrenBags: children,
	}
}

var childRegexp = regexp.MustCompile(`([0-9]) ([a-z ]*) bags?`)

func parseChildBag(childBagText string) string {
	matches := childRegexp.FindStringSubmatch(childBagText)
	return matches[2]
}
