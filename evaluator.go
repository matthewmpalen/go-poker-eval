package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//  This routine will search a deck for a specific card
//  (specified by rank/suit), and return the INDEX giving
//  the position of the found card.  If it is not found,
//  then it returns an error
func findCard(rank, suit int, deck []int) (int, error) {
	for i, card := range deck {
		if ((card & suit) != 0) && (Rank(card) == rank) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Could not find card: rank=%d, suit=%d", rank, suit)
}

//  This routine takes a deck and randomly mixes up
//  the order of the cards.
func shuffleDeck(deck []int) {
	temp := make([]int, len(deck))
	_ = copy(deck, temp)
	perm := rand.Perm(len(deck))

	for i, v := range perm {
		deck[i] = temp[v]
	}
}

func handRank(x int) int {
	switch {
	case x > 6185:
		return HighCard // 1277 high card
	case x > 3325:
		return OnePair // 2860 one pair
	case x > 2467:
		return TwoPair //  858 two pair
	case x > 1609:
		return ThreeOfAKind //  858 three-kind
	case x > 1599:
		return Straight //   10 straights
	case x > 322:
		return Flush // 1277 flushes
	case x > 166:
		return FullHouse //  156 full house
	case x > 10:
		return FourOfAKind //  156 four-kind
	default:
		return StraightFlush //   10 straight-flushes
	}
}

func eval5Hand(hand []int) int {
	c1, c2, c3, c4, c5 := hand[0], hand[1], hand[2], hand[3], hand[4]
	q := (c1 | c2 | c3 | c4 | c5) >> 16

	// check for Flushes and StraightFlushes
	if (c1 & c2 & c3 & c4 & c5 & 0xF000) != 0 {
		return flushes[q]
	}

	// check for Straights and HighCard hands
	if s := unique5[q]; s != 0 {
		return s
	}

	// let's do it the hard way (binary search)
	q = (c1 & 0xFF) * (c2 & 0xFF) * (c3 & 0xFF) * (c4 & 0xFF) * (c5 & 0xFF)
	searchFunc := func(i int) bool { return products[i] >= q }
	q = sort.Search(len(products), searchFunc)

	return values[q]
}
