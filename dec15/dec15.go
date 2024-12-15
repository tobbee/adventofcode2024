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
	warehouse, moves := parseInput(lines)
	//fmt.Println(warehouse)
	r, c := warehouse.Find("@")
	pos := u.Pos2D{r, c}
	for i := 0; i < len(moves); i++ {
		pos = move(warehouse, moves[i:i+1], pos)
	}
	return gpsSum(warehouse)
}

func task2(lines []string) int {
	w, moves := parseInput(lines)
	warehouse := doubleWarehouse(w)
	// fmt.Println(warehouse)
	r, c := warehouse.Find("@")
	pos := u.Pos2D{r, c}
	for i := 0; i < len(moves); i++ {
		// fmt.Printf("%d move: %s", i, moves[i:i+1])
		pos = move2(warehouse, moves[i:i+1], pos)
		// fmt.Println(warehouse)
		// fmt.Println()
	}
	return gpsSum(warehouse)
}

func parseInput(lines []string) (u.CharGrid, string) {
	breakLine := 0
	for i, line := range lines {
		if line == "" {
			breakLine = i
			break
		}
	}
	warehouse := u.CreateCharGridFromLines(lines[:breakLine])
	moves := ""
	for _, line := range lines[breakLine+1:] {
		moves += line
	}
	return warehouse, moves
}

func move(w u.CharGrid, m string, p u.Pos2D) u.Pos2D {
	if w.Grid[p.Row][p.Col] != "@" {
		fmt.Println(w)
		panic("Invalid position")
	}
	var dir u.Pos2D
	switch m {
	case "^":
		dir = u.Pos2D{-1, 0}
	case "v":
		dir = u.Pos2D{1, 0}
	case "<":
		dir = u.Pos2D{0, -1}
	case ">":
		dir = u.Pos2D{0, 1}
	}
	n := p.Add(dir)
	switch c := w.Grid[n.Row][n.Col]; c {
	case "#":
		return p
	case ".":
		return moveRobot(w, p, n)
	case "O":
		// Check if we can push one or more boxes
		e := n
	trainLoop:
		for {
			e = e.Add(dir)
			switch w.Grid[e.Row][e.Col] {
			case "O":
				continue
			case ".":
				break trainLoop
			default:
				return p // Can't push
			}
		}
		w.Grid[e.Row][e.Col] = "O" // Add a box at the end
		return moveRobot(w, p, n)
	default:
		panic("Invalid char '" + c + "'")
	}
}

func move2(w u.CharGrid, m string, p u.Pos2D) u.Pos2D {
	if w.Grid[p.Row][p.Col] != "@" {
		fmt.Println(w)
		panic("Invalid position")
	}
	var dir u.Pos2D
	switch m {
	case "^":
		dir = u.Pos2D{-1, 0}
	case "v":
		dir = u.Pos2D{1, 0}
	case "<":
		dir = u.Pos2D{0, -1}
	case ">":
		dir = u.Pos2D{0, 1}
	}
	n := p.Add(dir)
	switch nc := w.Grid[n.Row][n.Col]; nc {
	case "#":
		return p
	case ".":
		return moveRobot(w, p, n)
	case "[", "]":
		// Check if we can push one or more boxes
		// Need to distinguish between directions
		switch dir {
		case u.Pos2D{0, 1}: // Right
			if nc != "[" {
				panic("Invalid char '" + nc + "'")
			}
			nrBoxes := 1
		rightLoop:
			for i := n.Col + 2; i < w.Width; i += 2 {
				switch w.Grid[n.Row][i] {
				case "[":
					nrBoxes++
					continue // another box to the right
				case ".":
					break rightLoop
				case "#":
					return p // Can't push
				default:
					panic("Invalid char '" + w.Grid[n.Row][i] + "'")
				}
			}
			// Now move the boxes and the robot
			for i := 0; i < nrBoxes; i++ {
				w.Grid[n.Row][n.Col+1+2*i] = "["
				w.Grid[n.Row][n.Col+2+2*i] = "]"
			}
			return moveRobot(w, p, n)
		case u.Pos2D{0, -1}: // Left
			if nc != "]" {
				panic("Invalid char '" + nc + "'")
			}
			nrBoxes := 1
		leftLoop:
			for i := n.Col - 2; i >= 0; i -= 2 {
				switch w.Grid[n.Row][i] {
				case "]":
					nrBoxes++
					continue // another box to the left
				case ".":
					break leftLoop
				case "#":
					return p // Can't push
				default:
					panic("Invalid char '" + w.Grid[n.Row][i] + "'")
				}
			}
			// Now move the boxes and the robot
			for i := 0; i < nrBoxes; i++ {
				w.Grid[n.Row][n.Col-1-2*i] = "]"
				w.Grid[n.Row][n.Col-2-2*i] = "["
			}
			return moveRobot(w, p, n)
		case u.Pos2D{1, 0}, u.Pos2D{-1, 0}: // Down or up
			if canPushVertical(w, p, dir) {
				pushVertical(w, p, dir)
				return n
			} else {
				return p
			}
		default:
			panic("Invalid direction")
		}
	default:
		panic("Invalid char '" + nc + "'")
	}
	panic("Should not reach here")
}

func moveRobot(w u.CharGrid, present, new u.Pos2D) u.Pos2D {
	w.Grid[new.Row][new.Col] = "@"
	w.Grid[present.Row][present.Col] = "."
	return new
}

func canPushVertical(w u.CharGrid, p u.Pos2D, dir u.Pos2D) bool {
	n := p.Add(dir)
	switch ch := w.Grid[n.Row][n.Col]; ch {
	case ".":
		return true
	case "#":
		return false
	case "[":
		n2 := n.Add(u.Pos2D{0, 1})
		return canPushVertical(w, n, dir) && canPushVertical(w, n2, dir)
	case "]":
		n2 := n.Add(u.Pos2D{0, -1})
		return canPushVertical(w, n, dir) && canPushVertical(w, n2, dir)
	default:
		panic("Invalid char '" + ch + "'")
	}
}

// Push a vertical train of wide boxes
func pushVertical(w u.CharGrid, p u.Pos2D, dir u.Pos2D) {
	n := p.Add(dir)
	if w.Grid[n.Row][n.Col] == "." {
		w.Grid[n.Row][n.Col] = w.Grid[p.Row][p.Col]
		w.Grid[p.Row][p.Col] = "."
		return
	}
	// Widen p to two positions depending on edge
	var n1, n2 u.Pos2D
	switch w.Grid[n.Row][n.Col] {
	case "[":
		n1 = n
		n2 = n.Add(u.Pos2D{0, 1})
	case "]":
		n1 = n.Add(u.Pos2D{0, -1})
		n2 = n
	}

	pushVertical(w, n1, dir)
	pushVertical(w, n2, dir)
	// Now we should be able to move ourself
	w.Grid[n.Row][n.Col] = w.Grid[p.Row][p.Col]
	w.Grid[p.Row][p.Col] = "."
}

func gpsSum(w u.CharGrid) int {
	sum := 0
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Grid[i][j] == "O" || w.Grid[i][j] == "[" {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func doubleWarehouse(w u.CharGrid) u.CharGrid {
	w2 := u.CreateEmptyCharGrid(w.Width*2, w.Height)
	for r := 0; r < w.Height; r++ {
		for c := 0; c < w.Width; c++ {
			switch w.Grid[r][c] {
			case "#":
				w2.Grid[r][2*c] = "#"
				w2.Grid[r][2*c+1] = "#"
			case "O":
				w2.Grid[r][2*c] = "["
				w2.Grid[r][2*c+1] = "]"
			case ".":
				w2.Grid[r][2*c] = "."
				w2.Grid[r][2*c+1] = "."
			case "@":
				w2.Grid[r][2*c] = "@"
				w2.Grid[r][2*c+1] = "."
			default:
				panic("Invalid char '" + w.Grid[r][c] + "'")
			}
		}
	}
	return w2
}
