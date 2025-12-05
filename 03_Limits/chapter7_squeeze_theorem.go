package limits

import "math"


func SqueezeTheorem(f, lower, upper Function, a float64) (float64, bool) {
	L1, ok1 := Limit(lower, a)
	L2, ok2 := Limit(upper, a)

	if !ok1 || !ok2 {
		return 0, false
	}

	if math.Abs(L1-L2) > 1e-5 {
		return 0, false
	}

	h := 1e-5
	x := a + h
	if lower(x) > f(x) || f(x) > upper(x) {
		return 0, false
	}
	x = a - h
	if lower(x) > f(x) || f(x) > upper(x) {
		return 0, false
	}

	return L1, true
}
