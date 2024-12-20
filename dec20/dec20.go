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
		fmt.Println("task1: ", task1(lines, 100))
	} else {
		fmt.Println("task2: ", task2(lines, 20, 100))
	}
}

func task1(lines []string, minWin int) int {
	g := u.CreateCharGridFromLines(lines)
	//fmt.Println(g)
	r, c := g.Find("S")
	start := u.Pos2D{r, c}
	r, c = g.Find("E")
	end := u.Pos2D{r, c}
	endDists := findEndDists(g, start, end)
	nrCheats := 0
	for pos := range endDists {
		nrCheats += findGoodCheats1(g, endDists, pos, minWin)
	}
	return nrCheats
}

func task2(lines []string, maxCheat, minWin int) int {
	g := u.CreateCharGridFromLines(lines)
	// fmt.Println(g)
	r, c := g.Find("S")
	start := u.Pos2D{r, c}
	r, c = g.Find("E")
	end := u.Pos2D{r, c}
	endDists := findEndDists(g, start, end)
	nrCheats := 0
	for pos := range endDists {
		nrCheats += findGoodCheats(g, endDists, pos, maxCheat, minWin)
	}
	return nrCheats
}

func findEndDists(g u.CharGrid, start, end u.Pos2D) map[u.Pos2D]int {
	dists := make(map[u.Pos2D]int)
	pos := end
	dists[pos] = 0
	for {
		for _, dir := range u.Dirs2D {
			np := pos.Add(dir)
			if g.Grid[np.Row][np.Col] == "#" {
				continue
			}
			if _, ok := dists[np]; ok {
				continue
			}
			dists[np] = dists[pos] + 1
			if np == start {
				return dists
			}
			pos = np
			break
		}
	}
}

func findGoodCheats(g u.CharGrid, endDists map[u.Pos2D]int, start u.Pos2D, maxDist, minWin int) int {
	nrCheats := 0
	startDist := endDists[start]
	if startDist < minWin {
		return 0
	}

	for i := -maxDist; i <= maxDist; i++ {
		for j := -maxDist; j <= maxDist; j++ {
			p := start.Add(u.Pos2D{i, j})
			if !g.InBounds(p.Row, p.Col) {
				continue
			}
			mDist := abs(i) + abs(j)
			if mDist > maxDist {
				continue
			}
			if ed, ok := endDists[p]; ok {
				win := startDist - ed - mDist
				if win >= minWin {
					nrCheats++
				}
			}
		}
	}
	return nrCheats
}

func findGoodCheats1(g u.CharGrid, endDists map[u.Pos2D]int, pos u.Pos2D, minWin int) int {
	nrCheats := 0
	for _, dir := range u.Dirs2D {
		np := pos.Add(dir)
		np2 := np.Add(dir)
		np3 := np2.Add(dir)
		if g.Grid[np.Row][np.Col] == "#" {
			if g.InBounds(np2.Row, np2.Col) && g.Grid[np2.Row][np2.Col] != "#" {
				win := endDists[np2] - endDists[pos] - 2
				if win >= minWin {
					nrCheats++
				}
			} else if g.InBounds(np3.Row, np3.Col) && g.Grid[np3.Row][np3.Col] != "#" {
				win := endDists[np3] - endDists[pos] - 3
				if win >= minWin {
					nrCheats++
				}
			}
		}
	}
	return nrCheats
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
