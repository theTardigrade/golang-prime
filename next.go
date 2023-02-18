package prime

import (
	"math"
	"math/big"
)

// Next returns the next prime number after a given 64-bit number.
// It also returns a boolean value that is set to false if no
// prime number exists after the given number &mdash;
// not because there is no such number in the mathematical set of integers,
// but rather because there is no such number that can be represented
// by a 64-bit integer in Go.
func Next(n int64) (prime int64, exists bool) {
	if n <= 1 {
		prime, exists = 2, true
		return
	}

	prime = n

	for {
		if prime == math.MaxInt64 {
			prime = 0
			break
		}

		prime++

		if Is(prime) {
			exists = true
			break
		}
	}

	return
}

// BigNext returns the next prime number after a given number of arbitrary size.
func BigNext(bigN *big.Int) (bigPrime *big.Int) {
	if bigN.Cmp(bigOne) <= 0 {
		bigPrime = big.NewInt(2)
		return
	}

	bigPrime = (new(big.Int)).Set(bigN)

	for {
		bigPrime.Add(bigPrime, bigOne)

		if BigIs(bigPrime) {
			break
		}
	}

	return
}
