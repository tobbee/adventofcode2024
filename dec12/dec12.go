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

type regData struct {
	crop      string
	area      int
	perimeter int
	sides     int
}

func task1(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	regions := u.CreateGrid2D[int](g.Width, g.Height)
	regions.SetAll(-1)
	findRegions(g, regions)
	regionData := fillRegionData(g, regions)
	total := 0
	for _, rd := range regionData {
		total += rd.area * rd.perimeter
	}
	return total
}

func fillRegionData(g u.CharGrid, regions u.Grid2D[int]) map[int]*regData {
	regionData := make(map[int]*regData)
	for r := 0; r < regions.Height; r++ {
		for c := 0; c < regions.Width; c++ {
			region := regions.Get(r, c)
			if _, ok := regionData[region]; !ok {
				regionData[region] = &regData{crop: g.Grid[r][c]}
			}
			regionData[region].area++
			pos := u.Pos2D{Row: r, Col: c}
			for _, dir := range u.Dirs2D {
				np := pos.Add(dir)
				if !g.InBounds(np.Row, np.Col) || regions.Get(np.Row, np.Col) != region {
					regionData[region].perimeter++
					// Next check if there is a new side and count it
					var neighbor u.Pos2D
					switch dir {
					case u.Pos2D{-1, 0}, u.Pos2D{1, 0}: // up or down
						neighbor = pos.Add(u.Pos2D{0, -1}) // left
					case u.Pos2D{0, -1}, u.Pos2D{0, 1}: // left, right
						neighbor = pos.Add(u.Pos2D{-1, 0}) // above
					}
					switch {
					case !g.InBounds(neighbor.Row, neighbor.Col) || regions.Get(neighbor.Row, neighbor.Col) != region:
						regionData[region].sides++ // New side for sure
					case !hasSameEdge(regions, pos, neighbor, dir):
						regionData[region].sides++
					}
				}
			}
		}
	}
	return regionData
}

func hasSameEdge(regions u.Grid2D[int], pos, neighbor, edge u.Pos2D) bool {
	region := regions.Get(pos.Row, pos.Col)
	if !regions.InBounds(neighbor.Row, neighbor.Col) {
		return true
	}
	neighborRegion := regions.Get(neighbor.Row, neighbor.Col)
	if neighborRegion != region {
		return false
	}
	nextNb := neighbor.Add(edge)
	if !regions.InBounds(nextNb.Row, nextNb.Col) || regions.Get(nextNb.Row, nextNb.Col) != region {
		return true
	}
	return false
}

func findRegions(g u.CharGrid, regions u.Grid2D[int]) {
	region := 0
	for {
		pos := findNextUnknown(regions)
		if pos.Row == -1 {
			break
		}
		crop := g.Grid[pos.Row][pos.Col]
		growRegion(g, regions, pos, crop, region)
		region++
	}
}

func growRegion(g u.CharGrid, regions u.Grid2D[int], pos u.Pos2D, crop string, region int) {
	regions.Set(region, pos.Row, pos.Col)
	for _, dir := range u.Dirs2D {
		np := pos.Add(dir)
		if g.InBounds(np.Row, np.Col) && regions.Get(np.Row, np.Col) == -1 && g.Grid[np.Row][np.Col] == crop {
			growRegion(g, regions, np, crop, region)
		}
	}
}

func findNextUnknown(regions u.Grid2D[int]) u.Pos2D {
	for r := 0; r < regions.Height; r++ {
		for c := 0; c < regions.Width; c++ {
			if regions.Get(r, c) == -1 {
				return u.Pos2D{Row: r, Col: c}
			}
		}
	}
	return u.Pos2D{Row: -1, Col: -1}
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	regions := u.CreateGrid2D[int](g.Width, g.Height)
	regions.SetAll(-1)
	findRegions(g, regions)
	regionData := fillRegionData(g, regions)
	total := 0
	for _, rd := range regionData {
		total += rd.area * rd.sides
	}
	return total
}
