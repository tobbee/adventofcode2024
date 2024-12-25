package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	u "github.com/tobbee/adventofcode2024/utils"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines)
	require.Equal(t, 37327623, result)
}

func TestTask2(t *testing.T) {
	lines := []string{"1", "2", "3", "2024"}
	result := task2(lines)
	require.Equal(t, 23, result)
}