package main

import (
	"flag"
	"fmt"
	"os"

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

func task1(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	// fmt.Printf("%s\n\n", g)
	pos, dir := findStart(g)
	// fmt.Println("pos", pos, "dir", dir)
	g.Grid[pos.r][pos.c] = dirChar(dir)
	nrSteps := 0
	pathPos := u.Set[Pos]{}
	pathPos.Add(pos) // Add starting position
	for {
		r, c := pos.r+dir.r, pos.c+dir.c
		if !g.InBounds(r, c) {
			break
		}
		if g.At(r, c) == "#" {
			dir = turnRight(dir)
			continue
		}
		g.Grid[r][c] = dirChar(dir)
		pos = Pos{r, c}
		pathPos.Add(pos)
		nrSteps++
	}
	fmt.Println("nrSteps", nrSteps)
	fmt.Println("nrPos", len(pathPos))
	return len(pathPos) // One extra for the starting position
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	// fmt.Printf("%s\n\n", g)
	pos, dir := findStart(g)
	nrSteps := 0
	pathPos := u.Set[Pos]{}
	for {
		r, c := pos.r+dir.r, pos.c+dir.c
		if !g.InBounds(r, c) {
			break
		}
		if g.At(r, c) == "#" {
			dir = turnRight(dir)
			continue
		}
		g.Grid[r][c] = dirChar(dir)
		pos = Pos{r, c}
		pathPos.Add(pos)
		nrSteps++
	}
	nrGoodBlocks := 0
	for oPos := range pathPos {
		g := u.CreateCharGridFromLines(lines)
		pos, dir := findStart(g)
		g.Grid[oPos.r][oPos.c] = "O" // Insert obstacle
		nrSteps := 0
		visited := u.Set[state]{}
		visited.Add(state{pos, dir})
		for {
			r, c := pos.r+dir.r, pos.c+dir.c
			if !g.InBounds(r, c) {
				break
			}
			if g.At(r, c) == "#" || g.At(r, c) == "O" {
				dir = turnRight(dir)
				continue
			}
			pos = Pos{r, c}
			g.Grid[r][c] = dirChar(dir)
			if visited.Contains(state{pos, dir}) {
				nrGoodBlocks++
				break
			}
			visited.Add(state{pos, dir})
			nrSteps++
			if nrSteps > 10000 {
				fmt.Println("nrSteps", nrSteps)
				os.Exit(1)
			}
		}
	}
	return nrGoodBlocks
}

type state struct {
	pos, dir Pos
}

func dirChar(dir Pos) string {
	switch dir {
	case Pos{1, 0}:
		return "v"
	case Pos{0, 1}:
		return ">"
	case Pos{-1, 0}:
		return "^"
	case Pos{0, -1}:
		return "<"
	}
	return "?"
}

func findStart(g u.CharGrid) (pos, dir Pos) {
	for _, ch := range []string{"^", "v", "<", ">"} {
		r, c := g.Find(ch)
		if r != -1 {
			pos = Pos{r, c}
			switch ch {
			case ">":
				dir = Pos{0, 1}
			case "<":
				dir = Pos{0, -1}
			case "^":
				dir = Pos{-1, 0}
			case "v":
				dir = Pos{1, 0}
			}
			break
		}
	}
	return pos, dir
}

func turnRight(dir Pos) Pos {
	switch dir {
	case Pos{1, 0}:
		dir = Pos{0, -1}
	case Pos{0, 1}:
		dir = Pos{1, 0}
	case Pos{-1, 0}:
		dir = Pos{0, 1}
	case Pos{0, -1}:
		dir = Pos{-1, 0}
	}
	return dir
}
