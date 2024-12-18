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
		fmt.Println("task1: ", task1(lines, 71, 71, 1024))
	} else {
		fmt.Println("task1: ", task2(lines, 71, 71, 1024))
	}
}

func task1(lines []string, w, h, nr int) int {
	g := u.CreateEmptyCharGrid(w, h)
	g.SetAll(".")
	for i := 0; i < nr; i++ {
		coord := u.SplitToInts(lines[i])
		x, y := coord[0], coord[1]
		g.Grid[y][x] = "#"
	}
	fmt.Println(g)
	start := u.Pos2D{0, 0}
	end := u.Pos2D{h - 1, w - 1}
	_, minCost := traverse(g, start, end)
	return minCost
}

func task2(lines []string, w, h, nr int) string {
	g := u.CreateEmptyCharGrid(w, h)
	g.SetAll(".")
	for i := 0; i < nr+1; i++ {
		coord := u.SplitToInts(lines[i])
		x, y := coord[0], coord[1]
		g.Grid[y][x] = "#"
	}
	i := nr + 1
	start := u.Pos2D{0, 0}
	end := u.Pos2D{h - 1, w - 1}
	for {
		coord := u.SplitToInts(lines[i])
		x, y := coord[0], coord[1]
		g.Grid[y][x] = "#"
		ok := traverse2(g, start, end)
		if !ok {
			return fmt.Sprintf("%d,%d", x, y)
		}
		i++
	}
}

type state struct {
	pos  u.Pos2D
	from path
	Cost int
}

type path struct {
	pos u.Pos2D
	dir u.Pos2D
}

func traverse(g u.CharGrid, start u.Pos2D, end u.Pos2D) (map[u.Pos2D]int, int) {
	pq := u.NewHeap(func(s1, s2 state) int {
		return s1.Cost - s2.Cost
	})
	right := u.Pos2D{0, 1}
	pq.Push(state{pos: start, Cost: 0, from: path{start, right}})
	visited := u.CreateSet[state]()
	lowestCost := make(map[u.Pos2D]int)
	for {
		if pq.Len() == 0 {
			break
		}
		curr := pq.Pop()
		if _, ok := lowestCost[curr.pos]; ok {
			continue
		}
		visited.Add(curr)
		if _, ok := lowestCost[curr.pos]; !ok {
			lowestCost[curr.pos] = curr.Cost
		}
		if curr.pos == end {
			return lowestCost, curr.Cost
		}
		for _, dir := range u.Dirs2D {
			if free(g, curr.pos, dir) {
				new := curr.pos.Add(dir)
				pq.Push(state{pos: new, from: path{curr.pos, dir},
					Cost: curr.Cost + 1})
			}
		}
	}
	return nil, -1
}

func free(g u.CharGrid, pos, dir u.Pos2D) bool {
	n := pos.Add(dir)
	return g.InBounds(n.Row, n.Col) && g.At(n.Row, n.Col) != "#"
}

func traverse2(g u.CharGrid, start u.Pos2D, end u.Pos2D) bool {
	pq := u.NewHeap(func(s1, s2 state) int {
		return s1.Cost - s2.Cost
	})
	right := u.Pos2D{0, 1}
	pq.Push(state{pos: start, Cost: 0, from: path{start, right}})
	visited := u.CreateSet[state]()
	lowestCost := make(map[u.Pos2D]int)
	for {
		if pq.Len() == 0 {
			return false
		}
		curr := pq.Pop()
		if _, ok := lowestCost[curr.pos]; ok {
			continue
		}
		visited.Add(curr)
		if _, ok := lowestCost[curr.pos]; !ok {
			lowestCost[curr.pos] = curr.Cost
		}
		if curr.pos == end {
			return true
		}
		for _, dir := range u.Dirs2D {
			if free(g, curr.pos, dir) {
				new := curr.pos.Add(dir)
				pq.Push(state{pos: new, from: path{curr.pos, dir},
					Cost: curr.Cost + 1})
			}
		}
	}
	return false
}
