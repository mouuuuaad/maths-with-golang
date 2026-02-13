// 2026 Update: Limits Fundamentals
package calculus

import "math"

func absL(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Limit(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a + h)
	for i := 0; i < 20; i++ {
		h *= 0.1
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
	if okf && okg && absL(lg) > 1e-12 {
		return lf / lg, true
	}
	return 0, false
}

func LimitLeft(f Function, a float64) (float64, bool) {
	h := 0.1
	prev := f(a - h)
	for i := 0; i < 20; i++ {
		h *= 0.1
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
	for i := 0; i < 20; i++ {
		h *= 0.1
		curr := f(a + h)
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func TwoSidedLimit(f Function, a float64) (float64, bool) {
	l, okL := LimitLeft(f, a)
	r, okR := LimitRight(f, a)
	if okL && okR && absL(l-r) < 1e-8 {
		return (l + r) * 0.5, true
	}
	return 0, false
}

func LimitInfinity(f Function) (float64, bool) {
	prev := f(100)
	x := 200.0
	for i := 0; i < 20; i++ {
		curr := f(x)
		if absL(curr-prev) < 1e-8 {
			return curr, true
		}
		prev = curr
		x *= 2
	}
	return prev, false
}

func LimitNegInfinity(f Function) (float64, bool) {
	prev := f(-100)
	x := -200.0
	for i := 0; i < 20; i++ {
		curr := f(x)
		if absL(curr-prev) < 1e-8 {
			return curr, true
		}
		prev = curr
		x *= 2
	}
	return prev, false
}

func LimitAtInfinity(f Function, start float64, grow float64) (float64, bool) {
	if grow <= 1 {
		grow = 2
	}
	prev := f(start)
	x := start * grow
	for i := 0; i < 25; i++ {
		curr := f(x)
		if absL(curr-prev) < 1e-8 {
			return curr, true
		}
		prev = curr
		x *= grow
	}
	return prev, false
}

func LimitSequence(seq Sequence) (float64, bool) {
	prev := seq(1)
	for i := 2; i < 40; i++ {
		curr := seq(i)
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitBySequence(f Function, a float64, seq Sequence) (float64, bool) {
	prev := f(a + seq(1))
	for i := 2; i < 30; i++ {
		curr := f(a + seq(i))
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitEpsilonDelta(f Function, a, eps float64) float64 {
	delta := 1.0
	for i := 0; i < 30; i++ {
		ok := true
		for k := 1; k <= 10; k++ {
			x := a + delta*float64(k)/10.0
			if absL(f(x)-f(a)) >= eps {
				ok = false
				break
			}
			x = a - delta*float64(k)/10.0
			if absL(f(x)-f(a)) >= eps {
				ok = false
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

func LimitTable(f Function, a float64, steps int) []float64 {
	vals := make([]float64, steps)
	h := 0.1
	for i := 0; i < steps; i++ {
		vals[i] = f(a + h)
		h *= 0.5
	}
	return vals
}

func LimitTableBothSides(f Function, a float64, steps int) ([]float64, []float64) {
	left := make([]float64, steps)
	right := make([]float64, steps)
	h := 0.1
	for i := 0; i < steps; i++ {
		left[i] = f(a - h)
		right[i] = f(a + h)
		h *= 0.5
	}
	return left, right
}

func LimitAverage(f Function, a float64, steps int) float64 {
	left, right := LimitTableBothSides(f, a, steps)
	sum := 0.0
	for i := 0; i < steps; i++ {
		sum += 0.5 * (left[i] + right[i])
	}
	return sum / float64(steps)
}

func LimitRichardson(f Function, a float64, h float64, order int) float64 {
	if h == 0 {
		h = 1e-2
	}
	approx := make([]float64, order)
	for i := 0; i < order; i++ {
		ht := h / math.Pow(2, float64(i))
		approx[i] = f(a + ht)
	}
	for k := 1; k < order; k++ {
		for i := 0; i < order-k; i++ {
			approx[i] = approx[i+1] + (approx[i+1]-approx[i])/(math.Pow(2, float64(k))-1)
		}
	}
	return approx[0]
}

func LimitCauchy(f Function, a float64, steps int) bool {
	vals := LimitTable(f, a, steps)
	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			if absL(vals[i]-vals[j]) > 1e-6 {
				return false
			}
		}
	}
	return true
}

func LimitOscillation(f Function, a float64, steps int) float64 {
	vals := LimitTable(f, a, steps)
	minv := math.Inf(1)
	maxv := math.Inf(-1)
	for _, v := range vals {
		if v < minv {
			minv = v
		}
		if v > maxv {
			maxv = v
		}
	}
	return maxv - minv
}

func LimitExists(f Function, a float64) bool {
	l, okL := LimitLeft(f, a)
	r, okR := LimitRight(f, a)
	if okL && okR && absL(l-r) < 1e-8 {
		return true
	}
	return false
}

func LimitOfDifferenceQuotient(f Function, a float64) float64 {
	h := 1e-5
	return (f(a+h) - f(a)) / h
}

func LimitComposite(f Function, g Function, a float64) (float64, bool) {
	lg, okg := Limit(g, a)
	if !okg {
		return 0, false
	}
	return Limit(f, lg)
}

func LimitComparison(f, g Function, a float64) float64 {
	lf, okf := Limit(f, a)
	lg, okg := Limit(g, a)
	if okf && okg && lg != 0 {
		return lf / lg
	}
	return 0
}

func LimitSqueeze(f, g, h Function, a float64) (float64, bool) {
	lg, okg := Limit(g, a)
	lh, okh := Limit(h, a)
	if okg && okh && absL(lg-lh) < 1e-8 {
		return lg, true
	}
	return 0, false
}

func LimitPower(f Function, a float64, p float64) (float64, bool) {
	lf, ok := Limit(f, a)
	if !ok {
		return 0, false
	}
	return math.Pow(lf, p), true
}

func LimitRoot(f Function, a float64, p float64) (float64, bool) {
	lf, ok := Limit(f, a)
	if !ok {
		return 0, false
	}
	if lf < 0 && p != math.Trunc(p) {
		return 0, false
	}
	return math.Pow(lf, 1.0/p), true
}

func LimitRational(num, den Function, a float64) (float64, bool) {
	ln, okn := Limit(num, a)
	ld, okd := Limit(den, a)
	if okn && okd && absL(ld) > 1e-12 {
		return ln / ld, true
	}
	return 0, false
}

func LimitExponential(f Function, a float64) (float64, bool) {
	lf, ok := Limit(f, a)
	if !ok {
		return 0, false
	}
	return math.Exp(lf), true
}

func LimitLogarithm(f Function, a float64) (float64, bool) {
	lf, ok := Limit(f, a)
	if !ok || lf <= 0 {
		return 0, false
	}
	return math.Log(lf), true
}

func LimitAtPointGrid(f Function, a float64, offsets []float64) []float64 {
	vals := make([]float64, len(offsets))
	for i := range offsets {
		vals[i] = f(a + offsets[i])
	}
	return vals
}

func LimitDetectJump(f Function, a float64) float64 {
	l, _ := LimitLeft(f, a)
	r, _ := LimitRight(f, a)
	return r - l
}

func LimitFromSequence(f Function, a float64, seq []float64) float64 {
	if len(seq) == 0 {
		return 0
	}
	vals := make([]float64, len(seq))
	for i := range seq {
		vals[i] = f(a + seq[i])
	}
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

func LimitWeightedAverage(f Function, a float64, weights []float64) float64 {
	if len(weights) == 0 {
		return 0
	}
	h := 0.1
	sum := 0.0
	wsum := 0.0
	for _, w := range weights {
		v := f(a + h)
		sum += w * v
		wsum += w
		h *= 0.5
	}
	if wsum == 0 {
		return 0
	}
	return sum / wsum
}

func LimitDirectional(f Function, a float64, dir float64) (float64, bool) {
	h := 0.1
	prev := f(a + dir*h)
	for i := 0; i < 20; i++ {
		h *= 0.1
		curr := f(a + dir*h)
		if absL(curr-prev) < 1e-9 {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func LimitFromAbove(f Function, a float64) (float64, bool) {
	return LimitDirectional(f, a, 1)
}

func LimitFromBelow(f Function, a float64) (float64, bool) {
	return LimitDirectional(f, a, -1)
}
