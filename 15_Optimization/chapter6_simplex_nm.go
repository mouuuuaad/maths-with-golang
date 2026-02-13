// 2026 Update: Simplex And Direct Search
package optimization

import "math"

type NelderMeadSettings struct {
	MaxIter        int
	Tolerance      float64
	Alpha          float64
	Gamma          float64
	Rho            float64
	Sigma          float64
	SimplexStep    float64
	StallLimit     int
	MinImprovement float64
}

func DefaultNelderMeadSettings() NelderMeadSettings {
	return NelderMeadSettings{
		MaxIter:        1500,
		Tolerance:      1e-8,
		Alpha:          1.0,
		Gamma:          2.0,
		Rho:            0.5,
		Sigma:          0.5,
		SimplexStep:    0.5,
		StallLimit:     80,
		MinImprovement: 1e-9,
	}
}

func addVectors(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}

func subVectors(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] - b[i]
	}
	return out
}

func scaleVector(a []float64, s float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] * s
	}
	return out
}

func addScaled(base, dir []float64, scale float64) []float64 {
	out := make([]float64, len(base))
	for i := range base {
		out[i] = base[i] + scale*dir[i]
	}
	return out
}

func vectorNormInf(a []float64) float64 {
	maxv := 0.0
	for i := range a {
		av := absO(a[i])
		if av > maxv {
			maxv = av
		}
	}
	return maxv
}

func simplexCentroid(simplex [][]float64, exclude int) []float64 {
	n := len(simplex[0])
	out := make([]float64, n)
	count := 0
	for i := range simplex {
		if i == exclude {
			continue
		}
		count++
		for j := 0; j < n; j++ {
			out[j] += simplex[i][j]
		}
	}
	if count == 0 {
		return out
	}
	inv := 1.0 / float64(count)
	for j := 0; j < n; j++ {
		out[j] *= inv
	}
	return out
}

func argmin(values []float64) int {
	idx := 0
	for i := 1; i < len(values); i++ {
		if values[i] < values[idx] {
			idx = i
		}
	}
	return idx
}

func argmax(values []float64) int {
	idx := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[idx] {
			idx = i
		}
	}
	return idx
}

func argsecond(values []float64, worst int) int {
	idx := -1
	for i := 0; i < len(values); i++ {
		if i == worst {
			continue
		}
		if idx == -1 || values[i] > values[idx] {
			idx = i
		}
	}
	if idx == -1 {
		return 0
	}
	return idx
}

func simplexSpread(simplex [][]float64) float64 {
	if len(simplex) == 0 {
		return 0
	}
	center := simplexCentroid(simplex, -1)
	maxd := 0.0
	for i := range simplex {
		d := vectorNormInf(subVectors(simplex[i], center))
		if d > maxd {
			maxd = d
		}
	}
	return maxd
}

func NelderMead(f ObjectiveFunc, x0 []float64, tol float64) []float64 {
	settings := DefaultNelderMeadSettings()
	settings.Tolerance = tol
	return NelderMeadWithSettings(f, x0, settings)
}

func NelderMeadWithSettings(f ObjectiveFunc, x0 []float64, settings NelderMeadSettings) []float64 {
	n := len(x0)
	if n == 0 {
		return nil
	}
	simplex := make([][]float64, n+1)
	simplex[0] = cloneVector(x0)
	for i := 1; i <= n; i++ {
		p := cloneVector(x0)
		p[i-1] += settings.SimplexStep
		simplex[i] = p
	}
	values := make([]float64, n+1)
	for i := range simplex {
		values[i] = f(simplex[i])
	}
	stall := 0
	bestPrev := values[argmin(values)]
	for iter := 0; iter < settings.MaxIter; iter++ {
		best := argmin(values)
		worst := argmax(values)
		second := argsecond(values, worst)
		if absO(values[worst]-values[best]) < settings.Tolerance {
			return simplex[best]
		}
		if simplexSpread(simplex) < settings.Tolerance {
			return simplex[best]
		}
		centroid := simplexCentroid(simplex, worst)
		reflected := addScaled(centroid, subVectors(centroid, simplex[worst]), settings.Alpha)
		fr := f(reflected)
		if fr < values[second] && fr >= values[best] {
			copy(simplex[worst], reflected)
			values[worst] = fr
		} else if fr < values[best] {
			expanded := addScaled(centroid, subVectors(reflected, centroid), settings.Gamma)
			fe := f(expanded)
			if fe < fr {
				copy(simplex[worst], expanded)
				values[worst] = fe
			} else {
				copy(simplex[worst], reflected)
				values[worst] = fr
			}
		} else {
			contracted := addScaled(centroid, subVectors(simplex[worst], centroid), settings.Rho)
			fc := f(contracted)
			if fc < values[worst] {
				copy(simplex[worst], contracted)
				values[worst] = fc
			} else {
				bestPoint := simplex[best]
				for i := range simplex {
					if i == best {
						continue
					}
					for j := 0; j < n; j++ {
						simplex[i][j] = bestPoint[j] + settings.Sigma*(simplex[i][j]-bestPoint[j])
					}
					values[i] = f(simplex[i])
				}
			}
		}
		bestNow := values[argmin(values)]
		if bestPrev-bestNow < settings.MinImprovement {
			stall++
		} else {
			stall = 0
			bestPrev = bestNow
		}
		if stall >= settings.StallLimit {
			return simplex[argmin(values)]
		}
	}
	return simplex[argmin(values)]
}

func CoordinateSearch(f ObjectiveFunc, x0 []float64, step, tol float64, maxIter int) []float64 {
	x := cloneVector(x0)
	if step == 0 {
		step = 0.1
	}
	for iter := 0; iter < maxIter; iter++ {
		improved := false
		for i := range x {
			bestVal := f(x)
			cand := cloneVector(x)
			cand[i] += step
			candVal := f(cand)
			if candVal < bestVal {
				x = cand
				bestVal = candVal
				improved = true
			} else {
				cand[i] -= 2 * step
				candVal = f(cand)
				if candVal < bestVal {
					x = cand
					improved = true
				}
			}
		}
		if !improved {
			step *= 0.5
			if step < tol {
				return x
			}
		}
	}
	return x
}

func HookeJeeves(f ObjectiveFunc, x0 []float64, step, tol float64, maxIter int) []float64 {
	base := cloneVector(x0)
	best := cloneVector(x0)
	bestVal := f(best)
	if step == 0 {
		step = 0.5
	}
	for iter := 0; iter < maxIter; iter++ {
		next := cloneVector(base)
		for i := range next {
			cand := cloneVector(next)
			cand[i] += step
			candVal := f(cand)
			if candVal < bestVal {
				next = cand
				bestVal = candVal
				continue
			}
			cand[i] -= 2 * step
			candVal = f(cand)
			if candVal < bestVal {
				next = cand
				bestVal = candVal
			}
		}
		if vectorNormInf(subVectors(next, base)) < tol {
			return next
		}
		if f(next) < f(base) {
			pattern := subVectors(next, base)
			trial := addVectors(next, pattern)
			if f(trial) < f(next) {
				base = next
				best = trial
				bestVal = f(best)
			} else {
				base = next
				best = next
				bestVal = f(best)
			}
		} else {
			step *= 0.5
			if step < tol {
				return best
			}
		}
	}
	return best
}

func PowellDirectionSet(f ObjectiveFunc, x0 []float64, tol float64, maxIter int) []float64 {
	n := len(x0)
	x := cloneVector(x0)
	if n == 0 {
		return x
	}
	dirs := make([][]float64, n)
	for i := 0; i < n; i++ {
		d := make([]float64, n)
		d[i] = 1
		dirs[i] = d
	}
	for iter := 0; iter < maxIter; iter++ {
		start := cloneVector(x)
		startVal := f(start)
		bestDir := -1
		bestImprove := 0.0
		for i := 0; i < n; i++ {
			line := func(alpha float64) float64 {
				cand := addScaled(x, dirs[i], alpha)
				return f(cand)
			}
			alpha := lineSearchBrent(line, -1.0, 1.0, tol)
			x = addScaled(x, dirs[i], alpha)
			improve := startVal - f(x)
			if improve > bestImprove {
				bestImprove = improve
				bestDir = i
			}
		}
		if vectorNormInf(subVectors(x, start)) < tol {
			return x
		}
		dir := subVectors(x, start)
		if bestDir >= 0 {
			dirs[bestDir] = dir
		}
	}
	return x
}

func lineSearchBrent(f func(float64) float64, a, b, tol float64) float64 {
	x := a
	w := x
	v := x
	fx := f(x)
	fw := fx
	fv := fx
	d := 0.0
	e := 0.0
	for i := 0; i < 100; i++ {
		m := 0.5 * (a + b)
		tol1 := tol*absO(x) + 1e-12
		if absO(x-m) <= 2*tol1-0.5*(b-a) {
			return x
		}
		var p, q, r float64
		if absO(e) > tol1 {
			r = (x - w) * (fx - fv)
			q = (x - v) * (fx - fw)
			p = (x-v)*q - (x-w)*r
			q = 2 * (q - r)
			if q > 0 {
				p = -p
			} else {
				q = -q
			}
			if absO(p) < absO(0.5*q*e) && p > q*(a-x) && p < q*(b-x) {
				d = p / q
			} else {
				e = b - x
				if x > m {
					e = a - x
				}
				d = 0.3819660112501051 * e
			}
		} else {
			e = b - x
			if x > m {
				e = a - x
			}
			d = 0.3819660112501051 * e
		}
		u := x + d
		if absO(d) < tol1 {
			if d > 0 {
				u = x + tol1
			} else {
				u = x - tol1
			}
		}
		fu := f(u)
		if fu <= fx {
			if u < x {
				b = x
			} else {
				a = x
			}
			v, fv = w, fw
			w, fw = x, fx
			x, fx = u, fu
		} else {
			if u < x {
				a = u
			} else {
				b = u
			}
			if fu <= fw || w == x {
				v, fv = w, fw
				w, fw = u, fu
			} else if fu <= fv || v == x || v == w {
				v, fv = u, fu
			}
		}
	}
	return x
}

func AdaptiveDirectSearch(f ObjectiveFunc, x0 []float64, step, tol float64, maxIter int) []float64 {
	x := cloneVector(x0)
	if step == 0 {
		step = 0.25
	}
	directions := make([][]float64, len(x))
	for i := range directions {
		d := make([]float64, len(x))
		d[i] = 1
		directions[i] = d
	}
	for iter := 0; iter < maxIter; iter++ {
		improved := false
		bestVal := f(x)
		bestPoint := cloneVector(x)
		for _, d := range directions {
			cand := addScaled(x, d, step)
			val := f(cand)
			if val < bestVal {
				bestVal = val
				bestPoint = cand
				improved = true
				continue
			}
			cand = addScaled(x, d, -step)
			val = f(cand)
			if val < bestVal {
				bestVal = val
				bestPoint = cand
				improved = true
			}
		}
		if improved {
			x = bestPoint
		} else {
			step *= 0.5
			if step < tol {
				return x
			}
		}
	}
	return x
}

func BoxComplexMethod(f ObjectiveFunc, bounds [][2]float64, x0 []float64, tol float64, maxIter int) []float64 {
	n := len(x0)
	if n == 0 {
		return nil
	}
	m := n*2 + 1
	simplex := make([][]float64, m)
	simplex[0] = projectIntoBounds(x0, bounds)
	for i := 1; i < m; i++ {
		p := cloneVector(simplex[0])
		idx := (i - 1) % n
		shift := 0.5
		p[idx] += shift
		simplex[i] = projectIntoBounds(p, bounds)
	}
	values := make([]float64, m)
	for i := range simplex {
		values[i] = f(simplex[i])
	}
	for iter := 0; iter < maxIter; iter++ {
		worst := argmax(values)
		best := argmin(values)
		if absO(values[worst]-values[best]) < tol {
			return simplex[best]
		}
		centroid := simplexCentroid(simplex, worst)
		refl := addScaled(centroid, subVectors(centroid, simplex[worst]), 1.3)
		refl = projectIntoBounds(refl, bounds)
		fr := f(refl)
		if fr < values[worst] {
			copy(simplex[worst], refl)
			values[worst] = fr
		} else {
			for i := range simplex {
				if i == best {
					continue
				}
				for j := 0; j < n; j++ {
					simplex[i][j] = 0.5 * (simplex[i][j] + simplex[best][j])
				}
				values[i] = f(simplex[i])
			}
		}
	}
	return simplex[argmin(values)]
}

func projectIntoBounds(x []float64, bounds [][2]float64) []float64 {
	out := cloneVector(x)
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

func TrustRegionDirectSearch(f ObjectiveFunc, x0 []float64, radius, tol float64, maxIter int) []float64 {
	x := cloneVector(x0)
	if radius == 0 {
		radius = 1.0
	}
	directions := make([][]float64, len(x))
	for i := range directions {
		d := make([]float64, len(x))
		d[i] = 1
		directions[i] = d
	}
	for iter := 0; iter < maxIter; iter++ {
		bestVal := f(x)
		bestPoint := cloneVector(x)
		for _, d := range directions {
			cand := addScaled(x, d, radius)
			val := f(cand)
			if val < bestVal {
				bestVal = val
				bestPoint = cand
			}
			cand = addScaled(x, d, -radius)
			val = f(cand)
			if val < bestVal {
				bestVal = val
				bestPoint = cand
			}
		}
		if vectorNormInf(subVectors(bestPoint, x)) < tol {
			return bestPoint
		}
		if bestVal < f(x) {
			x = bestPoint
			radius *= 1.2
			if radius > 10 {
				radius = 10
			}
		} else {
			radius *= 0.5
			if radius < tol {
				return x
			}
		}
	}
	return x
}

func RandomRestartNelderMead(f ObjectiveFunc, seeds [][]float64, tol float64, maxIter int) []float64 {
	best := []float64{}
	bestVal := math.Inf(1)
	settings := DefaultNelderMeadSettings()
	settings.MaxIter = maxIter
	settings.Tolerance = tol
	for _, seed := range seeds {
		cand := NelderMeadWithSettings(f, seed, settings)
		val := f(cand)
		if val < bestVal {
			bestVal = val
			best = cand
		}
	}
	return best
}

func SimplexProjection(x []float64) []float64 {
	out := cloneVector(x)
	sum := 0.0
	for i := range out {
		if out[i] < 0 {
			out[i] = 0
		}
		sum += out[i]
	}
	if sum == 0 {
		inv := 1.0 / float64(len(out))
		for i := range out {
			out[i] = inv
		}
		return out
	}
	inv := 1.0 / sum
	for i := range out {
		out[i] *= inv
	}
	return out
}

func CoordinateDescentQuadratic(Q [][]float64, b []float64, x0 []float64, iters int) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < iters; iter++ {
		for i := range x {
			den := Q[i][i]
			if den == 0 {
				continue
			}
			sum := b[i]
			for j := range x {
				if j == i {
					continue
				}
				sum -= Q[i][j] * x[j]
			}
			x[i] = sum / den
		}
	}
	return x
}

func MirrorDescent(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, step float64, iters int) []float64 {
	x := cloneVector(x0)
	for iter := 0; iter < iters; iter++ {
		g := grad(x)
		for i := range x {
			x[i] = x[i] * math.Exp(-step*g[i])
		}
		x = SimplexProjection(x)
	}
	return x
}
