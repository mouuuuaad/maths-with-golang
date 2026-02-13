// 2026 Update: Differentiation
package calculus

import "math"

func Derivative(f Function, x float64) float64 {
	h := 1e-7
	return (f(x+h) - f(x-h)) / (2 * h)
}

func DerivativeForward(f Function, x float64) float64 {
	h := 1e-7
	return (f(x+h) - f(x)) / h
}

func DerivativeBackward(f Function, x float64) float64 {
	h := 1e-7
	return (f(x) - f(x-h)) / h
}

func SecondDerivative(f Function, x float64) float64 {
	h := 1e-5
	return (f(x+h) - 2*f(x) + f(x-h)) / (h * h)
}

func NthDerivative(f Function, x float64, n int) float64 {
	if n == 0 {
		return f(x)
	}
	if n == 1 {
		return Derivative(f, x)
	}
	df := func(t float64) float64 {
		return Derivative(f, t)
	}
	return NthDerivative(df, x, n-1)
}

func TangentLine(f Function, a float64) Function {
	fa := f(a)
	dfa := Derivative(f, a)
	return func(x float64) float64 {
		return fa + dfa*(x-a)
	}
}

func LinearApproximation(f Function, a, x float64) float64 {
	return f(a) + Derivative(f, a)*(x-a)
}

func QuadraticApproximation(f Function, a, x float64) float64 {
	h := x - a
	return f(a) + Derivative(f, a)*h + SecondDerivative(f, a)*h*h/2
}

func IsDifferentiableAt(f Function, a float64) bool {
	left := DerivativeBackward(f, a)
	right := DerivativeForward(f, a)
	return absL(left-right) < 1e-5
}

func CriticalPoints(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	for x := a; x <= b; x += step {
		d := Derivative(f, x)
		if absL(d) < 1e-6 {
			result = append(result, x)
		}
	}
	return result
}

func InflectionPoints(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	prev := SecondDerivative(f, a)
	for x := a + step; x <= b; x += step {
		curr := SecondDerivative(f, x)
		if prev*curr < 0 {
			result = append(result, x)
		}
		prev = curr
	}
	return result
}

func DerivativeRichardson(f Function, x float64, h float64) float64 {
	if h == 0 {
		h = 1e-3
	}
	d1 := (f(x+h) - f(x-h)) / (2 * h)
	d2 := (f(x+h/2) - f(x-h/2)) / (h)
	return d2 + (d2-d1)/3
}

func DerivativeFivePoint(f Function, x float64, h float64) float64 {
	if h == 0 {
		h = 1e-3
	}
	return (f(x-2*h) - 8*f(x-h) + 8*f(x+h) - f(x+2*h)) / (12 * h)
}

func SecondDerivativeFivePoint(f Function, x float64, h float64) float64 {
	if h == 0 {
		h = 1e-3
	}
	return (-f(x-2*h) + 16*f(x-h) - 30*f(x) + 16*f(x+h) - f(x+2*h)) / (12 * h * h)
}

func ComplexStepDerivative(f Function, x float64, h float64) float64 {
	if h == 0 {
		h = 1e-20
	}
	return (f(x+h) - f(x)) / h
}

func DerivativeAdaptive(f Function, x float64) float64 {
	h := 1e-2
	prev := DerivativeFivePoint(f, x, h)
	for i := 0; i < 10; i++ {
		h *= 0.5
		curr := DerivativeFivePoint(f, x, h)
		if absL(curr-prev) < 1e-10 {
			return curr
		}
		prev = curr
	}
	return prev
}

func DerivativeFromSamples(xs, ys []float64, x float64) float64 {
	if len(xs) < 2 || len(xs) != len(ys) {
		return 0
	}
	for i := 0; i < len(xs)-1; i++ {
		if x >= xs[i] && x <= xs[i+1] {
			dx := xs[i+1] - xs[i]
			if dx == 0 {
				return 0
			}
			return (ys[i+1] - ys[i]) / dx
		}
	}
	if x <= xs[0] {
		dx := xs[1] - xs[0]
		if dx == 0 {
			return 0
		}
		return (ys[1] - ys[0]) / dx
	}
	last := len(xs) - 1
	dx := xs[last] - xs[last-1]
	if dx == 0 {
		return 0
	}
	return (ys[last] - ys[last-1]) / dx
}

func DerivativeProduct(f, g Function, x float64) float64 {
	return Derivative(f, x)*g(x) + f(x)*Derivative(g, x)
}

func DerivativeQuotient(f, g Function, x float64) float64 {
	den := g(x)
	if den == 0 {
		return 0
	}
	return (Derivative(f, x)*g(x) - f(x)*Derivative(g, x)) / (den * den)
}

func DerivativeChain(f, g Function, x float64) float64 {
	return Derivative(f, g(x)) * Derivative(g, x)
}

func DirectionalDerivative(f MultiFunction, x []float64, v []float64) float64 {
	if len(x) != len(v) {
		return 0
	}
	h := 1e-6
	xp := make([]float64, len(x))
	xm := make([]float64, len(x))
	for i := range x {
		xp[i] = x[i] + h*v[i]
		xm[i] = x[i] - h*v[i]
	}
	return (f(xp) - f(xm)) / (2 * h)
}

func GradientApprox(f MultiFunction, x []float64, h float64) []float64 {
	if h == 0 {
		h = 1e-6
	}
	grad := make([]float64, len(x))
	for i := range x {
		xp := make([]float64, len(x))
		xm := make([]float64, len(x))
		copy(xp, x)
		copy(xm, x)
		xp[i] += h
		xm[i] -= h
		grad[i] = (f(xp) - f(xm)) / (2 * h)
	}
	return grad
}

func HessianDiagonalApprox(f MultiFunction, x []float64, h float64) []float64 {
	if h == 0 {
		h = 1e-4
	}
	diag := make([]float64, len(x))
	fx := f(x)
	for i := range x {
		xp := make([]float64, len(x))
		xm := make([]float64, len(x))
		copy(xp, x)
		copy(xm, x)
		xp[i] += h
		xm[i] -= h
		diag[i] = (f(xp) - 2*fx + f(xm)) / (h * h)
	}
	return diag
}

func NewtonStep1D(f, df Function, x float64) float64 {
	d := df(x)
	if d == 0 {
		return x
	}
	return x - f(x)/d
}

func SecantStep(f Function, x0, x1 float64) float64 {
	f0 := f(x0)
	f1 := f(x1)
	den := f1 - f0
	if den == 0 {
		return x1
	}
	return x1 - f1*(x1-x0)/den
}

func DerivativeStability(f Function, x float64, hs []float64) float64 {
	if len(hs) == 0 {
		return 0
	}
	values := make([]float64, len(hs))
	for i := range hs {
		values[i] = DerivativeFivePoint(f, x, hs[i])
	}
	mean := 0.0
	for _, v := range values {
		mean += v
	}
	mean /= float64(len(values))
	varsum := 0.0
	for _, v := range values {
		d := v - mean
		varsum += d * d
	}
	return math.Sqrt(varsum / float64(len(values)))
}

func HigherOrderCentral(f Function, x float64, h float64, order int) float64 {
	if order <= 1 {
		return DerivativeFivePoint(f, x, h)
	}
	if h == 0 {
		h = 1e-3
	}
	coeff := make([]float64, 2*order+1)
	for i := -order; i <= order; i++ {
		idx := i + order
		coeff[idx] = float64(i)
	}
	sum := 0.0
	for i := -order; i <= order; i++ {
		idx := i + order
		sum += coeff[idx] * f(x+float64(i)*h)
	}
	den := 0.0
	for i := -order; i <= order; i++ {
		den += coeff[i+order] * float64(i)
	}
	if den == 0 {
		return 0
	}
	return sum / (den * h)
}

func ForwardDifferenceTable(f Function, x float64, h float64, n int) []float64 {
	table := make([]float64, n)
	current := make([]float64, n)
	for i := 0; i < n; i++ {
		current[i] = f(x + float64(i)*h)
	}
	for k := 0; k < n; k++ {
		table[k] = current[0]
		for i := 0; i < n-1-k; i++ {
			current[i] = current[i+1] - current[i]
		}
	}
	return table
}

func BackwardDifferenceTable(f Function, x float64, h float64, n int) []float64 {
	table := make([]float64, n)
	current := make([]float64, n)
	for i := 0; i < n; i++ {
		current[i] = f(x - float64(i)*h)
	}
	for k := 0; k < n; k++ {
		table[k] = current[0]
		for i := 0; i < n-1-k; i++ {
			current[i] = current[i+1] - current[i]
		}
	}
	return table
}

func SymmetricDifferenceTable(f Function, x float64, h float64, n int) []float64 {
	table := make([]float64, n)
	current := make([]float64, n)
	for i := 0; i < n; i++ {
		off := float64(i) * h
		current[i] = f(x+off) - f(x-off)
	}
	for k := 0; k < n; k++ {
		table[k] = current[0]
		for i := 0; i < n-1-k; i++ {
			current[i] = current[i+1] - current[i]
		}
	}
	return table
}

func DerivativeErrorEstimate(f Function, x float64, h float64) float64 {
	if h == 0 {
		h = 1e-3
	}
	d1 := DerivativeFivePoint(f, x, h)
	d2 := DerivativeFivePoint(f, x, h/2)
	return absL(d2 - d1)
}

func MeanValueTheoremPoint(f Function, a, b float64) float64 {
	slope := (f(b) - f(a)) / (b - a)
	g := func(x float64) float64 {
		return Derivative(f, x) - slope
	}
	root, _ := Limit(g, (a+b)/2)
	return root
}

func RollePoint(f Function, a, b float64) float64 {
	if absL(f(a)-f(b)) > 1e-6 {
		return math.NaN()
	}
	candidates := CriticalPoints(f, a, b, 200)
	if len(candidates) == 0 {
		return math.NaN()
	}
	return candidates[len(candidates)/2]
}

func TotalVariationApprox(f Function, a, b float64, n int) float64 {
	if n <= 0 {
		n = 200
	}
	dx := (b - a) / float64(n)
	varsum := 0.0
	prev := f(a)
	for i := 1; i <= n; i++ {
		x := a + float64(i)*dx
		cur := f(x)
		varsum += math.Abs(cur - prev)
		prev = cur
	}
	return varsum
}

func DerivativeSignChanges(f Function, a, b float64, n int) int {
	if n <= 0 {
		n = 200
	}
	dx := (b - a) / float64(n)
	prev := Derivative(f, a)
	count := 0
	for i := 1; i <= n; i++ {
		x := a + float64(i)*dx
		cur := Derivative(f, x)
		if prev*cur < 0 {
			count++
		}
		prev = cur
	}
	return count
}

func DerivativeEnvelope(f Function, a, b float64, n int) (float64, float64) {
	if n <= 0 {
		n = 200
	}
	dx := (b - a) / float64(n)
	minv := math.Inf(1)
	maxv := math.Inf(-1)
	for i := 0; i <= n; i++ {
		x := a + float64(i)*dx
		v := Derivative(f, x)
		if v < minv {
			minv = v
		}
		if v > maxv {
			maxv = v
		}
	}
	return minv, maxv
}
