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

type state struct {
	pos  u.Pos2D
	dir  u.Pos2D
	from path
	Cost int
}

type path struct {
	pos u.Pos2D
	dir u.Pos2D
}

func task1(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	rs, cs := g.Find("S")
	start := u.Pos2D{rs, cs}
	re, ce := g.Find("E")
	end := u.Pos2D{re, ce}
	_, minCost := traverse(g, start, end)
	return minCost
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	rs, cs := g.Find("S")
	start := u.Pos2D{rs, cs}
	re, ce := g.Find("E")
	end := u.Pos2D{re, ce}
	pathCosts, minCost := traverse(g, start, end)
	fmt.Println("minCost: ", minCost)
	bestPos := u.CreateSet[u.Pos2D]()
	// Start with the two possible end paths
	bestPos.Add(end)
	backTrack(g, pathCosts, path{end, u.Pos2D{0, 1}}, minCost, bestPos)
	backTrack(g, pathCosts, path{end, u.Pos2D{-1, 0}}, minCost, bestPos)
	fmt.Printf("nr best pos: %d\n", len(bestPos))
	return len(bestPos)
}

func free(g u.CharGrid, pos, dir u.Pos2D) bool {
	n := pos.Add(dir)
	return g.InBounds(n.Row, n.Col) && g.At(n.Row, n.Col) != "#"
}

func traverse(g u.CharGrid, start u.Pos2D, end u.Pos2D) (map[path]int, int) {
	pq := u.NewHeap(func(s1, s2 state) int {
		return s1.Cost - s2.Cost
	})
	right := u.Pos2D{0, 1}
	pq.Push(state{pos: start, dir: right, Cost: 0, from: path{start, right}})
	visited := u.CreateSet[state]()
	lowestCost := make(map[path]int)
	for {
		if pq.Len() == 0 {
			break
		}
		curr := pq.Pop()
		if visited.Contains(curr) {
			continue
		}
		visited.Add(curr)
		if _, ok := lowestCost[path{curr.pos, curr.dir}]; !ok {
			lowestCost[path{curr.pos, curr.dir}] = curr.Cost
		}
		if curr.Cost < lowestCost[path{curr.pos, curr.dir}] {
			lowestCost[path{curr.pos, curr.dir}] = curr.Cost
		}
		if curr.pos == end {
			return lowestCost, curr.Cost
		}
		/*
			switch curr.dir {
			case u.Pos2D{0, 1}:
				g.Grid[curr.pos.Row][curr.pos.Col] = ">"
			case u.Pos2D{1, 0}:
				g.Grid[curr.pos.Row][curr.pos.Col] = "v"
			case u.Pos2D{0, -1}:
				g.Grid[curr.pos.Row][curr.pos.Col] = "<"
			case u.Pos2D{-1, 0}:
				g.Grid[curr.pos.Row][curr.pos.Col] = "^"
			}
			fmt.Println(g)
			fmt.Println("curr: ", curr)
			fmt.Println()
		*/
		if free(g, curr.pos, curr.dir) {
			next := curr.pos.Add(curr.dir)
			pq.Push(state{pos: next, dir: curr.dir, Cost: curr.Cost + 1,
				from: path{curr.pos, curr.dir}})
		}
		if free(g, curr.pos, curr.dir.Left()) {
			pq.Push(state{
				pos:  curr.pos,
				dir:  curr.dir.Left(),
				Cost: curr.Cost + 1000,
				from: path{curr.pos, curr.dir}})
		}
		if free(g, curr.pos, curr.dir.Right()) {
			pq.Push(state{
				pos:  curr.pos,
				dir:  curr.dir.Right(),
				Cost: curr.Cost + 1000,
				from: path{curr.pos, curr.dir}})
		}
	}
	return nil, -1
}

func backTrack(g u.CharGrid, pathCosts map[path]int, currPath path, currCost int, bestPos u.Set[u.Pos2D]) {
	if currCost == 0 {
		return
	}
	prevPath := path{currPath.pos.Sub(currPath.dir), currPath.dir}
	if pathCosts[prevPath] == currCost-1 {
		bestPos.Add(prevPath.pos)
		backTrack(g, pathCosts, prevPath, currCost-1, bestPos)
	}
	prevPath = path{currPath.pos, currPath.dir.Left()}
	if pathCosts[prevPath] == currCost-1000 {
		bestPos.Add(prevPath.pos)
		backTrack(g, pathCosts, prevPath, currCost-1000, bestPos)
	}
	prevPath = path{currPath.pos, currPath.dir.Right()}
	if pathCosts[prevPath] == currCost-1000 {
		bestPos.Add(prevPath.pos)
		backTrack(g, pathCosts, prevPath, currCost-1000, bestPos)
	}
}
