package calculus

type Function func(float64) float64

func DomainCheck(f Function, x float64) bool {
	val := f(x)
	return val == val && val < 1e308 && val > -1e308
}

func FindDomainApprox(f Function, start, end, step float64) (float64, float64) {
	minX := 1e308
	maxX := -1e308
	found := false
	for x := start; x <= end; x += step {
		if DomainCheck(f, x) {
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			found = true
		}
	}
	if !found {
		return 0, 0
	}
	return minX, maxX
}

func Evaluate(f Function, x float64) float64 {
	return f(x)
}

func Compose(f, g Function) Function {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func Add(f, g Function) Function {
	return func(x float64) float64 {
		return f(x) + g(x)
	}
}

func Multiply(f, g Function) Function {
	return func(x float64) float64 {
		return f(x) * g(x)
	}
}

func Scale(f Function, c float64) Function {
	return func(x float64) float64 {
		return c * f(x)
	}
}

func Translate(f Function, h float64) Function {
	return func(x float64) float64 {
		return f(x - h)
	}
}
