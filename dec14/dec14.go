package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	u "github.com/tobbee/adventofcode2024/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines, 101, 103))
	} else {
		fmt.Println("task2: ", task2(lines, 101, 103))
	}
}

type robot struct {
	pos u.Pos2D
	vel u.Pos2D
}

func task1(lines []string, width, height int) int {
	robots := getRobots(lines)
	for _, robot := range robots {
		fmt.Println(robot)
	}
	for i := 0; i < 100; i++ {
		moveRobots(width, height, robots)
	}
	//printGrid(width, height, robots)
	//fmt.Println()
	q := countQuadrants(width, height, robots)
	prod := 1
	for _, v := range q {
		prod *= v
	}

	return prod
}

func moveRobots(width, height int, robots []robot) {
	for i, robot := range robots {
		pos := robot.pos.Add(robot.vel)
		pos.Row = pos.Row % height
		if pos.Row < 0 {
			pos.Row += height
		}
		pos.Col = pos.Col % width
		if pos.Col < 0 {
			pos.Col += width
		}
		robots[i].pos = pos
	}
}

func printTree(width, height int, robots []robot) {
	g := u.CreateGrid2D[int](width, height)
	nr := 0
	for row := 0; row < height; row++ {
		mid := width / 2
		for col := mid - row; col <= mid+row; col++ {
			g.Set(1, row, col)
			nr++
		}
		fmt.Println("row", row, "nr: ", nr)
	}
}

// countTreePixels returns the number of pixels in a pyramid shape
// from the mid top of the grid to the bottom row.
func countTreePixels(width, _ int, robots []robot) int {
	mid := width / 2
	edgePositions := u.CreateSet[u.Pos2D]()
	for _, robot := range robots {
		row := robot.pos.Row
		col := robot.pos.Col
		leftBorder := mid - row/2
		rightBorder := mid + row/2
		if leftBorder <= col && col <= rightBorder {
			edgePositions.Add(robot.pos)
		}
	}
	return edgePositions.Size()
}

func countQuadrants(width, height int, robots []robot) []int {
	quadrants := make([]int, 4)
	midHeight := height / 2
	midWidth := width / 2
	for _, robot := range robots {
		quad := 0
		switch {
		case robot.pos.Row > midHeight:
			quad += 2
		case robot.pos.Row < midHeight:
			quad += 0
		default:
			continue // Skip robots on the middle row

		}
		switch {
		case robot.pos.Col > midWidth:
			quad++
		case robot.pos.Col < midWidth:
			quad += 0
		default:
			continue // Skip robots on the middle col
		}
		quadrants[quad]++
	}
	return quadrants
}

func getRobots(lines []string) []robot {
	robots := []robot{}
	for _, line := range lines {
		nrs := u.SplitToInts(line)
		robots = append(robots, robot{u.Pos2D{nrs[1], nrs[0]}, u.Pos2D{nrs[3], nrs[2]}})
	}
	return robots
}

func printGrid(width, height int, robots []robot) {
	g := u.CreateGrid2D[int](width, height)
	for _, robot := range robots {
		val := g.Get(robot.pos.Row, robot.pos.Col)
		g.Set(val+1, robot.pos.Row, robot.pos.Col)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			val := g.Get(i, j)
			if val == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("X")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func task2(lines []string, width, height int) int {
	robots := getRobots(lines)
	// The width and height are both prime numbers, so the motion is
	// periodic with period width*height.

	// We do not know the shape of the Christmas Tree but we stop
	// as the number of "robots" within a pyramid shape from the top
	// reaches a new maximum, and inspects the picture.
	// Hopefully, this reveals a christmas tree.

	nrSeconds := 0
	maxTreePixels := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to continue after each stop")
	reader.ReadString('\n')
	for i := 0; i < width*height; i++ {
		moveRobots(width, height, robots)
		nrSeconds++
		treePixels := countTreePixels(width, height, robots)
		if treePixels > maxTreePixels {
			maxTreePixels = treePixels
			fmt.Printf("Seconds: %d, nr tree pixels: %d\n", nrSeconds, treePixels)
			printGrid(width, height, robots)
			reader.ReadString('\n')
		}
	}
	return nrSeconds
}
