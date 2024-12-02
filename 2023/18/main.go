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
	part2()
}

const testInput1 = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func part1() {
	inp := testInput1
	lines := strings.Split(inp, "\n")
	curr := aoc.Coord2D{X: 0, Y: 0}
	var maximX, maximY int
	for _, l := range lines {
		ps := strings.Split(l, " ")
		switch ps[0] {
		case "U":
			curr.Y -= aoc.Atoi(ps[1])
		case "L":
			curr.X -= aoc.Atoi(ps[1])
		case "D":
			curr.Y += aoc.Atoi(ps[1])
		case "R":
			curr.X += aoc.Atoi(ps[1])
		}
		maximX = max(maximX, curr.X)
		maximY = max(maximY, curr.Y)
	}

	fmt.Println(maximX, maximY)
}

func part2() {

}
