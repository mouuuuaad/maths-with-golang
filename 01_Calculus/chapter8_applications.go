// 2026 Update: Calculus Applications
package calculus

import "math"

type DifferentialEquation func(x, y float64) float64

type ODESystem func(x float64, y []float64) []float64

func EulerMethod(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		y += step * f(x, y)
		x += step
	}
	return y
}

func HeunMethod(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		k1 := f(x, y)
		k2 := f(x+step, y+step*k1)
		y += step * (k1 + k2) / 2
		x += step
	}
	return y
}

func MidpointMethod(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		k1 := f(x, y)
		k2 := f(x+step/2, y+step*k1/2)
		y += step * k2
		x += step
	}
	return y
}

func RungeKutta4(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		k1 := step * f(x, y)
		k2 := step * f(x+step/2, y+k1/2)
		k3 := step * f(x+step/2, y+k2/2)
		k4 := step * f(x+step, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
		x += step
	}
	return y
}

func RungeKutta4Full(f DifferentialEquation, x0, y0, xEnd, step float64) ([]float64, []float64) {
	xs := []float64{x0}
	ys := []float64{y0}
	x := x0
	y := y0
	for x < xEnd {
		k1 := step * f(x, y)
		k2 := step * f(x+step/2, y+k1/2)
		k3 := step * f(x+step/2, y+k2/2)
		k4 := step * f(x+step, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
		x += step
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func AdaptiveRK45(f DifferentialEquation, x0, y0, xEnd, tol float64) float64 {
	x := x0
	y := y0
	h := (xEnd - x0) / 100
	for x < xEnd {
		k1 := h * f(x, y)
		k2 := h * f(x+h/4, y+k1/4)
		k3 := h * f(x+3*h/8, y+3*k1/32+9*k2/32)
		k4 := h * f(x+12*h/13, y+1932*k1/2197-7200*k2/2197+7296*k3/2197)
		k5 := h * f(x+h, y+439*k1/216-8*k2+3680*k3/513-845*k4/4104)
		k6 := h * f(x+h/2, y-8*k1/27+2*k2-3544*k3/2565+1859*k4/4104-11*k5/40)
		y4 := y + 25*k1/216 + 1408*k3/2565 + 2197*k4/4104 - k5/5
		y5 := y + 16*k1/135 + 6656*k3/12825 + 28561*k4/56430 - 9*k5/50 + 2*k6/55
		err := absL(y5 - y4)
		if err < tol {
			y = y5
			x += h
			if err > 0 {
				h *= 0.9 * powerS(tol/err, 0.2)
			}
		} else {
			h *= 0.9 * powerS(tol/err, 0.25)
		}
		if h < 1e-8 {
			break
		}
	}
	return y
}

func EulerSystem(f ODESystem, x0 float64, y0 []float64, xEnd, step float64) ([]float64, []float64) {
	x := x0
	y := make([]float64, len(y0))
	copy(y, y0)
	xs := []float64{x}
	for x < xEnd {
		dy := f(x, y)
		for i := range y {
			y[i] += step * dy[i]
		}
		x += step
		xs = append(xs, x)
	}
	return xs, y
}

func RK4System(f ODESystem, x0 float64, y0 []float64, xEnd, step float64) ([]float64, []float64) {
	x := x0
	y := make([]float64, len(y0))
	copy(y, y0)
	xs := []float64{x}
	for x < xEnd {
		k1 := f(x, y)
		y2 := addVecSys(y, scaleVecSys(k1, step/2))
		k2 := f(x+step/2, y2)
		y3 := addVecSys(y, scaleVecSys(k2, step/2))
		k3 := f(x+step/2, y3)
		y4 := addVecSys(y, scaleVecSys(k3, step))
		k4 := f(x+step, y4)
		for i := range y {
			y[i] += step * (k1[i] + 2*k2[i] + 2*k3[i] + k4[i]) / 6
		}
		x += step
		xs = append(xs, x)
	}
	return xs, y
}

func AdamsBashforth2(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	k1 := f(x, y)
	y1 := y + step*k1
	x1 := x + step
	k2 := f(x1, y1)
	x = x1
	y = y1
	for x < xEnd {
		yNext := y + step*(1.5*k2-0.5*k1)
		xNext := x + step
		k1 = k2
		k2 = f(xNext, yNext)
		x = xNext
		y = yNext
	}
	return y
}

func AdamsMoulton2(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	k1 := f(x, y)
	y = y + step*k1
	x = x + step
	for x < xEnd {
		k2 := f(x, y)
		y = y + step*0.5*(k2+k1)
		k1 = k2
		x += step
	}
	return y
}

func PredictorCorrector(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	k1 := f(x, y)
	y = y + step*k1
	x = x + step
	for x < xEnd {
		k2 := f(x, y)
		yPred := y + step*k2
		k3 := f(x+step, yPred)
		y = y + step*0.5*(k2+k3)
		x += step
	}
	return y
}

func ExplicitEulerStability(lambda float64, step float64) bool {
	return absL(1+step*lambda) < 1
}

func LogisticGrowth(r, K float64) DifferentialEquation {
	return func(x, y float64) float64 {
		_ = x
		return r * y * (1 - y/K)
	}
}

func LotkaVolterra(alpha, beta, delta, gamma float64) ODESystem {
	return func(x float64, y []float64) []float64 {
		_ = x
		if len(y) < 2 {
			return []float64{0, 0}
		}
		dx := alpha*y[0] - beta*y[0]*y[1]
		dy := delta*y[0]*y[1] - gamma*y[1]
		return []float64{dx, dy}
	}
}

func DampedOscillator(omega, damping float64) ODESystem {
	return func(x float64, y []float64) []float64 {
		_ = x
		if len(y) < 2 {
			return []float64{0, 0}
		}
		return []float64{y[1], -2*damping*y[1] - omega*omega*y[0]}
	}
}

func EnergyOscillator(y []float64, omega float64) float64 {
	if len(y) < 2 {
		return 0
	}
	return 0.5 * (y[1]*y[1] + omega*omega*y[0]*y[0])
}

func ShootingMethod(f DifferentialEquation, x0, xEnd, y0, target, guess1, guess2, step float64) float64 {
	g := func(slope float64) float64 {
		y := y0
		x := x0
		v := slope
		for x < xEnd {
			k1 := step * v
			l1 := step * f(x, y)
			k2 := step * (v + 0.5*l1)
			l2 := step * f(x+0.5*step, y+0.5*k1)
			k3 := step * (v + 0.5*l2)
			l3 := step * f(x+0.5*step, y+0.5*k2)
			k4 := step * (v + l3)
			l4 := step * f(x+step, y+k3)
			y += (k1 + 2*k2 + 2*k3 + k4) / 6
			v += (l1 + 2*l2 + 2*l3 + l4) / 6
			x += step
		}
		return y - target
	}
	for i := 0; i < 20; i++ {
		f1 := g(guess1)
		f2 := g(guess2)
		if f2-f1 == 0 {
			break
		}
		guess3 := guess2 - f2*(guess2-guess1)/(f2-f1)
		guess1, guess2 = guess2, guess3
		if absL(g(guess2)) < 1e-6 {
			break
		}
	}
	return guess2
}

func StabilityRegionRK4(lambda float64, step float64) float64 {
	z := lambda * step
	return 1 + z + z*z/2 + z*z*z/6 + z*z*z*z/24
}

func ExplicitEulerLocalError(f DifferentialEquation, x, y, step float64) float64 {
	return absL(step * step * Derivative(func(t float64) float64 { return f(t, y) }, x))
}

func addVecSys(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}

func scaleVecSys(a []float64, s float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] * s
	}
	return out
}

func ODEErrorEstimate(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	y1 := RungeKutta4(f, x0, y0, xEnd, step)
	y2 := RungeKutta4(f, x0, y0, xEnd, step/2)
	return absL(y2 - y1)
}

func TimeGrid(x0, xEnd, step float64) []float64 {
	count := int(math.Ceil((xEnd-x0)/step)) + 1
	grid := make([]float64, count)
	for i := 0; i < count; i++ {
		grid[i] = x0 + float64(i)*step
	}
	return grid
}

func EulerFull(f DifferentialEquation, x0, y0, xEnd, step float64) ([]float64, []float64) {
	x := x0
	y := y0
	xs := []float64{x}
	ys := []float64{y}
	for x < xEnd {
		y += step * f(x, y)
		x += step
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func HeunFull(f DifferentialEquation, x0, y0, xEnd, step float64) ([]float64, []float64) {
	x := x0
	y := y0
	xs := []float64{x}
	ys := []float64{y}
	for x < xEnd {
		k1 := f(x, y)
		k2 := f(x+step, y+step*k1)
		y += step * (k1 + k2) / 2
		x += step
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func MidpointFull(f DifferentialEquation, x0, y0, xEnd, step float64) ([]float64, []float64) {
	x := x0
	y := y0
	xs := []float64{x}
	ys := []float64{y}
	for x < xEnd {
		k1 := f(x, y)
		k2 := f(x+step/2, y+step*k1/2)
		y += step * k2
		x += step
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func RichardsonExtrapolation(y1, y2 float64, order int) float64 {
	factor := math.Pow(2, float64(order))
	return y2 + (y2-y1)/(factor-1)
}

func RK4WithExtrapolation(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	y1 := RungeKutta4(f, x0, y0, xEnd, step)
	y2 := RungeKutta4(f, x0, y0, xEnd, step/2)
	return RichardsonExtrapolation(y1, y2, 4)
}

func EulerStabilityTest(f DifferentialEquation, x0, y0, xEnd float64, step float64, reference float64) bool {
	y := EulerMethod(f, x0, y0, xEnd, step)
	return absL(y-reference) < 1e-3
}
