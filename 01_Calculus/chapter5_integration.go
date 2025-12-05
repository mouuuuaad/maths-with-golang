package calculus

func RiemannSumLeft(f Function, a, b float64, n int) float64 {
	dx := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + float64(i)*dx
		sum += f(x)
	}
	return sum * dx
}

func RiemannSumRight(f Function, a, b float64, n int) float64 {
	dx := (b - a) / float64(n)
	sum := 0.0
	for i := 1; i <= n; i++ {
		x := a + float64(i)*dx
		sum += f(x)
	}
	return sum * dx
}

func RiemannSumMidpoint(f Function, a, b float64, n int) float64 {
	dx := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + (float64(i)+0.5)*dx
		sum += f(x)
	}
	return sum * dx
}

func TrapezoidalRule(f Function, a, b float64, n int) float64 {
	dx := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2
	for i := 1; i < n; i++ {
		x := a + float64(i)*dx
		sum += f(x)
	}
	return sum * dx
}

func SimpsonRule(f Function, a, b float64, n int) float64 {
	if n%2 != 0 {
		n++
	}
	dx := (b - a) / float64(n)
	sum := f(a) + f(b)
	for i := 1; i < n; i++ {
		x := a + float64(i)*dx
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}
	return sum * dx / 3
}

func Simpson38Rule(f Function, a, b float64, n int) float64 {
	for n%3 != 0 {
		n++
	}
	dx := (b - a) / float64(n)
	sum := f(a) + f(b)
	for i := 1; i < n; i++ {
		x := a + float64(i)*dx
		if i%3 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 3 * f(x)
		}
	}
	return sum * 3 * dx / 8
}

func BooleRule(f Function, a, b float64) float64 {
	h := (b - a) / 4
	return (b - a) / 90 * (7*f(a) + 32*f(a+h) + 12*f(a+2*h) + 32*f(a+3*h) + 7*f(b))
}

func MonteCarlo(f Function, a, b float64, n int) float64 {
	sum := 0.0
	seed := uint64(12345)
	for i := 0; i < n; i++ {
		seed = seed*1103515245 + 12345
		r := float64(seed%1000000) / 1000000.0
		x := a + r*(b-a)
		sum += f(x)
	}
	return (b - a) * sum / float64(n)
}

func ImproperIntegral(f Function, a float64, n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		b := a + float64(i*i)
		sum += TrapezoidalRule(f, a, b, 100)
		a = b
	}
	return sum
}
