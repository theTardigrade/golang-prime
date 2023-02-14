package prime

// Prev returns the previous prime number before a given number.
// It also returns a boolean value that is set to false if no
// prime number exists before the given number.
func Prev(n int64) (prime int64, exists bool) {
	prime = n

	for {
		if prime--; prime <= 1 {
			prime = 0
			break
		}

		if Is(prime) {
			exists = true
			break
		}
	}

	return
}
