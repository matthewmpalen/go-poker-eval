package main

const (
	DeckSize = 52
	HandSize = 5
)

const (
	Spade uint = 0x1000 << iota
	Heart
	Diamond
	Club
)

const (
	Deuce = iota
	Trey
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	StraightFlush = iota + 1
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

const (
	rankStr = "23456789TJQKA"
)

var ValueStr = []string{
	"",
	"Straight Flush",
	"Four of a Kind",
	"Full House",
	"Flush",
	"Straight",
	"Three of a Kind",
	"Two Pair",
	"One Pair",
	"High Card",
}

func Rank(x int) int {
	return ((x >> 8) & 0xF)
}
