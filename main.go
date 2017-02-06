package main

import "fmt"

var (
	deck = make([]int, DeckSize)
	hand = make([]int, HandSize)
	freq = make([]int, 10)
)

//   This routine initializes the deck.  A deck of cards is
//   simply an integer array of length 52 (no jokers).  This
//   array is populated with each card, using the following
//   scheme:
//
//   An integer is made up of four bytes.  The high-order
//   bytes are used to hold the rank bit pattern, whereas
//   the low-order bytes hold the suit/rank/prime value
//   of the card.
//
//   +--------+--------+--------+--------+
//   |xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
//   +--------+--------+--------+--------+
//
//   p = prime number of rank (deuce=2,trey=3,four=5,five=7,...,ace=41)
//   r = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
//   cdhs = suit of card
//   b = bit turned on depending on rank of card
//
func init() {
	// Initialize deck
	n := 0
	suit := Spade
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			j1 := uint(j)
			card := primes[j] | (j1 << 8) | suit | (1 << (16 + j1))
			deck[n] = int(card)
			n++
		}
		suit <<= 1
	}
}

func main() {
	// Loop over every possible five-card hand.
	for a := 0; a < 48; a++ {
		hand[0] = deck[a]

		for b := a + 1; b < 49; b++ {
			hand[1] = deck[b]

			for c := b + 1; c < 50; c++ {
				hand[2] = deck[c]

				for d := c + 1; d < 51; d++ {
					hand[3] = deck[d]

					for e := d + 1; e < 52; e++ {
						hand[4] = deck[e]

						i := eval5Hand(hand)
						j := handRank(i)
						freq[j]++
					}
				}
			}
		}
	}

	for i := 1; i <= 9; i++ {
		fmt.Printf("%15s: %8d\n", ValueStr[i], freq[i])
	}
}
