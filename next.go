package prime

import "math"

// Next returns the next prime number after a given number.
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
