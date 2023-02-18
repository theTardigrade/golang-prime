package prime

import (
	"math"
	"math/big"
)

// Is determines whether a given 64-bit number is prime.
func Is(n int64) bool {
	if n >= int64(math.Sqrt(math.MaxInt64)) {
		bigN := big.NewInt(n)

		return BigIs(bigN)
	}

	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

// BigIs determines whether a given number of arbitrary size is prime.
func BigIs(bigN *big.Int) bool {
	if bigN.Cmp(bigOne) <= 0 {
		return false
	}

	if bigN.Cmp(bigThree) <= 0 {
		return true
	}

	bigTemp := new(big.Int)

	if bigTemp.Mod(bigN, bigTwo).Cmp(bigZero) == 0 || bigTemp.Mod(bigN, bigThree).Cmp(bigZero) == 0 {
		return false
	}

	for bigI := big.NewInt(5); bigTemp.Mul(bigI, bigI).Cmp(bigN) <= 0; bigI.Add(bigI, bigSix) {
		if bigTemp.Mod(bigN, bigI).Cmp(bigZero) == 0 {
			return false
		}

		bigTemp.Add(bigI, bigTwo)

		if bigTemp.Mod(bigN, bigTemp).Cmp(bigZero) == 0 {
			return false
		}
	}

	return true
}

// IsAdditive determines whether a given 64-bit number is additively prime.
// This means that the number itself must be prime and the sum of
// its decimal digits must also be prime.
func IsAdditive(n int64) bool {
	bigN := big.NewInt(n)

	return BigIsAdditive(bigN)
}

// BigIsAdditive determines whether a given number of arbitrary size
// is additively prime.
// This means that the number itself must be prime and the sum of
// its decimal digits must also be prime.
func BigIsAdditive(bigN *big.Int) bool {
	if !BigIs(bigN) {
		return false
	}

	bigDigitSum := new(big.Int)
	bigTemp := new(big.Int)

	for bigI := (new(big.Int)).Set(bigN); bigI.Cmp(bigZero) > 0; bigI.Div(bigI, bigTen) {
		bigDigitSum.Add(bigDigitSum, bigTemp.Mod(bigI, bigTen))
	}

	return BigIs(bigDigitSum)
}
