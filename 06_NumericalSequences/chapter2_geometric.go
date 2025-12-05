package sequences

func absS(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func powerS(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	if base <= 0 {
		return 0
	}
	return expS(exp * lnS(base))
}

func expS(x float64) float64 {
	if x < 0 {
		return 1.0 / expS(-x)
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 50; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func lnS(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	for i := 0; i < 50; i++ {
		ey := expS(y)
		diff := ey - x
		if absS(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
	}
	return y
}

type GeometricSequence struct {
	Start float64
	Ratio float64
}

func (s GeometricSequence) NthTerm(n int) float64 {
	return s.Start * powerS(s.Ratio, float64(n))
}

func (s GeometricSequence) SumN(n int) float64 {
	if s.Ratio == 1 {
		return s.Start * float64(n+1)
	}
	return s.Start * (1 - powerS(s.Ratio, float64(n+1))) / (1 - s.Ratio)
}

func (s GeometricSequence) SumInfinite() float64 {
	if absS(s.Ratio) >= 1 {
		return 0
	}
	return s.Start / (1 - s.Ratio)
}

func (s GeometricSequence) Generate(n int) []float64 {
	result := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = s.NthTerm(i)
	}
	return result
}

func GeometricMean(a, b float64) float64 {
	if a*b < 0 {
		return 0
	}
	return sqrtS(a * b)
}

func sqrtS(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}

func InsertGeometricMeans(a, b float64, n int) []float64 {
	if a <= 0 || b <= 0 {
		return nil
	}
	r := powerS(b/a, 1.0/float64(n+1))
	result := make([]float64, n+2)
	result[0] = a
	for i := 1; i <= n; i++ {
		result[i] = a * powerS(r, float64(i))
	}
	result[n+1] = b
	return result
}
