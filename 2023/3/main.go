package main

import (
	_ "embed"
	"fmt"
	adventofc2023 "github.com/wesrobin/adventofc/2023"
	"math"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

var testInput1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var testInput2 = ``

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(inp string) any {
	var sum int
	lines := strings.Split(inp, "\n")
	for y := 0; y < len(lines); y++ {
		var (
			num string
			adj bool
		)
		for x := 0; x < len(lines[y]); x++ {
			curr := rune(lines[y][x])
			if unicode.IsDigit(curr) {
				num += string(curr)
				if !adj && checkAdj(lines, x, y) {
					adj = true
				}
			} else if num != "" {
				if adj {
					sum += adventofc2023.Atoi(num)
				}
				num = ""
				adj = false
			}
		}
		if num != "" && adj {
			sum += adventofc2023.Atoi(num)
		}
	}
	return sum
}

func checkAdj(lines []string, x int, y int) bool {
	xLower, xUpper := int(math.Max(float64(x-1), 0)), int(math.Min(float64(x+1), float64(len(lines[0])-1)))
	yLower, yUpper := int(math.Max(float64(y-1), 0)), int(math.Min(float64(y+1), float64(len(lines)-1)))
	for yIdx := yLower; yIdx <= yUpper; yIdx++ {
		line := ""
		for xIdx := xLower; xIdx <= xUpper; xIdx++ {
			curr := rune(lines[yIdx][xIdx])
			line += string(curr)
			if curr != '.' && !unicode.IsDigit(curr) {
				return true
			}
		}
	}
	return false
}

func part2(inp string) any {
	var sum int
	lines := strings.Split(inp, "\n")
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			curr := rune(lines[y][x])
			if curr != '*' {
				continue
			}
			if nums, ok := has2Adj(lines, x, y); ok {
				sum += nums[0] * nums[1]
			}
		}
	}
	//fmt.Println(numGroups)
	return sum
}

func has2Adj(lines []string, x int, y int) ([]int, bool) {
	var found []int
	xLower, xUpper := int(math.Max(float64(x-1), 0)), int(math.Min(float64(x+1), float64(len(lines[0])-1)))
	yLower, yUpper := int(math.Max(float64(y-1), 0)), int(math.Min(float64(y+1), float64(len(lines)-1)))
	for yIdx := yLower; yIdx <= yUpper; yIdx++ {
		for xIdx := xLower; xIdx <= xUpper; xIdx++ {
			curr := rune(lines[yIdx][xIdx])
			if unicode.IsDigit(curr) {
				var num int
				num, xIdx = findNum(lines[yIdx], xIdx)
				found = append(found, num)
				if len(found) > 2 {
					return nil, false
				}
				xIdx++ // skip over the next char because it must be nondigit
			}
		}
	}
	if len(found) < 2 {
		return nil, false
	}

	return found, true
}

func findNum(line string, idx int) (int, int) {
	i := idx
	num := string(line[idx])

	for {
		i--
		if i < 0 {
			break
		}
		curr := rune(line[i])
		if !unicode.IsDigit(curr) {
			break
		}
		num = string(curr) + num
	}
	i = idx
	for {
		i++
		if i >= len(line) {
			break
		}
		curr := rune(line[i])
		if !unicode.IsDigit(curr) {
			break
		}
		num += string(curr)
	}
	return adventofc2023.Atoi(num), i - 1
}
