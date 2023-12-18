package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	aoc "github.com/wesrobin/adventofc/2023"
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

const testInput1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

const testInput2 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

const testInput3 = `#..#.##.###
.##..##.#.#
.....#..#.#
#..##....#.
....#####..
.##.#......
....#..#...
.##.##....#
#..##.###.#
.......#.##
....#..#.##
#..##.###.#
.##.##....#`

const testInput4 = `#......##..
.#.##.#..##
.#....#..##
.#....#.#..
####.###.##
.######....
###..######
..####..#..
..#..#..#..
##....##...
##....##...
########.##
##....##...`

func part1() {
	inp := input
	sections := strings.Split(inp, "\n\n")
	var sum int
	for _, s := range sections {
		lines := strings.Split(s, "\n")
		cols := make([]string, len(lines[0]))
		for _, li := range lines {
			for x, c := range li {
				cols[x] += string(c)
			}
		}
		if n, ok := try(lines); ok {
			sum += 100 * n
		} else if n, ok := try(cols); ok {
			sum += n
		}
	}
	fmt.Println(sum)
}

func printLines(cols []string, n int) {
	for i, li := range cols {
		fmt.Println(li)
		if i == n {
			fmt.Println("-----")
		}
	}
}

func try(lines []string) (int, bool) {
	s := aoc.NewStack[string]()
	for i, li := range lines {
		if sli, ok := s.Top(); ok && sli == li {
			// found
			found := true
			for j, k := i, 0; j < len(lines) && k < len(*s); j++ {
				if l, ok := s.Peek(k); !ok || l != lines[j] {
					found = false
				}
				k++
			}
			if found {
				return i, true
			}
		}
		s.Push(li)
	}
	return 0, false
}

func part2() {
	// Fails:
	// 38733 too low
	// 42691 too low
	inp := input
	sections := strings.Split(inp, "\n\n")
	var sum int
	for _, s := range sections {
		lines := strings.Split(s, "\n")
		cols := make([]string, len(lines[0]))
		for _, li := range lines {
			for x, c := range li {
				cols[x] += string(c)
			}
		}
		if n, ok := try2(lines); ok {
			sum += 100 * n
		} else if n, ok := try2(cols); ok {
			sum += n
		}
	}
	fmt.Println(sum)

}

func try2(lines []string) (int, bool) {
	s := aoc.NewStack[string]()
	for i, li := range lines {
		sli, ok := s.Top()
		if ok {
			eq, sm := smudgeEq(sli, li, true)
			if eq {
				// found
				found := true
				smChecked := sm
				for j, k := i+1, 1; j < len(lines) && k < len(*s); j++ {
					l, ok := s.Peek(k)
					if !ok {
						found = false
						break
					}
					if smChecked {
						if l != lines[j] {
							found = false
							break
						}
					} else {
						eq, sm := smudgeEq(l, lines[j], true)
						if !eq {
							found = false
							break
						}
						smChecked = sm
					}
					k++
				}
				if found && smChecked {
					return i, true
				}
			}
		}
		s.Push(li)
	}
	return 0, false

}

func smudgeEq(sli, li string, checkSmudge bool) (bool, bool) {
	if sli == li {
		return true, false
	}
	if !checkSmudge {
		return false, false
	}
	var diff int
	for i := range sli {
		if sli[i] != li[i] {
			diff++
		}
	}
	return diff == 1, true
	// if diff > 1 {
	// 	return false
	// } else if diff == 0 {
	// 	return true
	// }
	//
	// for i := 0; i < len(sli); i++ {
	// 	if sli[i] != li[i] {
	//
	// 	}
	// }
	//
	// return false
}
