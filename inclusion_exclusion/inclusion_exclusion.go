package inclusionexclusion

func solve(n, r int) int {
	p := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			p = append(p, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		p = append(p, n)
	}

	sum := 0
	for msk := 1; msk < (1 << len(p)); msk++ {
		mult, bits := 1, 0
		for i := 0; i < len(p); i++ {
			if msk&(1<<i) == 1 {
				bits++
				mult *= p[i]
			}
		}
		cur := r / mult
		if bits%2 == 1 {
			sum += cur
		} else {
			sum -= cur
		}
	}

	return r - sum
}
