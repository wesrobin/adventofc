package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var testInput1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type locashe string

func part1(inp string) any {
	inst, lines, _ := strings.Cut(inp, "\n\n")
	m := make(map[locashe][2]locashe)
	for _, line := range strings.Split(lines, "\n") {
		l, i, _ := strings.Cut(line, " = ")
		m[locashe(l)] = [2]locashe{locashe(i[1:4]), locashe(i[6:9])}
	}

	var (
		i, count int
		curr     = locashe("AAA")
	)

	for curr != "ZZZ" {
		if inst[i] == 'L' {
			curr = m[curr][0]
		} else {
			curr = m[curr][1]
		}
		i++
		count++
		i %= len(inst)
	}
	return count
}

type node struct {
	l     string
	left  *node
	right *node
}

func printGraph(curr *node, level int, visited map[*node]bool) {
	if curr == nil {
		return
	}

	fmt.Printf("%*sNode %s\n", level*2, "", curr.l)
	visited[curr] = true

	if curr.left != nil {
		fmt.Printf("%*s -> Left %s\n", level*2, "", curr.left.l)
		if !visited[curr.left] {
			printGraph(curr.left, level+1, visited)
		} else {
			fmt.Printf("%*s -> (Already Visited)\n", level*2, "")
		}
	}
	if curr.right != nil {
		fmt.Printf("%*s -> Right %s\n", level*2, "", curr.right.l)
		if !visited[curr.right] {
			printGraph(curr.right, level+1, visited)
		} else {
			fmt.Printf("%*s -> (Already Visited)\n", level*2, "")
		}
	}
}

var testInput2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func part2(inp string) any {
	inst, lines, _ := strings.Cut(inp, "\n\n")
	seen := make(map[string]*node)
	var roots []*node
	for _, line := range strings.Split(lines, "\n") {
		l, i, _ := strings.Cut(line, " = ")
		var (
			base, left, right *node
		)
		if _, ok := seen[l]; ok {
			base = seen[l]
		} else {
			base = &node{l: l}
			seen[l] = base
		}
		if _, ok := seen[i[1:4]]; ok {
			left = seen[i[1:4]]
		} else {
			left = &node{l: i[1:4]}
			seen[i[1:4]] = left
		}
		if _, ok := seen[i[6:9]]; ok {
			right = seen[i[6:9]]
		} else {
			right = &node{l: i[6:9]}
			seen[i[6:9]] = right
		}
		base.left = left
		base.right = right
		if base.l[2] == 'A' {
			roots = append(roots, base)
		}
	}

	var paths []int
	for _, c := range roots {
		var path int
		curr := new(node)
		*curr = *c
		// First find the path to the Z
		for i := 0; ; i = (i + 1) % len(inst) {
			if curr.l[2] == 'Z' {
				break
			}
			path++
			curr = next(curr, inst[i])
		}
		paths = append(paths, path)
	}

	fmt.Println(paths)

	return multipleLCM(paths)
}

func next(curr *node, dir uint8) *node {
	if dir == 'L' {
		return curr.left
	}
	return curr.right
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(float64(a*b) / float64(gcd(a, b)))
}

func multipleLCM(ns []int) int {
	if len(ns) < 2 {
		panic("At least two numbers are required to find the LCM.")
	}

	res := lcm(ns[0], ns[1])

	for i := 2; i < len(ns); i++ {
		res = lcm(res, ns[i])
	}

	return res
}
