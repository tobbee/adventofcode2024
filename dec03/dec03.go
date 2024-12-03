package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

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
	var rex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	for _, line := range lines {
		matches := rex.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			sum += a * b
		}
	}
	return sum
}

func task2(lines []string) int {
	var rex = regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)

	sum := 0
	enabled := true
	for _, line := range lines {
		matches := rex.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			switch m[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					a, _ := strconv.Atoi(m[1])
					b, _ := strconv.Atoi(m[2])
					sum += a * b
				}
			}
		}
	}
	return sum
}
