package derivatives

func NewtonMethod(f Function, x0 float64, maxIter int) float64 {
	x := x0
	for i := 0; i < maxIter; i++ {
		fx := f(x)
		if absD(fx) < 1e-12 {
			return x
		}
		dfx := Derivative(f, x)
		if absD(dfx) < 1e-15 {
			break
		}
		x -= fx / dfx
	}
	return x
}

func GradientDescent(f MultiFunction, x0 []float64, alpha float64, maxIter int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for i := 0; i < maxIter; i++ {
		grad := Gradient(f, x)
		norm := 0.0
		for _, g := range grad {
			norm += g * g
		}
		if sqrtD(norm) < 1e-8 {
			break
		}
		for j := range x {
			x[j] -= alpha * grad[j]
		}
	}
	return x
}

func GradientDescentWithMomentum(f MultiFunction, x0 []float64, alpha, beta float64, maxIter int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	v := make([]float64, len(x0))
	for i := 0; i < maxIter; i++ {
		grad := Gradient(f, x)
		for j := range x {
			v[j] = beta*v[j] + alpha*grad[j]
			x[j] -= v[j]
		}
		norm := 0.0
		for _, g := range grad {
			norm += g * g
		}
		if sqrtD(norm) < 1e-8 {
			break
		}
	}
	return x
}

func Bisection(f Function, a, b float64, tol float64) float64 {
	fa := f(a)
	for i := 0; i < 100; i++ {
		c := (a + b) / 2
		fc := f(c)
		if absD(fc) < tol || (b-a)/2 < tol {
			return c
		}
		if fa*fc < 0 {
			b = c
		} else {
			a = c
			fa = fc
		}
	}
	return (a + b) / 2
}

func Secant(f Function, x0, x1 float64, maxIter int) float64 {
	for i := 0; i < maxIter; i++ {
		f0 := f(x0)
		f1 := f(x1)
		if absD(f1-f0) < 1e-15 {
			break
		}
		x2 := x1 - f1*(x1-x0)/(f1-f0)
		if absD(x2-x1) < 1e-12 {
			return x2
		}
		x0 = x1
		x1 = x2
	}
	return x1
}

func FindExtremum(f Function, start float64) float64 {
	df := func(x float64) float64 { return Derivative(f, x) }
	return NewtonMethod(df, start, 100)
}
