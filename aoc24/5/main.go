package main

import (
	_ "embed"
	"fmt"
	"github.com/wesrobin/adventofc/aoc24"
	"slices"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

var testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func main() {
	inp := input
	if false {
		inp = testInput
	}

	rs, ps, _ := strings.Cut(inp, "\n\n")

	rules := strings.Split(rs, "\n")
	pages := strings.Split(ps, "\n")

	fmt.Println(p1(rules, pages))
	fmt.Println(p2(rules, pages))
}

func p2(rules, pages []string) int {
	before := make(map[int][]int)
	for _, rule := range rules {
		b, a, _ := strings.Cut(rule, "|")
		before[aoc24.Atoi(b)] = append(before[aoc24.Atoi(b)], aoc24.Atoi(a))
	}

	var cnt int
	for _, p := range pages {
		pSplit := strings.Split(p, ",")
		sortFunc := func(i, j int) bool {
			return slices.Contains(before[aoc24.Atoi(pSplit[i])], aoc24.Atoi(pSplit[j]))
		}
		if !sort.SliceIsSorted(pSplit, sortFunc) {
			sort.SliceStable(pSplit, sortFunc)
			cnt += aoc24.Atoi(pSplit[(len(pSplit)-1)/2])
		}
	}

	return cnt
}

func p1(rules, pages []string) int {
	before := make(map[int][]int)
	for _, rule := range rules {
		b, a, _ := strings.Cut(rule, "|")
		before[aoc24.Atoi(b)] = append(before[aoc24.Atoi(b)], aoc24.Atoi(a))
	}

	var cnt int
	for _, p := range pages {
		pSplit := strings.Split(p, ",")
		idxs := make(map[int]int)
		for i, n := range pSplit {
			idxs[aoc24.Atoi(n)] = i
		}
		if check(before, idxs) {
			cnt += aoc24.Atoi(pSplit[(len(pSplit)-1)/2])
		}
	}

	return cnt
}

func check(before map[int][]int, idxs map[int]int) bool {
	for k, v := range idxs {
		if befores, ok := before[k]; ok {
			for _, b := range befores {
				if bIdx, bOk := idxs[b]; bOk && bIdx < v {
					return false
				}
			}
		}
	}

	return true
}
