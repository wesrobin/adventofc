package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	aoc "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var input string

func main() {
	inp := input
	t0 := time.Now()
	part1(inp)
	fmt.Println(time.Since(t0))
	//part2(inp)
	//
	//fmt.Println("")
	//fmt.Println(".###???????? 3,2,1 4", validUntil(strings.Split(".###????????", ""), []int{3, 2, 1}, 4))
	//fmt.Println(".###.??????? 3,2,1 5", validUntil(strings.Split(".###????????", ""), []int{3, 2, 1}, 4))
	//fmt.Println(".###.##..... 3,2,1 11", validUntil(strings.Split(".###.##.....", ""), []int{3, 2, 1}, 11))
	//fmt.Println(".###..##.#.? 3,2,1 10", validUntil(strings.Split(".###..##.#.?", ""), []int{3, 2, 1}, 11))
	//fmt.Println(".###.......# 3,2,1 12", validUntil(strings.Split(".###.......#", ""), []int{3, 2, 1}, 11))
	//fmt.Println("#.#.### 1,1,3", validUntil(strings.Split("#.#.###", ""), []int{1, 1, 3}, 6))
	//fmt.Println(".###......... 3,2,1", validUntil(strings.Split(".###......... 1,1,3", ""), []int{3, 2, 1}, 80))
	//fmt.Println("#.#.### 2,3", validUntil(strings.Split("#.#.###", ""), []int{2, 3}, 6))
	//fmt.Println("##..### 1,1,3", validUntil(strings.Split("##..###", ""), []int{1, 1, 3}, 6))
	//fmt.Println("##..### 2,3", validUntil(strings.Split("##..###", ""), []int{2, 3}, 6))
	//fmt.Println("##.###.## 2,3 1", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 1))
	//fmt.Println("##.###.## 2,3 2", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 2))
	//fmt.Println("##.###.## 2,3 3", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 3))
	//fmt.Println("##.###.## 2,3 4", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 4))
	//fmt.Println("##.###.## 2,3 5", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 5))
	//fmt.Println("##.###.## 2,3 6", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 6))
	//fmt.Println("##.###.## 2,3 7", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 7))
	//fmt.Println("##.###.## 2,3 8", validUntil(strings.Split("##.###.##", ""), []int{2, 3}, 8))
	//fmt.Println("##.###. 2,3,1", validUntil(strings.Split("##.###.", ""), []int{2, 3, 1}, 6))

	//fmt.Println(".###...##.#. 3,2,1", valid(strings.Split(".###...##.#.", ""), []int{3, 2, 1}, 1))
	//fmt.Println("##.###.## 2,3 1", valid(strings.Split("##.###.##", ""), []int{2, 3}, 1))
	//fmt.Println("##.###.## 2,3 2", valid(strings.Split("##.###.##", ""), []int{2, 3}, 2))
	//fmt.Println("##.###.## 2,3 3", valid(strings.Split("##.###.##", ""), []int{2, 3}, 3))
	//fmt.Println("##.###.## 2,3 4", valid(strings.Split("##.###.##", ""), []int{2, 3}, 4))
	//fmt.Println("##.###.## 2,3 5", valid(strings.Split("##.###.##", ""), []int{2, 3}, 5))
	//fmt.Println("##.###.## 2,3 6", valid(strings.Split("##.###.##", ""), []int{2, 3}, 6))
	//fmt.Println("##.###.## 2,3 7", valid(strings.Split("##.###.##", ""), []int{2, 3}, 7))
	//fmt.Println("##.###.## 2,3 8", valid(strings.Split("##.###.##", ""), []int{2, 3}, 8))
	//fmt.Println(".###???????? 3,2,1 4", valid(strings.Split(".###????????", ""), []int{3, 2, 1}, 4))
	//fmt.Println(".###.??????? 3,2,1 5", valid(strings.Split(".###????????", ""), []int{3, 2, 1}, 4))
	//fmt.Println(".###.##..... 3,2,1 11", valid(strings.Split(".###.##.....", ""), []int{3, 2, 1}, 11))
	//fmt.Println(".###..##.#.? 3,2,1 10", valid(strings.Split(".###..##.#.?", ""), []int{3, 2, 1}, 11))
	//fmt.Println(".###.......# 3,2,1 12", valid(strings.Split(".###.......#", ""), []int{3, 2, 1}, 11))
	//fmt.Println("#.#.### 1,1,3", valid(strings.Split("#.#.###", ""), []int{1, 1, 3}, 6))
	//fmt.Println(".###......... 3,2,1", valid(strings.Split(".###......... 1,1,3", ""), []int{3, 2, 1}, 80))
	//fmt.Println("#.#.### 2,3", valid(strings.Split("#.#.###", ""), []int{2, 3}, 6))
	//fmt.Println("##..### 1,1,3", valid(strings.Split("##..###", ""), []int{1, 1, 3}, 6))
	//fmt.Println("##..### 2,3", valid(strings.Split("##..###", ""), []int{2, 3}, 6))
	//fmt.Println("##.###. 2,3,1", validUntil(strings.Split("##.###.", ""), []int{2, 3, 1}, 6))

}

const testInput1 = `???.### 1,1,3`

const testInput2 = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

const testInput3 = `?###???????? 3,2,1`

func part1(inp string) {
	lines := strings.Split(inp, "\n")
	var sum int
	for _, l := range lines {
		l, r, _ := strings.Cut(l, " ")
		amt := countPerms(strings.Split(l, ""), aoc.MapIntsToStr(strings.Split(r, ",")))
		//fmt.Println(l, r, "=", amt)
		sum += amt
	}
	fmt.Println(sum)
}

func key(ss []string, newC string, idx int) string {
	return fmt.Sprintf("%s %s %d", strings.Join(ss, ""), newC, idx)
}

func countPerms(s []string, i []int) int {
	m := make(map[string]bool)
	added := make(map[string]bool)

	var sum int
	inclS := make([]string, len(s))
	copy(inclS, s)
	sum += countRecurse(inclS, i, "#", 0, m, added)
	exclS := make([]string, len(s))
	copy(exclS, s)
	sum += countRecurse(exclS, i, ".", 0, m, added)
	return sum
}

func countRecurse(ss []string, ns []int, newC string, idx int, seen, added map[string]bool) int {
	// For now don't try stop recursing when it's 'already invalid'. It's confusing af
	if seen[key(ss, newC, idx)] {
		return 0
	} else {
		seen[key(ss, newC, idx)] = true
	}
	if idx == len(ss) {
		joinS := strings.Join(ss, "")
		if added[joinS] {
			return 0
		}
		if valid(ss, ns) {
			added[joinS] = true
			return 1
		} else {
			return 0
		}
	}

	var sum int
	for _, c := range []string{".", "#"} {
		newS := make([]string, len(ss))
		copy(newS, ss)
		if ss[idx] == "?" {
			newS[idx] = newC
		}
		sum += countRecurse(newS, ns, c, idx+1, seen, added)
		//if shouldContinue(newS, ns) {
		//	sum += countRecurse(newS, ns, c, idx+1, seen, added)
		//} else {
		//	fmt.Println("sc", newS, ns)
		//}
	}
	return sum
}

func shouldContinue(ss []string, ns []int) bool {
	var (
		check []int
		count int
	)
	for i, s := range ss {
		if s == "#" {
			count++
		} else if s == "." && count > 0 {
			check = append(check, count)
			count = 0
		} else if s == "?" {
			// Try a 'lookahead' recursion for all '?'s to determine if it's possible. Cap at 5 so it doesn't do the whole damn thing
			lastIdx := i + 1
			for ; lastIdx < i+5; lastIdx++ {
				if ss[lastIdx] != "?" {
					lastIdx--
					break
				}
			}

			break
		}
	}
	if count > 0 {
		check = append(check, count)
	}

	for i, c := range check {
		if c != ns[i] {
			return false
		}
	}
	return true
}

func valid(ss []string, ns []int) bool {
	var (
		check []int
		count int
	)
	for _, s := range ss {
		if s == "#" {
			count++
		} else if count > 0 {
			check = append(check, count)
			if len(check) > len(ns) || count != ns[len(check)-1] {
				return false
			}
			count = 0
		}
	}
	if count > 0 {
		check = append(check, count)
	}

	if len(check) != len(ns) {
		return false
	}
	for i, c := range ns {
		if c != check[i] {
			return false
		}
	}
	return true
}

func part2(inp string) {
	lines := strings.Split(inp, "\n")
	var sum int
	for _, li := range lines {
		l, r, _ := strings.Cut(li, " ")
		l, r = expand(l, r)
		amt := countPerms(strings.Split(l, ""), aoc.MapIntsToStr(strings.Split(r, ",")))
		//fmt.Println(l, r, "=", amt)
		sum += amt
	}
	// I think this should memoize like: given input ..##.???? the following permutations are
	// possible -> [x, y, z]
	// If validating the solutions against some ns is fast, this should significantly speed up
	// the process.
	// Will need to split up the generation and validation/counting though.
	fmt.Println(sum)
}

func expand(l string, r string) (string, string) {
	var newL, newS []string
	for i := 0; i < 5; i++ {
		newL = append(newL, l)
		newS = append(newS, r)
	}
	return strings.Join(newL, "?"), strings.Join(newS, ",")
}
