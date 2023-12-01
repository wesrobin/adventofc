package adventofc2023

import "strconv"

func Atoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
