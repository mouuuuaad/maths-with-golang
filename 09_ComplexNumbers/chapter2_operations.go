package complexnums

func (c ComplexNumber) Add(other ComplexNumber) ComplexNumber {
	return ComplexNumber{R: c.R + other.R, I: c.I + other.I}
}

func (c ComplexNumber) Subtract(other ComplexNumber) ComplexNumber {
	return ComplexNumber{R: c.R - other.R, I: c.I - other.I}
}

func (c ComplexNumber) Multiply(other ComplexNumber) ComplexNumber {
	return ComplexNumber{
		R: c.R*other.R - c.I*other.I,
		I: c.R*other.I + c.I*other.R,
	}
}

func (c ComplexNumber) Divide(other ComplexNumber) ComplexNumber {
	denom := other.R*other.R + other.I*other.I
	if denom == 0 {
		return ComplexNumber{0, 0}
	}
	return ComplexNumber{
		R: (c.R*other.R + c.I*other.I) / denom,
		I: (c.I*other.R - c.R*other.I) / denom,
	}
}

func (c ComplexNumber) Conjugate() ComplexNumber {
	return ComplexNumber{R: c.R, I: -c.I}
}

func (c ComplexNumber) Inverse() ComplexNumber {
	denom := c.R*c.R + c.I*c.I
	if denom == 0 {
		return ComplexNumber{0, 0}
	}
	return ComplexNumber{R: c.R / denom, I: -c.I / denom}
}

func (c ComplexNumber) Abs() float64 {
	return sqrtC(c.R*c.R + c.I*c.I)
}

func (c ComplexNumber) AbsSquared() float64 {
	return c.R*c.R + c.I*c.I
}

func (c ComplexNumber) Scale(s float64) ComplexNumber {
	return ComplexNumber{R: c.R * s, I: c.I * s}
}

func (c ComplexNumber) Negate() ComplexNumber {
	return ComplexNumber{R: -c.R, I: -c.I}
}

func sqrtC(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}
