package complexnums

type ComplexNumber struct {
	R float64
	I float64
}

func New(real, imag float64) ComplexNumber {
	return ComplexNumber{R: real, I: imag}
}

func (c ComplexNumber) Real() float64 {
	return c.R
}

func (c ComplexNumber) Imag() float64 {
	return c.I
}

func (c ComplexNumber) IsPureReal() bool {
	return absC(c.I) < 1e-9
}

func (c ComplexNumber) IsPureImaginary() bool {
	return absC(c.R) < 1e-9
}

func (c ComplexNumber) IsZero() bool {
	return absC(c.R) < 1e-9 && absC(c.I) < 1e-9
}

func absC(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
