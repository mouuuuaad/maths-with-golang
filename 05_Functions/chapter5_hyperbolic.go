package functions

func Sinh(x float64) float64 {
	ex := Exp(x)
	emx := Exp(-x)
	return (ex - emx) / 2
}

func Cosh(x float64) float64 {
	ex := Exp(x)
	emx := Exp(-x)
	return (ex + emx) / 2
}

func Tanh(x float64) float64 {
	ex := Exp(x)
	emx := Exp(-x)
	return (ex - emx) / (ex + emx)
}

func SinhDef(x float64) float64 {
	return Sinh(x)
}

func CoshDef(x float64) float64 {
	return Cosh(x)
}

func Sech(x float64) float64 {
	return 1.0 / Cosh(x)
}

func Csch(x float64) float64 {
	return 1.0 / Sinh(x)
}

func Coth(x float64) float64 {
	return Cosh(x) / Sinh(x)
}

func SinhTaylor(x float64, terms int) float64 {
	sum := 0.0
	term := x
	x2 := x * x
	for i := 1; i <= terms; i++ {
		sum += term
		term *= x2 / float64(2*i*(2*i+1))
	}
	return sum
}

func CoshTaylor(x float64, terms int) float64 {
	sum := 0.0
	term := 1.0
	x2 := x * x
	for i := 1; i <= terms; i++ {
		sum += term
		term *= x2 / float64((2*i-1)*(2*i))
	}
	return sum
}
