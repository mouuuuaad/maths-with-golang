package calculus

func absL(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Limit(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a + h)
	for i := 0; i < 15; i++ {
		h /= 10
		curr := f(a + h)
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitSum(f, g Function, a float64) (float64, bool) {
	lf, okf := Limit(f, a)
	lg, okg := Limit(g, a)
	if okf && okg {
		return lf + lg, true
	}
	return 0, false
}

func LimitProduct(f, g Function, a float64) (float64, bool) {
	lf, okf := Limit(f, a)
	lg, okg := Limit(g, a)
	if okf && okg {
		return lf * lg, true
	}
	return 0, false
}

func LimitQuotient(f, g Function, a float64) (float64, bool) {
	lf, okf := Limit(f, a)
	lg, okg := Limit(g, a)
	if okf && okg && absL(lg) > 1e-9 {
		return lf / lg, true
	}
	return 0, false
}

func LimitLeft(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a - h)
	for i := 0; i < 15; i++ {
		h /= 10
		curr := f(a - h)
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitRight(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a + h)
	for i := 0; i < 15; i++ {
		h /= 10
		curr := f(a + h)
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitInfinity(f Function) (float64, bool) {
	prev := f(1000)
	for _, x := range []float64{10000, 100000, 1000000} {
		curr := f(x)
		if absL(curr-prev) < 1e-7 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitNegInfinity(f Function) (float64, bool) {
	prev := f(-1000)
	for _, x := range []float64{-10000, -100000, -1000000} {
		curr := f(x)
		if absL(curr-prev) < 1e-7 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}
