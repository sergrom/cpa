package alg

func modPow(x, e, m int) int {
    if m == 1 {
        return 0
    }
    if e == 0 {
        return 1 % m
    }

    x %= m

    p := modPow(x, e/2, m)
    p = (p * p) % m

    if e&1 != 0 {
        return (p * x) % m
    }
    return p
}

func modPow64(x, e, m int64) int64 {
    x %= m
    res := int64(1)

    for e > 0 {
        if e&1 != 0 {
            res = res * x % m
        }
        x = x * x % m
        e >>= 1
    }

    return res
}
