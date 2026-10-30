package alg

import "math"

func getDivisors(n int) []int {
	sq := int(math.Sqrt(float64(n)))
	divs := []int{}
	for i := 1; i <= sq; i++ {
		if n%i == 0 {
			if n/i == i {
				divs = append(divs, i)
			} else {
				divs = append(divs, i, n/i)
			}
		}
	}

	return divs
}
