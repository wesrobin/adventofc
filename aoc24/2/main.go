package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wesrobin/adventofc/aoc24"
)

//go:embed input.txt
var inp string

func main() {
	lines := strings.Split(inp, "\n")

	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}

func p1(lines []string) int {
	var is [][]int
	for _, line := range lines {
		is = append(is, aoc24.MapSlice(strings.Fields(line), aoc24.Atoi))
	}

	var c int
	for _, iss := range is {
		safe := true
		var inc bool
		curr := iss[0]
		for i := 1; i < len(iss); i++ {
			if i == 1 {
				inc = iss[i] > curr
			}
			if !inc && (iss[i]-curr >= 0 || iss[i]-curr < -3) {
				safe = false
				break
			}
			if inc && (iss[i]-curr <= 0 || iss[i]-curr > 3) {
				safe = false
				break
			}
			curr = iss[i]
		}
		if safe {
			//fmt.Println(iss, safe)
			c++
		}
	}
	return c
}

func p2(lines []string) int {
	var is [][]int
	for _, line := range lines {
		is = append(is, aoc24.MapSlice(strings.Fields(line), aoc24.Atoi))
	}

	var c int
	for _, iss := range is {
		if valid(iss) {
			c++
		} else {
			for i := range iss {
				newISS := make([]int, len(iss))
				copy(newISS, iss)

				if valid(append(newISS[:i], newISS[i+1:]...)) {
					c++
					break
				}
			}
		}
	}
	return c
}

func valid(is []int) bool {
	validPos := map[int]bool{
		1: true,
		2: true,
		3: true,
	}
	validNeg := map[int]bool{
		-1: true,
		-2: true,
		-3: true,
	}
	for i := range len(is) - 1 {
		diff := is[i+1] - is[i]
		validPos[diff] = true
		validNeg[diff] = true
	}

	return len(validPos) == 3 || len(validNeg) == 3
}
