package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	aoc "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var input string

func main() {
	inp := input
	part1(inp)
	part2(inp)
}

const testInput1 = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

/**
[false false true false false true false false true false]
[false false false true false false false true false false]

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
*/

func part1(inp string) {
	var grid [][]string
	for _, l := range strings.Split(inp, "\n") {
		lSplit := strings.Split(l, "")
		grid = append(grid, lSplit)
	}
	grid = grow(grid)

	var stars []aoc.Coord2D
	for y, l := range grid {
		for x, c := range l {
			if c == "#" {
				stars = append(stars, aoc.Coord2D{X: x, Y: y})
			}
		}
	}
	var sum int
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			l := int(math.Abs(float64(stars[j].X-stars[i].X)) + math.Abs(float64(stars[j].Y-stars[i].Y)))
			sum += l
		}
	}
	fmt.Println(sum)
}

func grow(grid [][]string) [][]string {
	vertEmpty := make([]bool, len(grid))
	for i := range vertEmpty {
		vertEmpty[i] = true
	}
	var horizEmpty []bool
	var vertEmptyCount int
	for _, l := range grid {
		lineEmpty := true
		for x, c := range l {
			if lineEmpty && c != "." {
				lineEmpty = false
			}
			if vertEmpty[x] && c != "." {
				vertEmpty[x] = false
			}
		}
		horizEmpty = append(horizEmpty, lineEmpty)
	}

	for _, v := range vertEmpty {
		if v {
			vertEmptyCount++
		}
	}

	horizLen := len(grid[0]) + vertEmptyCount
	var newGrid [][]string
	var yCount int
	for y, l := range grid {
		if horizEmpty[y] {
			yCount++
			newGrid = append(newGrid, strings.Split(strings.Repeat(".", horizLen), ""))
		}
		newGrid = append(newGrid, []string{})
		for x, c := range l {
			if vertEmpty[x] {
				newGrid[y+yCount] = append(newGrid[y+yCount], ".")
			}
			newGrid[y+yCount] = append(newGrid[y+yCount], c)
		}
	}
	return newGrid
}

func part2(inp string) {
	var grid [][]string
	for _, l := range strings.Split(inp, "\n") {
		lSplit := strings.Split(l, "")
		grid = append(grid, lSplit)
	}

	vertEmpty := make([]bool, len(grid[0]))
	for i := range vertEmpty {
		vertEmpty[i] = true
	}
	var horizEmpty []bool
	for _, l := range grid {
		lineEmpty := true
		for x, c := range l {
			if lineEmpty && c != "." {
				lineEmpty = false
			}
			if vertEmpty[x] && c != "." {
				vertEmpty[x] = false
			}
		}
		horizEmpty = append(horizEmpty, lineEmpty)
	}

	growAmt := 999999
	vertEmptyCount := make([]int, len(grid[0]))
	horizEmptyCount := make([]int, len(grid))
	for i, e := range vertEmpty {
		if e {
			if i == 0 {
				vertEmptyCount[i] = growAmt
			} else {
				vertEmptyCount[i] = vertEmptyCount[i-1] + growAmt
			}
		} else {
			if i == 0 {
				vertEmptyCount[i] = 0
			} else {
				vertEmptyCount[i] = vertEmptyCount[i-1]
			}
		}
	}
	for i, e := range horizEmpty {
		if e {
			if i == 0 {
				horizEmptyCount[i] = growAmt
			} else {
				horizEmptyCount[i] = horizEmptyCount[i-1] + growAmt
			}
		} else {
			if i == 0 {
				horizEmptyCount[i] = 0
			} else {
				horizEmptyCount[i] = horizEmptyCount[i-1]
			}
		}
	}

	var stars []aoc.Coord2D
	for y, l := range grid {
		for x, c := range l {
			if c == "#" {
				stars = append(stars, grownCoord(aoc.Coord2D{X: x, Y: y}, horizEmptyCount[y], vertEmptyCount[x]))
			}
		}
	}
	var sum int
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			l := int(math.Abs(float64(stars[j].X-stars[i].X)) + math.Abs(float64(stars[j].Y-stars[i].Y)))
			sum += l
		}
	}
	fmt.Println(sum)
}

func grownCoord(a aoc.Coord2D, horizEmptyCount, vertEmptyCount int) aoc.Coord2D {
	return aoc.Coord2D{X: a.X + vertEmptyCount, Y: a.Y + horizEmptyCount}
}
