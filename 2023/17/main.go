package main

import (
	"container/heap"
	_ "embed"
	"fmt"
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

	populateCosts(grid)
}

type traversal struct {
	last pathStep
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

func populateCosts(grid [][]int) {
	start := aoc.Coord2D{}
	end := aoc.Coord2D{X: len(grid[0]) - 1, Y: len(grid) - 1}
	down := pathStep{
		c:     start,
		depth: 1,
		dir:   aoc.Down,
	}
	right := pathStep{
		c:     start,
		depth: 1,
		dir:   aoc.Right,
	}

	// Idea: Sorting the queue _significantly_ cuts down iterations (assuming because it optimises the memoization but
	// tbh not sure). But appending the next node and then sorting nLogn every iteration is very slow. What about smth
	// like a BST, where insertion is logn, and we always have a sorted queue?
	queue := TraversalQueue{}
	queue.Push(&traversal{last: down})
	queue.Push(&traversal{last: right})

	heap.Init(&queue)

	//minAt := map[key]int{
	//	down.key():  grid[0][0],
	//	right.key(): grid[0][0],
	//}
	seen := map[key]bool{}
	minEnd := 1_000_000_000
	var count int
	t0 := time.Now()

OuterLoop:
	for len(queue) > 0 {
		count++
		if count%1_000_000 == 0 {
			fmt.Println("ql", len(queue))
			fmt.Println(time.Since(t0))
		}
		curr := heap.Pop(&queue).(*traversal)

		for _, n := range next(curr.last, grid) {
			if !n.valid(grid) {
				continue
			}

			if n.c == end {
				if n.cumulative < minEnd {
					minEnd = n.cumulative
				}
				break OuterLoop
			}
			if seen[n.key()] {
				continue
			}
			seen[n.key()] = true

			heap.Push(&queue, &traversal{last: n})
		}
	}
	fmt.Println(count)
	fmt.Println(minEnd)
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
	if p.depth >= 3 {
		return false
	}
	if (p.c.X == 0 || p.c.X == len(grid[0])-1) && p.dir == aoc.Up {
		return false
	}
	if (p.c.Y == 0 || p.c.Y == len(grid)-1) && p.dir == aoc.Left {
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
