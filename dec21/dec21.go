package main

import (
	"flag"
	"fmt"
	"math"

	u "github.com/tobbee/adventofcode2024/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task(lines, 2))
	} else {
		fmt.Println("task2: ", task(lines, 25))
	}
}

type keyMap map[string]u.Pos2D
type reverseMap map[u.Pos2D]string

var numGrid = keyMap{
	"7": u.Pos2D{0, 0},
	"8": u.Pos2D{0, 1},
	"9": u.Pos2D{0, 2},
	"4": u.Pos2D{1, 0},
	"5": u.Pos2D{1, 1},
	"6": u.Pos2D{1, 2},
	"1": u.Pos2D{2, 0},
	"2": u.Pos2D{2, 1},
	"3": u.Pos2D{2, 2},
	"0": u.Pos2D{3, 1},
	"A": u.Pos2D{3, 2},
}

var numGridReverse = reverseMap{
	u.Pos2D{0, 0}: "7",
	u.Pos2D{0, 1}: "8",
	u.Pos2D{0, 2}: "9",
	u.Pos2D{1, 0}: "4",
	u.Pos2D{1, 1}: "5",
	u.Pos2D{1, 2}: "6",
	u.Pos2D{2, 0}: "1",
	u.Pos2D{2, 1}: "2",
	u.Pos2D{2, 2}: "3",
	u.Pos2D{3, 1}: "0",
	u.Pos2D{3, 2}: "A",
}

var dirGrid = keyMap{
	"^": u.Pos2D{0, 1},
	"A": u.Pos2D{0, 2},
	"<": u.Pos2D{1, 0},
	"v": u.Pos2D{1, 1},
	">": u.Pos2D{1, 2},
}

var dirGridReverse = reverseMap{
	u.Pos2D{0, 1}: "^",
	u.Pos2D{0, 2}: "A",
	u.Pos2D{1, 0}: "<",
	u.Pos2D{1, 1}: "v",
	u.Pos2D{1, 2}: ">",
}

var dirToButton = map[u.Pos2D]string{
	u.Pos2D{0, 1}:  ">",
	u.Pos2D{0, -1}: "<",
	u.Pos2D{1, 0}:  "v",
	u.Pos2D{-1, 0}: "^",
}

// stepsFromPath returns a a list of steps to take to to visit all the
// buttons in the code. Each step is a two-character string, where
// the first character is the start position and the second character
// is the end position.
func stepsFromPath(code string) (steps []string) {
	for i := 0; i < len(code); i++ {
		switch i {
		case 0:
			steps = append(steps, "A"+code[0:1])
		default:
			steps = append(steps, code[i-1:i+1])
		}
	}
	return steps
}

func getKeys(m map[string]u.Pos2D) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// findShortetPaths returns a map of the shortest sequendes to go from
// one key to another and press it. It always end by an "A".
func findShortestSequences(keyPad keyMap, reverse reverseMap) map[string][]string {
	shortestPaths := make(map[string][]string)
	keys := getKeys(keyPad)
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(keys); j++ {
			start := keys[i]
			end := keys[j]
			paths, _ := findAllShortest(start, end, keyPad, reverse, 0)
			shortestPaths[start+end] = paths
		}
	}
	return shortestPaths
}

func findAllShortest(start, end string, keyPad keyMap, reverse reverseMap, length int) ([]string, int) {
	minLen := math.MaxInt
	shortest := make([]string, 0)
	switch {
	case start == end:
		return []string{"A"}, length + 1
	case length > 5:
		return nil, math.MaxInt
	default:
		pos := keyPad[start]
		// Breadth first search
		for _, dir := range u.Dirs2D {
			newPos := pos.Add(dir)
			if _, ok := reverse[newPos]; !ok {
				continue
			}
			newChar := reverse[newPos]
			if newChar == end {
				return []string{dirToButton[dir] + "A"}, length + 2
			}
		}
		for _, dir := range u.Dirs2D {
			newPos := pos.Add(dir)
			if _, ok := reverse[newPos]; !ok {
				continue
			}
			endPos := keyPad[end]
			newDiff := endPos.Sub(newPos)
			oldDiff := endPos.Sub(pos)
			if newDiff.Manhattan() > oldDiff.Manhattan() {
				continue
			}
			newChar := reverse[newPos]
			paths, pathLen := findAllShortest(newChar, end, keyPad, reverse, length+1)
			switch {
			case pathLen > minLen:
				continue
			case pathLen < minLen:
				shortest = shortest[:0] // Forget what we had
				for _, path := range paths {
					p := dirToButton[dir] + path
					shortest = append(shortest, p)
					minLen = pathLen
				}
			default:
				for _, path := range paths {
					p := dirToButton[dir] + path
					shortest = append(shortest, p)
				}
			}
		}
	}
	return shortest, minLen
}

func task(lines []string, nrLevels int) int {
	shortestNumPaths := findShortestSequences(numGrid, numGridReverse)
	shortestDirPaths := findShortestSequences(dirGrid, dirGridReverse)
	totProd := 0
	for _, code := range lines {
		steps := stepsFromPath(code)
		fmt.Println("code", code, "steps", steps)
		shortestPaths := []string{""}
		for _, step := range steps {
			paths := shortestNumPaths[step]
			newShortestPaths := make([]string, 0, len(paths)*len(shortestPaths))
			for _, sp := range shortestPaths {
				for _, path := range paths {
					newShortestPaths = append(newShortestPaths, sp+path)
				}
			}
			shortestPaths = newShortestPaths
		}
		minNrPresses := math.MaxInt
		for _, sp := range shortestPaths {
			nrPresses := pressButtons(sp, nrLevels, shortestDirPaths)
			if nrPresses < minNrPresses {
				minNrPresses = nrPresses
			}
		}
		nr := u.SplitToInts(code)[0]
		prod := nr * minNrPresses
		totProd += prod
	}

	return totProd
}

type entry struct {
	path  string
	level int
}

var cache = make(map[entry]int)

func pressButtons(code string, level int, shortestDirPaths map[string][]string) int {
	if val, ok := cache[entry{code, level}]; ok {
		return val
	}
	if level == 0 {
		return len(code)
	}
	length := 0
	steps := stepsFromPath(code)

	for _, step := range steps {
		minLength := math.MaxInt
		subPaths := shortestDirPaths[step]
		for _, sp := range subPaths {
			l := pressButtons(sp, level-1, shortestDirPaths)
			if l < minLength {
				minLength = l
			}
		}
		length += minLength
	}
	cache[entry{code, level}] = length
	return length
}
