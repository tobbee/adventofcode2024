package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	u "github.com/tobbee/adventofcode2024/utils"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task(lines, 2)
	require.Equal(t, 126384, result)
}

func TestAllShortest(t *testing.T) {
	cases := []struct {
		start         string
		end           string
		kMap          keyMap
		iMap          reverseMap
		wantedLength  int
		wantedNrPaths int
		wantedPaths   []string
	}{
		{start: "A", end: "^", kMap: dirGrid, iMap: dirGridReverse, wantedLength: 2, wantedNrPaths: 1,
			wantedPaths: []string{"v<<A", "<v<A"}},
		{start: "A", end: "<", kMap: dirGrid, iMap: dirGridReverse, wantedLength: 4, wantedNrPaths: 2,
			wantedPaths: []string{"<A"}},
		{start: "A", end: "7", kMap: numGrid, iMap: numGridReverse, wantedLength: 6, wantedNrPaths: 9,
			wantedPaths: []string{
				"^^^<<A", "^^<^<A", "^^<<^A",
				"^<^^<A", "^<^<^A", "^<<^^A",
				"<^^^<A", "<^^<^A", "<^<^^A"}},
		{start: "5", end: "9", kMap: numGrid, iMap: numGridReverse, wantedLength: 3, wantedNrPaths: 2,
			wantedPaths: []string{">^A", "^>A"}},
		{start: "5", end: "5", kMap: numGrid, iMap: numGridReverse, wantedLength: 1, wantedNrPaths: 1,
			wantedPaths: []string{"A"}},
		{start: "5", end: "8", kMap: numGrid, iMap: numGridReverse, wantedLength: 2, wantedNrPaths: 1,
			wantedPaths: []string{"^A"}},
		{start: "1", end: "0", kMap: numGrid, iMap: numGridReverse, wantedLength: 3, wantedNrPaths: 1},
	}
	for _, c := range cases {
		paths, length := findAllShortest(c.start, c.end, c.kMap, c.iMap, 0)
		require.Equal(t, c.wantedLength, length)
		require.Equal(t, c.wantedNrPaths, len(paths), "nr paths not matching")
		if len(c.wantedPaths) > 0 {
			require.Equal(t, c.wantedPaths, paths)
		}
	}
}
