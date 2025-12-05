package calculus

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

func Divergence(F VectorFunction, x []float64) float64 {
	n := len(x)
	div := 0.0
	h := 1e-7
	for i := 0; i < n; i++ {
		xPlus := make([]float64, n)
		xMinus := make([]float64, n)
		copy(xPlus, x)
		copy(xMinus, x)
		xPlus[i] += h
		xMinus[i] -= h
		div += (F(xPlus)[i] - F(xMinus)[i]) / (2 * h)
	}
	return div
}

func Curl3D(F VectorFunction, x []float64) []float64 {
	if len(x) != 3 {
		return nil
	}
	h := 1e-7
	partial := func(comp, var_ int) float64 {
		xPlus := make([]float64, 3)
		xMinus := make([]float64, 3)
		copy(xPlus, x)
		copy(xMinus, x)
		xPlus[var_] += h
		xMinus[var_] -= h
		return (F(xPlus)[comp] - F(xMinus)[comp]) / (2 * h)
	}
	return []float64{
		partial(2, 1) - partial(1, 2),
		partial(0, 2) - partial(2, 0),
		partial(1, 0) - partial(0, 1),
	}
}

func Laplacian(f MultiFunction, x []float64) float64 {
	n := len(x)
	h := 1e-5
	lap := 0.0
	for i := 0; i < n; i++ {
		xPlus := make([]float64, n)
		xMinus := make([]float64, n)
		copy(xPlus, x)
		copy(xMinus, x)
		xPlus[i] += h
		xMinus[i] -= h
		lap += (f(xPlus) - 2*f(x) + f(xMinus)) / (h * h)
	}
	return lap
}
