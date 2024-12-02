package aoc24

import (
	"strconv"
	"strings"
)

func Atoi64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func Atoi(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

type Coord2D struct {
	X, Y int
}

func (c Coord2D) WithinPositive(xMax, yMax int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < xMax && c.Y < yMax
}

type Coord3D struct {
	X, Y, Z int
}

type Dir int

const (
	Up    Dir = 0
	Left  Dir = 1
	Down  Dir = 2
	Right Dir = 3
)

func (d Dir) Invert() Dir {
	return (d + 2) % 4
}

func (d Dir) String() string {
	switch d {
	case Up:
		return "up"
	case Left:
		return "le"
	case Down:
		return "do"
	case Right:
		return "ri"
	}
	return "none"
}

func (d Dir) Char() rune {
	switch d {
	case Up:
		return '^'
	case Left:
		return '<'
	case Down:
		return 'v'
	case Right:
		return '>'
	}
	return '?'
}

func PrevCoord(c Coord2D, d Dir) Coord2D {
	return NextCoord(c, d.Invert())
}

func NextCoord(c Coord2D, d Dir) Coord2D {
	x := c.X
	y := c.Y
	switch d {
	case Up:
		y -= 1
	case Left:
		x -= 1
	case Down:
		y += 1
	case Right:
		x += 1
	}
	return Coord2D{X: x, Y: y}
}

// MapSlice accepts a slice of type T1 and returns a slice of type T2, executing fn on each
// element of T1. Returns nil if input is empty or nil.
func MapSlice[T1, T2 any](s []T1, fn func(T1) T2) []T2 {
	if len(s) == 0 {
		return nil
	}
	res := make([]T2, len(s))
	for i, elem := range s {
		res[i] = fn(elem)
	}
	return res
}

func MapIntsToStr(ss []string) []int {
	return MapSlice(ss, Atoi)
}

// MapSliceErr is similar to MapSliceErr, except it will return an error state if the loop had a failure.
// This is useful when converting something like protobuf files to normal go structs, the protocp convert methods all
// return errors.
func MapSliceErr[T1, T2 any](s []T1, fn func(T1) (T2, error)) ([]T2, error) {
	if len(s) == 0 {
		return nil, nil
	}

	res := make([]T2, len(s))
	for i, elem := range s {
		r, err := fn(elem)
		if err != nil {
			return nil, err
		}

		res[i] = r
	}

	return res, nil
}

func CutRight(s, cut string) string {
	_, res, _ := strings.Cut(s, cut)
	return res
}

func CutLeft(s, cut string) string {
	res, _, _ := strings.Cut(s, cut)
	return res
}

type Stack[T any] []T

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		return *(new(T)), false
	}
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e, true
}

func (s *Stack[T]) Top() (T, bool) {
	return s.Peek(0)
}

func (s *Stack[T]) Peek(depth int) (T, bool) {
	if len(*s) == 0 || depth >= len(*s) {
		return *(new(T)), false
	}
	return (*s)[len(*s)-1-depth], true
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
