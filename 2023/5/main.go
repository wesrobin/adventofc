package main

import (
	_ "embed"
	"fmt"
	adventofc2023 "github.com/wesrobin/adventofc/2023"
	"strings"
)

//go:embed input.txt
var input string

/**
[50..97] +2> [52..99] <> [52..54] -15> [37..39] <> [37..39] -11> [26..28] <>
*/

var testInput1 = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 20`

var testInput2 = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4`

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type dict struct {
	entries []entry
}

func (d dict) getMapping(src int) int {
	for _, e := range d.entries {
		if src >= e.low && src <= e.high {
			return src + e.add
		}
	}
	return src
}

type entry struct {
	low, high int
	add       int
}

func (e entry) contains(i int) bool {
	return i >= e.low && i <= e.high
}

func getSeeds(inp string) []int {
	_, seedsStr, _ := strings.Cut(strings.Split(inp, "\n")[0], ": ")
	seeds := adventofc2023.MapSlice(strings.Split(seedsStr, " "), func(s string) int {
		return adventofc2023.Atoi(s)
	})
	return seeds
}

func makeDict(inp string) []dict {
	ds := make([]dict, len(strings.Split(inp, "\n\n"))-1)

	for i, section := range strings.Split(inp, "\n\n") {
		if i == 0 {
			continue
		}
		for _, line := range strings.Split(adventofc2023.CutRight(section, ":\n"), "\n") {
			parts := adventofc2023.MapSlice(strings.Split(line, " "), adventofc2023.Atoi)
			ds[i-1].entries = append(ds[i-1].entries, entry{
				low:  parts[1],
				high: parts[1] + parts[2],
				add:  parts[0] - parts[1],
			})
		}
	}
	return ds
}

func part1(inp string) any {
	seeds := getSeeds(inp)
	ds := makeDict(inp)

	smallest := -1
	for _, s := range seeds {
		curr := s
		for _, m := range ds {
			curr = m.getMapping(curr)
		}
		if smallest == -1 {
			smallest = curr
		} else {
			smallest = min(smallest, curr)
		}
	}

	return smallest
}

type rnge struct {
	l, h int
}

func (r rnge) isZero() bool {
	return r.l == 0 && r.h == 0
}

func (r rnge) apply(a int) rnge {
	return rnge{
		l: r.l + a,
		h: r.h + a,
	}
}

func overlap(a, b rnge) bool {
	return !(a.l > b.h || a.h < b.l || b.l > a.h || b.h < a.l)
}

func clamp(a, b rnge) rnge {
	return rnge{
		l: max(a.l, b.l),
		h: min(a.h, b.h),
	}
}

func part2(inp string) any {
	seedNums := getSeeds(inp)
	var seeds []rnge
	for i := range seedNums {
		if i%2 == 1 {
			seeds = append(seeds, rnge{l: seedNums[i-1], h: seedNums[i-1] + seedNums[i]})
		}
	}

	ds := makeDict(inp)

	for _, d := range ds {
		seeds = applyLayer(seeds, d)
		//fmt.Println(i, seeds)
	}
	var m int
	for i, s := range seeds {
		if i == 0 {
			m = s.l
		} else {
			m = min(m, s.l)
		}
	}

	return m
}

func applyLayer(seeds []rnge, d dict) []rnge {
	var res []rnge
	toCheck := make([]rnge, len(seeds))
	copy(toCheck, seeds)

	for _, e := range d.entries {
		var nextCheck []rnge
		r := rnge{l: e.low, h: e.high}
		//fmt.Println("c", r, "tc", toCheck, "r", res, "add", e.add)
		for _, tc := range toCheck {
			if tc.isZero() {
				continue
			}
			if !overlap(tc, r) {
				// No matches, we just pass through the seed range because no mapping means n -> n
				nextCheck = append(nextCheck, tc)
				continue
			}
			be, mid, af := split(tc, r)
			//if !mid.isZero() {
			//	fmt.Println("match", mid, "add", mid.apply(e.add))
			//}
			res = append(res, mid.apply(e.add))
			if !be.isZero() {
				nextCheck = append(nextCheck, be)
			}
			if !af.isZero() {
				nextCheck = append(nextCheck, af)
			}
		}
		toCheck = nextCheck
	}

	return append(res, toCheck...)
}

func split(tc rnge, r rnge) (be rnge, mid rnge, af rnge) {
	if !overlap(tc, r) {
		return tc, rnge{}, rnge{}
	}
	mid = clamp(tc, r)
	if tc.l < r.l {
		be = rnge{l: min(tc.l, r.l), h: max(tc.l, r.l)}
	}
	if tc.h > r.h {
		af = rnge{l: min(tc.h, r.h), h: max(tc.h, r.h)}
	}
	return be, mid, af
}
