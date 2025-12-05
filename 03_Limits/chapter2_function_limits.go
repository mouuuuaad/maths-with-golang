package limits

type Function func(float64) float64

func Limit(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a + h)
	for i := 0; i < 15; i++ {
		h /= 10
		curr := f(a + h)
		if absLim(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitLeft(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a - h)
	for i := 0; i < 15; i++ {
		h /= 10
		curr := f(a - h)
		if absLim(curr-prev) < 1e-9 {
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
		if absLim(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitBothSides(f Function, a float64) (float64, bool) {
	left, lOk := LimitLeft(f, a)
	right, rOk := LimitRight(f, a)
	if lOk && rOk && absLim(left-right) < 1e-9 {
		return left, true
	}
	return 0, false
}

func LimitExists(f Function, a float64) bool {
	_, ok := LimitBothSides(f, a)
	return ok
}

func LimitPosInfinity(f Function) (float64, bool) {
	prev := f(1000)
	for _, x := range []float64{10000, 100000, 1000000} {
		curr := f(x)
		if absLim(curr-prev) < 1e-7 {
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
		if absLim(curr-prev) < 1e-7 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitComposite(f, g Function, a float64) (float64, bool) {
	gLim, gOk := Limit(g, a)
	if !gOk {
		return 0, false
	}
	return Limit(f, gLim)
}
