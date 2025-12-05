package sequences

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibonacciClosedForm(n int) float64 {
	sqrt5 := sqrtS(5)
	phi := (1 + sqrt5) / 2
	psi := (1 - sqrt5) / 2
	return (powerS(phi, float64(n)) - powerS(psi, float64(n))) / sqrt5
}

func Catalan(n int) int {
	if n <= 0 {
		return 1
	}
	c := 1
	for i := 1; i <= n; i++ {
		c = c * 2 * (2*i - 1) / (i + 1)
	}
	return c
}

func Lucas(n int) int {
	if n == 0 {
		return 2
	}
	if n == 1 {
		return 1
	}
	a, b := 2, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

func Pell(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, 2*b+a
	}
	return b
}

func Bernoulli(n int) float64 {
	B := make([]float64, n+1)
	for m := 0; m <= n; m++ {
		B[m] = 1.0 / float64(m+1)
		for j := m; j >= 1; j-- {
			B[j-1] = float64(j) * (B[j-1] - B[j])
		}
	}
	return B[0]
}

func Euler(n int) float64 {
	E := make([]float64, n+1)
	E[0] = 1
	for m := 2; m <= n; m += 2 {
		sum := 0.0
		for k := 0; k < m; k += 2 {
			binomial := 1.0
			for j := 1; j <= k; j++ {
				binomial *= float64(m-j+1) / float64(j)
			}
			sum += binomial * E[k]
		}
		E[m] = -sum
	}
	return E[n]
}
