package adventofc2023

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

type Coord3D struct {
	X, Y, Z int
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
