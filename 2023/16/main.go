package main

import (
	_ "embed"
	"fmt"
	"strings"

	aoc "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var inpoot string

func input() string {
	return strings.TrimSpace(inpoot)
}

func main() {
	part1()
	part2() // 8755 h,
}

const testInput1 = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

type beamNode struct {
	pos aoc.Coord2D
	dir aoc.Dir
}

func (b beamNode) string() string {
	var d string
	switch b.dir {
	case aoc.Up:
		d = "up"
	case aoc.Down:
		d = "down"
	case aoc.Left:
		d = "left"
	case aoc.Right:
		d = "right"
	}
	return fmt.Sprintf("%v - %s", b.pos, d)
}

func (b beamNode) nextWithDir(d aoc.Dir) beamNode {
	var c aoc.Coord2D
	if d == aoc.Up {
		c = aoc.Coord2D{X: b.pos.X, Y: b.pos.Y - 1}
	} else if d == aoc.Down {
		c = aoc.Coord2D{X: b.pos.X, Y: b.pos.Y + 1}
	} else if d == aoc.Left {
		c = aoc.Coord2D{X: b.pos.X - 1, Y: b.pos.Y}
	} else { // Right
		c = aoc.Coord2D{X: b.pos.X + 1, Y: b.pos.Y}
	}
	return beamNode{pos: c, dir: d}

}

func nextNode(n beamNode, m mirror) []beamNode {
	var next []beamNode
	if a, b, ok := m.split(n.dir); ok {
		next = append(next, n.nextWithDir(a), n.nextWithDir(b))
	} else {
		d := m.pass(n.dir)
		next = append(next, n.nextWithDir(d))
	}
	return next
}

type mirror rune

const (
	ns   mirror = '|'
	we   mirror = '-'
	nesw mirror = '/'
	nwse mirror = '\\'
	dot  mirror = '.'
)

func (m mirror) split(d aoc.Dir) (aoc.Dir, aoc.Dir, bool) {
	switch d {
	case aoc.Up:
		return aoc.Left, aoc.Right, m == we
	case aoc.Down:
		return aoc.Left, aoc.Right, m == we
	case aoc.Left:
		return aoc.Up, aoc.Down, m == ns
	case aoc.Right:
		return aoc.Up, aoc.Down, m == ns
	}
	return d, d, false
}

func (m mirror) pass(d aoc.Dir) aoc.Dir {
	switch d {
	case aoc.Up:
		if m == nwse {
			return aoc.Left
		} else if m == nesw {
			return aoc.Right
		}

	case aoc.Down:
		if m == nwse {
			return aoc.Right
		} else if m == nesw {
			return aoc.Left
		}

	case aoc.Left:
		if m == nwse {
			return aoc.Up
		} else if m == nesw {
			return aoc.Down
		}

	case aoc.Right:
		if m == nwse {
			return aoc.Down
		} else if m == nesw {
			return aoc.Up
		}

	}
	return d
}

func part1() {
	// inp := testInput1
	inp := input()
	lines := strings.Split(inp, "\n")
	grid := make(map[aoc.Coord2D]mirror)

	for y, l := range lines {
		for x, c := range l {
			coord := aoc.Coord2D{X: x, Y: y}
			grid[coord] = mirror(c)
		}
	}

	xLen = len(lines[0])
	yLen = len(lines)
	tot := maxDepthBFS(beamNode{pos: aoc.Coord2D{X: 0, Y: 0}, dir: aoc.Right}, grid)

	fmt.Println(tot)
}

func part2() {
	// inp := testInput1
	inp := input()
	lines := strings.Split(inp, "\n")
	grid := make(map[aoc.Coord2D]mirror)
	for y, l := range lines {
		for x, c := range l {
			grid[aoc.Coord2D{X: x, Y: y}] = mirror(c)
		}
	}
	xLen = len(lines[0])
	yLen = len(lines)

	var longest int

	for y := range lines {
		longest = max(longest, maxDepthBFS(beamNode{pos: aoc.Coord2D{X: 0, Y: y}, dir: aoc.Right}, grid))
		longest = max(longest, maxDepthBFS(beamNode{pos: aoc.Coord2D{X: xLen - 1, Y: y}, dir: aoc.Left}, grid))
	}

	for x := range lines[0] {
		longest = max(longest, maxDepthBFS(beamNode{pos: aoc.Coord2D{X: x, Y: 0}, dir: aoc.Down}, grid))
		longest = max(longest, maxDepthBFS(beamNode{pos: aoc.Coord2D{X: x, Y: yLen - 1}, dir: aoc.Up}, grid))
	}

	fmt.Println(longest)
}

var (
	xLen, yLen int
)

func maxDepthBFS(start beamNode, grid map[aoc.Coord2D]mirror) int {
	// fmt.Println(start.string())
	beams := []beamNode{start}
	path := make(map[aoc.Coord2D]bool)
	seen := make(map[beamNode]bool)

	for len(beams) > 0 {
		curr := beams[0]
		beams = beams[1:]
		if seen[curr] {
			continue
		}

		seen[curr] = true
		path[curr.pos] = true
		m := grid[curr.pos]
		for _, n := range nextNode(curr, m) {
			// Check if within bounds
			if n.pos.WithinPositive(xLen, yLen) {
				beams = append(beams, n)
			}
		}
	}
	return len(path)
}
