package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

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
	values, ops := parseInput(lines)
	result := calc(values, ops)
	return result
}

func calc(values map[string]int, ops []op) int {
	unused := u.CreateSet[op]()
	for _, o := range ops {
		unused.Add(o)
	}
	for {
		used := u.CreateSet[op]()
		for o := range unused {
			if values[o.in1] != -1 && values[o.in2] != -1 {
				switch o.f {
				case "AND":
					values[o.out] = values[o.in1] & values[o.in2]
				case "OR":
					values[o.out] = values[o.in1] | values[o.in2]
				case "XOR":
					values[o.out] = values[o.in1] ^ values[o.in2]
				}
				used.Add(o)
			}
		}
		unused.Subtract(used)
		if len(unused) == 0 {
			break
		}
	}
	var zs []string
	for v := range values {
		if strings.HasPrefix(v, "z") {
			zs = append(zs, v)
		}
	}
	sort.Strings(zs)
	nr := 0
	for i := len(zs) - 1; i >= 0; i-- {
		nr = nr<<1 | values[zs[i]]
	}
	return nr
}

type conf struct {
	in1, in2, f string
}

func confFrom(a, b, op string) conf {
	if a < b {
		return conf{a, b, op}
	}
	return conf{b, a, op}
}

type pair struct {
	a, b string
}

func toValues(x, y int) map[string]int {
	m := map[string]int{}
	for i := 0; i < 8; i++ {
		m[fmt.Sprintf("x%d", i)] = x >> uint(7-i) & 1
		m[fmt.Sprintf("y%d", i)] = y >> uint(7-i) & 1
	}
	return m
}

func swapPairs(opsMap map[conf]string, swaps []pair) {
	for _, swap := range swaps {
		keys := make([]conf, 0, len(opsMap))
		for k, v := range opsMap {
			if v == swap.a {
				keys = append(keys, k)
			}
			if v == swap.b {
				keys = append(keys, k)
			}
		}
		if opsMap[keys[0]] == swap.a {
			opsMap[keys[0]] = swap.b
			opsMap[keys[1]] = swap.a
		} else {
			opsMap[keys[0]] = swap.a
			opsMap[keys[1]] = swap.b
		}
	}
}

// Full adder for each bit and carry
// There are 44 bits in the inputs
// z0 = x0 ^ y0
// c0 = x0 & y0
// z1 = (x1 ^ y1) ^ c0
// c1 = (x1 & y1) | (c0 & (x1 ^ y1))
// ...
// z44 = (x44 ^ y44) ^ c43
// c44 = (x44 & y44) | (c43 & (x44 ^ y44))
// z45 = c44
func task2(lines []string) string {
	_, ops := parseInput(lines)
	opsMap := make(map[conf]string)
	for _, o := range ops {
		opsMap[conf{o.in1, o.in2, o.f}] = o.out
	}
	// The swaps are manually determined by fixing
	// the process every time it fails.
	// At the end, there should be four stops
	// according to the description of the task.
	swaps := []pair{{"qjj", "gjc"}, {"wmp", "z17"},
		{"z26", "gvm"}, {"z39", "qsb"}}
	swapPairs(opsMap, swaps)
	if opsMap[conf{"x00", "y00", "XOR"}] != "z00" {
		fmt.Println("failed for z00")
	}
	delete(opsMap, conf{"x00", "y00", "XOR"})
	carry := opsMap[conf{"x00", "y00", "AND"}]
	if carry == "" {
		fmt.Println("failed for c00")
	}
	fmt.Printf("%02d Carry = %s\n", 0, carry)
	delete(opsMap, conf{"x00", "y00", "AND"})
	for i := 1; i < 44; i++ {
		x := fmt.Sprintf("x%02d", i)
		y := fmt.Sprintf("y%02d", i)
		z := fmt.Sprintf("z%02d", i)
		xXorY := opsMap[conf{x, y, "XOR"}]
		xAndY := opsMap[conf{x, y, "AND"}]
		if xXorY == "" || xAndY == "" {
			fmt.Println("failed for xXorY, xAndY", z, i)
			break
		}
		delete(opsMap, conf{x, y, "XOR"})
		delete(opsMap, conf{x, y, "AND"})
		cf := confFrom(xXorY, carry, "XOR")
		v := opsMap[cf]
		if v != z {
			fmt.Printf("failed for z%02d, x^y^c=%s, x^y=%s, c=%s\n",
				i, v, xXorY, carry)
			break
		}
		delete(opsMap, cf)
		// carry1 = (x1&y1) | (c0&(x1^y1))
		// carry1 = xAndY | (carry & xXorY)
		// carry1 = xAndY | rh
		rhConf := confFrom(carry, xXorY, "AND")
		rc := opsMap[rhConf]
		if rc == "" {
			fmt.Println("failed for rh", i)
			break
		}
		delete(opsMap, rhConf)
		lc := confFrom(xAndY, rc, "OR")
		newCarry := opsMap[lc]
		if newCarry == "" {
			fmt.Println("failed for newCarry", i)
		}
		delete(opsMap, lc)

		fmt.Printf("%02d: %s <- %s AND %s, %s <- %s XOR %s Carry = %s <- %s OR ( %s AND %s) \n",
			i, xAndY, x, y, xXorY, x, y, newCarry, xAndY, xXorY, carry)
		carry = newCarry
	}
	names := make([]string, 0, 8)
	for _, p := range swaps {
		names = append(names, p.a, p.b)
	}
	sort.Strings(names)
	answer := strings.Join(names, ",")
	return answer
}

type op struct {
	in1, in2, f, out string
}

func (o op) String() string {
	return fmt.Sprintf("%s %s %s -> %s", o.in1, o.f, o.in2, o.out)
}

func parseInput(lines []string) (map[string]int, []op) {
	m := make(map[string]int)
	var ops []op
	before := true
	for _, line := range lines {
		if line == "" {
			before = false
			continue
		}
		if before {
			parts := strings.Split(line, ": ")
			k := parts[0]
			v := u.Atoi(parts[1])
			m[k] = v
			continue
		}
		o := op{}
		parts := strings.Split(line, " -> ")
		o.out = parts[1]
		parts = strings.Split(parts[0], " ")
		if parts[2] < parts[0] {
			// Sort inputs in alphabetical order
			parts[0], parts[2] = parts[2], parts[0]
		}
		o.in1 = parts[0]
		o.f = parts[1]
		o.in2 = parts[2]
		ops = append(ops, o)
		m[o.out] = -1
		if _, ok := m[o.in1]; !ok {
			m[o.in1] = -1
		}
		if _, ok := m[o.in2]; !ok {
			m[o.in2] = -1
		}
	}
	sort.Slice(ops, func(i, j int) bool {
		if ops[i].in1 < ops[j].in1 {
			return true
		}
		if ops[i].in1 > ops[j].in1 {
			return false
		}
		if ops[i].in2 < ops[j].in2 {
			return true
		}
		if ops[i].in2 > ops[j].in2 {
			return false
		}
		return ops[i].f < ops[j].f
	})
	return m, ops
}
