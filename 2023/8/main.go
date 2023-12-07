package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

var testInput1 = ``

var testInput2 = ``

func main() {
	fmt.Println(part1(input)) // 253410429 too high, 252960917 too low, 253000392 bad
	fmt.Println(part2(input))
}

func part1(inp string) any {
	return 0
}

func part2(inp string) any {
	return 0
}
