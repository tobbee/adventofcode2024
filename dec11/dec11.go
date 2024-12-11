package main

import (
	"flag"
	"fmt"

	u "github.com/tobbee/adventofcode2024/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines, 25))
	} else {
		fmt.Println("task2: ", task2(lines, 75))
	}
}

func task1(lines []string, nrIterations int) int {
	line := lines[0]
	stones := u.SplitToInts((line))
	for i := 0; i < nrIterations; i++ {
		stones = blink(stones)
	}

	return len(stones)
}

func blink(stones []int) []int {
	new := make([]int, 0, len(stones))
	for _, stone := range stones {
		n := blinkOne(stone)
		new = append(new, n...)
	}
	return new
}

func blinkOne(stone int) []int {
	new := make([]int, 0, 1)
	switch {
	case stone == 0:
		new = append(new, 1)
	case evenNrDigits(stone):
		a, b := splitNr(stone)
		new = append(new, a, b)
	default:
		new = append(new, stone*2024)
	}
	return new
}

func evenNrDigits(nr int) bool {
	return len(fmt.Sprintf("%d", nr))%2 == 0
}

func splitNr(nr int) (int, int) {
	str := fmt.Sprintf("%d", nr)
	a, b := str[:len(str)/2], str[len(str)/2:]
	return u.Atoi(a), u.Atoi(b)
}

func task2(lines []string, nrIterations int) int {
	line := lines[0]
	stones := u.SplitToInts((line))
	bins := make(map[int]int)
	for _, stone := range stones {
		bins[stone]++
	}
	for i := 0; i < nrIterations; i++ {
		newBin := make(map[int]int)
		for stone, nr := range bins {
			new := blinkOne(stone)
			for _, n := range new {
				newBin[n] += nr
			}
		}
		bins = newBin
		nrStones := 0
		for _, nr := range bins {
			nrStones += nr
		}
		fmt.Println(i, len(bins), nrStones)
	}
	nrStones := 0
	for _, nr := range bins {
		nrStones += nr
	}
	return nrStones
}
