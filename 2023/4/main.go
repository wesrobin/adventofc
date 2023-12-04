package main

import (
	_ "embed"
	"fmt"
	adventofc2023 "github.com/wesrobin/adventofc/2023"
	"strings"
)

//go:embed input.txt
var input string

var testInput1 = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

var testInput2 = ``

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(inp string) any {
	lines := strings.Split(inp, "\n")
	var sum int
	for _, l := range lines {
		_, game, _ := strings.Cut(l, ": ")
		hand, winning, _ := strings.Cut(game, " | ")
		shmeh := make(map[int]bool)
		for _, c := range strings.Split(strings.TrimSpace(hand), " ") {
			if c == " " {
				continue
			}
			shmeh[adventofc2023.Atoi(c)] = true
		}
		var tot int
		for _, c := range strings.Split(strings.TrimSpace(winning), " ") {
			if c == " " || c == "" {
				continue
			}
			if shmeh[adventofc2023.Atoi(c)] {
				if tot == 0 {
					tot++
				} else {
					tot *= 2
				}
			}
		}
		sum += tot
	}
	return sum
}

// 136188 x

func part2(inp string) any {
	lines := strings.Split(inp, "\n")
	cardCounter := make(map[int]int)
	for i, l := range lines {
		cardCounter[i+1]++
		_, game, _ := strings.Cut(l, ": ")
		hand, winning, _ := strings.Cut(game, " | ")
		dict := make(map[int]bool)
		for _, c := range strings.Split(strings.TrimSpace(hand), " ") {
			if c == " " {
				continue
			}
			dict[adventofc2023.Atoi(c)] = true
		}
		var tot int
		for _, c := range strings.Split(strings.TrimSpace(winning), " ") {
			if c == " " || c == "" {
				continue
			}
			if dict[adventofc2023.Atoi(c)] {
				tot++
			}
		}
		for j := 0; j < tot; j++ {
			idx := 2 + i + j
			cardCounter[idx] += cardCounter[i+1]
		}
	}
	var sum int
	for _, v := range cardCounter {
		sum += v
	}
	return sum
}
