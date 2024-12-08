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
		fmt.Println("task1: ", task1(lines))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

type Pos struct {
	r, c int
}

func sub(p1, p2 Pos) Pos {
	return Pos{p1.r - p2.r, p1.c - p2.c}
}

func add(p1, p2 Pos) Pos {
	return Pos{p1.r + p2.r, p1.c + p2.c}
}

func calcAntiNodes(ant1, ant2 Pos) (Pos, Pos) {
	d := sub(ant1, ant2)
	a1 := add(ant1, d)
	a2 := sub(ant2, d)
	return a1, a2
}

func task1(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	antennas := findAntennas(g)
	allAntiNodes := u.Set[Pos]{}
	for a := range antennas {
		aNodes := findAntiNodes(g, antennas[a])
		allAntiNodes.Extend(aNodes)
	}
	return len(allAntiNodes)
}

func findAntennas(g u.CharGrid) map[string][]Pos {
	ants := map[string][]Pos{}
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			char := g.Grid[r][c]
			if char != "." {
				ants[char] = append(ants[char], Pos{r, c})
			}
		}
	}
	return ants
}

func findAntiNodes(g u.CharGrid, ap []Pos) u.Set[Pos] {
	antiNodes := u.Set[Pos]{}
	// Find all pairs of antennas and the corresponding anti-nodes
	for i := 0; i < len(ap); i++ {
		for j := i + 1; j < len(ap); j++ {
			a1, a2 := calcAntiNodes(ap[i], ap[j])
			if g.InBounds(a1.r, a1.c) {
				//g.Grid[a1.r][a1.c] = "X"
				antiNodes.Add(a1)
			}
			if g.InBounds(a2.r, a2.c) {
				//g.Grid[a2.r][a2.c] = "X"
				antiNodes.Add(a2)
			}
			//fmt.Printf("%s\n", g)
		}
	}
	return antiNodes
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	antennas := findAntennas(g)
	nodes := u.Set[Pos]{}
	for a := range antennas {
		aNodes := findAllNodes(g, antennas[a])
		nodes.Extend(aNodes)
	}
	return len(nodes)
}

func findAllNodes(g u.CharGrid, ap []Pos) u.Set[Pos] {
	antiNodes := u.Set[Pos]{}
	for i := 0; i < len(ap); i++ {
		for j := i + 1; j < len(ap); j++ {
			step := sub(ap[i], ap[j])
			// Minimize step
			gcd := u.GCD(step.r, step.c)
			step.r /= gcd
			step.c /= gcd
			pos := ap[i]
			for {
				pos = add(pos, step)
				if !g.InBounds(pos.r, pos.c) {
					break
				}
				antiNodes.Add(pos)
			}
			for {
				pos = sub(pos, step)
				if !g.InBounds(pos.r, pos.c) {
					break
				}
				antiNodes.Add(pos)
			}
		}
	}
	return antiNodes
}
