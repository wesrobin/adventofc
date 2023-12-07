package main

import (
	_ "embed"
	"fmt"
	adventofc2023 "github.com/wesrobin/adventofc/2023"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

var testInput1 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var testInput2 = ``

func main() {
	fmt.Println(part1(input)) // 253410429 too high, 252960917 too low, 253000392 bad
	fmt.Println(part2(input))
}

type hand struct {
	cards []card
}

func (h hand) beats(a hand) bool {
	hs := fromHand(h)
	as := fromHand(a)
	if hs > as {
		return true
	} else if hs < as {
		return false
	}

	for i := range h.cards {
		if h.cards[i] > a.cards[i] {
			return true
		} else if h.cards[i] < a.cards[i] {
			return false
		}
	}

	return false
}

func handFromStr(s string) hand {
	var cs []card
	for _, r := range s {
		cs = append(cs, dict[r])
	}
	return hand{cards: cs}
}

type special int

const (
	fiveKind  special = 6
	fourKind  special = 5
	fullHouse special = 4
	threekind special = 3
	twoPair   special = 2
	onePair   special = 1
	none      special = 0
)

func fromHand(h hand) special {
	uniques := make([]card, 0)
	counts := make(map[card]int)
	for _, c := range h.cards {
		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c]++
			uniques = append(uniques, c)
		}
	}
	if len(counts) == 1 {
		return fiveKind
	}
	if len(counts) == 2 {
		// Full house or four of a kind
		if counts[uniques[0]] == 2 || counts[uniques[0]] == 3 {
			return fullHouse
		}
		return fourKind
	}
	if len(counts) == 3 {
		if counts[uniques[0]] == 3 || counts[uniques[1]] == 3 || counts[uniques[2]] == 3 {
			return threekind
		}
		if (counts[uniques[0]] == 2 && counts[uniques[1]] == 2) ||
			(counts[uniques[1]] == 2 && counts[uniques[2]] == 2) ||
			(counts[uniques[2]] == 2 && counts[uniques[0]] == 2) {
			return twoPair
		}
	}
	if len(counts) == 4 {
		return onePair
	}
	return none
}

type card int

const (
	ace   card = 14
	king  card = 13
	queen card = 12
	jack  card = 11
	ten   card = 10
	nine  card = 9
	eight card = 8
	seven card = 7
	six   card = 6
	five  card = 5
	four  card = 4
	three card = 3
	two   card = 2
	joker card = 1
)

func (c card) equal(a card) bool {
	return c == a
}

func (c card) beats(a card) bool {
	return c > a
}

var dict = map[rune]card{
	'A': ace,
	'K': king,
	'Q': queen,
	'J': jack,
	'T': ten,
	'9': nine,
	'8': eight,
	'7': seven,
	'6': six,
	'5': five,
	'4': four,
	'3': three,
	'2': two,
}

func part1(inp string) any {
	type bet struct {
		h   hand
		bet int
	}
	var bets []bet
	for _, line := range strings.Split(inp, "\n") {
		h, b, _ := strings.Cut(line, " ")
		bets = append(bets, bet{h: handFromStr(h), bet: adventofc2023.Atoi(b)})
	}

	sort.Slice(bets, func(i, j int) bool {
		return bets[j].h.beats(bets[i].h)
	})

	var sum int
	for i, b := range bets {
		sum += (i + 1) * b.bet
	}
	return sum
}

func part2(inp string) any {
	type bet struct {
		h   hand
		bet int
	}
	var bets []bet
	for _, line := range strings.Split(inp, "\n") {
		h, b, _ := strings.Cut(line, " ")
		bets = append(bets, bet{h: handFromStr(h), bet: adventofc2023.Atoi(b)})
	}

	sort.Slice(bets, func(i, j int) bool {
		return bets[j].h.beats2(bets[i].h)
	})

	var sum int
	for i, b := range bets {
		sum += (i + 1) * b.bet
	}
	return sum
}

func (h hand) beats2(a hand) bool {
	hs := fromHand2(h)
	as := fromHand2(a)
	if hs > as {
		return true
	} else if hs < as {
		return false
	}

	for i := range h.cards {
		hC, aC := h.cards[i], a.cards[i]
		if hC == jack {
			hC = joker
		}
		if aC == jack {
			aC = joker
		}
		if hC > aC {
			return true
		} else if hC < aC {
			return false
		}
	}

	return false
}

func fromHand2(h hand) special {
	uniques := make([]card, 0)
	counts := make(map[card]int)
	for _, c := range h.cards {
		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c]++
			uniques = append(uniques, c)
		}
	}

	// Do Joker stuffs
	if _, ok := counts[jack]; ok {
		var (
			most    int
			c       card
			jackPos int
		)
		for i, uniq := range uniques {
			if uniq == jack {
				jackPos = i
				continue
			}
			if counts[uniq] > most {
				c = uniq
				most = counts[uniq]
			}
		}
		counts[c] += counts[jack]
		delete(counts, jack)
		uniques = append(uniques[:jackPos], uniques[jackPos+1:]...)
	}

	if len(counts) == 1 {
		return fiveKind
	}
	if len(counts) == 2 {
		// Full house or four of a kind
		if counts[uniques[0]] == 2 || counts[uniques[0]] == 3 {
			return fullHouse
		}
		return fourKind
	}
	if len(counts) == 3 {
		if counts[uniques[0]] == 3 || counts[uniques[1]] == 3 || counts[uniques[2]] == 3 {
			return threekind
		}
		if (counts[uniques[0]] == 2 && counts[uniques[1]] == 2) ||
			(counts[uniques[1]] == 2 && counts[uniques[2]] == 2) ||
			(counts[uniques[2]] == 2 && counts[uniques[0]] == 2) {
			return twoPair
		}
	}
	if len(counts) == 4 {
		return onePair
	}
	return none
}
