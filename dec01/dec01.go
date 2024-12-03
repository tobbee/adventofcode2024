package main

import (
	"flag"
	"fmt"
	"sort"

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
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		numbers := u.SplitToInts(line)
		left = append(left, numbers[0])
		right = append(right, numbers[1])
	}
	sort.Ints(left)
	sort.Ints(right)
	totDist := 0
	for i := 0; i < len(left); i++ {
		dist := u.Abs(right[i] - left[i])
		totDist += dist
	}
	return totDist
}

func task2(lines []string) int {
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		numbers := u.SplitToInts(line)
		left = append(left, numbers[0])
		right = append(right, numbers[1])
	}
	freq := make(map[int]int)
	for _, r := range right {
		freq[r]++
	}
	totSimilarity := 0
	for i := 0; i < len(left); i++ {
		l := left[i]
		totSimilarity += l * freq[l]
	}
	return totSimilarity
}
