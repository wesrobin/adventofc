package main

import (
	_ "embed"
	"fmt"
	adventofc2023 "github.com/wesrobin/adventofc/2023"
	"strings"
)

//go:embed input.txt
var input string

var testInput1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var testInput2 = ``

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(inp string) int {
	lines := strings.Split(inp, "\n")
	var sum int
	for i, l := range lines {
		g := strings.Split(l, ": ")[1]
		valid := true
		for _, round := range strings.Split(g, "; ") {
			draw := strings.Split(round, ", ")
			for _, d := range draw {
				str := strings.Split(d, " ")
				n, c := str[0], str[1]
				nInt := adventofc2023.Atoi(n)
				if c == "red" && nInt > 12 {
					valid = false
					break
				} else if c == "blue" && nInt > 14 {
					valid = false
					break
				} else if c == "green" && nInt > 13 {
					valid = false
					break
				}
			}
		}
		if valid {
			sum += i + 1
		}
	}
	return sum
}

func part2(inp string) int64 {
	lines := strings.Split(inp, "\n")
	var sum int64
	for _, l := range lines {
		g := strings.Split(l, ": ")[1]
		mins := make([]int64, 3)
		for _, round := range strings.Split(g, "; ") {
			draw := strings.Split(round, ", ")
			for _, d := range draw {
				str := strings.Split(d, " ")
				n, c := str[0], str[1]
				nInt := adventofc2023.Atoi(n)
				if c == "red" && (nInt > mins[0]) {
					mins[0] = nInt
					continue
				} else if c == "blue" && (nInt > mins[1]) {
					mins[1] = nInt
					continue
				} else if c == "green" && (nInt > mins[2]) {
					mins[2] = nInt
					continue
				}
			}
		}
		sum += mins[0] * mins[1] * mins[2]
	}
	return sum
}
