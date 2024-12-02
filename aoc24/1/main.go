package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/wesrobin/adventofc/aoc2024"
)

//go:embed input.txt
var inp string

func main() {
	var lList, rList []int
	for _, line := range strings.Split(inp, "\n") {
		l, r, _ := strings.Cut(line, "   ")
		lList = append(lList, aoc24.Atoi(l))
		rList = append(rList, aoc24.Atoi(r))
	}

	fmt.Println(p1(lList, rList))
	fmt.Println(p2(lList, rList))
}

func p1(lList, rList []int) int {
	sort.Ints(lList)
	sort.Ints(rList)

	var dist int
	for i := range len(lList) {
		dist += int(math.Abs(float64(lList[i]) - float64(rList[i])))
	}

	return dist
}

func p2(lList, rList []int) int {
	lMap := make(map[int]int)
	rMap := make(map[int]int)

	for i := range lList {
		lMap[lList[i]]++
		rMap[rList[i]]++
	}

	var similar int
	for k, v := range lMap {
		toAdd := k * rMap[k]
		similar += v * toAdd
	}

	return similar
}
