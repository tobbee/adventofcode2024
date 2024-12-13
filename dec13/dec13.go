package main

import (
	"flag"
	"fmt"
	"math"
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

type unit struct {
	X, Y int
}

type machine struct {
	a     unit
	b     unit
	prize unit
}

func task1(lines []string) int {
	machines := parseInput(lines, 0)
	total := 0
	for _, m := range machines {
		fmt.Println(m)
		cost, ok := findCheapestSolution(m)
		fmt.Println(cost, ok)
		if ok {
			total += cost
		}
	}
	return total
}

func findCheapestSolution(m machine) (int, bool) {
	// Brute force a maximum of 100 times on each button
	lowestCost := math.MaxInt
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			x := m.a.X*i + m.b.X*j
			y := m.a.Y*i + m.b.Y*j
			if x == m.prize.X && y == m.prize.Y {
				cost := 3*i + j
				if cost < lowestCost {
					lowestCost = cost
				}
			}
		}
	}
	if lowestCost == math.MaxInt {
		return 0, false
	}
	return lowestCost, true
}

func task2(lines []string) int {
	machines := parseInput(lines, 10000000000000)
	total := 0
	for _, m := range machines {
		fmt.Println(m)
		cost, ok := findCheapestSolution2(m)
		fmt.Println(cost, ok)
		if ok {
			total += cost
		}
	}
	return total
}

func findCheapestSolution2(m machine) (int, bool) {
	// Numerator and denominator of the linear equation
	// solution for press on A button pA = (pANum/pADen)
	pANum := m.prize.X*m.b.Y - m.prize.Y*m.b.X
	pADen := m.a.X*m.b.Y - m.a.Y*m.b.X
	if pADen == 0 {
		return 0, false
	}
	gcd := u.GCD(pANum, pADen)
	pANum /= gcd
	pADen /= gcd
	if pADen != 1 {
		return 0, false
	}
	pA := pANum // An integer solution for pA

	pBNum := m.prize.Y - m.a.Y*pA
	pBDen := m.b.Y
	gcd = u.GCD(pBNum, pBDen)
	pBNum /= gcd
	pBDen /= gcd
	if pBDen != 1 {
		return 0, false
	}
	pB := pBNum // An integer solution for pB

	cost := 3*pA + pB
	return cost, true
}

func parseInput(lines []string, prizeOffset int) []machine {
	var machines []machine
	m := machine{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		switch {
		case strings.HasSuffix(parts[0], "A"):
			steps := u.SplitToInts(parts[1])
			m.a.X = steps[0]
			m.a.Y = steps[1]
		case strings.HasSuffix(parts[0], "B"):
			steps := u.SplitToInts(parts[1])
			m.b.X = steps[0]
			m.b.Y = steps[1]
		case strings.HasPrefix(parts[0], "Prize"):
			prizePos := u.SplitToInts(parts[1])
			m.prize.X = prizePos[0] + prizeOffset
			m.prize.Y = prizePos[1] + prizeOffset
			machines = append(machines, m)
		}
	}
	return machines
}
