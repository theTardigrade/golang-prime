package prime

func Next(n int) int {
	if n <= 1 {
		return 2
	}

	prime := n

	for {
		prime++

		if Is(prime) {
			break
		}
	}

	return prime
}
