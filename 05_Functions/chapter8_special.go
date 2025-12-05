package functions

func Gamma(x float64) float64 {
	p := []float64{
		76.18009172947146,
		-86.50532032941677,
		24.01409824083091,
		-1.231739572450155,
		0.1208650973866179e-2,
		-0.5395239384953e-5,
	}
	if x < 0.5 {
		return Pi / (Sin(Pi*x) * Gamma(1-x))
	}
	x -= 1
	tmp := x + 5.5
	tmp = (x+0.5)*Ln(tmp) - tmp
	ser := 1.000000000190015
	for i := 0; i < len(p); i++ {
		ser += p[i] / (x + float64(i) + 1)
	}
	return Exp(tmp + Ln(2.5066282746310005*ser))
}

func Beta(x, y float64) float64 {
	return Gamma(x) * Gamma(y) / Gamma(x+y)
}

func Factorial(n int) float64 {
	if n <= 1 {
		return 1
	}
	res := 1.0
	for i := 2; i <= n; i++ {
		res *= float64(i)
	}
	return res
}

func BesselJ0(x float64) float64 {
	sum := 0.0
	term := 1.0
	x2_4 := x * x / 4
	for k := 0; k < 30; k++ {
		sum += term
		term *= -x2_4 / float64((k+1)*(k+1))
	}
	return sum
}

func BesselJ1(x float64) float64 {
	sum := 0.0
	term := x / 2
	x2_4 := x * x / 4
	for k := 0; k < 30; k++ {
		sum += term
		term *= -x2_4 / float64((k+1)*(k+2))
	}
	return sum
}

func BesselJn(n int, x float64) float64 {
	if n == 0 {
		return BesselJ0(x)
	}
	if n == 1 {
		return BesselJ1(x)
	}
	if n < 0 {
		if n%2 == 0 {
			return BesselJn(-n, x)
		}
		return -BesselJn(-n, x)
	}
	j0 := BesselJ0(x)
	j1 := BesselJ1(x)
	for k := 1; k < n; k++ {
		j0, j1 = j1, float64(2*k)/x*j1-j0
	}
	return j1
}

func Erf(x float64) float64 {
	a1 := 0.254829592
	a2 := -0.284496736
	a3 := 1.421413741
	a4 := -1.453152027
	a5 := 1.061405429
	p := 0.3275911
	sign := 1.0
	if x < 0 {
		sign = -1
		x = -x
	}
	t := 1.0 / (1.0 + p*x)
	y := 1.0 - (((((a5*t+a4)*t)+a3)*t+a2)*t+a1)*t*Exp(-x*x)
	return sign * y
}

func Erfc(x float64) float64 {
	return 1 - Erf(x)
}

func Zeta(s float64, terms int) float64 {
	if s <= 1 {
		return 0
	}
	sum := 0.0
	for n := 1; n <= terms; n++ {
		sum += 1.0 / Power(float64(n), s)
	}
	return sum
}

func DiGamma(x float64) float64 {
	result := 0.0
	for x < 6 {
		result -= 1.0 / x
		x += 1
	}
	result += Ln(x) - 1.0/(2*x) - 1.0/(12*x*x) + 1.0/(120*x*x*x*x)
	return result
}
