package alg

import "fmt"

func sieve(n int) {
	prime := make([]bool, n+1)
	for i := range prime {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if prime[p] {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	// Print all prime numbers
	for p := 2; p <= n; p++ {
		if prime[p] {
			fmt.Print(p, " ")
		}
	}

	fmt.Println()
}
