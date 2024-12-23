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
	conns := parseConnections(lines)
	//fmt.Println(conns)
	lans := findAllLan3s(conns)
	nrWithT := 0
	for lan := range lans {
		for _, cpu := range lan {
			if cpu[0] == 't' {
				nrWithT++
				break
			}
		}
	}
	return nrWithT
}

func parseConnections(lines []string) map[string]u.Set[string] {
	conns := make(map[string]u.Set[string])
	for _, line := range lines {
		part0 := line[:2]
		part1 := line[3:5]
		if _, ok := conns[part0]; !ok {
			conns[part0] = u.CreateSet[string]()
		}
		conns[part0].Add(part1)
		if _, ok := conns[part1]; !ok {
			conns[part1] = u.CreateSet[string]()
		}
		conns[part1].Add(part0)
	}
	return conns
}

type lan3 [3]string

func (l *lan3) sort() {
	if l[0] > l[1] {
		l[0], l[1] = l[1], l[0]
	}
	if l[0] > l[2] {
		l[0], l[2] = l[2], l[0]
	}
	if l[1] > l[2] {
		l[1], l[2] = l[2], l[1]
	}
}

func findAllLan3s(conns map[string]u.Set[string]) u.Set[lan3] {
	lans := u.CreateSet[lan3]()
	cpus := make([]string, 0, len(conns))
	lanMembers := u.CreateSet[string]()
	for cpu := range conns {
		lanMembers.Add(cpu)
	}

	for cpu := range conns {
		cpus = append(cpus, cpu)
	}
	sort.Strings(cpus)
	for _, cpu0 := range cpus {
		conns0 := conns[cpu0]
		for cpu1 := range conns0 {
			for cpu2 := range conns[cpu1] {
				if conns0.Contains(cpu2) {
					lan := lan3{cpu0, cpu1, cpu2}
					lan.sort()
					lans.Add(lan)
				}
			}
		}
	}
	return lans
}

func task2(lines []string) string {
	conns := parseConnections(lines)
	cliques := findCliques(conns)
	maxLen := 0
	maxClique := ""
	for clique := range cliques {
		if len(clique) > maxLen {
			maxLen = len(clique)
			maxClique = clique
		}
	}
	parts := strings.Split(maxClique, ",")
	sort.Strings(parts)
	return strings.Join(parts, ",")
}

// findCliques finds all maximal cliques in a graph using the Bron-Kerbosch algorithm.
// https://en.wikipedia.org/wiki/Bron-Kerbosch_algorithm
func findCliques(conns map[string]u.Set[string]) u.Set[string] {
	p := u.CreateSet[string]()
	for v := range conns {
		p.Add(v)
	}
	r := u.CreateSet[string]()
	x := u.CreateSet[string]()
	cliques := u.CreateSet[string]()
	vertices := orderByNrOfNeighbors(conns)
	for _, v := range vertices {
		neighbors := conns[v]
		rx := r.Clone()
		rx.Add(v)
		px := p.Clone()
		px.Intersect(neighbors)
		xx := x.Clone()
		xx.Intersect(neighbors)
		findCliquesPivot(conns, rx, px, xx, cliques)
		p.Remove(v)
		x.Add(v)
	}
	return cliques
}

func orderByNrOfNeighbors(conns map[string]u.Set[string]) []string {
	vertices := make([]string, 0, len(conns))
	for v := range conns {
		vertices = append(vertices, v)
	}
	sort.Slice(vertices, func(i, j int) bool {
		return conns[vertices[i]].Size() < conns[vertices[j]].Size()
	})
	return vertices
}

func findCliquesPivot(conns map[string]u.Set[string], r, p, x u.Set[string], cliques u.Set[string]) {
	if p.Size() == 0 && x.Size() == 0 {
		sv := sort.StringSlice(r.Values())
		cliques.Add(strings.Join(sv, ","))
		return
	}
	px := p.Clone()
	px.Extend(x)
	vertices := px.Values()
	for _, u := range vertices {
		pm := p.Clone()
		pm.Subtract(conns[u])
		for v := range pm {
			neighbors := conns[v]
			rx := r.Clone()
			rx.Add(v)
			px := p.Clone()
			px.Intersect(neighbors)
			xx := x.Clone()
			xx.Intersect(neighbors)
			findCliquesPivot(conns, rx, px, xx, cliques)
			p.Remove(v)
			x.Add(v)
		}
	}
}
