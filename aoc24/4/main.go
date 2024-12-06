package main

import (
	_ "embed"
	"fmt"
	"github.com/wesrobin/adventofc/aoc24"
	"strings"
)

//go:embed input.txt
var input string

var testInp = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func main() {
	inp := input
	testing := false
	if testing {
		inp = testInp
	}

	lines := strings.Split(inp, "\n")
	var chs [][]string
	for _, l := range lines {
		chs = append(chs, strings.Split(l, ""))
	}

	fmt.Println(p1(chs))
	fmt.Println(p2(chs))
}

func p2(chs [][]string) int {
	var coords []aoc24.Coord2D
	for y, chss := range chs {
		for x, ch := range chss {
			if ch == "A" {
				coords = append(coords, aoc24.Coord2D{X: x, Y: y})
			}
		}
	}

	check := func(c aoc24.Coord2D) bool {
		validStrs := map[string]bool{
			"MMSS": true,
			"MSSM": true,
			"SSMM": true,
			"SMMS": true,
		}
		var str string
		for _, xC := range xCoords(c) {
			if xC.X < 0 || xC.X >= len(chs[0]) || xC.Y < 0 || xC.Y >= len(chs) {
				continue
			}
			str += chs[xC.Y][xC.X]
		}
		return validStrs[str]
	}

	var count int
	for _, c := range coords {
		if check(c) {
			//fmt.Println(c)
			count++
		}
	}

	return count
}

func xCoords(c aoc24.Coord2D) []aoc24.Coord2D {
	return []aoc24.Coord2D{
		aoc24.NextCoord(c, aoc24.NorthWest),
		aoc24.NextCoord(c, aoc24.SouthWest),
		aoc24.NextCoord(c, aoc24.SouthEast),
		aoc24.NextCoord(c, aoc24.NorthEast),
	}
}

var nextCh = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
	"S": "$", // Some end char
}

func p1(chs [][]string) int {
	type traversal struct {
		nextChar string
		currPos  aoc24.Coord2D
		dir      aoc24.Dir
		seen     []aoc24.Coord2D
	}
	var que []traversal

	step := func(t traversal) {
		if t.dir == 0 {
			cs, ds := nextCs(t.currPos)
			for i, next := range cs {
				if next.X < 0 || next.X >= len(chs[0]) || next.Y < 0 || next.Y >= len(chs) {
					continue
				}
				if chs[next.Y][next.X] == t.nextChar {
					que = append(que, traversal{
						nextChar: nextCh[t.nextChar],
						currPos:  next,
						dir:      ds[i],
						seen:     append(t.seen, next),
					})
				}
			}
		}
		nextCoord := aoc24.NextCoord(t.currPos, t.dir)
		if nextCoord.X < 0 || nextCoord.X >= len(chs[0]) || nextCoord.Y < 0 || nextCoord.Y >= len(chs) {
			return
		}
		if chs[nextCoord.Y][nextCoord.X] == t.nextChar {
			que = append(que, traversal{
				nextChar: nextCh[t.nextChar],
				currPos:  nextCoord,
				dir:      t.dir,
				seen:     append(t.seen, nextCoord),
			})
		}
	}

	for y, chss := range chs {
		for x, ch := range chss {
			if ch == "X" {
				tvs := traversal{
					nextChar: "M",
					currPos:  aoc24.Coord2D{X: x, Y: y},
					seen: []aoc24.Coord2D{
						{X: x, Y: y},
					},
				}
				que = append(que, tvs)
			}
		}
	}

	var count int
	for len(que) > 0 {
		//fmt.Println(que)
		curr := que[0]
		que = que[1:]
		if curr.nextChar == "$" {
			count++
			//fmt.Println(curr.seen, curr.dir.String())
		} else {
			step(curr)
		}
	}

	return count
}

func nextCs(c aoc24.Coord2D) ([]aoc24.Coord2D, []aoc24.Dir) {
	return []aoc24.Coord2D{
			aoc24.NextCoord(c, aoc24.North),
			aoc24.NextCoord(c, aoc24.South),
			aoc24.NextCoord(c, aoc24.West),
			aoc24.NextCoord(c, aoc24.East),
			aoc24.NextCoord(c, aoc24.NorthWest),
			aoc24.NextCoord(c, aoc24.SouthWest),
			aoc24.NextCoord(c, aoc24.SouthEast),
			aoc24.NextCoord(c, aoc24.NorthEast),
		}, []aoc24.Dir{
			aoc24.North,
			aoc24.South,
			aoc24.West,
			aoc24.East,
			aoc24.NorthWest,
			aoc24.SouthWest,
			aoc24.SouthEast,
			aoc24.NorthEast,
		}
}
