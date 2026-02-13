// 2026 Update: Unconstrained Optimization
package optimization

import "math"

func absO(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

type ObjectiveFunc func(x []float64) float64

type UnconstrainedSettings struct {
	MaxIter   int
	Tol       float64
	Step      float64
	Decay     float64
	Momentum  float64
	Nesterov  bool
	LineC1    float64
	LineTau   float64
	Restart   int
	Seed      uint64
	Direction []float64
}

func DefaultUnconstrainedSettings() UnconstrainedSettings {
	return UnconstrainedSettings{
		MaxIter:  1000,
		Tol:      1e-8,
		Step:     0.1,
		Decay:    0.0,
		Momentum: 0.0,
		Nesterov: false,
		LineC1:   1e-4,
		LineTau:  0.5,
		Restart:  0,
		Seed:     42,
	}
}

func GoldenSection(f func(float64) float64, a, b, tol float64) float64 {
	phi := (1 + 2.2360679774997896) / 2
	x1 := b - (b-a)/phi
	x2 := a + (b-a)/phi
	f1, f2 := f(x1), f(x2)
	for absO(b-a) > tol {
		if f1 < f2 {
			b = x2
			x2 = x1
			f2 = f1
			x1 = b - (b-a)/phi
			f1 = f(x1)
		} else {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + (b-a)/phi
			f2 = f(x2)
		}
	}
	return (a + b) / 2
}

func GradientDescent(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, lr float64, iters int) []float64 {
	settings := DefaultUnconstrainedSettings()
	settings.Step = lr
	settings.MaxIter = iters
	return GradientDescentWithSettings(f, grad, x0, settings)
}

func GradientDescentWithSettings(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings UnconstrainedSettings) []float64 {
	x := cloneVector(x0)
	velocity := make([]float64, len(x))
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		if vecNormU(g) < settings.Tol {
			break
		}
		step := settings.Step
		if settings.Decay > 0 {
			step = settings.Step / (1 + settings.Decay*float64(iter))
		}
		if settings.Momentum > 0 {
			for i := range x {
				velocity[i] = settings.Momentum*velocity[i] + step*g[i]
				x[i] -= velocity[i]
			}
		} else {
			for i := range x {
				x[i] -= step * g[i]
			}
		}
	}
	return x
}

func NesterovDescent(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings UnconstrainedSettings) []float64 {
	x := cloneVector(x0)
	v := make([]float64, len(x))
	for iter := 0; iter < settings.MaxIter; iter++ {
		step := settings.Step
		if settings.Decay > 0 {
			step = settings.Step / (1 + settings.Decay*float64(iter))
		}
		look := make([]float64, len(x))
		for i := range x {
			look[i] = x[i] - settings.Momentum*v[i]
		}
		g := grad(look)
		if vecNormU(g) < settings.Tol {
			break
		}
		for i := range x {
			v[i] = settings.Momentum*v[i] + step*g[i]
			x[i] -= v[i]
		}
	}
	return x
}

func NewtonMethod1D(f, df func(float64) float64, x0, tol float64) float64 {
	x := x0
	for i := 0; i < 100; i++ {
		fx := f(x)
		if absO(fx) < tol {
			return x
		}
		dfx := df(x)
		if dfx == 0 {
			return x
		}
		x = x - fx/dfx
	}
	return x
}

func NewtonMethodMulti(f ObjectiveFunc, grad func([]float64) []float64, hess func([]float64) [][]float64, x0 []float64, settings UnconstrainedSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		if vecNormU(g) < settings.Tol {
			break
		}
		H := hess(x)
		step := solveDiagonal(H, g)
		for i := range x {
			x[i] -= settings.Step * step[i]
		}
	}
	return x
}

func solveDiagonal(H [][]float64, g []float64) []float64 {
	out := make([]float64, len(g))
	for i := range g {
		den := 1.0
		if i < len(H) && i < len(H[i]) && H[i][i] != 0 {
			den = H[i][i]
		}
		out[i] = g[i] / den
	}
	return out
}

func CoordinateDescentUnconstrained(f ObjectiveFunc, x0 []float64, step float64, settings UnconstrainedSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		improved := false
		for i := range x {
			best := x[i]
			bestVal := f(x)
			x[i] += step
			if f(x) < bestVal {
				bestVal = f(x)
				best = x[i]
				improved = true
			} else {
				x[i] -= 2 * step
				if f(x) < bestVal {
					bestVal = f(x)
					best = x[i]
					improved = true
				}
			}
			x[i] = best
		}
		if !improved {
			step *= 0.5
			if step < settings.Tol {
				break
			}
		}
	}
	return x
}

func BacktrackingLineSearch(f ObjectiveFunc, x, p []float64, c1, tau float64) float64 {
	alpha := 1.0
	fx := f(x)
	g := finiteDiffGradU(f, x, 1e-6)
	gp := 0.0
	for i := range x {
		gp += g[i] * p[i]
	}
	for i := 0; i < 20; i++ {
		xNew := addVecU(x, scaleVecU(p, alpha))
		if f(xNew) <= fx+c1*alpha*gp {
			return alpha
		}
		alpha *= tau
	}
	return alpha
}

func SteepestDescentLineSearch(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings UnconstrainedSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		if vecNormU(g) < settings.Tol {
			break
		}
		p := scaleVecU(g, -1)
		alpha := BacktrackingLineSearch(f, x, p, settings.LineC1, settings.LineTau)
		x = addVecU(x, scaleVecU(p, alpha))
	}
	return x
}

func RandomRestartGradientDescent(f ObjectiveFunc, grad func([]float64) []float64, seeds [][]float64, settings UnconstrainedSettings) []float64 {
	best := cloneVector(seeds[0])
	bestVal := f(best)
	for _, s := range seeds {
		cand := GradientDescentWithSettings(f, grad, s, settings)
		val := f(cand)
		if val < bestVal {
			bestVal = val
			best = cand
		}
	}
	return best
}

func finiteDiffGradU(f ObjectiveFunc, x []float64, h float64) []float64 {
	if h == 0 {
		h = 1e-6
	}
	grad := make([]float64, len(x))
	for i := range x {
		xp := cloneVector(x)
		xm := cloneVector(x)
		xp[i] += h
		xm[i] -= h
		grad[i] = (f(xp) - f(xm)) / (2 * h)
	}
	return grad
}

func vecNormU(v []float64) float64 {
	sum := 0.0
	for i := range v {
		sum += v[i] * v[i]
	}
	return math.Sqrt(sum)
}

func cloneVector(v []float64) []float64 {
	out := make([]float64, len(v))
	copy(out, v)
	return out
}

func addVecU(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}

func scaleVecU(a []float64, s float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] * s
	}
	return out
}

func PolyakStepGradientDescent(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, fStar float64, settings UnconstrainedSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		gn := vecNormU(g)
		if gn < settings.Tol {
			break
		}
		step := (f(x) - fStar) / (gn * gn)
		for i := range x {
			x[i] -= step * g[i]
		}
	}
	return x
}

type TrustRegionSettings struct {
	MaxIter int
	Radius  float64
	Tol     float64
}

func DefaultTrustRegionSettings() TrustRegionSettings {
	return TrustRegionSettings{
		MaxIter: 200,
		Radius:  1.0,
		Tol:     1e-8,
	}
}

func TrustRegionCauchy(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings TrustRegionSettings) []float64 {
	x := cloneVector(x0)
	radius := settings.Radius
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		gn := vecNormU(g)
		if gn < settings.Tol {
			break
		}
		step := scaleVecU(g, -radius/gn)
		trial := addVecU(x, step)
		if f(trial) < f(x) {
			x = trial
			radius *= 1.2
			if radius > 10 {
				radius = 10
			}
		} else {
			radius *= 0.5
			if radius < settings.Tol {
				break
			}
		}
	}
	return x
}

func RandomDirectionSearch(f ObjectiveFunc, x0 []float64, step float64, iters int, seed uint64) []float64 {
	rng := NewRNG(seed)
	x := cloneVector(x0)
	best := cloneVector(x0)
	bestVal := f(best)
	for iter := 0; iter < iters; iter++ {
		dir := make([]float64, len(x))
		for i := range dir {
			dir[i] = rng.Float64()*2 - 1
		}
		trial := addVecU(x, scaleVecU(dir, step))
		val := f(trial)
		if val < bestVal {
			bestVal = val
			best = trial
			x = trial
		}
		step *= 0.99
		if step < 1e-8 {
			break
		}
	}
	return best
}

func EstimateHessianDiag(f ObjectiveFunc, x []float64, h float64) []float64 {
	if h == 0 {
		h = 1e-4
	}
	diag := make([]float64, len(x))
	fx := f(x)
	for i := range x {
		xp := cloneVector(x)
		xm := cloneVector(x)
		xp[i] += h
		xm[i] -= h
		diag[i] = (f(xp) - 2*fx + f(xm)) / (h * h)
	}
	return diag
}
