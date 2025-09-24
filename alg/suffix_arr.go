package alg

func suffixArr(s string) []int {
	n := len(s)
	alphabet := 256

	p := make([]int, n)
	c := make([]int, n)
	cnt := make([]int, max(alphabet, n))

	// initial counting sort by single characters
	for i := range n {
		cnt[s[i]]++
	}
	for i := 1; i < alphabet; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := range n {
		cnt[s[i]]--
		p[cnt[s[i]]] = i
	}
	c[p[0]] = 0
	classes := 1
	for i := 1; i < n; i++ {
		if s[p[i]] != s[p[i-1]] {
			classes++
		}
		c[p[i]] = classes - 1
	}

	pn := make([]int, n)
	cn := make([]int, n)

	for h := 0; (1 << h) < n; h++ {
		for i := range n {
			pn[i] = p[i] - (1 << h)
			if pn[i] < 0 {
				pn[i] += n
			}
		}

		for i := 0; i < classes; i++ {
			cnt[i] = 0
		}
		for i := range n {
			cnt[c[pn[i]]]++
		}
		for i := 1; i < classes; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			cl := c[pn[i]]
			cnt[cl]--
			p[cnt[cl]] = pn[i]
		}

		cn[p[0]] = 0
		classes = 1
		for i := 1; i < n; i++ {
			cur := [2]int{c[p[i]], c[(p[i]+(1<<h))%n]}
			prev := [2]int{c[p[i-1]], c[(p[i-1]+(1<<h))%n]}
			if cur != prev {
				classes++
			}
			cn[p[i]] = classes - 1
		}
		copy(c, cn)
	}

	return p
}func sortCyclicShifts(s string) []int {
	n := len(s)
	alphabet := 256

	p := make([]int, n)
	c := make([]int, n)
	cnt := make([]int, max(alphabet, n))

	// initial counting sort by single characters
	for i := range n {
		cnt[s[i]]++
	}
	for i := 1; i < alphabet; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := range n {
		cnt[s[i]]--
		p[cnt[s[i]]] = i
	}
	c[p[0]] = 0
	classes := 1
	for i := 1; i < n; i++ {
		if s[p[i]] != s[p[i-1]] {
			classes++
		}
		c[p[i]] = classes - 1
	}

	pn := make([]int, n)
	cn := make([]int, n)

	for h := 0; (1 << h) < n; h++ {
		for i := range n {
			pn[i] = p[i] - (1 << h)
			if pn[i] < 0 {
				pn[i] += n
			}
		}

		for i := 0; i < classes; i++ {
			cnt[i] = 0
		}
		for i := range n {
			cnt[c[pn[i]]]++
		}
		for i := 1; i < classes; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			cl := c[pn[i]]
			cnt[cl]--
			p[cnt[cl]] = pn[i]
		}

		cn[p[0]] = 0
		classes = 1
		for i := 1; i < n; i++ {
			cur := [2]int{c[p[i]], c[(p[i]+(1<<h))%n]}
			prev := [2]int{c[p[i-1]], c[(p[i-1]+(1<<h))%n]}
			if cur != prev {
				classes++
			}
			cn[p[i]] = classes - 1
		}
		copy(c, cn)
	}

	return p
}
