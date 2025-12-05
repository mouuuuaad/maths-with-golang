package limits

func PInfinity(f Function, a float64) bool {
	left, _ := LimitLeft(f, a)
	right, _ := LimitRight(f, a)
	return left > 1e10 || right > 1e10
}

func NInfinity(f Function, a float64) bool {
	left, _ := LimitLeft(f, a)
	right, _ := LimitRight(f, a)
	return left < -1e10 || right < -1e10
}

func DetectAsymptote(f Function, a float64) string {
	left, lOk := LimitLeft(f, a)
	right, rOk := LimitRight(f, a)
	if !lOk && !rOk {
		return "undefined"
	}
	if absLim(left) > 1e10 || absLim(right) > 1e10 {
		return "vertical"
	}
	return "none"
}

func HorizontalAsymptote(f Function) (float64, bool) {
	return LimitPosInfinity(f)
}

func ObliqueAsymptote(f Function) (float64, float64, bool) {
	m := func(x float64) float64 {
		return f(x) / x
	}
	slope, mOk := LimitPosInfinity(m)
	if !mOk {
		return 0, 0, false
	}
	c := func(x float64) float64 {
		return f(x) - slope*x
	}
	intercept, cOk := LimitPosInfinity(c)
	if !cOk {
		return 0, 0, false
	}
	return slope, intercept, true
}

func FindVerticalAsymptotes(f Function, start, end float64, n int) []float64 {
	result := []float64{}
	step := (end - start) / float64(n)
	for x := start; x <= end; x += step {
		val := f(x)
		if val != val || absLim(val) > 1e10 {
			result = append(result, x)
		}
	}
	return result
}
