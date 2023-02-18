package prime

import "math/big"

// Prev returns the previous prime number before a given 64-bit number.
// It also returns a boolean value that is set to false if no
// prime number exists before the given number.
func Prev(n int64) (prime int64, exists bool) {
	prime = n

	for {
		if prime <= 1 {
			prime = 0
			break
		}

		prime--

		if Is(prime) {
			exists = true
			break
		}
	}

	return
}

// BigPrev returns the previous prime number before
// a given number of arbitrary size.
// It also returns a boolean value that is set to false if no
// prime number exists before the given number.
func BigPrev(bigN *big.Int) (bigPrime *big.Int, exists bool) {
	bigPrime = (new(big.Int)).Set(bigN)

	for {
		if bigPrime.Cmp(bigOne) <= 0 {
			bigPrime = nil
			break
		}

		bigPrime.Sub(bigPrime, bigOne)

		if BigIs(bigPrime) {
			exists = true
			break
		}
	}

	return
}

// Prev returns the previous additive prime number before a given 64-bit number.
// It also returns a boolean value that is set to false if no additive
// prime number exists before the given number.
func PrevAdditive(n int64) (additivePrime int64, exists bool) {
	additivePrime = n

	for {
		if additivePrime <= 1 {
			additivePrime = 0
			break
		}

		additivePrime--

		if IsAdditive(additivePrime) {
			exists = true
			break
		}
	}

	return
}

// BigPrevAdditive returns the previous additive prime number before
// a given number of arbitrary size.
// It also returns a boolean value that is set to false if no additive
// prime number exists before the given number.
func BigPrevAdditive(bigN *big.Int) (bigAdditivePrime *big.Int, exists bool) {
	bigAdditivePrime = (new(big.Int)).Set(bigN)

	for {
		if bigAdditivePrime.Cmp(bigOne) <= 0 {
			bigAdditivePrime = nil
			break
		}

		bigAdditivePrime.Sub(bigAdditivePrime, bigOne)

		if BigIsAdditive(bigAdditivePrime) {
			exists = true
			break
		}
	}

	return
}
