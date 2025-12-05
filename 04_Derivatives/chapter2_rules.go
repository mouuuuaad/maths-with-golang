package derivatives

func SumRule(f, g Function, x float64) float64 {
	return Derivative(f, x) + Derivative(g, x)
}

func ProductRule(f, g Function, x float64) float64 {
	return Derivative(f, x)*g(x) + f(x)*Derivative(g, x)
}

func QuotientRule(f, g Function, x float64) float64 {
	gx := g(x)
	if absD(gx) < 1e-12 {
		return 0
	}
	return (Derivative(f, x)*g(x) - f(x)*Derivative(g, x)) / (gx * gx)
}

func ChainRule(f, g Function, x float64) float64 {
	return Derivative(f, g(x)) * Derivative(g, x)
}

func PowerRule(n float64, x float64) float64 {
	return n * powerD(x, n-1)
}

func ConstantMultiple(c float64, f Function, x float64) float64 {
	return c * Derivative(f, x)
}

func DifferenceRule(f, g Function, x float64) float64 {
	return Derivative(f, x) - Derivative(g, x)
}

func InverseRule(f Function, y float64) float64 {
	_ = f(y)
	df := Derivative(f, y)
	if absD(df) < 1e-12 {
		return 0
	}
	return 1.0 / df
}

func LogarithmicDerivative(f Function, x float64) float64 {
	fx := f(x)
	if absD(fx) < 1e-12 {
		return 0
	}
	return Derivative(f, x) / fx
}
