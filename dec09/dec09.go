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
	memory := initMemory(lines[0])
	//fmt.Println(memory)
	updateMemory(memory)
	//fmt.Println(memory)
	checksum := calculateChecksum(memory)
	return checksum
}

func initMemory(line string) []int {
	digits := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		digits[i] = int(line[i] - '0')
	}
	memory := make([]int, 0, len(digits))
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		nr := -1
		if i%2 == 0 {
			nr = i / 2
		}
		for j := 0; j < digit; j++ {
			memory = append(memory, nr)
		}
	}
	return memory
}

func task2(lines []string) int {
	memory := initMemory(lines[0])
	//fmt.Println(memory)
	updateMemory2B(memory)
	//fmt.Println(memory)
	checksum := calculateChecksum(memory)
	return checksum
}

func updateMemory(memory []int) {
	nextEmpty := findNextEmpty(memory, -1)
	for i := len(memory) - 1; i >= 0; i-- {
		if nextEmpty > i {
			break
		}
		if memory[i] >= 0 {
			memory[nextEmpty] = memory[i]
			memory[i] = -1
			nextEmpty = findNextEmpty(memory, nextEmpty)
		}
	}
}

func findNextEmpty(memory []int, pos int) int {
	pos++
	for i := pos; i < len(memory); i++ {
		if memory[i] == -1 {
			return i
		}
	}
	return -1
}

func updateMemory2B(memory []int) {
	endPos := len(memory) - 1
	for {
		startPos, length := findLastInterval(memory, endPos)
		emptyStart := findFirstEmptyInterval(memory, length)
		if emptyStart >= 0 && emptyStart < startPos {
			copy(memory[emptyStart:emptyStart+length], memory[startPos:startPos+length])
			fillStart := max(startPos, emptyStart+length)
			for i := fillStart; i < startPos+length; i++ {
				memory[i] = -1
			}
		}
		endPos = startPos - 1
		if endPos < 0 {
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findFirstEmptyInterval(memory []int, length int) (start int) {
	insideEmpty := false
	emptyLen := 0
	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			if !insideEmpty {
				insideEmpty = true
				start = i
			}
			emptyLen++
			if emptyLen == length {
				return start
			}
			continue
		}
		insideEmpty = false
		emptyLen = 0
	}
	return -1
}

// findInterval returns the start and length of the interval of equal non-negative
// numbers starting at ending at startPos or before.
func findLastInterval(memory []int, startPos int) (start, length int) {
	end := -1
	for j := startPos; j >= 0; j-- {
		if memory[j] != -1 {
			end = j
			break
		}
	}
	if end == -1 {
		return 0, 0
	}
	value := memory[end]
	start = 0
	for j := end - 1; j >= 0; j-- {
		if memory[j] != value {
			start = j + 1
			break
		}
	}
	return start, end - start + 1
}

func calculateChecksum(memory []int) int {
	checksum := 0
	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			continue
		}
		checksum += memory[i] * i
	}
	return checksum
}
