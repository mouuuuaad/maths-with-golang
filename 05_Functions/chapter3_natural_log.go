package functions

func Ln(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	if x > 2 {
		y = x / 2.0
	}
	for i := 0; i < 100; i++ {
		ey := Exp(y)
		diff := ey - x
		if Abs(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
	}
	return y
}

func LnMercator(x float64, terms int) float64 {
	if x <= -1 || x > 1 {
		return 0
	}
	sum := 0.0
	term := x
	sign := 1.0
	for i := 1; i <= terms; i++ {
		sum += sign * term / float64(i)
		term *= x
		sign = -sign
	}
	return sum
}

func LnSeries(x float64, terms int) float64 {
	if x <= 0 {
		return -1e308
	}
	y := (x - 1) / (x + 1)
	y2 := y * y
	sum := 0.0
	power := y
	for i := 0; i < terms; i++ {
		sum += power / float64(2*i+1)
		power *= y2
	}
	return 2 * sum
}
