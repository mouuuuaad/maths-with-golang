package limits

import "math"

type IndeterminateType int

const (
	None IndeterminateType = iota
	ZeroOverZero
	InfOverInf
	ZeroTimesInf
	InfMinusInf
)

func CheckIndeterminate(f, g Function, a float64) IndeterminateType {
	fa := f(a)
	ga := g(a)

	if math.Abs(fa) < 1e-9 && math.Abs(ga) < 1e-9 {
		return ZeroOverZero
	}
	if math.IsInf(fa, 0) && math.IsInf(ga, 0) {
		return InfOverInf
	}
	return None
}
