package main

import (
	_ "embed"
	"fmt"
	"maps"
	"sort"
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
	// part2()
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

func part1() {
	inp := testInput1
	//inp := input()
	lines := strings.Split(inp, "\n")
	grid := make([][]int, len(lines))
	costGrid := make([][]int, len(lines))
	for y, l := range lines {
		for _, c := range l {
			grid[y] = append(grid[y], aoc.Atoi(string(c)))
			costGrid[y] = append(costGrid[y], aoc.Atoi(string(c)))
		}
	}

	// Start from bottom right, bfs to all other nodes in the grid, and
	// incrementally set the minimum 'cost' of that node.

	populateCosts(grid)
	//for y, l := range costGrid {
	//	for x := range l {
	//		fmt.Printf("%4d", costGrid[y][x])
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	//for y, l := range grid {
	//	for x := range l {
	//		fmt.Printf("%4d", grid[y][x])
	//	}
	//	fmt.Println()
	//}

	// seen := make(map[pathStep]int)
	// weight, path := findPath(grid, seen,
	// 	pathStep{c: aoc.Coord2D{X: 0, Y: 0}, dir: aoc.Down, depth: 0},
	// 	grid[0][0], []pathStep{})
	// fmt.Println(weight)
}

func dfSearch(grid [][]int, path []pathStep, curr pathStep, seen map[aoc.Coord2D]bool) int {
	return 0
}

type traversal struct {
	last pathStep
	seen map[aoc.Coord2D]bool // Can we remove and replace with a memoizer
}

//	func (t *traversal) last() pathStep {
//		return t.steps[len(t.steps)-1]
//	}

func (t *traversal) duplicate() traversal {
	var cp traversal
	cp.last = t.last
	cp.seen = maps.Clone(t.seen)
	return cp
}

//
//func (t *traversal) string() string {
//	var s []string
//	for _, step := range t.steps {
//		s = append(s, step.string())
//	}
//	return strings.Join(s, " -> ")
//}

func populateCosts(grid [][]int) {
	start := aoc.Coord2D{}
	end := aoc.Coord2D{X: len(grid[0]) - 1, Y: len(grid) - 1}
	down := pathStep{
		c:          start,
		depth:      1,
		cumulative: grid[0][0],
		dir:        aoc.Down,
	}
	right := pathStep{
		c:          start,
		depth:      1,
		cumulative: grid[0][0],
		dir:        aoc.Right,
	}

	// Idea: Sorting the queue _significantly_ cuts down iterations (assuming because it optimises the memoization but
	// tbh not sure). But appending the next node and then sorting nLogn every iteration is very slow. What about smth
	// like a BST, where insertion is logn, and we always have a sorted queue?
	queue := []traversal{
		{last: down, seen: map[aoc.Coord2D]bool{}},
		{last: right, seen: map[aoc.Coord2D]bool{}},
	}
	minAt := map[string]int{
		down.key():  grid[0][0],
		right.key(): grid[0][0],
	}
	minEnd := 1_000_000_000
	var count int
	t0 := time.Now()
	for len(queue) > 0 {
		count++
		if count%1_000_000 == 0 {
			fmt.Println("ql", len(queue))
			fmt.Println(time.Since(t0))
		}
		curr := queue[0]
		queue = queue[1:]
		if curr.seen[curr.last.c] {
			continue
		} else {
			curr.seen[curr.last.c] = true
		}
		//fmt.Println(curr.string())

		for _, n := range next(curr.last, grid) {
			if !n.valid(grid) {
				continue
			}
			if v, ok := minAt[n.key()]; ok && v < n.cumulative {
				continue
			} else {
				minAt[n.key()] = n.cumulative
			}
			if n.c == end {
				fmt.Println("ma", n.cumulative)
				minEnd = min(minEnd, n.cumulative)
				continue
			}
			nt := curr.duplicate()
			nt.last = n
			queue = append(queue, nt)
		}
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].last.cumulative < queue[j].last.cumulative
		})
	}
	//for y, l := range grid {
	//	for x := range l {
	//		fmt.Printf("%4d", minAt[aoc.Coord2D{x, y}])
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	fmt.Println(count)
	fmt.Println(minEnd)
	//fmt.Println(minAt)
	//type key struct {
	//	c   aoc.Coord2D
	//	d   aoc.Dir
	//	dep int
	//}
	//minAt := make(map[key]int)

	//for len(queue) > 0 {
	//	// fmt.Println()
	//	curr := queue[0]
	//	// fmt.Println(curr, grid[curr.c.Y][curr.c.X])
	//	queue = queue[1:]
	//
	//	newCost := grid[curr.c.Y][curr.c.X] + costGrid[prevC.Y][prevC.X]
	//	// fmt.Println(newCost)
	//	if unchanged(grid, costGrid, curr.c) || costGrid[curr.c.Y][curr.c.X] > newCost {
	//		costGrid[curr.c.Y][curr.c.X] = newCost
	//	}
	//	for _, n := range next(curr) {
	//		if !n.valid(grid) {
	//			continue
	//		}
	//		queue = append(queue, n)
	//	}
	//}
}

func unchanged(grid, costGrid [][]int, c aoc.Coord2D) bool {
	// fmt.Println("check:", grid[c.Y][c.X], costGrid[c.Y][c.X])
	return grid[c.Y][c.X] == costGrid[c.Y][c.X]
}

type pathStep struct {
	c          aoc.Coord2D
	depth      int
	cumulative int
	dir        aoc.Dir
}

func (p pathStep) key() string {
	return fmt.Sprintf("%d,%d,%d,%d", p.c.X, p.c.Y, p.depth, p.dir)
}

func (p pathStep) valid(grid [][]int) bool {
	if p.c.X == 0 && p.c.Y == 0 {
		return false
	}
	if !p.c.WithinPositive(len(grid[0]), len(grid)) {
		return false
	}
	if p.depth >= 3 {
		return false
	}
	return true
}

func (p pathStep) string() string {
	return fmt.Sprintf("(%d;%d), dep:%d, cum:%d, dir:%s", p.c.X, p.c.Y, p.depth, p.cumulative, p.dir.String())
}

func next(curr pathStep, grid [][]int) []pathStep {
	var ps []pathStep
	cl := aoc.NextCoord(curr.c, (curr.dir-1+4)%4)
	cs := aoc.NextCoord(curr.c, curr.dir)
	cr := aoc.NextCoord(curr.c, (curr.dir+1+4)%4)
	if cl.WithinPositive(len(grid[0]), len(grid)) {
		ps = append(ps, pathStep{
			c:          cl,
			depth:      0,
			dir:        (curr.dir - 1 + 4) % 4,
			cumulative: curr.cumulative + grid[cl.Y][cl.X],
		})
	}
	if cs.WithinPositive(len(grid[0]), len(grid)) {
		ps = append(ps, pathStep{
			c:          cs,
			depth:      curr.depth + 1,
			dir:        curr.dir,
			cumulative: curr.cumulative + grid[cs.Y][cs.X],
		})
	}
	if cr.WithinPositive(len(grid[0]), len(grid)) {
		ps = append(ps, pathStep{
			c:          cr,
			depth:      0,
			dir:        (curr.dir + 1 + 4) % 4,
			cumulative: curr.cumulative + grid[cr.Y][cr.X],
		})
	}
	return ps
}

func part2() {
}
