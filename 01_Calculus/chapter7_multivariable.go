// 2026 Update: Multivariable Calculus
package calculus

import "math"

type MultiFunction func([]float64) float64
type VectorFunction func([]float64) []float64

func PartialDerivative(f MultiFunction, x []float64, i int) float64 {
	h := 1e-7
	xPlus := make([]float64, len(x))
	xMinus := make([]float64, len(x))
	copy(xPlus, x)
	copy(xMinus, x)
	xPlus[i] += h
	xMinus[i] -= h
	return (f(xPlus) - f(xMinus)) / (2 * h)
}

func Gradient(f MultiFunction, x []float64) []float64 {
	n := len(x)
	grad := make([]float64, n)
	for i := 0; i < n; i++ {
		grad[i] = PartialDerivative(f, x, i)
	}
	return grad
}

func Hessian(f MultiFunction, x []float64) [][]float64 {
	n := len(x)
	H := make([][]float64, n)
	for i := range H {
		H[i] = make([]float64, n)
	}
	h := 1e-5
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			xpp := make([]float64, n)
			xpm := make([]float64, n)
			xmp := make([]float64, n)
			xmm := make([]float64, n)
			copy(xpp, x)
			copy(xpm, x)
			copy(xmp, x)
			copy(xmm, x)
			xpp[i] += h
			xpp[j] += h
			xpm[i] += h
			xpm[j] -= h
			xmp[i] -= h
			xmp[j] += h
			xmm[i] -= h
			xmm[j] -= h
			H[i][j] = (f(xpp) - f(xpm) - f(xmp) + f(xmm)) / (4 * h * h)
		}
	}
	return H
}

func Jacobian(F VectorFunction, x []float64) [][]float64 {
	n := len(x)
	y := F(x)
	m := len(y)
	J := make([][]float64, m)
	for i := range J {
		J[i] = make([]float64, n)
	}
	h := 1e-7
	for j := 0; j < n; j++ {
		xPlus := make([]float64, n)
		xMinus := make([]float64, n)
		copy(xPlus, x)
		copy(xMinus, x)
		xPlus[j] += h
		xMinus[j] -= h
		yPlus := F(xPlus)
		yMinus := F(xMinus)
		for i := 0; i < m; i++ {
			J[i][j] = (yPlus[i] - yMinus[i]) / (2 * h)
		}
	}
	return J
}

func Laplacian(f MultiFunction, x []float64) float64 {
	diag := HessianDiagonalApprox(f, x, 1e-4)
	sum := 0.0
	for _, v := range diag {
		sum += v
	}
	return sum
}

func Divergence(F VectorFunction, x []float64) float64 {
	J := Jacobian(F, x)
	sum := 0.0
	for i := 0; i < len(J) && i < len(J[i]); i++ {
		sum += J[i][i]
	}
	return sum
}

func Curl2D(F VectorFunction, x []float64) float64 {
	J := Jacobian(F, x)
	if len(J) < 2 || len(J[0]) < 2 || len(J[1]) < 2 {
		return 0
	}
	return J[1][0] - J[0][1]
}

func LineIntegral2D(F VectorFunction, curve func(float64) []float64, a, b float64, n int) float64 {
	if n <= 0 {
		n = 200
	}
	dt := (b - a) / float64(n)
	sum := 0.0
	prev := curve(a)
	for i := 1; i <= n; i++ {
		t := a + float64(i)*dt
		curr := curve(t)
		mid := []float64{0.5 * (prev[0] + curr[0]), 0.5 * (prev[1] + curr[1])}
		v := F(mid)
		dx := curr[0] - prev[0]
		dy := curr[1] - prev[1]
		sum += v[0]*dx + v[1]*dy
		prev = curr
	}
	return sum
}

func SurfaceIntegralScalar(f MultiFunction, xRange, yRange [2]float64, n int) float64 {
	if n <= 0 {
		n = 100
	}
	x0, x1 := xRange[0], xRange[1]
	y0, y1 := yRange[0], yRange[1]
	dx := (x1 - x0) / float64(n)
	dy := (y1 - y0) / float64(n)
	sum := 0.0
	for i := 0; i <= n; i++ {
		x := x0 + float64(i)*dx
		for j := 0; j <= n; j++ {
			y := y0 + float64(j)*dy
			w := 1.0
			if i == 0 || i == n {
				w *= 0.5
			}
			if j == 0 || j == n {
				w *= 0.5
			}
			sum += w * f([]float64{x, y})
		}
	}
	return sum * dx * dy
}

func GradientDescentMulti(f MultiFunction, x0 []float64, step float64, iters int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for iter := 0; iter < iters; iter++ {
		g := Gradient(f, x)
		for i := range x {
			x[i] -= step * g[i]
		}
	}
	return x
}

func NewtonStepMulti(f MultiFunction, x []float64) []float64 {
	H := Hessian(f, x)
	g := Gradient(f, x)
	step := solveLinearSystem(H, g)
	out := make([]float64, len(x))
	for i := range x {
		out[i] = x[i] - step[i]
	}
	return out
}

func solveLinearSystem(A [][]float64, b []float64) []float64 {
	n := len(b)
	M := make([][]float64, n)
	for i := 0; i < n; i++ {
		M[i] = make([]float64, n+1)
		copy(M[i], A[i])
		M[i][n] = b[i]
	}
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if math.Abs(M[j][i]) > math.Abs(M[pivot][i]) {
				pivot = j
			}
		}
		M[i], M[pivot] = M[pivot], M[i]
		if M[i][i] == 0 {
			continue
		}
		inv := 1.0 / M[i][i]
		for k := i; k <= n; k++ {
			M[i][k] *= inv
		}
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			factor := M[j][i]
			for k := i; k <= n; k++ {
				M[j][k] -= factor * M[i][k]
			}
		}
	}
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = M[i][n]
	}
	return x
}

func GradientNorm(f MultiFunction, x []float64) float64 {
	g := Gradient(f, x)
	sum := 0.0
	for i := range g {
		sum += g[i] * g[i]
	}
	return math.Sqrt(sum)
}

func DirectionalDerivativeMulti(f MultiFunction, x []float64, dir []float64) float64 {
	if len(x) != len(dir) {
		return 0
	}
	h := 1e-6
	xp := make([]float64, len(x))
	xm := make([]float64, len(x))
	for i := range x {
		xp[i] = x[i] + h*dir[i]
		xm[i] = x[i] - h*dir[i]
	}
	return (f(xp) - f(xm)) / (2 * h)
}

func DirectionalDerivativeUnit(f MultiFunction, x []float64, dir []float64) float64 {
	norm := 0.0
	for i := range dir {
		norm += dir[i] * dir[i]
	}
	if norm == 0 {
		return 0
	}
	norm = math.Sqrt(norm)
	unit := make([]float64, len(dir))
	for i := range dir {
		unit[i] = dir[i] / norm
	}
	return DirectionalDerivativeMulti(f, x, unit)
}

func GradientProjection(f MultiFunction, x []float64, v []float64) float64 {
	g := Gradient(f, x)
	if len(g) != len(v) {
		return 0
	}
	sum := 0.0
	for i := range g {
		sum += g[i] * v[i]
	}
	return sum
}

func HessianTrace(f MultiFunction, x []float64) float64 {
	H := Hessian(f, x)
	sum := 0.0
	for i := 0; i < len(H); i++ {
		sum += H[i][i]
	}
	return sum
}

func HessianDeterminant2D(f MultiFunction, x []float64) float64 {
	H := Hessian(f, x)
	if len(H) < 2 || len(H[0]) < 2 {
		return 0
	}
	return H[0][0]*H[1][1] - H[0][1]*H[1][0]
}

func JacobianDeterminant2D(F VectorFunction, x []float64) float64 {
	J := Jacobian(F, x)
	if len(J) < 2 || len(J[0]) < 2 {
		return 0
	}
	return J[0][0]*J[1][1] - J[0][1]*J[1][0]
}

func GradientFlowStep(f MultiFunction, x []float64, step float64) []float64 {
	g := Gradient(f, x)
	out := make([]float64, len(x))
	for i := range x {
		out[i] = x[i] - step*g[i]
	}
	return out
}

func LevelSet(f MultiFunction, x0 []float64, level float64, step float64, iters int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for iter := 0; iter < iters; iter++ {
		g := Gradient(f, x)
		norm := 0.0
		for i := range g {
			norm += g[i] * g[i]
		}
		if norm == 0 {
			break
		}
		norm = math.Sqrt(norm)
		for i := range x {
			x[i] -= step * (f(x) - level) * g[i] / norm
		}
	}
	return x
}

func SurfaceAreaGraph(f MultiFunction, xRange, yRange [2]float64, n int) float64 {
	if n <= 0 {
		n = 100
	}
	x0, x1 := xRange[0], xRange[1]
	y0, y1 := yRange[0], yRange[1]
	dx := (x1 - x0) / float64(n)
	dy := (y1 - y0) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := x0 + float64(i)*dx
		for j := 0; j < n; j++ {
			y := y0 + float64(j)*dy
			grad := Gradient(f, []float64{x, y})
			factor := math.Sqrt(1 + grad[0]*grad[0] + grad[1]*grad[1])
			sum += factor
		}
	}
	return sum * dx * dy
}

func PolarToCartesian(r, theta float64) []float64 {
	return []float64{r * math.Cos(theta), r * math.Sin(theta)}
}

func DivergenceTheorem2D(F VectorFunction, xRange, yRange [2]float64, n int) float64 {
	if n <= 0 {
		n = 100
	}
	x0, x1 := xRange[0], xRange[1]
	y0, y1 := yRange[0], yRange[1]
	dx := (x1 - x0) / float64(n)
	dy := (y1 - y0) / float64(n)
	sum := 0.0
	for i := 0; i <= n; i++ {
		x := x0 + float64(i)*dx
		for j := 0; j <= n; j++ {
			y := y0 + float64(j)*dy
			sum += Divergence(F, []float64{x, y})
		}
	}
	return sum * dx * dy
}

func GreenTheoremApprox(F VectorFunction, xRange, yRange [2]float64, n int) float64 {
	return DivergenceTheorem2D(func(x []float64) []float64 {
		return []float64{-F(x)[1], F(x)[0]}
	}, xRange, yRange, n)
}

func LineIntegralScalar(f MultiFunction, curve func(float64) []float64, a, b float64, n int) float64 {
	if n <= 0 {
		n = 200
	}
	dt := (b - a) / float64(n)
	sum := 0.0
	prev := curve(a)
	for i := 1; i <= n; i++ {
		t := a + float64(i)*dt
		curr := curve(t)
		mid := []float64{0.5 * (prev[0] + curr[0]), 0.5 * (prev[1] + curr[1])}
		dx := curr[0] - prev[0]
		dy := curr[1] - prev[1]
		len := math.Sqrt(dx*dx + dy*dy)
		sum += f(mid) * len
		prev = curr
	}
	return sum
}

func JacobianDeterminant3D(F VectorFunction, x []float64) float64 {
	J := Jacobian(F, x)
	if len(J) < 3 || len(J[0]) < 3 {
		return 0
	}
	return J[0][0]*(J[1][1]*J[2][2]-J[1][2]*J[2][1]) - J[0][1]*(J[1][0]*J[2][2]-J[1][2]*J[2][0]) + J[0][2]*(J[1][0]*J[2][1]-J[1][1]*J[2][0])
}
