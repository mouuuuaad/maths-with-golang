package complexnums

func EulerFormula(x float64) ComplexNumber {
	return ComplexNumber{
		R: cosC(x),
		I: sinC(x),
	}
}

func (c ComplexNumber) Exp() ComplexNumber {
	ea := expC(c.R)
	return ComplexNumber{
		R: ea * cosC(c.I),
		I: ea * sinC(c.I),
	}
}

func (c ComplexNumber) Log() ComplexNumber {
	return ComplexNumber{
		R: lnC(c.Abs()),
		I: c.Argument(),
	}
}

func Pow(base, exponent ComplexNumber) ComplexNumber {
	if base.Abs() == 0 {
		return New(0, 0)
	}
	lnBase := base.Log()
	prod := exponent.Multiply(lnBase)
	return prod.Exp()
}

func (c ComplexNumber) Sqrt() ComplexNumber {
	p := c.ToPolar()
	return PolarForm{Radius: sqrtC(p.Radius), Theta: p.Theta / 2}.ToComplex()
}

func (c ComplexNumber) PowerN(n int) ComplexNumber {
	if n == 0 {
		return New(1, 0)
	}
	if n < 0 {
		return c.Inverse().PowerN(-n)
	}
	result := New(1, 0)
	base := c
	for n > 0 {
		if n%2 == 1 {
			result = result.Multiply(base)
		}
		base = base.Multiply(base)
		n /= 2
	}
	return result
}

func expC(x float64) float64 {
	if x < 0 {
		return 1.0 / expC(-x)
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 100; i++ {
		term *= x / float64(i)
		sum += term
		if absC(term) < 1e-15 {
			break
		}
	}
	return sum
}

func lnC(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	for i := 0; i < 100; i++ {
		ey := expC(y)
		diff := ey - x
		if absC(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
	}
	return y
}
