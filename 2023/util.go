package adventofc2023

import "strconv"

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
