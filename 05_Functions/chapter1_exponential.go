package functions

const Pi = 3.14159265358979323846
const E = 2.71828182845904523536

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Exp(x float64) float64 {
	if x < 0 {
		return 1.0 / Exp(-x)
	}
	n := 0
	for x > 1 {
		x /= 2
		n++
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 100; i++ {
		term *= x / float64(i)
		sum += term
		if Abs(term) < 1e-15 {
			break
		}
	}
	for i := 0; i < n; i++ {
		sum *= sum
	}
	return sum
}

func Power(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	if exp == float64(int(exp)) {
		return powerInt(base, int(exp))
	}
	if base < 0 {
		return 0
	}
	return Exp(exp * Ln(base))
}

func powerInt(base float64, exp int) float64 {
	if exp < 0 {
		return 1.0 / powerInt(base, -exp)
	}
	res := 1.0
	for exp > 0 {
		if exp%2 == 1 {
			res *= base
		}
		base *= base
		exp /= 2
	}
	return res
}

func ExpTaylor(x float64, terms int) float64 {
	sum := 1.0
	term := 1.0
	for i := 1; i < terms; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func PowerRecursive(base float64, exp int) float64 {
	if exp == 0 {
		return 1
	}
	if exp < 0 {
		return 1.0 / PowerRecursive(base, -exp)
	}
	if exp%2 == 0 {
		half := PowerRecursive(base, exp/2)
		return half * half
	}
	return base * PowerRecursive(base, exp-1)
}
