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
walk:
	for {
		r, c := pos.r+dir.r, pos.c+dir.c
		if !g.InBounds(r, c) {
			break walk
		}
		if g.At(r, c) == "#" {
			// fmt.Println(g)
			// fmt.Println()
			dir = turnRight(dir)

			continue walk
		}
		g.Grid[r][c] = dirChar(dir)
		pos = Pos{r, c}
		nrSteps++
	}
	count := 0
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			switch g.Grid[r][c] {
			case "^", "v", "<", ">":
				count++
			}
		}
	}
	fmt.Println("nrSteps", nrSteps)
	// fmt.Println(g)
	return count
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	// fmt.Printf("%s\n\n", g)
	pos, dir := findStart(g)
	startPos := pos
	// fmt.Println("pos", pos, "dir", dir)
	nrLoops := 0
	nrSteps := 0
	goodPos := make(map[Pos]bool)
	for {
		r, c := pos.r+dir.r, pos.c+dir.c
		if !g.InBounds(r, c) {
			break
		}
		if g.At(r, c) == "#" {
			// fmt.Println(g)
			// fmt.Println()
			dir = turnRight(dir)
			continue
		}
		g2 := g.Copy()
		if r == startPos.r && c == startPos.c {
			fmt.Println("back at start pos")
			os.Exit(1)
			continue
		}
		if g.Grid[r][c] == "#" {
			fmt.Println("bad state found")
			os.Exit(2)
		}
		g2.Grid[r][c] = "O" // Insert obstacle
		if checkLoop(g2, pos, dir) {
			nrLoops++
			goodPos[Pos{r, c}] = true
			fmt.Println("nrLoops", nrLoops, "nrSteps", nrSteps, "nrObstacles", len(goodPos))
		}
		g.Grid[r][c] = dirChar(dir)
		pos = Pos{r, c}
		nrSteps++
	}
	return len(goodPos)
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

func checkLoop(g u.CharGrid, pos, dir Pos) bool {
	nrSteps := 0
	visited := make(map[state]bool)
	for {
		if _, ok := visited[state{pos, dir}]; ok {
			return true
		}
		r, c := pos.r+dir.r, pos.c+dir.c
		if !g.InBounds(r, c) {
			return false
		}
		visited[state{pos, dir}] = true
		if g.At(r, c) == "O" || g.At(r, c) == "#" {
			// fmt.Println(g)
			// fmt.Println()
			dir = turnRight(dir)
			continue
		}
		pos = Pos{r, c}
		if g.Grid[pos.r][pos.c] == "#" || g.Grid[pos.r][pos.c] == "O" {
			fmt.Println("bad state found")
		}
		g.Grid[r][c] = dirChar(dir)
		//fmt.Printf("%s\n", g)
		nrSteps++
	}
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
