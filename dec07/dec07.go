package main

import (
	"flag"
	"fmt"
	"strconv"

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
	total := 0
	for _, line := range lines {
		nrs := u.SplitToInts(line)
		target := nrs[0]
		nrs = nrs[1:]
		curr := nrs[0]
		if check(target, curr, nrs[1:]) {
			total += target
		}
	}
	return total
}

func check(target, prev int, nrs []int) bool {
	if len(nrs) == 1 {
		if prev+nrs[0] == target {
			return true
		}
		if prev*nrs[0] == target {
			return true
		}
		return false
	}
	new := prev + nrs[0]
	if new <= target && new > 0 { // just check sign for overflow
		if check(target, new, nrs[1:]) {
			return true
		}
	}
	new = prev * nrs[0]
	if new <= target && new > 0 { // just check sign for overflow
		if check(target, new, nrs[1:]) {
			return true
		}
	}
	return false
}

func check2(target, prev int, nrs []int) bool {
	// fmt.Println("check2", target, prev, nrs)
	if len(nrs) == 1 {
		nr := nrs[0]
		if prev+nr == target {
			return true
		}
		if prev*nr == target {
			return true
		}
		concat, err := combine(prev, nr)
		if err == nil && concat == target {
			return true
		}
		return false
	}
	new := prev + nrs[0]
	if new <= target && new > 0 { // just check sign for overflow
		if check2(target, new, nrs[1:]) {
			return true
		}
	}
	new = prev * nrs[0]
	if new <= target && new > 0 { // just check sign for overflow
		if check2(target, new, nrs[1:]) {
			return true
		}
	}
	var err error
	new, err = combine(prev, nrs[0])
	if err == nil && new <= target && new > 0 { // just check sign for overflow
		if check2(target, new, nrs[1:]) {
			return true
		}
	}
	return false
}

func task2(lines []string) int {
	total := 0
	for _, line := range lines {
		nrs := u.SplitToInts(line)
		target := nrs[0]
		nrs = nrs[1:]
		curr := nrs[0]
		if check2(target, curr, nrs[1:]) {
			total += target
		}
	}
	return total
}

func combine(a, b int) (int, error) {
	joined := fmt.Sprintf("%d%d", a, b)
	return strconv.Atoi(joined)
}
