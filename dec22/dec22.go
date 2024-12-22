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

type generator struct {
	lastSequence sequence
	start        int
	last         int
	lastDigit    int
	bananas      map[sequence]int
}

func newGenerator(start int) *generator {
	return &generator{
		start:     start,
		last:      start,
		lastDigit: start % 10,
		bananas:   make(map[sequence]int),
	}
}

func (g *generator) next() int {
	secret := g.last
	n := secret << 6                 // multiply by 64
	secret = (secret ^ n) & 0xffffff // xor and keep 24 bits
	n = secret >> 5                  // divide by 32
	secret = (secret ^ n) & 0xffffff // xor and keep 24 bits
	n = secret << 11                 // multiply by 2048
	secret = (secret ^ n) & 0xffffff // xor and keep 24 bits
	g.last = secret
	return secret
}

func (g *generator) nextSequence() sequence {
	nr := g.next()
	digit := nr % 10
	diff := digit - g.lastDigit
	g.lastDigit = digit
	g.lastSequence.add(diff)
	return g.lastSequence
}

// findMatchingSequence searches for first  sequence of 4 numbers that matches the given sequence.
// If a match is found, the function returns the price.
// If no match is found, the function returns -1.
func (g *generator) findMatchingSequence(seq sequence) (price int) {
	bananas, ok := g.bananas[seq]
	if ok {
		return bananas
	}
	return -1
}

func (g *generator) genMatches() {
	for i := 0; i < 2000; i++ {
		seq := g.nextSequence()
		if i >= 3 {
			if _, ok := g.bananas[seq]; !ok {
				g.bananas[seq] = g.lastDigit
			}
		}
	}
}

func task1(lines []string) int {
	sum := 0
	for _, line := range lines {
		nr := u.Atoi(line)
		g := newGenerator(nr)
		for i := 0; i < 2000; i++ {
			nr = g.next()
		}
		fmt.Println(nr)
		sum += nr
	}
	return sum
}

type sequence [4]int

func (s *sequence) add(nr int) {
	s[0] = s[1]
	s[1] = s[2]
	s[2] = s[3]
	s[3] = nr
}

func task2(lines []string) int {
	seqs := u.CreateSet[sequence]()
	nrs := make([]int, 0, len(lines))
	for _, line := range lines {
		nr := u.Atoi(line)
		nrs = append(nrs, nr)
	}
	generators := make([]*generator, 0, len(nrs))
	fmt.Println("nr monkeys", len(nrs))
	for _, nr := range nrs {
		g := newGenerator(nr)
		g.genMatches()
		for seq := range g.bananas {
			seqs.Add(seq)
		}
		generators = append(generators, g)
	}
	fmt.Println("gathered sequences", seqs.Size())
	best := 0
	sNr := 0
	nrSeqs := seqs.Size()
	for seq := range seqs {
		nrBananas := 0
		for _, g := range generators {
			price := g.findMatchingSequence(seq)
			if price > 0 {
				nrBananas += price
			}
		}
		if nrBananas > best {
			best = nrBananas
			fmt.Println("seq", sNr, seq, "of", nrSeqs, "best", best)
		}
		sNr++
	}

	return best
}
