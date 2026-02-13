// 2026 Update: Continuity
package calculus

import "math"

type ContinuityReport struct {
	Points    []float64
	IsCont    []bool
	MaxJump   float64
	MaxOsc    float64
	AllCont   bool
	CheckedAt []float64
}

func IsContinuousAt(f Function, a, epsilon float64) bool {
	delta := epsilon
	for i := 0; i < 12; i++ {
		h := delta / 2
		if absL(f(a+h)-f(a)) < epsilon && absL(f(a-h)-f(a)) < epsilon {
			return true
		}
		delta *= 0.5
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
	delta := epsilon
	for k := 0; k < 20; k++ {
		ok := true
		samples := 50
		step := (b - a) / float64(samples)
		for i := 0; i < samples; i++ {
			x := a + float64(i)*step
			for j := i + 1; j < samples; j++ {
				y := a + float64(j)*step
				if absL(x-y) < delta && absL(f(x)-f(y)) >= epsilon {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			return delta
		}
		delta *= 0.5
	}
	return delta
}

func LipschitzConstantApprox(f Function, a, b float64, n int) float64 {
	if n <= 1 {
		n = 100
	}
	maxL := 0.0
	step := (b - a) / float64(n)
	for i := 0; i < n; i++ {
		x := a + float64(i)*step
		y := x + step
		dx := y - x
		if dx == 0 {
			continue
		}
		L := absL(f(y)-f(x)) / absL(dx)
		if L > maxL {
			maxL = L
		}
	}
	return maxL
}

func ModulusOfContinuity(f Function, a, b float64, delta float64, n int) float64 {
	if n <= 1 {
		n = 100
	}
	maxOsc := 0.0
	step := (b - a) / float64(n)
	for i := 0; i < n; i++ {
		x := a + float64(i)*step
		y := x + delta
		if y > b {
			break
		}
		osc := absL(f(y) - f(x))
		if osc > maxOsc {
			maxOsc = osc
		}
	}
	return maxOsc
}

func OscillationAtPoint(f Function, a float64, steps int) float64 {
	left, right := LimitTableBothSides(f, a, steps)
	minv := math.Inf(1)
	maxv := math.Inf(-1)
	for _, v := range left {
		if v < minv {
			minv = v
		}
		if v > maxv {
			maxv = v
		}
	}
	for _, v := range right {
		if v < minv {
			minv = v
		}
		if v > maxv {
			maxv = v
		}
	}
	return maxv - minv
}

func JumpDiscontinuitySize(f Function, a float64) float64 {
	l, _ := LimitLeft(f, a)
	r, _ := LimitRight(f, a)
	return r - l
}

func EssentialDiscontinuity(f Function, a float64, steps int) bool {
	osc := OscillationAtPoint(f, a, steps)
	return osc > 1e-2
}

func ContinuityReportInterval(f Function, a, b float64, n int) ContinuityReport {
	points := make([]float64, n+1)
	isCont := make([]bool, n+1)
	step := (b - a) / float64(n)
	all := true
	maxJump := 0.0
	maxOsc := 0.0
	for i := 0; i <= n; i++ {
		x := a + float64(i)*step
		points[i] = x
		c := IsContinuousAt(f, x, 1e-6)
		isCont[i] = c
		if !c {
			all = false
		}
		jump := absL(JumpDiscontinuitySize(f, x))
		if jump > maxJump {
			maxJump = jump
		}
		osc := OscillationAtPoint(f, x, 8)
		if osc > maxOsc {
			maxOsc = osc
		}
	}
	return ContinuityReport{Points: points, IsCont: isCont, MaxJump: maxJump, MaxOsc: maxOsc, AllCont: all, CheckedAt: points}
}

func ContinuousExtension(f Function, a float64) (Function, bool) {
	lim, ok := RemovableDiscontinuity(f, a)
	if !ok {
		return nil, false
	}
	return func(x float64) float64 {
		if absL(x-a) < 1e-12 {
			return lim
		}
		return f(x)
	}, true
}

func UniformContinuityCheck(f Function, a, b float64, eps float64, n int) bool {
	delta := UniformContinuity(f, a, b, eps)
	if delta <= 0 {
		return false
	}
	step := (b - a) / float64(n)
	for i := 0; i < n; i++ {
		x := a + float64(i)*step
		y := x + 0.5*delta
		if y > b {
			break
		}
		if absL(f(x)-f(y)) >= eps {
			return false
		}
	}
	return true
}

func HeineCantorBound(f Function, a, b float64) float64 {
	L := LipschitzConstantApprox(f, a, b, 200)
	if L == 0 {
		return 0
	}
	return L
}

func ContinuityModulusSequence(f Function, a, b float64, deltas []float64) []float64 {
	mods := make([]float64, len(deltas))
	for i := range deltas {
		mods[i] = ModulusOfContinuity(f, a, b, deltas[i], 200)
	}
	return mods
}

func OscillationInterval(f Function, a, b float64, n int) float64 {
	minv := math.Inf(1)
	maxv := math.Inf(-1)
	step := (b - a) / float64(n)
	for i := 0; i <= n; i++ {
		x := a + float64(i)*step
		v := f(x)
		if v < minv {
			minv = v
		}
		if v > maxv {
			maxv = v
		}
	}
	return maxv - minv
}

func ContinuityAtSamples(f Function, xs []float64, eps float64) []bool {
	res := make([]bool, len(xs))
	for i := range xs {
		res[i] = IsContinuousAt(f, xs[i], eps)
	}
	return res
}

func ContinuityStrength(f Function, a, b float64, n int) float64 {
	step := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i <= n; i++ {
		x := a + float64(i)*step
		l, okL := LimitLeft(f, x)
		r, okR := LimitRight(f, x)
		if okL && okR {
			sum += 1.0 / (1.0 + absL(l-r))
		}
	}
	return sum / float64(n+1)
}

func ContinuityGap(f Function, a, b float64, n int) float64 {
	step := (b - a) / float64(n)
	maxGap := 0.0
	for i := 0; i <= n; i++ {
		x := a + float64(i)*step
		jump := absL(JumpDiscontinuitySize(f, x))
		if jump > maxGap {
			maxGap = jump
		}
	}
	return maxGap
}

func ContinuityNear(f Function, a float64, eps float64) bool {
	left, okL := LimitLeft(f, a)
	right, okR := LimitRight(f, a)
	if !okL || !okR {
		return false
	}
	return absL(left-right) < eps
}

func UniformContinuityEstimate(f Function, a, b float64) float64 {
	L := LipschitzConstantApprox(f, a, b, 200)
	if L == 0 {
		return math.Inf(1)
	}
	return 1 / L
}

func ContinuityIndicator(f Function, x float64) float64 {
	l, okL := LimitLeft(f, x)
	r, okR := LimitRight(f, x)
	if okL && okR {
		return 1.0 / (1.0 + absL(l-r))
	}
	return 0
}

func ContinuityIndicatorSamples(f Function, xs []float64) []float64 {
	vals := make([]float64, len(xs))
	for i := range xs {
		vals[i] = ContinuityIndicator(f, xs[i])
	}
	return vals
}

func ContinuityBand(f Function, a, b float64, n int) (float64, float64) {
	minv := math.Inf(1)
	maxv := math.Inf(-1)
	step := (b - a) / float64(n)
	for i := 0; i <= n; i++ {
		x := a + float64(i)*step
		v := f(x)
		if v < minv {
			minv = v
		}
		if v > maxv {
			maxv = v
		}
	}
	return minv, maxv
}

func ContinuityWithinBand(f Function, a, b float64, n int, band float64) bool {
	minv, maxv := ContinuityBand(f, a, b, n)
	return (maxv - minv) <= band
}
