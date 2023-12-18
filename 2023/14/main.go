package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	t0 := time.Now()
	part1()
	fmt.Println(time.Since(t0))
	t0 = time.Now()
	part2()
	fmt.Println(time.Since(t0))
}

const testInput1 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

type column struct {
	base int
	num  int
}

func part1() {
	// fails:
	// 110805 too high
	inp := input

	lines := strings.Split(inp, "\n")
	weights := make([][]column, len(lines[0]))
	for y, li := range lines {
		for x, c := range li {
			// Do we want to check these separately? Probably not
			if y == 0 && c != '#' {
				weights[x] = append(weights[x], column{base: y})
			}
			if c == '#' {
				weights[x] = append(weights[x], column{base: y + 1})
			}
			if c == 'O' {
				weights[x][len(weights[x])-1].num++
			}
		}
	}

	var sum int
	for _, col := range weights {
		// fmt.Println(col)
		for _, c := range col {
			// fmt.Println(c)
			weight := len(lines) - c.base
			for i := 0; i < c.num; i++ {
				//	fmt.Print(weight-i, "+")
				sum += weight - i
			}
			// fmt.Println()
		}
	}
	fmt.Println(sum)
}

const testInput2 = `......
......
..#...
......
......`

func key(lines []string) string {
	return strings.Join(lines, "")
}

func part2() {
	// Check for cycles
	// Move north, rotate, move north etc
	// Can I use the 'cols' solution above to figure out the rotation?
	inp := input
	maxIter := 1000000000
	lines := strings.Split(inp, "\n")
	seen := make(map[string]int)
	var permutations [][]string
	var (
		t0 = time.Now()
		i  int
	)
	for ; i < maxIter; i++ {
		k := key(lines)
		if _, ok := seen[k]; ok {
			// We have a loop. Use the permutations cache to figure out the end state
			i = maxIter - (maxIter % i)
			break
		} else {
			seen[k] = i
			permutations = append(permutations, lines)
		}
		lines = north(lines)
		lines = rotate(lines)
		lines = north(lines)
		lines = rotate(lines)
		lines = north(lines)
		lines = rotate(lines)
		lines = north(lines)
		lines = rotate(lines)
		if i%10_000_000 == 0 {
			fmt.Println(i, time.Since(t0))
			t0 = time.Now()
		}
	}

	k := key(lines)
	fmt.Println(seen[k])           // start of cycle
	fmt.Println(len(permutations)) // end of cycle
	// cycle = (end -> start)
	fmt.Println((maxIter-seen[k])%(len(permutations)-seen[k]) + seen[k])
	printLines(permutations[(maxIter-seen[k])%(len(permutations)-seen[k])+seen[k]])
	lines = permutations[(maxIter-seen[k])%(len(permutations)-seen[k])+seen[k]]

	var sum int
	for y, l := range lines {
		for _, c := range l {
			if c == 'O' {
				sum += len(lines) - y
			}
		}
	}
	fmt.Println(sum)
}

func printLines(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
	fmt.Println()
}

func rotate(lines []string) []string {
	newLines := make([]string, len(lines[0]))
	for _, l := range lines {
		//fmt.Println(l)
		for x, c := range l {
			newLines[x] = string(c) + newLines[x]
		}
	}
	return newLines
}

func north(lines []string) []string {
	weights := make([][]column, len(lines[0]))

	for y, li := range lines {
		for x, c := range li {
			if y == 0 && c != '#' {
				weights[x] = append(weights[x], column{base: y})
			}
			if c == '#' {
				weights[x] = append(weights[x], column{base: y + 1})
			}
			if c == 'O' {
				weights[x][len(weights[x])-1].num++
			}
		}
	}
	newLines := make([]string, len(lines))
	for x := 0; x < len(lines[0]); x++ {
		var base, toAdd, weightIdx int
		if len(weights[x]) > weightIdx {
			base = weights[x][weightIdx].base
			toAdd = weights[x][weightIdx].num
			weightIdx++
		}
		for y := 0; y < len(lines); y++ {
			if y == base-1 {
				newLines[y] += "#"
				for j := y; j < base-1+toAdd; j++ {
					if j < base+toAdd-1 { // kill me
						y++
					}
					newLines[y] += "O"
				}
				if len(weights[x]) > weightIdx {
					base = weights[x][weightIdx].base
					toAdd = weights[x][weightIdx].num
					weightIdx++
				}
			} else if y == 0 {
				if toAdd > 0 {
					for j := y; j < base+toAdd; j++ {
						newLines[y] += "O"
						if j < base+toAdd-1 { // kill me
							y++
						}
					}
					if len(weights[x]) > weightIdx {
						base = weights[x][weightIdx].base
						toAdd = weights[x][weightIdx].num
						weightIdx++
					}
				} else {
					newLines[y] += "."
					if len(weights[x]) > weightIdx {
						base = weights[x][weightIdx].base
						toAdd = weights[x][weightIdx].num
						weightIdx++
					}
				}
			} else {
				newLines[y] += "."
			}
		}
	}
	return newLines
}

func eqSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
