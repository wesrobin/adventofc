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
	North     Dir = 1
	West      Dir = 2
	South     Dir = 3
	East      Dir = 4
	NorthWest Dir = 5
	SouthWest Dir = 6
	SouthEast Dir = 7
	NorthEast Dir = 8
)

func (d Dir) String() string {
	switch d {
	case North:
		return "n"
	case West:
		return "w"
	case South:
		return "s"
	case East:
		return "e"
	case NorthWest:
		return "nw"
	case SouthWest:
		return "sw"
	case SouthEast:
		return "se"
	case NorthEast:
		return "ne"
	}
	return "none"
}

func (d Dir) Char() rune {
	switch d {
	case North:
		return '^'
	case West:
		return '<'
	case South:
		return 'v'
	case East:
		return '>'
	}
	return '?'
}

//func PrevCoord(c Coord2D, d Dir) Coord2D {
//	return NextCoord(c, d.Invert())
//}

func NextCoord(c Coord2D, d Dir) Coord2D {
	x := c.X
	y := c.Y
	switch d {
	case North:
		y -= 1
	case West:
		x -= 1
	case South:
		y += 1
	case East:
		x += 1
	case NorthWest:
		y -= 1
		x -= 1
	case SouthWest:
		y += 1
		x -= 1
	case SouthEast:
		y += 1
		x += 1
	case NorthEast:
		y -= 1
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
