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
	nrSafe := 0
	for _, line := range lines {
		values := u.SplitToInts(line)
		if isSafe(values) {
			nrSafe++
		}
	}
	return nrSafe
}

func isSafe(values []int) bool {
	inc := false
	lastV := 0
	for i, v := range values {
		if i == 0 {
			continue
		}
		lastV = values[i-1]
		if i == 1 {
			if v > lastV {
				inc = true
			}
		}
		if inc {
			if v < lastV+1 || v > lastV+3 {
				return false
			}
			continue
		}
		if v < lastV-3 || v > lastV-1 {
			return false
		}
	}
	return true
}

func firstBadIndex(values []int) int {
	inc := false
	lastV := 0
	for i, v := range values {
		if i == 0 {
			continue
		}
		lastV = values[i-1]
		if i == 1 {
			if v > lastV {
				inc = true
			}
		}
		if inc {
			if v < lastV+1 || v > lastV+3 {
				return i
			}
			continue
		}
		if v < lastV-3 || v > lastV-1 {
			return i
		}
	}
	return -1

}

func task2(lines []string) int {
	nrSafe := 0
	for _, line := range lines {
		values := u.SplitToInts(line)
		badIndex := firstBadIndex(values)
		if badIndex == -1 {
			nrSafe++
			continue
		}
		dampened := make([]int, 0, len(values)-1)
		dampened = append(dampened, values[:badIndex]...)
		dampened = append(dampened, values[badIndex+1:]...)
		if isSafe(dampened) {
			nrSafe++
			continue
		}
		// Try to remove the first value
		copy(dampened, values[1:])
		if isSafe(dampened) {
			nrSafe++
			continue
		}
		// Try to remove the second value
		dampened = dampened[:0]
		dampened = append(dampened, values[0])
		dampened = append(dampened, values[2:]...)
		if isSafe(dampened) {
			nrSafe++
			continue
		}
	}
	return nrSafe
}
