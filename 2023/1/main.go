package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	adventofc2023 "github.com/wesrobin/adventofc/2023"
)

/**
The newly-improved calibration document consists of lines of text; each line originally contained a
specific calibration value that the Elves now need to recover. On each line, the calibration value can
be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
*/

//go:embed input.txt
var input string

var testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int64 {
	var tot int64
	for _, line := range strings.Split(input, "\n") {
		var num string
		for i := 0; i < len(line); i++ {
			if line[i]-48 < 10 {
				num += string(line[i])
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i]-48 < 10 {
				num += string(line[i])
				break
			}
		}
		tot += adventofc2023.Atoi(num)
	}
	return tot
}

var dict = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type queue []string

func (q *queue) push(r string) {
	*q = append(*q, r)
	if len(*q) > 5 {
		*q = (*q)[1:]
	}
}

func (q *queue) match() (int, bool) {
	word := strings.Join(*q, "")
	for k, v := range dict {
		if strings.Contains(word, k) {
			return v, true
		}
	}
	return 0, false
}

type stack []string

func (s *stack) push(r string) {
	*s = append([]string{r}, *s...)
	if len(*s) > 5 {
		*s = (*s)[:5]
	}
}

func (s *stack) match() (int, bool) {
	var word string
	for i := 0; i < len(*s); i++ {
		word += (*s)[i]
		if num, ok := dict[word]; ok {
			return num, true
		}
	}
	return 0, false
}

func part2() int64 {
	var tot int64
	for _, line := range strings.Split(input, "\n") {
		var (
			num string
			q   queue
			s   stack
		)
		for i := 0; i < len(line); i++ {
			if line[i]-48 < 10 {
				num += string(line[i])
				break
			} else {
				q.push(string(line[i]))
				if n, ok := q.match(); ok {
					num += strconv.Itoa(n)
					break
				}
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i]-48 < 10 {
				num += string(line[i])
				break
			} else {
				s.push(string(line[i]))
				if n, ok := s.match(); ok {
					num += strconv.Itoa(n)
					break
				}
			}
		}
		tot += adventofc2023.Atoi(num)
	}
	return tot
}
