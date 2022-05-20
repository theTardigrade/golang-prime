package prime

func Next(n int) int {
	if n <= 1 {
		return 2
	}

	prime := n
	found := false

	for !found {
		prime++

		if Is(prime) {
			found = true
		}
	}

	return prime
}
