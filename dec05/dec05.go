package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	u "github.com/tobbee/adventofcode2024/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string) int {
	rules, pages := parseInput(lines)
	middleSum := 0
	for _, page := range pages {
		if isValid(rules, page) {
			middle := page[len(page)/2]
			middleSum += middle
		}
	}

	return middleSum
}

func isValid(rules map[int][]int, page []int) bool {
	for i := 0; i < len(page); i++ {
		curr := page[i]
		for j := 0; j < i; j++ {
			prev := page[j]
			rule, ok := rules[curr]
			if !ok {
				continue
			}
			for _, p := range rule {
				if p == prev {
					// Broken rule, something that should come after comes before
					return false
				}
			}
		}
	}
	return true
}

func parseInput(lines []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	pages := make([][]int, 0)
	for _, line := range lines {
		ruleParts := strings.Split(line, "|")
		if len(ruleParts) == 2 {
			first, err := strconv.Atoi(ruleParts[0])
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.Atoi(ruleParts[1])
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := rules[first]; !ok {
				rules[first] = make([]int, 0, 1)
			}
			rules[first] = append(rules[first], second)
			continue
		}
		pageParts := strings.Split(line, ",")
		if len(pageParts) > 1 {
			page := make([]int, 0, len(pageParts))
			for _, p := range pageParts {
				pInt, err := strconv.Atoi(p)
				if err != nil {
					log.Fatal(err)
				}
				page = append(page, pInt)
			}
			pages = append(pages, page)
		}
	}
	return rules, pages
}

func reorderPage(rules map[int][]int, page []int) {
	i := 0
outerLoop:
	for {
		if i == len(page) {
			return
		}
		curr := page[i]
		for j := 0; j < i; j++ {
			prev := page[j]
			rule, ok := rules[curr]
			if !ok {
				continue
			}
			for _, p := range rule {
				if p == prev {
					// Broken rule, swap elements and check again
					page[i], page[j] = page[j], page[i]
					i = j
					continue outerLoop
				}
			}
		}
		i++
	}
}

func task2(lines []string) int {
	rules, pages := parseInput(lines)
	middleSum := 0
	for _, page := range pages {
		if !isValid(rules, page) {
			reorderPage(rules, page)
			middle := page[len(page)/2]
			middleSum += middle
		}
	}

	return middleSum
}
