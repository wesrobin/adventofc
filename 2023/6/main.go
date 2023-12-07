package main

import (
	_ "embed"
	"fmt"
	adventofc2023 "github.com/wesrobin/adventofc/2023"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input2.txt
var input2 string

var testInput1 = `Time:      7  15   30
Distance:  9  40  200`

var testInput2 = `Time:      71530
Distance:  940200`

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input2))
}

type race struct {
	t int
	d int
}

func part1(inp string) any {
	tStr, dStr, _ := strings.Cut(inp, "\n")
	ts, ds := strings.Fields(tStr)[1:], strings.Fields(dStr)[1:]
	var races []race
	for i := 0; i < len(ts); i++ {
		races = append(races, race{t: adventofc2023.Atoi(ts[i]), d: adventofc2023.Atoi(ds[i])})
	}

	var poss []int
	for _, r := range races {
		var wins int
		for j := 1; j < r.t-1; j++ {
			if j*(r.t-j) > r.d {
				wins++
			}
		}
		poss = append(poss, wins)
	}

	prod := 1
	for _, p := range poss {
		prod *= p
	}

	return prod
}

func part2(inp string) any {
	tStr, dStr, _ := strings.Cut(inp, "\n")
	ts, ds := strings.Fields(tStr)[1:], strings.Fields(dStr)[1:]
	var races []race
	for i := 0; i < len(ts); i++ {
		races = append(races, race{t: adventofc2023.Atoi(ts[i]), d: adventofc2023.Atoi(ds[i])})
	}

	var poss []int64
	for _, r := range races {
		var wins int64
		for j := 1; j < r.t-1; j++ {
			if j*(r.t-j) > r.d {
				wins++
			}
		}
		poss = append(poss, wins)
	}

	prod := int64(1)
	for _, p := range poss {
		prod *= p
	}

	return prod
}
