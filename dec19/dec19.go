package main

import (
	"flag"
	"fmt"
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
	patterns, designs := parseInput(lines)
	count := 0
	for _, design := range designs {
		if possible(design, patterns) {
			count++
		}
	}
	return count
}

func possible(design string, patterns []string) bool {
	if len(design) == 0 {
		return true
	}
	{
		for _, p := range patterns {
			if strings.HasPrefix(design, p) {
				poss := possible(design[len(p):], patterns)
				if poss {
					return true
				}
			}
		}
	}
	return false
}

func task2(lines []string) int {
	patterns, designs := parseInput(lines)
	var working []string
	for _, design := range designs {
		if possible(design, patterns) {
			working = append(working, design)
		}
	}
	pat := patternsByLength(patterns)
	count := 0
	cache := make(map[string]int)
	for _, design := range working {
		count += nrWays(design, pat, cache)
	}
	return count
}

func patternsByLength(patterns []string) []u.Set[string] {
	maxLen := 0
	for _, p := range patterns {
		l := len(p)
		if l > maxLen {
			maxLen = l
		}
	}
	pat := make([]u.Set[string], 0)
	for i := 0; i <= maxLen; i++ {
		pat = append(pat, u.CreateSet[string]())
	}
	for _, p := range patterns {
		pat[len(p)].Add(p)
	}
	return pat
}

func nrWays(design string, pat []u.Set[string], cache map[string]int) int {
	if cached, ok := cache[design]; ok {
		return cached
	}
	if len(design) == 0 {
		return 1
	}
	count := 0
	for i := 1; i < len(pat) && i <= len(design); i++ {
		sp := design[:i]
		pp := pat[i]
		if pp.Contains(sp) {
			nr := nrWays(design[i:], pat, cache)
			count += nr
		}
	}
	cache[design] = count
	return count
}

func parseInput(lines []string) (patterns, designs []string) {
	for i, line := range lines {
		switch i {
		case 0:
			patterns = strings.Split(line, ", ")
		case 1:
			continue
		default:
			designs = append(designs, line)
		}
	}
	return patterns, designs
}
