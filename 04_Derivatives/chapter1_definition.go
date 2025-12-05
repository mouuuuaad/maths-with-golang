package derivatives

type Function func(float64) float64

func absD(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

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

func IsDifferentiable(f Function, x float64) bool {
	left := DerivativeBackward(f, x)
	right := DerivativeForward(f, x)
	return absD(left-right) < 1e-5
}

func DerivativeHighPrecision(f Function, x float64) float64 {
	h := 1e-4
	return (-f(x+2*h) + 8*f(x+h) - 8*f(x-h) + f(x-2*h)) / (12 * h)
}

func DerivativeRichardson(f Function, x, h float64, n int) float64 {
	D := make([][]float64, n)
	for i := range D {
		D[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		hi := h / powerD(2, float64(i))
		D[i][0] = (f(x+hi) - f(x-hi)) / (2 * hi)
	}
	for j := 1; j < n; j++ {
		for i := j; i < n; i++ {
			factor := powerD(4, float64(j))
			D[i][j] = (factor*D[i][j-1] - D[i-1][j-1]) / (factor - 1)
		}
	}
	return D[n-1][n-1]
}

func powerD(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
}
