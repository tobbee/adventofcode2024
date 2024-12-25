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
	keys, locks := parseInput(lines)
	fmt.Println("nrKeys: ", len(keys), " nrLocks: ", len(locks))
	nrCompatible := 0
	for _, key := range keys {
		for _, lock := range locks {
			if compatible(key, lock) {
				nrCompatible++
			}
		}
	}

	return nrCompatible
}

func compatible(key, lock [5]int) bool {
	for i := 0; i < 5; i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}

func task2(lines []string) int {
	return 0
}

func parseInput(lines []string) ([][5]int, [][5]int) {
	keys := make([][5]int, 0, 100)
	locks := make([][5]int, 0, 100)
	group := make([]string, 0, 7)
	for _, line := range lines {
		switch line {
		case "":
			if group[0] == "#####" {
				locks = append(locks, makeLock(group[1:6]))
			} else {
				keys = append(keys, makeKey(group[1:6]))
			}
			group = group[:0]
		default:
			group = append(group, line)
		}
	}
	if group[0] == "#####" {
		locks = append(locks, makeLock(group[1:6]))
	} else {
		keys = append(keys, makeKey(group[1:6]))
	}
	return keys, locks
}

func makeLock(group []string) [5]int {
	lock := [5]int{5, 5, 5, 5, 5}
	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if group[row][col] == '.' {
				lock[col] = row
				break
			}
		}
	}
	return lock
}

func makeKey(group []string) [5]int {
	key := [5]int{5, 5, 5, 5, 5}
	for col := 0; col < 5; col++ {
		for row := 4; row >= 0; row-- {
			if group[row][col] == '.' {
				key[col] = 4 - row
				break
			}
		}
	}
	return key
}
