package prime

import (
	"math/big"
)

// Is determines whether a given number is prime.
func Is(n int64) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	bigN := big.NewInt(n)
	bigI := big.NewInt(5)
	bigZero := big.NewInt(0)
	bigTwo := big.NewInt(2)
	bigSix := big.NewInt(6)
	bigTemp := new(big.Int)

	for bigTemp.Mul(bigI, bigI).Cmp(bigN) <= 0 {
		if bigTemp.Mod(bigN, bigI) == bigZero {
			return false
		}

		bigTemp.Add(bigI, bigTwo)

		if bigTemp.Mod(bigN, bigTemp) == bigZero {
			return false
		}

		bigI.Add(bigI, bigSix)
	}

	return true
}
