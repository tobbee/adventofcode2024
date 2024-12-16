package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	u "github.com/tobbee/adventofcode2024/utils"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines)
	require.Equal(t, 7036, result)
}

func TestTask1b(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput2")
	result := task1(lines)
	require.Equal(t, 11048, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task2(lines)
	require.Equal(t, 45, result)
}

func TestTask2b(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput2")
	result := task2(lines)
	require.Equal(t, 64, result)
}
