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
	grid := u.CreateCharGridFromLines(lines)
	nrXmas := 0
	chars := getChars("XMAS")
	dirs := [][2]int{{1, 0}, {1, 1}, {1, -1}, {-1, 0}, {-1, 1}, {-1, -1}, {0, 1}, {0, -1}}
	for row := 0; row < grid.Height; row++ {
		for col := 0; col < grid.Width; col++ {
			if grid.Grid[row][col] == "X" {
				for _, dir := range dirs {
					if checkDirection(grid, row, col, dir, chars) {
						nrXmas++
					}
				}
			}
		}
	}
	return nrXmas
}

func checkDirection(grid u.CharGrid, row, col int, dir [2]int, chars []string) bool {
	for i := 0; i < len(chars); i++ {
		if !grid.InBounds(row, col) {
			return false
		}
		if grid.Grid[row][col] != chars[i] {
			return false
		}
		row += dir[0]
		col += dir[1]
	}
	return true
}

func task2(lines []string) int {
	grid := u.CreateCharGridFromLines(lines)
	nrXmas := 0
	chars := getChars("MAS")
	diag1Dirs := [][2]int{{1, 1}, {-1, -1}}
	diag2Dirs := [][2]int{{1, -1}, {-1, 1}}

	for row := 1; row < grid.Height-1; row++ {
		for col := 1; col < grid.Width-1; col++ {
			if grid.Grid[row][col] == "A" {
				for _, dir := range diag1Dirs {
					if checkDiagonal(grid, row, col, dir, chars) {
						for _, dir2 := range diag2Dirs {
							if checkDiagonal(grid, row, col, dir2, chars) {
								nrXmas++
								continue
							}
						}
					}
				}
			}
		}
	}
	return nrXmas
}

func checkDiagonal(grid u.CharGrid, row, col int, dir [2]int, chars []string) bool {
	for i := -1; i <= 1; i++ {
		r := row + i*dir[0]
		c := col + i*dir[1]
		if !grid.InBounds(r, c) {
			continue
		}
		if grid.Grid[r][c] != chars[i+1] {
			return false
		}
	}
	return true
}

func getChars(word string) []string {
	chars := make([]string, 0, len(word))
	for i := range word {
		chars = append(chars, string(word[i]))
	}
	return chars
}
