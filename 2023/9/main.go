package main

import (
	_ "embed"
	"fmt"
	"strings"

	adventofc2023 "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

var testInput1 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

var ti3 = "12 18 47 119 263 528 1003 1849 3355 6042 10851 19460 34787 61785 108811 190335 332858 588110 1061583 1971129 3760859"

func part1(inp string) any {
	var sum int
	for _, line := range strings.Split(inp, "\n") {
		if line == "" {
			continue
		}
		lineNums := adventofc2023.MapSlice(strings.Fields(line),
			adventofc2023.Atoi)
		ds := recurse(lineNums, [][]int{})
		sum += findNext(ds[:len(ds)-1]) + lineNums[len(lineNums)-1]
	}
	return sum
}

func recurse(nums []int, diffs [][]int) [][]int {
	if allZero(nums) {
		return diffs
	}
	var ds []int
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		ds = append(ds, d)
	}
	return recurse(ds, append(diffs, ds))
}

func findNext(diffs [][]int) int {
	curr := diffs[len(diffs)-1][0]
	for i := len(diffs) - 2; i >= 0; i-- {
		n := diffs[i][len(diffs[i])-1]
		curr += n
	}
	return curr
}

func allZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

var testInput2 = ``

func part2(inp string) any {
	var sum int
	for _, line := range strings.Split(inp, "\n") {
		if line == "" {
			continue
		}
		lineNums := adventofc2023.MapSlice(strings.Fields(line),
			adventofc2023.Atoi)
		ds := recurse(lineNums, [][]int{})

		sum +=  lineNums[0] - findNext2(ds[:len(ds)-1])
	}
	return sum
}

func findNext2(diffs [][]int) int {
	curr := diffs[len(diffs)-1][0]
	for i := len(diffs) - 2; i >= 0; i-- {
		n := diffs[i][0]
		curr = n - curr
	}
	return curr
}
