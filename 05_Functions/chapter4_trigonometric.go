package functions

func Sin(x float64) float64 {
	k := int((x + Pi) / (2 * Pi))
	x -= float64(k) * 2 * Pi
	sum := 0.0
	term := x
	x2 := x * x
	for i := 1; i < 50; i++ {
		sum += term
		term *= -x2 / float64(2*i*(2*i+1))
		if Abs(term) < 1e-15 {
			break
		}
	}
	return sum
}

func Cos(x float64) float64 {
	k := int((x + Pi) / (2 * Pi))
	x -= float64(k) * 2 * Pi
	sum := 0.0
	term := 1.0
	x2 := x * x
	for i := 1; i < 50; i++ {
		sum += term
		term *= -x2 / float64((2*i-1)*(2*i))
		if Abs(term) < 1e-15 {
			break
		}
	}
	return sum
}

func Tan(x float64) float64 {
	c := Cos(x)
	if Abs(c) < 1e-15 {
		if Sin(x) > 0 {
			return 1e308
		}
		return -1e308
	}
	return Sin(x) / c
}

func SinTaylor(x float64, terms int) float64 {
	sum := 0.0
	term := x
	x2 := x * x
	for i := 1; i <= terms; i++ {
		sum += term
		term *= -x2 / float64(2*i*(2*i+1))
	}
	return sum
}

func CosTaylor(x float64, terms int) float64 {
	sum := 0.0
	term := 1.0
	x2 := x * x
	for i := 1; i <= terms; i++ {
		sum += term
		term *= -x2 / float64((2*i-1)*(2*i))
	}
	return sum
}

func Sec(x float64) float64 {
	return 1.0 / Cos(x)
}

func Csc(x float64) float64 {
	return 1.0 / Sin(x)
}

func Cot(x float64) float64 {
	return Cos(x) / Sin(x)
}
