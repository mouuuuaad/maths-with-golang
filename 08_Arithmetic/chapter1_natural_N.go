package arithmetic

type Natural uint64

func IsNatural(n int64) bool {
	return n >= 0
}

func FactorialN(n Natural) Natural {
	if n <= 1 {
		return 1
	}
	result := Natural(1)
	for i := Natural(2); i <= n; i++ {
		result *= i
	}
	return result
}

func PermutationN(n, k Natural) Natural {
	if k > n {
		return 0
	}
	result := Natural(1)
	for i := n; i > n-k; i-- {
		result *= i
	}
	return result
}

func CombinationN(n, k Natural) Natural {
	if k > n {
		return 0
	}
	if k > n/2 {
		k = n - k
	}
	result := Natural(1)
	for i := Natural(1); i <= k; i++ {
		result = result * (n - i + 1) / i
	}
	return result
}

func PowerN(base, exp Natural) Natural {
	if exp == 0 {
		return 1
	}
	result := Natural(1)
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func SumFirstN(n Natural) Natural {
	return n * (n + 1) / 2
}

func SumSquaresFirstN(n Natural) Natural {
	return n * (n + 1) * (2*n + 1) / 6
}

func SumCubesFirstN(n Natural) Natural {
	s := SumFirstN(n)
	return s * s
}

func FibonacciN(n Natural) Natural {
	if n <= 1 {
		return n
	}
	a, b := Natural(0), Natural(1)
	for i := Natural(2); i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func CatalanN(n Natural) Natural {
	if n == 0 {
		return 1
	}
	c := Natural(1)
	for i := Natural(1); i <= n; i++ {
		c = c * 2 * (2*i - 1) / (i + 1)
	}
	return c
}
