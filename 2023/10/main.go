package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"time"

	adventofc2023 "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var input string

func main() {
	//inp := input // 681
	inp := input
    t0 := time.Now()
	seen := part1(inp)
    fmt.Println(time.Since(t0))
    t1 := time.Now()
	part2(inp, seen)
    fmt.Println(time.Since(t1))
}

type dir int

const (
	up    dir = 1
	down  dir = 2
	left  dir = 3
	right dir = 4
)

func (d dir) str() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}
	return "no_dir_found"
}

func getNextC(c adventofc2023.Coord2D, d dir, xLen, yLen int) (adventofc2023.Coord2D, dir, error) {
	var c2 adventofc2023.Coord2D
	var fromDir dir
	switch d {
	case up:
		c2 = adventofc2023.Coord2D{X: c.X, Y: c.Y - 1}
		fromDir = down
	case down:
		c2 = adventofc2023.Coord2D{X: c.X, Y: c.Y + 1}
		fromDir = up
	case left:
		c2 = adventofc2023.Coord2D{X: c.X - 1, Y: c.Y}
		fromDir = right
	case right:
		c2 = adventofc2023.Coord2D{X: c.X + 1, Y: c.Y}
		fromDir = left
	}
	if c2.X > xLen || c2.Y > yLen || c2.X < 0 || c2.Y < 0 {
		return adventofc2023.Coord2D{}, dir(0), errors.New("bad")
	}
	return c2, fromDir, nil
}

type symbol rune

const (
	vert   symbol = '|'
	hori   symbol = '-'
	ne     symbol = 'L'
	nw     symbol = 'J'
	sw     symbol = '7'
	se     symbol = 'F'
	ground symbol = '.'
)

func (s symbol) pretty() string {
    switch s {
        case vert:
            return "┃"
        case hori:
            return "━"
        case ne:
            return "┗"
        case nw:
            return "┛"
        case se:
            return "┏"
        case sw:
            return "┓"
    }
    return ""
}

func (s symbol) getNextDir(from dir) (dir, bool) {
	switch from {
	case up:
		if s == vert {
			return down, true
		}
		if s == nw {
			return left, true
		}
		if s == ne {
			return right, true
		}

	case down:
		if s == vert {
			return up, true
		}
		if s == sw {
			return left, true
		}
		if s == se {
			return right, true
		}

	case left:
		if s == hori {
			return right, true
		}
		if s == sw {
			return down, true
		}
		if s == nw {
			return up, true
		}
	case right:
		if s == hori {
			return left, true
		}
		if s == se {
			return down, true
		}
		if s == ne {
			return up, true
		}
	}
	return dir(0), false
}

const testInput1 = `.....
.S-7.
.|.|.
.L-J.
.....`

const testInput2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const testInput3 = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const testInput4 = `...........
.S-------7.
.|F-----7|.
.||.....|L7
.||.....|FJ
FJL-7.F-J|.
L7..|.|..|.
.L--J.L--J.
...........`

const testInput5 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func part1(inp string) map[adventofc2023.Coord2D]bool {
	lines := strings.Split(inp, "\n")
	xLen, yLen := len(lines[0]), len(lines)
	var start adventofc2023.Coord2D
	for y, l := range lines {
		for x, c := range l {
			if c == 'S' {
				start = adventofc2023.Coord2D{X: x, Y: y}
				break
			}
		}
	}

	type step struct {
		c     adventofc2023.Coord2D
		from  dir
		depth int
	}
	queue := []step{{c: start, depth: 0}}
	seen := make(map[adventofc2023.Coord2D]bool)

	for len(queue) > 0 {
		curr := queue[0]
		currSym := symbol(lines[curr.c.Y][curr.c.X])
		queue = queue[1:]
		if seen[curr.c] {
			printGrid(lines, seen)
			fmt.Println(curr.depth)
			return seen
		} else if currSym == ground {
			continue
		}

		seen[curr.c] = true
		if curr.c == start {
			if upC, _, err := getNextC(curr.c, up, xLen, yLen); err == nil {
				if sym := symbol(lines[upC.Y][upC.X]); sym == vert || sym == sw || sym == se {
					queue = append(queue, step{c: upC, from: down, depth: curr.depth + 1})
				}
			}
			if doC, _, err := getNextC(curr.c, down, xLen, yLen); err == nil {
				if sym := symbol(lines[doC.Y][doC.X]); sym == vert || sym == nw || sym == ne {
					queue = append(queue, step{c: doC, from: up, depth: curr.depth + 1})
				}
			}
			if leC, _, err := getNextC(curr.c, left, xLen, yLen); err == nil {
				if sym := symbol(lines[leC.Y][leC.X]); sym == hori || sym == se || sym == ne {
					queue = append(queue, step{c: leC, from: right, depth: curr.depth + 1})
				}
			}
			if riC, _, err := getNextC(curr.c, right, xLen, yLen); err == nil {
				if sym := symbol(lines[riC.Y][riC.X]); sym == hori || sym == sw || sym == nw {
					queue = append(queue, step{c: riC, from: left, depth: curr.depth + 1})
				}
			}
			continue
		}

		nextD, ok := currSym.getNextDir(curr.from)
		if !ok {
			continue
		}

		c, f, err := getNextC(curr.c, nextD, xLen, yLen)
		if err != nil {
			continue
		}

		queue = append(queue, step{c: c, from: f, depth: curr.depth + 1})
	}
	return nil
}

func printGrid(lines []string, seen map[adventofc2023.Coord2D]bool) {
	for y, l := range lines {
		var prLine string
		for x, c := range l {
			coord := adventofc2023.Coord2D{X: x, Y: y}
			if seen[coord] {
                sym := symbol(lines[y][x]).pretty()
				prLine += sym
			} else {
				prLine += string(c)
			}
		}
		fmt.Println(prLine)
	}
}

func part2(inp string, seen map[adventofc2023.Coord2D]bool) {
	lines := strings.Split(inp, "\n")
	// Based on output from previous, element in corners are not enclosed by the pipe
    var count int
    counted := make(map[adventofc2023.Coord2D]bool)
    for y, line := range lines {
        var in bool
        var cornerPair symbol
        for x := range line {
            coord := adventofc2023.Coord2D{X: x, Y: y}
            if seen[coord] {
                s := symbol(lines[y][x])
                if s == vert {
                    in = !in
                } else if s == se {
                    cornerPair = se
                } else if s == ne {
                    cornerPair = ne
                } else if s == sw {
                    if cornerPair == ne {
                        in = !in
                    }
                    cornerPair = symbol(0)
                } else if s == nw {
                    if cornerPair == se {
                        in = !in
                    }
                    cornerPair = symbol(0)
                }
            } else if in {
                counted[coord] = true
                count++
            }
        }
    }
    printGrid2(lines, seen, counted)
	fmt.Println(count)
}

func adj(c adventofc2023.Coord2D, xLen, yLen int) []adventofc2023.Coord2D {
	var cs []adventofc2023.Coord2D
	up := adventofc2023.Coord2D{X: c.X, Y: c.Y - 1}
	if inBounds(up, xLen, yLen) {
		cs = append(cs, up)
	}
	down := adventofc2023.Coord2D{X: c.X, Y: c.Y + 1}
	if inBounds(down, xLen, yLen) {
		cs = append(cs, down)
	}
	left := adventofc2023.Coord2D{X: c.X - 1, Y: c.Y}
	if inBounds(left, xLen, yLen) {
		cs = append(cs, left)
	}
	right := adventofc2023.Coord2D{X: c.X + 1, Y: c.Y}
	if inBounds(right, xLen, yLen) {
		cs = append(cs, right)
	}
	return cs
}

func inBounds(c adventofc2023.Coord2D, xLen, yLen int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < xLen && c.Y < yLen
}

func flood(lines []string, seen, floodSeen map[adventofc2023.Coord2D]bool, start adventofc2023.Coord2D) {
    if seen[start] || floodSeen[start] {
        return
    }
	queue := []adventofc2023.Coord2D{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, c := range adj(curr, len(lines[0]), len(lines)) {
			if seen[c] || floodSeen[c] {
				continue
			}
            floodSeen[c] = true
			queue = append(queue, c)
		}
	}
}

func printGrid2(lines []string, seen, counted map[adventofc2023.Coord2D]bool) {
	for y, l := range lines {
		var prLine string
		for x, c := range l {
			coord := adventofc2023.Coord2D{X: x, Y: y}
			if seen[coord] {
                sym := symbol(lines[y][x]).pretty()
				prLine += sym
			}  else if counted[coord] {
                prLine += "#"
            } else {
				prLine += string(c)
			}
		}
		fmt.Println(prLine)
	}
}
