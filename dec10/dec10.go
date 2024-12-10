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

func task1(lines []string) int {
	g := u.CreateDigitGridFromLines(lines)
	starts := findStarts(g)
	totNrPeaks := 0
	for _, start := range starts {
		peaks := u.CreateSet[u.Pos2D]()
		findPeaks(g, start, peaks)
		totNrPeaks += len(peaks)
	}
	return totNrPeaks
}

func task2(lines []string) int {
	g := u.CreateDigitGridFromLines(lines)
	starts := findStarts(g)
	totNrPeaks := 0
	for _, start := range starts {
		nrPeaks := getNrPaths(g, start)
		totNrPeaks += nrPeaks
	}
	return totNrPeaks
}

func findStarts(g u.DigitGrid) []u.Pos2D {
	starts := []u.Pos2D{}
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == 0 {
				starts = append(starts, u.Pos2D{r, c})
			}
		}
	}
	return starts
}

func findPeaks(g u.DigitGrid, start u.Pos2D, peaks u.Set[u.Pos2D]) {
	startLevel := g.Grid[start.Row][start.Col]
	for _, dir := range u.Dirs2D {
		neighbor := start.Add(dir)
		if g.InBounds(neighbor.Row, neighbor.Col) && g.Grid[neighbor.Row][neighbor.Col] == startLevel+1 {
			if g.Grid[neighbor.Row][neighbor.Col] == 9 {
				peaks.Add(neighbor)
				continue
			}
			findPeaks(g, neighbor, peaks)
		}
	}
}

func getNrPaths(g u.DigitGrid, start u.Pos2D) int {
	nrPeaks := 0
	startLevel := g.Grid[start.Row][start.Col]
	for _, dir := range u.Dirs2D {
		neighbor := start.Add(dir)
		if g.InBounds(neighbor.Row, neighbor.Col) && g.Grid[neighbor.Row][neighbor.Col] == startLevel+1 {
			if g.Grid[neighbor.Row][neighbor.Col] == 9 {
				nrPeaks++
				continue
			}
			nrPeaks += getNrPaths(g, neighbor)
		}
	}
	return nrPeaks
}
