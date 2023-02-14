package prime

// Next returns the next prime number after a given number.
func Next(n int) (prime int) {
	if n <= 1 {
		return 2
	}

	prime = n

	for {
		prime++

		if Is(prime) {
			break
		}
	}

	return
}
