package calculus

func IsContinuousAt(f Function, a, epsilon float64) bool {
	delta := epsilon
	for i := 0; i < 10; i++ {
		h := delta / 2
		if absL(f(a+h)-f(a)) < epsilon && absL(f(a-h)-f(a)) < epsilon {
			return true
		}
		delta /= 10
	}
	return false
}

func IsContinuousOnInterval(f Function, a, b float64, n int) bool {
	step := (b - a) / float64(n)
	epsilon := 1e-6
	for x := a; x <= b; x += step {
		if !IsContinuousAt(f, x, epsilon) {
			return false
		}
	}
	return true
}

func IntermediateValueTheorem(f Function, a, b, val float64) (float64, bool) {
	fa := f(a)
	fb := f(b)
	if (fa < val && fb < val) || (fa > val && fb > val) {
		return 0, false
	}
	low, high := a, b
	if fa > fb {
		low, high = b, a
	}
	for i := 0; i < 100; i++ {
		mid := (low + high) / 2
		fmid := f(mid)
		if absL(fmid-val) < 1e-9 {
			return mid, true
		}
		if fmid < val {
			if fa < fb {
				low = mid
			} else {
				high = mid
			}
		} else {
			if fa < fb {
				high = mid
			} else {
				low = mid
			}
		}
	}
	return (low + high) / 2, true
}

func FindDiscontinuities(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	epsilon := 1e-4
	for x := a; x <= b; x += step {
		if !IsContinuousAt(f, x, epsilon) {
			result = append(result, x)
		}
	}
	return result
}

func RemovableDiscontinuity(f Function, a float64) (float64, bool) {
	left, lOk := LimitLeft(f, a)
	right, rOk := LimitRight(f, a)
	if lOk && rOk && absL(left-right) < 1e-9 {
		return left, true
	}
	return 0, false
}

func UniformContinuity(f Function, a, b, epsilon float64) float64 {
	delta := (b - a) / 100
	for delta > 1e-10 {
		uniform := true
		for x := a; x < b; x += delta {
			for y := x; y < b && y-x < delta; y += delta / 10 {
				if absL(f(x)-f(y)) > epsilon {
					uniform = false
					break
				}
			}
			if !uniform {
				break
			}
		}
		if uniform {
			return delta
		}
		delta /= 2
	}
	return delta
}
