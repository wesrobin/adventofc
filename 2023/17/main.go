package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"maps"
	"strings"
	"time"

	aoc "github.com/wesrobin/adventofc/2023"
)

//go:embed input.txt
var inpoot string

func input() string {
	return strings.TrimSpace(inpoot)
}

func main() {
	t0 := time.Now()
	part1()
	fmt.Println(time.Since(t0))
	t0 = time.Now()
	part2() // 1243 H, 1227 H (same as someone else), 1219
	fmt.Println(time.Since(t0))
}

const testInput1 = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

const testInput2 = `563
735
533`

const testInput3 = `123
456
789`

const testInput4 = `24134
32154`

const testInput5 = `111111111111
999999999991
999999999991
999999999991
999999999991`

func part1() {
	//inp := testInput1
	inp := input()
	lines := strings.Split(inp, "\n")
	grid := make([][]int, len(lines))
	costGrid := make([][]int, len(lines))
	for y, l := range lines {
		for _, c := range l {
			grid[y] = append(grid[y], aoc.Atoi(string(c)))
			costGrid[y] = append(costGrid[y], aoc.Atoi(string(c)))
		}
	}
	down := pathStep{
		c:     aoc.Coord2D{},
		depth: 1,
		dir:   aoc.Down,
	}
	right := pathStep{
		c:     aoc.Coord2D{},
		depth: 1,
		dir:   aoc.Right,
	}

	populateCosts(grid, down, right)
}

type traversal struct {
	last pathStep
	hist map[aoc.Coord2D]rune
}

type TraversalQueue []*traversal

func (t *TraversalQueue) Push(x any) {
	item := x.(*traversal)
	*t = append(*t, item)
}

func (t *TraversalQueue) Pop() any {
	old := *t
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*t = old[0 : n-1]
	return item
}

func (t TraversalQueue) Len() int {
	return len(t)
}

func (t TraversalQueue) Less(i, j int) bool {
	return t[i].last.cumulative < t[j].last.cumulative
}

func (t TraversalQueue) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func populateCosts(grid [][]int, start1, start2 pathStep) {
	end := aoc.Coord2D{X: len(grid[0]) - 1, Y: len(grid) - 1}

	// Idea: Sorting the queue _significantly_ cuts down iterations (assuming because it optimises the memoization but
	// tbh not sure). But appending the next node and then sorting nLogn every iteration is very slow. What about smth
	// like a BST, where insertion is logn, and we always have a sorted queue?
	queue := TraversalQueue{}
	queue.Push(&traversal{last: start1, hist: map[aoc.Coord2D]rune{}})
	queue.Push(&traversal{last: start2, hist: map[aoc.Coord2D]rune{}})

	heap.Init(&queue)

	//minAt := map[key]int{
	//	down.key():  grid[0][0],
	//	right.key(): grid[0][0],
	//}
	seen := map[key]bool{}
	minEnd := 1_000_000_000

OuterLoop:
	for len(queue) > 0 {
		curr := heap.Pop(&queue).(*traversal)
		curr.hist[curr.last.c] = curr.last.dir.Char()

		var added int
		for _, n := range next(curr.last) {
			if !n.valid(grid) {
				continue
			}
			n.cumulative = curr.last.cumulative + grid[n.c.Y][n.c.X]

			if !curr.last.isUltra && n.c == end || curr.last.isUltra && n.c == end && n.depth > 3 {
				if n.cumulative < minEnd {
					minEnd = n.cumulative
				}
				printHist(curr.hist, grid)
				break OuterLoop
			}
			if seen[n.key()] {
				continue
			}
			seen[n.key()] = true
			newHist := make(map[aoc.Coord2D]rune)
			maps.Copy(newHist, curr.hist)

			heap.Push(&queue, &traversal{last: n, hist: newHist})
			added++
		}
		//if added == 0 {
		//	printHist(curr.hist, grid)
		//	fmt.Println()
		//}
	}
	fmt.Println(minEnd)
}

func printHist(hist map[aoc.Coord2D]rune, grid [][]int) {
	for y, l := range grid {
		for x := range l {
			c := aoc.Coord2D{X: x, Y: y}
			if ch, ok := hist[c]; ok {
				fmt.Print(string(ch))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type key struct {
	c     aoc.Coord2D
	depth int
	dir   aoc.Dir
}

type pathStep struct {
	c          aoc.Coord2D
	depth      int
	cumulative int
	dir        aoc.Dir
	isUltra    bool
}

func (p pathStep) key() key {
	return key{c: aoc.Coord2D{X: p.c.X, Y: p.c.Y}, depth: p.depth, dir: p.dir}
}

func (p pathStep) valid(grid [][]int) bool {
	if p.c.X == 0 && p.c.Y == 0 {
		return false
	}
	if !p.c.WithinPositive(len(grid[0]), len(grid)) {
		return false
	}
	if !p.isUltra && p.depth > 3 {
		return false
	} else if p.isUltra && p.depth > 10 {
		return false
	}
	return true
}

func (p pathStep) string() string {
	return fmt.Sprintf("(%d;%d), dep:%d, cum:%d, dir:%s", p.c.X, p.c.Y, p.depth, p.cumulative, p.dir.String())
}

func next(curr pathStep) []pathStep {
	var ps []pathStep

	cs := aoc.NextCoord(curr.c, curr.dir)
	ps = append(ps, pathStep{
		c:       cs,
		depth:   curr.depth + 1,
		dir:     curr.dir,
		isUltra: curr.isUltra,
	})

	if !curr.isUltra || curr.depth >= 4 {
		cl := aoc.NextCoord(curr.c, (curr.dir-1+4)%4)
		ps = append(ps, pathStep{
			c:       cl,
			depth:   1,
			dir:     (curr.dir - 1 + 4) % 4,
			isUltra: curr.isUltra,
		})
		cr := aoc.NextCoord(curr.c, (curr.dir+1+4)%4)
		ps = append(ps, pathStep{
			c:       cr,
			depth:   1,
			dir:     (curr.dir + 1 + 4) % 4,
			isUltra: curr.isUltra,
		})
	}

	return ps
}

func part2() {
	//inp := testInput1
	inp := input()
	lines := strings.Split(inp, "\n")
	grid := make([][]int, len(lines))
	costGrid := make([][]int, len(lines))
	for y, l := range lines {
		for _, c := range l {
			grid[y] = append(grid[y], aoc.Atoi(string(c)))
			costGrid[y] = append(costGrid[y], aoc.Atoi(string(c)))
		}
	}
	down := pathStep{
		c:       aoc.Coord2D{},
		depth:   1,
		dir:     aoc.Down,
		isUltra: true,
	}
	right := pathStep{
		c:       aoc.Coord2D{},
		depth:   1,
		dir:     aoc.Right,
		isUltra: true,
	}

	populateCosts(grid, down, right)
}
