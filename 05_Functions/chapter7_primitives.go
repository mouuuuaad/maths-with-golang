package functions

type Function func(float64) float64

func Integral(f Function, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x0 := a + float64(i)*h
		x1 := x0 + h
		sum += (f(x0) + f(x1)) / 2 * h
	}
	return sum
}

func SimpsonIntegral(f Function, a, b float64, n int) float64 {
	if n%2 != 0 {
		n++
	}
	h := (b - a) / float64(n)
	sum := f(a) + f(b)
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}
	return sum * h / 3
}

func Primitive(f Function, a, x float64, n int) float64 {
	return Integral(f, a, x, n)
}

func GaussLegendre5(f Function, a, b float64) float64 {
	x := []float64{0.0, 0.538469310105683, -0.538469310105683, 0.906179845938664, -0.906179845938664}
	w := []float64{0.568888888888889, 0.478628670499366, 0.478628670499366, 0.236926885056189, 0.236926885056189}
	mid := (b + a) / 2
	halfLen := (b - a) / 2
	sum := 0.0
	for i := 0; i < 5; i++ {
		sum += w[i] * f(mid+halfLen*x[i])
	}
	return sum * halfLen
}

func AdaptiveSimpson(f Function, a, b, eps float64) float64 {
	return adaptiveSimpsonRecursive(f, a, b, eps, simpsonOneStep(f, a, b), 10)
}

func simpsonOneStep(f Function, a, b float64) float64 {
	c := (a + b) / 2
	return (b - a) / 6 * (f(a) + 4*f(c) + f(b))
}

func adaptiveSimpsonRecursive(f Function, a, b, eps, whole float64, depth int) float64 {
	c := (a + b) / 2
	left := simpsonOneStep(f, a, c)
	right := simpsonOneStep(f, c, b)
	if depth <= 0 || Abs(left+right-whole) < 15*eps {
		return left + right + (left+right-whole)/15
	}
	return adaptiveSimpsonRecursive(f, a, c, eps/2, left, depth-1) +
		adaptiveSimpsonRecursive(f, c, b, eps/2, right, depth-1)
}

func RombergIntegration(f Function, a, b float64, n int) float64 {
	R := make([][]float64, n)
	for i := range R {
		R[i] = make([]float64, n)
	}
	h := b - a
	R[0][0] = h / 2 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		h /= 2
		sum := 0.0
		for k := 1; k <= (1 << (i - 1)); k++ {
			sum += f(a + float64(2*k-1)*h)
		}
		R[i][0] = R[i-1][0]/2 + h*sum
		for j := 1; j <= i; j++ {
			factor := Power(4, float64(j))
			R[i][j] = (factor*R[i][j-1] - R[i-1][j-1]) / (factor - 1)
		}
	}
	return R[n-1][n-1]
}
