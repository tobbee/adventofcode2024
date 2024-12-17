package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	u "github.com/tobbee/adventofcode2024/utils"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines)
	require.Equal(t, "4,6,3,5,6,3,5,2,1,0", result)
}

func TestTask2(t *testing.T) {
	// Tests that the right value was found
	program := []int{2, 4, 1, 3, 7, 5, 0, 3, 4, 3, 1, 5, 5, 5, 3, 0}
	a := 236580836040301
	o := runProgram(program, a)
	result := joinInts(o)
	require.Equal(t, "2,4,1,3,7,5,0,3,4,3,1,5,5,5,3,0", result)
}

func TestTwoLast(t *testing.T) {
	program := []int{2, 4, 1, 3, 7, 5, 0, 3, 4, 3, 1, 5, 5, 5, 3, 0}
	for a := 0; a < (1 << 21); a++ {
		o := runProgram(program, a)
		result := joinInts(o)
		if result == "0" {
			fmt.Printf("0: %6b\n", a)
			t.Logf("a: %d", a)
		}
		if result == "3,0" {
			fmt.Printf("3, 0: %6b\n", a)
			t.Logf("a: %d", a)
		}
		if result == "5,3,0" {
			fmt.Printf("5,3,0: %9b\n", a)
			t.Logf("a: %d", a)
		}
		if result == "5,5,3,0" {
			fmt.Printf("5,5,3,0: %12b\n", a)
			t.Logf("a: %d", a)
		}
		if result == "5,5,5,3,0" {
			fmt.Printf("5,5,5,3,0: %12b\n", a)
			t.Logf("a: %d", a)
		}
		if result == "1,5,5,5,3,0" {
			fmt.Printf("1,5,5,5,3,0: %12b\n", a)
			t.Logf("a: %d", a)
		}
	}

}
