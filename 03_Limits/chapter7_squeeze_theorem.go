package limits

import "math"

// SqueezeTheorem verifies lim(x->a) f(x) = L using bounding functions lower(x) and upper(x)
// Returns true if lower <= f <= upper near a, and limits match.
func SqueezeTheorem(f, lower, upper Function, a float64) (float64, bool) {
	L1, ok1 := Limit(lower, a)
	L2, ok2 := Limit(upper, a)

	if !ok1 || !ok2 {
		return 0, false
	}

	if math.Abs(L1-L2) > 1e-5 {
		return 0, false // Limits do not match
	}

	// Check ordering near a
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
