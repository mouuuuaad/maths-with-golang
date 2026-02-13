// 2026 Update: Constrained Optimization
package optimization

import "math"

type ConstraintSettings struct {
	MaxIter int
	Tol     float64
	Step    float64
}

func DefaultConstraintSettings() ConstraintSettings {
	return ConstraintSettings{
		MaxIter: 1000,
		Tol:     1e-8,
		Step:    0.01,
	}
}

func LagrangeMultiplier(f, g ObjectiveFunc, gradF, gradG func([]float64) []float64, x0 []float64, lambda0, tol float64) ([]float64, float64) {
	settings := DefaultConstraintSettings()
	settings.Tol = tol
	return LagrangeMultiplierWithSettings(f, g, gradF, gradG, x0, lambda0, settings)
}

func LagrangeMultiplierWithSettings(f, g ObjectiveFunc, gradF, gradG func([]float64) []float64, x0 []float64, lambda0 float64, settings ConstraintSettings) ([]float64, float64) {
	x := cloneVector(x0)
	lambda := lambda0
	for iter := 0; iter < settings.MaxIter; iter++ {
		gf := gradF(x)
		gg := gradG(x)
		gVal := g(x)
		maxGrad := 0.0
		for i := range x {
			grad := gf[i] - lambda*gg[i]
			if absO(grad) > maxGrad {
				maxGrad = absO(grad)
			}
			x[i] -= settings.Step * grad
		}
		lambda += settings.Step * gVal
		if maxGrad < settings.Tol && absO(gVal) < settings.Tol {
			break
		}
	}
	return x, lambda
}

func PenaltyMethod(f, g ObjectiveFunc, x0 []float64, rho, tol float64) []float64 {
	settings := DefaultConstraintSettings()
	settings.Tol = tol
	return PenaltyMethodWithSettings(f, g, x0, rho, settings)
}

func PenaltyMethodWithSettings(f, g ObjectiveFunc, x0 []float64, rho float64, settings ConstraintSettings) []float64 {
	x := cloneVector(x0)
	for k := 0; k < 20; k++ {
		penalty := func(xx []float64) float64 {
			gVal := g(xx)
			return f(xx) + rho*gVal*gVal
		}
		x = NelderMead(penalty, x, settings.Tol)
		if absO(g(x)) < settings.Tol {
			break
		}
		rho *= 10
	}
	return x
}

func BarrierMethod(f ObjectiveFunc, inequalities []ObjectiveFunc, x0 []float64, mu, tol float64) []float64 {
	settings := DefaultConstraintSettings()
	settings.Tol = tol
	return BarrierMethodWithSettings(f, inequalities, x0, mu, settings)
}

func BarrierMethodWithSettings(f ObjectiveFunc, inequalities []ObjectiveFunc, x0 []float64, mu float64, settings ConstraintSettings) []float64 {
	x := cloneVector(x0)
	for k := 0; k < 20; k++ {
		barrier := func(xx []float64) float64 {
			val := f(xx)
			for _, g := range inequalities {
				gVal := g(xx)
				if gVal <= 0 {
					return 1e18
				}
				val -= mu * math.Log(gVal)
			}
			return val
		}
		x = NelderMead(barrier, x, settings.Tol)
		mu *= 0.1
		if mu < settings.Tol {
			break
		}
	}
	return x
}

type AugmentedLagrangianSettings struct {
	MaxIter int
	Tol     float64
	Rho     float64
	Step    float64
}

func DefaultAugmentedLagrangianSettings() AugmentedLagrangianSettings {
	return AugmentedLagrangianSettings{
		MaxIter: 200,
		Tol:     1e-8,
		Rho:     1.0,
		Step:    0.05,
	}
}

func AugmentedLagrangian(f ObjectiveFunc, g ObjectiveFunc, gradF, gradG func([]float64) []float64, x0 []float64, settings AugmentedLagrangianSettings) []float64 {
	x := cloneVector(x0)
	lambda := 0.0
	for iter := 0; iter < settings.MaxIter; iter++ {
		grad := make([]float64, len(x))
		gf := gradF(x)
		gg := gradG(x)
		gVal := g(x)
		for i := range x {
			grad[i] = gf[i] + (lambda+settings.Rho*gVal)*gg[i]
		}
		for i := range x {
			x[i] -= settings.Step * grad[i]
		}
		lambda += settings.Rho * gVal
		if absO(gVal) < settings.Tol && vecNormCons(grad) < settings.Tol {
			break
		}
		settings.Rho *= 1.2
	}
	return x
}

func ProjectedGradient(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, bounds [][2]float64, settings ConstraintSettings) []float64 {
	x := projectBounds(x0, bounds)
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		for i := range x {
			x[i] -= settings.Step * g[i]
		}
		x = projectBounds(x, bounds)
		if vecNormCons(g) < settings.Tol {
			break
		}
	}
	return x
}

func FeasibleDirection(f ObjectiveFunc, grad func([]float64) []float64, ineq []ObjectiveFunc, x0 []float64, settings ConstraintSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		g := grad(x)
		for i := range x {
			x[i] -= settings.Step * g[i]
		}
		if satisfiesIneq(ineq, x, settings.Tol) {
			if vecNormCons(g) < settings.Tol {
				break
			}
		} else {
			for i := range x {
				x[i] += 0.5 * settings.Step * g[i]
			}
		}
	}
	return x
}

func satisfiesIneq(ineq []ObjectiveFunc, x []float64, tol float64) bool {
	for _, g := range ineq {
		if g(x) > tol {
			return false
		}
	}
	return true
}

func projectBounds(x []float64, bounds [][2]float64) []float64 {
	out := cloneVector(x)
	if bounds == nil || len(bounds) != len(x) {
		return out
	}
	for i := range out {
		low := bounds[i][0]
		high := bounds[i][1]
		if low > high {
			low, high = high, low
		}
		if out[i] < low {
			out[i] = low
		}
		if out[i] > high {
			out[i] = high
		}
	}
	return out
}

func QuadraticPenalty(f ObjectiveFunc, constraints []ObjectiveFunc, x0 []float64, rho float64, settings ConstraintSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		penalty := func(xx []float64) float64 {
			val := f(xx)
			for _, g := range constraints {
				gv := g(xx)
				val += rho * gv * gv
			}
			return val
		}
		x = NelderMead(penalty, x, settings.Tol)
		if maxConstraintViolation(constraints, x) < settings.Tol {
			break
		}
		rho *= 2
	}
	return x
}

func maxConstraintViolation(constraints []ObjectiveFunc, x []float64) float64 {
	maxv := 0.0
	for _, g := range constraints {
		v := absO(g(x))
		if v > maxv {
			maxv = v
		}
	}
	return maxv
}

func SequentialPenalty(f ObjectiveFunc, g ObjectiveFunc, x0 []float64, rho float64, settings ConstraintSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		obj := func(xx []float64) float64 {
			v := g(xx)
			return f(xx) + rho*v*v
		}
		x = NelderMead(obj, x, settings.Tol)
		if absO(g(x)) < settings.Tol {
			break
		}
		rho *= 3
	}
	return x
}

func ConstraintResidual(g ObjectiveFunc, x []float64) float64 {
	return g(x)
}

func vecNormCons(v []float64) float64 {
	sum := 0.0
	for i := range v {
		sum += v[i] * v[i]
	}
	return math.Sqrt(sum)
}

func FiniteDifferenceGrad(f ObjectiveFunc, x []float64, h float64) []float64 {
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

type BarrierGDSettings struct {
	MaxIter int
	Step    float64
	Mu      float64
	MuDecay float64
	Tol     float64
}

func DefaultBarrierGDSettings() BarrierGDSettings {
	return BarrierGDSettings{
		MaxIter: 500,
		Step:    0.02,
		Mu:      1.0,
		MuDecay: 0.5,
		Tol:     1e-8,
	}
}

func LogBarrierGradientDescent(f ObjectiveFunc, ineq []ObjectiveFunc, x0 []float64, settings BarrierGDSettings) []float64 {
	x := cloneVector(x0)
	mu := settings.Mu
	for iter := 0; iter < settings.MaxIter; iter++ {
		barrier := func(xx []float64) float64 {
			val := f(xx)
			for _, g := range ineq {
				gv := g(xx)
				if gv <= 0 {
					return 1e18
				}
				val -= mu * math.Log(gv)
			}
			return val
		}
		grad := FiniteDifferenceGrad(barrier, x, 1e-6)
		for i := range x {
			x[i] -= settings.Step * grad[i]
		}
		if vecNormCons(grad) < settings.Tol {
			mu *= settings.MuDecay
			if mu < settings.Tol {
				break
			}
		}
	}
	return x
}

func EqualityResidual(g ObjectiveFunc, x []float64) float64 {
	return absO(g(x))
}

func KKTResidual(gradF []float64, gradG []float64, lambda float64) float64 {
	n := len(gradF)
	if len(gradG) != n {
		return math.Inf(1)
	}
	res := 0.0
	for i := 0; i < n; i++ {
		d := gradF[i] - lambda*gradG[i]
		res += d * d
	}
	return math.Sqrt(res)
}

func PenaltySchedule(rho float64, factor float64, steps int) []float64 {
	values := make([]float64, steps)
	cur := rho
	for i := 0; i < steps; i++ {
		values[i] = cur
		cur *= factor
	}
	return values
}

func SequentialBarrier(f ObjectiveFunc, ineq []ObjectiveFunc, x0 []float64, mu float64, steps int, tol float64) []float64 {
	x := cloneVector(x0)
	for i := 0; i < steps; i++ {
		x = BarrierMethod(f, ineq, x, mu, tol)
		mu *= 0.2
		if mu < tol {
			break
		}
	}
	return x
}

func SoftConstraint(f ObjectiveFunc, g ObjectiveFunc, x []float64, rho float64) float64 {
	gv := g(x)
	if gv <= 0 {
		return f(x)
	}
	return f(x) + rho*gv*gv
}

func ProjectedNewtonStep(grad []float64, hessian [][]float64, bounds [][2]float64) []float64 {
	n := len(grad)
	step := make([]float64, n)
	for i := 0; i < n; i++ {
		if i < len(hessian) && i < len(hessian[i]) && hessian[i][i] != 0 {
			step[i] = -grad[i] / hessian[i][i]
		} else {
			step[i] = -grad[i]
		}
	}
	return projectBounds(step, bounds)
}

func AugmentedInequalityPenalty(f ObjectiveFunc, ineq []ObjectiveFunc, x0 []float64, rho float64, settings ConstraintSettings) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < settings.MaxIter; iter++ {
		obj := func(xx []float64) float64 {
			val := f(xx)
			for _, g := range ineq {
				gv := g(xx)
				if gv > 0 {
					val += rho * gv * gv
				}
			}
			return val
		}
		x = NelderMead(obj, x, settings.Tol)
		if maxConstraintViolation(ineq, x) < settings.Tol {
			break
		}
		rho *= 1.5
	}
	return x
}
