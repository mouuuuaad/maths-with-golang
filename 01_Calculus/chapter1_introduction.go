// 2026 Update: Calculus Introduction
package calculus

import "math"

type Function func(float64) float64

type Condition func(float64) bool

type Piece struct {
	Cond Condition
	Func Function
}

type RNG struct {
	state uint64
}

func NewRNG(seed uint64) *RNG {
	if seed == 0 {
		seed = 1
	}
	return &RNG{state: seed}
}

func (r *RNG) Next() uint64 {
	r.state = r.state*6364136223846793005 + 1442695040888963407
	return r.state
}

func (r *RNG) Float64() float64 {
	return float64(r.Next()>>11) / float64(1<<53)
}

func (r *RNG) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.Next() % uint64(n))
}

func DomainCheck(f Function, x float64) bool {
	val := f(x)
	return val == val && val < 1e308 && val > -1e308
}

func DomainCheckFinite(f Function, x float64) bool {
	val := f(x)
	return !math.IsNaN(val) && !math.IsInf(val, 0)
}

func DomainCheckInterval(f Function, a, b float64, samples int) bool {
	if samples <= 0 {
		samples = 100
	}
	step := (b - a) / float64(samples)
	for i := 0; i <= samples; i++ {
		x := a + float64(i)*step
		if !DomainCheckFinite(f, x) {
			return false
		}
	}
	return true
}

func FindDomainApprox(f Function, start, end, step float64) (float64, float64) {
	minX := 1e308
	maxX := -1e308
	found := false
	for x := start; x <= end; x += step {
		if DomainCheck(f, x) {
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			found = true
		}
	}
	if !found {
		return 0, 0
	}
	return minX, maxX
}

func Evaluate(f Function, x float64) float64 {
	return f(x)
}

func Compose(f, g Function) Function {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func ComposeMany(funcs ...Function) Function {
	return func(x float64) float64 {
		v := x
		for i := len(funcs) - 1; i >= 0; i-- {
			v = funcs[i](v)
		}
		return v
	}
}

func Add(f, g Function) Function {
	return func(x float64) float64 {
		return f(x) + g(x)
	}
}

func Subtract(f, g Function) Function {
	return func(x float64) float64 {
		return f(x) - g(x)
	}
}

func Multiply(f, g Function) Function {
	return func(x float64) float64 {
		return f(x) * g(x)
	}
}

func Divide(f, g Function) Function {
	return func(x float64) float64 {
		den := g(x)
		if den == 0 {
			return 0
		}
		return f(x) / den
	}
}

func Scale(f Function, c float64) Function {
	return func(x float64) float64 {
		return c * f(x)
	}
}

func ScaleX(f Function, s float64) Function {
	if s == 0 {
		s = 1
	}
	return func(x float64) float64 {
		return f(x / s)
	}
}

func Translate(f Function, h float64) Function {
	return func(x float64) float64 {
		return f(x - h)
	}
}

func TranslateY(f Function, k float64) Function {
	return func(x float64) float64 {
		return f(x) + k
	}
}

func ReflectX(f Function) Function {
	return func(x float64) float64 {
		return f(-x)
	}
}

func ReflectY(f Function) Function {
	return func(x float64) float64 {
		return -f(x)
	}
}

func Power(f Function, p float64) Function {
	return func(x float64) float64 {
		return math.Pow(f(x), p)
	}
}

func PowerN(f Function, n int) Function {
	return func(x float64) float64 {
		v := f(x)
		if n == 0 {
			return 1
		}
		res := 1.0
		absn := n
		if absn < 0 {
			absn = -absn
		}
		for i := 0; i < absn; i++ {
			res *= v
		}
		if n < 0 {
			if res == 0 {
				return 0
			}
			return 1 / res
		}
		return res
	}
}

func Reciprocal(f Function) Function {
	return func(x float64) float64 {
		v := f(x)
		if v == 0 {
			return 0
		}
		return 1 / v
	}
}

func ConstantFunction(c float64) Function {
	return func(float64) float64 {
		return c
	}
}

func LinearFunction(a, b float64) Function {
	return func(x float64) float64 {
		return a*x + b
	}
}

func QuadraticFunction(a, b, c float64) Function {
	return func(x float64) float64 {
		return a*x*x + b*x + c
	}
}

func PolynomialFunction(coeffs []float64) Function {
	return func(x float64) float64 {
		acc := 0.0
		for i := len(coeffs) - 1; i >= 0; i-- {
			acc = acc*x + coeffs[i]
		}
		return acc
	}
}

func PiecewiseFunction(pieces []Piece, fallback Function) Function {
	return func(x float64) float64 {
		for _, p := range pieces {
			if p.Cond != nil && p.Cond(x) {
				return p.Func(x)
			}
		}
		if fallback != nil {
			return fallback(x)
		}
		return 0
	}
}

func ClampOutput(f Function, low, high float64) Function {
	return func(x float64) float64 {
		v := f(x)
		if v < low {
			return low
		}
		if v > high {
			return high
		}
		return v
	}
}

func MapFunction(f Function, xs []float64) []float64 {
	ys := make([]float64, len(xs))
	for i := range xs {
		ys[i] = f(xs[i])
	}
	return ys
}

func SampleFunction(f Function, start, end float64, n int) ([]float64, []float64) {
	if n <= 1 {
		return []float64{start}, []float64{f(start)}
	}
	xs := make([]float64, n)
	ys := make([]float64, n)
	step := (end - start) / float64(n-1)
	for i := 0; i < n; i++ {
		x := start + float64(i)*step
		xs[i] = x
		ys[i] = f(x)
	}
	return xs, ys
}

func RangeApprox(f Function, start, end float64, n int) (float64, float64) {
	if n <= 0 {
		n = 100
	}
	minv := math.Inf(1)
	maxv := math.Inf(-1)
	step := (end - start) / float64(n)
	for i := 0; i <= n; i++ {
		x := start + float64(i)*step
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

func ArgMinApprox(f Function, start, end float64, n int) float64 {
	if n <= 0 {
		n = 100
	}
	step := (end - start) / float64(n)
	bestX := start
	bestV := f(start)
	for i := 1; i <= n; i++ {
		x := start + float64(i)*step
		v := f(x)
		if v < bestV {
			bestV = v
			bestX = x
		}
	}
	return bestX
}

func ArgMaxApprox(f Function, start, end float64, n int) float64 {
	if n <= 0 {
		n = 100
	}
	step := (end - start) / float64(n)
	bestX := start
	bestV := f(start)
	for i := 1; i <= n; i++ {
		x := start + float64(i)*step
		v := f(x)
		if v > bestV {
			bestV = v
			bestX = x
		}
	}
	return bestX
}

func ApproxIntegralTrapezoid(f Function, a, b float64, n int) float64 {
	if n <= 0 {
		n = 100
	}
	dx := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		x := a + float64(i)*dx
		sum += f(x)
	}
	return sum * dx
}

func ApproxDerivativeCentral(f Function, x, h float64) float64 {
	if h == 0 {
		h = 1e-6
	}
	return (f(x+h) - f(x-h)) / (2 * h)
}

func FunctionEqualApprox(f, g Function, start, end float64, n int, tol float64) bool {
	step := (end - start) / float64(n)
	for i := 0; i <= n; i++ {
		x := start + float64(i)*step
		if math.Abs(f(x)-g(x)) > tol {
			return false
		}
	}
	return true
}

func IsEvenFunction(f Function, start, end float64, n int, tol float64) bool {
	step := (end - start) / float64(n)
	for i := 0; i <= n; i++ {
		x := start + float64(i)*step
		if math.Abs(f(x)-f(-x)) > tol {
			return false
		}
	}
	return true
}

func IsOddFunction(f Function, start, end float64, n int, tol float64) bool {
	step := (end - start) / float64(n)
	for i := 0; i <= n; i++ {
		x := start + float64(i)*step
		if math.Abs(f(x)+f(-x)) > tol {
			return false
		}
	}
	return true
}

func NormalizeFunction(f Function, a, b, outMin, outMax float64, samples int) Function {
	minv, maxv := RangeApprox(f, a, b, samples)
	if maxv == minv {
		return func(x float64) float64 { return outMin }
	}
	return func(x float64) float64 {
		v := f(x)
		scaled := (v - minv) / (maxv - minv)
		return outMin + scaled*(outMax-outMin)
	}
}

func BlendFunctions(f, g Function, t float64) Function {
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	return func(x float64) float64 {
		return (1-t)*f(x) + t*g(x)
	}
}

func SmoothStep(x float64) float64 {
	if x <= 0 {
		return 0
	}
	if x >= 1 {
		return 1
	}
	return x * x * (3 - 2*x)
}

func Sigmoid(x float64) float64 {
	if x >= 0 {
		z := math.Exp(-x)
		return 1 / (1 + z)
	}
	z := math.Exp(x)
	return z / (1 + z)
}

func FunctionFromSamples(xs, ys []float64) Function {
	return func(x float64) float64 {
		if len(xs) == 0 {
			return 0
		}
		if x <= xs[0] {
			return ys[0]
		}
		if x >= xs[len(xs)-1] {
			return ys[len(ys)-1]
		}
		for i := 0; i < len(xs)-1; i++ {
			if x >= xs[i] && x <= xs[i+1] {
				t := (x - xs[i]) / (xs[i+1] - xs[i])
				return ys[i] + t*(ys[i+1]-ys[i])
			}
		}
		return ys[len(ys)-1]
	}
}

func InverseApprox(f Function, y, start, end float64, n int) float64 {
	bestX := start
	bestErr := math.Abs(f(start) - y)
	step := (end - start) / float64(n)
	for i := 1; i <= n; i++ {
		x := start + float64(i)*step
		err := math.Abs(f(x) - y)
		if err < bestErr {
			bestErr = err
			bestX = x
		}
	}
	return bestX
}

func ShiftedFunction(f Function, shiftX, shiftY float64) Function {
	return func(x float64) float64 {
		return f(x-shiftX) + shiftY
	}
}

func ClipInput(f Function, low, high float64) Function {
	return func(x float64) float64 {
		if x < low {
			x = low
		}
		if x > high {
			x = high
		}
		return f(x)
	}
}

func WrapInput(f Function, period float64) Function {
	if period == 0 {
		return f
	}
	return func(x float64) float64 {
		v := math.Mod(x, period)
		if v < 0 {
			v += period
		}
		return f(v)
	}
}

func LinearInterpolation(a, b float64) Function {
	return func(t float64) float64 {
		return a + (b-a)*t
	}
}

func QuadraticInterpolation(p0, p1, p2 float64) Function {
	return func(t float64) float64 {
		u := 1 - t
		return u*u*p0 + 2*u*t*p1 + t*t*p2
	}
}

func CubicInterpolation(p0, p1, p2, p3 float64) Function {
	return func(t float64) float64 {
		u := 1 - t
		return u*u*u*p0 + 3*u*u*t*p1 + 3*u*t*t*p2 + t*t*t*p3
	}
}

func ArcLengthApprox(f Function, a, b float64, n int) float64 {
	if n <= 0 {
		n = 200
	}
	dx := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x1 := a + float64(i)*dx
		x2 := x1 + dx
		y1 := f(x1)
		y2 := f(x2)
		dxv := x2 - x1
		dy := y2 - y1
		sum += math.Sqrt(dxv*dxv + dy*dy)
	}
	return sum
}

func AreaBetween(f, g Function, a, b float64, n int) float64 {
	if n <= 0 {
		n = 200
	}
	dx := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + float64(i)*dx
		sum += math.Abs(f(x)-g(x)) * dx
	}
	return sum
}

func StepFunction(threshold float64, low, high float64) Function {
	return func(x float64) float64 {
		if x < threshold {
			return low
		}
		return high
	}
}

func RampFunction(start, end float64) Function {
	return func(x float64) float64 {
		if x <= start {
			return 0
		}
		if x >= end {
			return 1
		}
		return (x - start) / (end - start)
	}
}

func SaturatingFunction(k float64) Function {
	return func(x float64) float64 {
		return x / (1 + k*math.Abs(x))
	}
}

func SoftClip(f Function, k float64) Function {
	return func(x float64) float64 {
		v := f(x)
		return v / (1 + k*math.Abs(v))
	}
}
