package arithmetic

type Real float64

const Epsilon = 1e-9
const PiR Real = 3.14159265358979323846
const ER Real = 2.71828182845904523536
const Phi Real = 1.61803398874989484820

func absF(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func (r Real) IsEqual(other Real) bool {
	return absF(float64(r-other)) < Epsilon
}

func (r Real) Sqrt() Real {
	if r < 0 {
		return 0
	}
	if r == 0 {
		return 0
	}
	x := float64(r)
	z := x
	for i := 0; i < 50; i++ {
		next := 0.5 * (z + x/z)
		if absF(next-z) < 1e-15 {
			return Real(next)
		}
		z = next
	}
	return Real(z)
}

func (r Real) Power(exp Real) Real {
	return Real(expF(float64(exp) * lnF(float64(r))))
}

func expF(x float64) float64 {
	if x < 0 {
		return 1.0 / expF(-x)
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 100; i++ {
		term *= x / float64(i)
		sum += term
		if absF(term) < 1e-15 {
			break
		}
	}
	return sum
}

func lnF(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	for i := 0; i < 100; i++ {
		ey := expF(y)
		diff := ey - x
		if absF(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
	}
	return y
}

func (r Real) Abs() Real {
	if r < 0 {
		return -r
	}
	return r
}

func (r Real) Floor() Integer {
	if r >= 0 {
		return Integer(r)
	}
	i := Integer(r)
	if Real(i) == r {
		return i
	}
	return i - 1
}

func (r Real) Ceil() Integer {
	if r >= 0 {
		i := Integer(r)
		if Real(i) == r {
			return i
		}
		return i + 1
	}
	return Integer(r)
}

func (r Real) Round() Integer {
	if r >= 0 {
		return Integer(r + 0.5)
	}
	return Integer(r - 0.5)
}

func NewtonSqrt(x Real, iterations int) Real {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < iterations; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func CubeRoot(x Real) Real {
	if x == 0 {
		return 0
	}
	sign := Real(1)
	if x < 0 {
		sign = -1
		x = -x
	}
	z := x
	for i := 0; i < 50; i++ {
		next := (2*z + Real(float64(x)/(float64(z)*float64(z)))) / 3
		if (next - z).Abs() < Real(1e-15) {
			return sign * next
		}
		z = next
	}
	return sign * z
}
