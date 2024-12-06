package main

import (
	_ "embed"
	"fmt"
	"github.com/wesrobin/adventofc/aoc24"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

var testLines = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func main() {
	debug := false
	inp := input
	if debug {
		inp = testLines
	}

	lines := strings.Split(inp, "\n")

	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}

func p1(lines []string) int {
	re := regexp.MustCompile(`mul\(-?\d+,-?\d+\)`)

	var tot int
	for _, l := range lines {
		muls := re.FindAllString(l, -1)
		for _, m := range muls {
			tot += mul(m)
		}
	}

	return tot
}

func p2(lines []string) int {
	re := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\(-?\d+,-?\d+\))`)

	var tot int
	doMul := true
	for _, l := range lines {
		muls := re.FindAllString(l, -1)
		fmt.Println(muls)
		for _, m := range muls {
			if m == "don't()" {
				doMul = false
				continue
			} else if m == "do()" {
				doMul = true
				continue
			}
			if doMul {
				//fmt.Println(m)
				tot += mul(m)
			} else {
				//fmt.Println("ignoring", m)
			}
		}
	}

	return tot
}

func mul(s string) int {
	s = s[len("mul(") : len(s)-1]
	i1s, i2s, _ := strings.Cut(s, ",")
	//fmt.Println(i1s, "*", i2s)
	return aoc24.Atoi(i1s) * aoc24.Atoi(i2s)
}
