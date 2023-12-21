package main

import (
	_ "embed"
	"fmt"
	"strings"

	aoc "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var input string

func main() {
	// part1()
	part2()
}

const testInput1 = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func part1() {
	inp := input

	var sum int
	for _, p := range strings.Split(inp, ",") {
		sum += hash(p)
	}

	fmt.Println(sum)
}

func hash(inp string) int {
	var h int
	for _, c := range inp {
		if c == '\n' {
			continue
		}
		h += int(c)
		h *= 17
		h %= 256
	}
	return h
}

type box struct {
	lenses map[string]int
	queue  []string
}

func (b *box) add(label string, value int) {
	if b.lenses == nil {
		b.lenses = make(map[string]int)
	}
	if _, ok := b.lenses[label]; ok {
		b.lenses[label] = value
	} else {
		b.lenses[label] = value
		b.queue = append(b.queue, label)
	}
}

func (b *box) remove(label string) {
	if _, ok := b.lenses[label]; !ok {
		return
	}
	delete(b.lenses, label)
	var i int
	for i = range b.queue {
		if b.queue[i] == label {
			break
		}
	}
	b.queue = append(b.queue[:i], b.queue[i+1:]...)
}

func (b *box) isEmpty() bool {
	return len(b.queue) == 0
}

func part2() {
	inp := input

	boxes := make([]box, 256)
	for _, instr := range strings.Split(inp, ",") {
		if instr == "\n" {
			continue
		}
		if strings.Contains(instr, "-") {
			label, _, _ := strings.Cut(instr, "-")
			h := hash(label)
			boxes[h].remove(label)
		} else {
			label, val, _ := strings.Cut(instr, "=")
			h := hash(label)
			boxes[h].add(label, aoc.Atoi(val))
		}
	}

	var sum int
	for i, b := range boxes {
		if !b.isEmpty() {
			for j, label := range b.queue {
				boxVal := (i + 1) * (j + 1) * b.lenses[label]
				sum += boxVal
			}
		}
	}
	fmt.Println(sum)

}
