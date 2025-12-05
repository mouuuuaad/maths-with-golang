package derivatives

type MultiFunction func([]float64) float64

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

func DirectionalDerivative(f MultiFunction, x, v []float64) float64 {
	grad := Gradient(f, x)
	norm := 0.0
	for _, vi := range v {
		norm += vi * vi
	}
	norm = sqrtD(norm)
	sum := 0.0
	for i := range grad {
		sum += grad[i] * v[i] / norm
	}
	return sum
}

func MixedPartial(f MultiFunction, x []float64, i, j int) float64 {
	h := 1e-5
	xpp := make([]float64, len(x))
	xpm := make([]float64, len(x))
	xmp := make([]float64, len(x))
	xmm := make([]float64, len(x))
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
	return (f(xpp) - f(xpm) - f(xmp) + f(xmm)) / (4 * h * h)
}

func Hessian(f MultiFunction, x []float64) [][]float64 {
	n := len(x)
	H := make([][]float64, n)
	for i := range H {
		H[i] = make([]float64, n)
		for j := range H[i] {
			H[i][j] = MixedPartial(f, x, i, j)
		}
	}
	return H
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

func sqrtD(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}
