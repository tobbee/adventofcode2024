package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"strings"

	"github.com/Eyevinn/mp4ff/bits"
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

type machine struct {
	A, B, C int
	ptr     int
	output  []int
}

func (m *machine) exec(instr, op int) {
	switch instr {
	case 0: // adv
		m.A = m.A >> m.combo(op)
	case 1: // bxl
		m.B = m.B ^ op
	case 2: // bst
		m.B = m.combo(op) & 0x7
	case 3: // jnz
		if m.A == 0 {
			m.ptr += 2
			return
		}
		m.ptr = op
		return
	case 4: // bxc
		m.B = m.B ^ m.C
	case 5: // out
		val := m.combo(op) & 0x7
		m.output = append(m.output, val)
	case 6: // bdv
		m.B = m.A >> m.combo(op)
	case 7: // cdv
		m.C = m.A >> m.combo(op)

	}
	m.ptr += 2
}

func (m *machine) combo(op int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return m.A
	case 5:
		return m.B
	case 6:
		return m.C
	case 7:
		panic("Invalid combo 7")
	default:
		panic("Invalid combo > 7")
	}
}

func (m *machine) execute(program []int) {
	for {
		if m.ptr >= len(program) {
			break
		}
		instr := program[m.ptr]
		op := program[m.ptr+1]
		m.exec(instr, op)
	}
}

func joinInts(ints []int) string {
	var s []string
	for _, i := range ints {
		s = append(s, fmt.Sprintf("%d", i))
	}
	return strings.Join(s, ",")
}

func task1(lines []string) string {
	var A int
	var program []int
	for i, line := range lines {
		nrs := u.SplitToInts(line)
		switch i {
		case 0:
			A = nrs[0]
		case 4:
			program = nrs
		}
	}
	output := runProgram(program, A)
	return joinInts(output)
}

func runProgram(program []int, a int) []int {
	m := machine{A: a, B: 0, C: 0}
	m.execute(program)
	return m.output
}

func task2(lines []string) int {
	var program []int
	for i, line := range lines {
		nrs := u.SplitToInts(line)
		switch i {
		case 4:
			program = nrs
		}
	}

	// Find the slice of 3-input bits that generate the digits
	// The machine eats 3 bits per digit.
	digitBytes := make([]byte, 0)
	digitPos := len(program) - 1
	digitBytes, ok := addNextBits(program, digitPos, digitBytes)
	if !ok {
		panic("Could not find digit bits")
	}
	for i, b := range digitBytes {
		fmt.Printf("Digit bits %2d: %03b\n", i, b)
	}
	nr := makeNumber(digitBytes)
	return nr
}

// addNextBits recursively tries to add bits to the digitBytes slice growing
// from the end of the program. If the full slice is correct, it will return true.
func addNextBits(program []int, digitPos int, digitBytes []byte) ([]byte, bool) {
	if digitPos < 0 {
		return digitBytes, true
	}
	nrDigitBytes := len(digitBytes)
	digit := program[digitPos]
	nextDigitPos := digitPos - 1
	start := 0
	if digitPos == len(program)-1 && program[digitPos] == 0 {
		start = 1
	}
	for i := start; i <= 0x7; i++ {
		digitBytes = append(digitBytes, byte(i))
		nr := makeNumber(digitBytes)
		m := machine{A: int(nr), B: 0, C: 0}
		m.execute(program)

		if len(m.output) == len(digitBytes) && m.output[0] == digit {
			var ok bool
			digitBytes, ok = addNextBits(program, nextDigitPos, digitBytes)
			if ok {
				return digitBytes, true
			}
		}
		digitBytes = digitBytes[:nrDigitBytes]
	}
	return digitBytes, false
}

func makeNumber(digitBytes []byte) int {
	w := bytes.Buffer{}
	bw := bits.NewWriter(&w)
	bw.Write(0, 64-3*(len(digitBytes)))
	for i := 0; i < len(digitBytes); i++ {
		bw.Write(uint(digitBytes[i]), 3)
	}
	data := w.Bytes()
	value := binary.BigEndian.Uint64(data)
	return int(value)
}
