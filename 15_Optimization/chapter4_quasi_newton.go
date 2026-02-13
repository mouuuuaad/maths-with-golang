// 2026 Update: Quasi-Newton And Conjugate Gradient
package optimization

import "math"

type CGSettings struct {
	MaxIter int
	Tol     float64
}

func DefaultCGSettings() CGSettings {
	return CGSettings{
		MaxIter: 1000,
		Tol:     1e-8,
	}
}

func ConjugateGradient(A [][]float64, b []float64, tol float64) []float64 {
	settings := DefaultCGSettings()
	settings.Tol = tol
	return ConjugateGradientWithSettings(A, b, settings)
}

func ConjugateGradientWithSettings(A [][]float64, b []float64, settings CGSettings) []float64 {
	n := len(b)
	x := make([]float64, n)
	r := make([]float64, n)
	copy(r, b)
	p := make([]float64, n)
	copy(p, r)
	rsOld := dotProd(r, r)
	if math.Sqrt(rsOld) <= settings.Tol {
		return x
	}
	for i := 0; i < settings.MaxIter; i++ {
		Ap := matVec(A, p)
		den := dotProd(p, Ap)
		if den == 0 {
			break
		}
		alpha := rsOld / den
		for j := 0; j < n; j++ {
			x[j] += alpha * p[j]
			r[j] -= alpha * Ap[j]
		}
		rsNew := dotProd(r, r)
		if math.Sqrt(rsNew) <= settings.Tol {
			break
		}
		beta := rsNew / rsOld
		for j := 0; j < n; j++ {
			p[j] = r[j] + beta*p[j]
		}
		rsOld = rsNew
	}
	return x
}

func ConjugateGradientDiagonalPrecond(A [][]float64, b []float64, settings CGSettings) []float64 {
	n := len(b)
	x := make([]float64, n)
	M := make([]float64, n)
	for i := 0; i < n; i++ {
		d := A[i][i]
		if d == 0 {
			d = 1
		}
		M[i] = 1 / d
	}
	r := make([]float64, n)
	copy(r, b)
	z := make([]float64, n)
	for i := 0; i < n; i++ {
		z[i] = M[i] * r[i]
	}
	p := make([]float64, n)
	copy(p, z)
	rzOld := dotProd(r, z)
	if math.Sqrt(rzOld) <= settings.Tol {
		return x
	}
	for i := 0; i < settings.MaxIter; i++ {
		Ap := matVec(A, p)
		den := dotProd(p, Ap)
		if den == 0 {
			break
		}
		alpha := rzOld / den
		for j := 0; j < n; j++ {
			x[j] += alpha * p[j]
			r[j] -= alpha * Ap[j]
			z[j] = M[j] * r[j]
		}
		rzNew := dotProd(r, z)
		if math.Sqrt(rzNew) <= settings.Tol {
			break
		}
		beta := rzNew / rzOld
		for j := 0; j < n; j++ {
			p[j] = z[j] + beta*p[j]
		}
		rzOld = rzNew
	}
	return x
}

type NLCGSettings struct {
	MaxIter  int
	Tol      float64
	LineC1   float64
	LineTau  float64
	MethodPR bool
}

func DefaultNLCGSettings() NLCGSettings {
	return NLCGSettings{
		MaxIter:  800,
		Tol:      1e-8,
		LineC1:   1e-4,
		LineTau:  0.5,
		MethodPR: true,
	}
}

func NonlinearConjugateGradient(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings NLCGSettings) []float64 {
	x := cloneVec(x0)
	g := grad(x)
	p := make([]float64, len(x))
	for i := range x {
		p[i] = -g[i]
	}
	for iter := 0; iter < settings.MaxIter; iter++ {
		if vecNorm(g) <= settings.Tol {
			break
		}
		alpha := lineSearchArmijo(f, x, p, settings.LineC1, settings.LineTau)
		for i := range x {
			x[i] += alpha * p[i]
		}
		gNew := grad(x)
		beta := 0.0
		if settings.MethodPR {
			diff := make([]float64, len(x))
			for i := range x {
				diff[i] = gNew[i] - g[i]
			}
			beta = dotProd(gNew, diff) / math.Max(dotProd(g, g), 1e-12)
			if beta < 0 {
				beta = 0
			}
		} else {
			beta = dotProd(gNew, gNew) / math.Max(dotProd(g, g), 1e-12)
		}
		for i := range x {
			p[i] = -gNew[i] + beta*p[i]
		}
		g = gNew
	}
	return x
}

type BFGSSettings struct {
	MaxIter int
	Tol     float64
	LineC1  float64
	LineTau float64
}

func DefaultBFGSSettings() BFGSSettings {
	return BFGSSettings{
		MaxIter: 1000,
		Tol:     1e-8,
		LineC1:  1e-4,
		LineTau: 0.5,
	}
}

func BFGS(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, tol float64) []float64 {
	settings := DefaultBFGSSettings()
	settings.Tol = tol
	return BFGSWithSettings(f, grad, x0, settings)
}

func BFGSWithSettings(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings BFGSSettings) []float64 {
	n := len(x0)
	x := cloneVec(x0)
	H := identityMat(n)
	g := grad(x)
	for iter := 0; iter < settings.MaxIter; iter++ {
		if vecNorm(g) < settings.Tol {
			break
		}
		p := make([]float64, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				p[i] -= H[i][j] * g[j]
			}
		}
		alpha := lineSearchArmijo(f, x, p, settings.LineC1, settings.LineTau)
		s := make([]float64, n)
		for i := 0; i < n; i++ {
			s[i] = alpha * p[i]
			x[i] += s[i]
		}
		gNew := grad(x)
		y := make([]float64, n)
		for i := 0; i < n; i++ {
			y[i] = gNew[i] - g[i]
		}
		rhoDen := dotProd(y, s)
		if rhoDen != 0 {
			rho := 1.0 / rhoDen
			H = updateBFGS(H, s, y, rho)
		}
		g = gNew
	}
	return x
}

func DFP(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings BFGSSettings) []float64 {
	n := len(x0)
	x := cloneVec(x0)
	H := identityMat(n)
	g := grad(x)
	for iter := 0; iter < settings.MaxIter; iter++ {
		if vecNorm(g) < settings.Tol {
			break
		}
		p := matVec(H, scaleVec(g, -1))
		alpha := lineSearchArmijo(f, x, p, settings.LineC1, settings.LineTau)
		s := scaleVec(p, alpha)
		x = addVec(x, s)
		gNew := grad(x)
		y := subVec(gNew, g)
		sy := dotProd(s, y)
		yHy := dotProd(y, matVec(H, y))
		if sy != 0 {
			term1 := outerVec(s, s)
			term2 := outerVec(matVec(H, y), matVec(H, y))
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					H[i][j] += term1[i][j]/sy - term2[i][j]/math.Max(yHy, 1e-12)
				}
			}
		}
		g = gNew
	}
	return x
}

func SR1(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings BFGSSettings) []float64 {
	n := len(x0)
	x := cloneVec(x0)
	H := identityMat(n)
	g := grad(x)
	for iter := 0; iter < settings.MaxIter; iter++ {
		if vecNorm(g) < settings.Tol {
			break
		}
		p := matVec(H, scaleVec(g, -1))
		alpha := lineSearchArmijo(f, x, p, settings.LineC1, settings.LineTau)
		s := scaleVec(p, alpha)
		x = addVec(x, s)
		gNew := grad(x)
		y := subVec(gNew, g)
		Hs := matVec(H, y)
		u := subVec(s, Hs)
		den := dotProd(u, y)
		if absO(den) > 1e-12 {
			outer := outerVec(u, u)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					H[i][j] += outer[i][j] / den
				}
			}
		}
		g = gNew
	}
	return x
}

type LBFGSSettings struct {
	MaxIter int
	Tol     float64
	Memory  int
	LineC1  float64
	LineTau float64
}

func DefaultLBFGSSettings() LBFGSSettings {
	return LBFGSSettings{
		MaxIter: 1000,
		Tol:     1e-8,
		Memory:  6,
		LineC1:  1e-4,
		LineTau: 0.5,
	}
}

func LBFGS(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, settings LBFGSSettings) []float64 {
	x := cloneVec(x0)
	g := grad(x)
	sList := make([][]float64, 0, settings.Memory)
	yList := make([][]float64, 0, settings.Memory)
	rhoList := make([]float64, 0, settings.Memory)
	for iter := 0; iter < settings.MaxIter; iter++ {
		if vecNorm(g) < settings.Tol {
			break
		}
		q := cloneVec(g)
		alpha := make([]float64, len(sList))
		for i := len(sList) - 1; i >= 0; i-- {
			alpha[i] = rhoList[i] * dotProd(sList[i], q)
			q = subVec(q, scaleVec(yList[i], alpha[i]))
		}
		r := scaleVec(q, -1)
		for i := 0; i < len(sList); i++ {
			beta := rhoList[i] * dotProd(yList[i], r)
			r = addVec(r, scaleVec(sList[i], alpha[i]-beta))
		}
		p := r
		step := lineSearchArmijo(f, x, p, settings.LineC1, settings.LineTau)
		s := scaleVec(p, step)
		x = addVec(x, s)
		gNew := grad(x)
		y := subVec(gNew, g)
		sy := dotProd(s, y)
		if sy != 0 {
			if len(sList) == settings.Memory {
				sList = sList[1:]
				yList = yList[1:]
				rhoList = rhoList[1:]
			}
			sList = append(sList, s)
			yList = append(yList, y)
			rhoList = append(rhoList, 1.0/sy)
		}
		g = gNew
	}
	return x
}

func lineSearchArmijo(f ObjectiveFunc, x, p []float64, c1, tau float64) float64 {
	alpha := 1.0
	fx := f(x)
	g := make([]float64, len(x))
	for i := range x {
		g[i] = 0
	}
	for i := 0; i < 10; i++ {
		xNew := addVec(x, scaleVec(p, alpha))
		if f(xNew) <= fx+c1*alpha*dotProd(g, p) {
			return alpha
		}
		alpha *= tau
	}
	return alpha
}

func identityMat(n int) [][]float64 {
	I := make([][]float64, n)
	for i := 0; i < n; i++ {
		I[i] = make([]float64, n)
		I[i][i] = 1
	}
	return I
}

func updateBFGS(H [][]float64, s, y []float64, rho float64) [][]float64 {
	n := len(s)
	I := identityMat(n)
	sy := outerVec(s, y)
	ss := outerVec(s, s)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			I[i][j] -= rho * sy[i][j]
		}
	}
	H = matMat(I, matMat(H, transposeMat(I)))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			H[i][j] += rho * ss[i][j]
		}
	}
	return H
}

func outerVec(a, b []float64) [][]float64 {
	n := len(a)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i] * b[j]
		}
	}
	return result
}

func transposeMat(m [][]float64) [][]float64 {
	n := len(m)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func matMat(a, b [][]float64) [][]float64 {
	n := len(a)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func dotProd(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func matVec(m [][]float64, v []float64) []float64 {
	result := make([]float64, len(v))
	for i := range m {
		for j := range v {
			result[i] += m[i][j] * v[j]
		}
	}
	return result
}

func vecNorm(v []float64) float64 {
	return math.Sqrt(dotProd(v, v))
}

func cloneVec(v []float64) []float64 {
	out := make([]float64, len(v))
	copy(out, v)
	return out
}

func addVec(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}

func subVec(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] - b[i]
	}
	return out
}

func scaleVec(a []float64, s float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] * s
	}
	return out
}
