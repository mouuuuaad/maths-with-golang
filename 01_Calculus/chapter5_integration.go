// 2026 Update: Integration
package calculus

import "math"

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

func CompositeTrapezoidal(f Function, a, b float64, n int) float64 {
	return TrapezoidalRule(f, a, b, n)
}

func CompositeSimpson(f Function, a, b float64, n int) float64 {
	return SimpsonRule(f, a, b, n)
}

func RombergIntegration(f Function, a, b float64, levels int) float64 {
	R := make([][]float64, levels)
	for i := 0; i < levels; i++ {
		R[i] = make([]float64, levels)
	}
	for i := 0; i < levels; i++ {
		n := int(math.Pow(2, float64(i)))
		R[i][0] = TrapezoidalRule(f, a, b, n)
		for k := 1; k <= i; k++ {
			factor := math.Pow(4, float64(k))
			R[i][k] = (factor*R[i][k-1] - R[i-1][k-1]) / (factor - 1)
		}
	}
	return R[levels-1][levels-1]
}

func AdaptiveSimpson(f Function, a, b float64, tol float64, maxRec int) float64 {
	fa := f(a)
	fb := f(b)
	fm := f((a + b) / 2)
	s := simpsonSegment(fa, fm, fb, a, b)
	return adaptiveSimpsonRec(f, a, b, fa, fm, fb, s, tol, maxRec)
}

func simpsonSegment(fa, fm, fb, a, b float64) float64 {
	return (b - a) * (fa + 4*fm + fb) / 6
}

func adaptiveSimpsonRec(f Function, a, b, fa, fm, fb, whole, tol float64, depth int) float64 {
	m := (a + b) / 2
	lm := (a + m) / 2
	rm := (m + b) / 2
	flm := f(lm)
	frm := f(rm)
	left := simpsonSegment(fa, flm, fm, a, m)
	right := simpsonSegment(fm, frm, fb, m, b)
	if depth <= 0 || math.Abs(left+right-whole) < 15*tol {
		return left + right + (left+right-whole)/15
	}
	return adaptiveSimpsonRec(f, a, m, fa, flm, fm, left, tol/2, depth-1) + adaptiveSimpsonRec(f, m, b, fm, frm, fb, right, tol/2, depth-1)
}

func GaussianLegendre2(f Function, a, b float64) float64 {
	m := (a + b) / 2
	h := (b - a) / 2
	x1 := -1 / math.Sqrt(3)
	x2 := 1 / math.Sqrt(3)
	return h * (f(m+h*x1) + f(m+h*x2))
}

func GaussianLegendre3(f Function, a, b float64) float64 {
	m := (a + b) / 2
	h := (b - a) / 2
	x1 := -math.Sqrt(3.0 / 5.0)
	x2 := 0.0
	x3 := math.Sqrt(3.0 / 5.0)
	w1 := 5.0 / 9.0
	w2 := 8.0 / 9.0
	w3 := 5.0 / 9.0
	return h * (w1*f(m+h*x1) + w2*f(m+h*x2) + w3*f(m+h*x3))
}

func GaussianLegendre4(f Function, a, b float64) float64 {
	m := (a + b) / 2
	h := (b - a) / 2
	x1 := 0.3399810435848563
	x2 := 0.8611363115940526
	w1 := 0.6521451548625461
	w2 := 0.3478548451374538
	return h * (w1*(f(m-h*x1)+f(m+h*x1)) + w2*(f(m-h*x2)+f(m+h*x2)))
}

func MidpointRule(f Function, a, b float64) float64 {
	m := (a + b) / 2
	return (b - a) * f(m)
}

func LeftEndpointRule(f Function, a, b float64) float64 {
	return (b - a) * f(a)
}

func RightEndpointRule(f Function, a, b float64) float64 {
	return (b - a) * f(b)
}

func MonteCarloIntegration(f Function, a, b float64, samples int, seed uint64) float64 {
	rng := NewRNG(seed)
	sum := 0.0
	for i := 0; i < samples; i++ {
		x := a + rng.Float64()*(b-a)
		sum += f(x)
	}
	return (b - a) * sum / float64(samples)
}

func MonteCarloImportance(f Function, sampler func(*RNG) float64, weight func(float64) float64, samples int, seed uint64) float64 {
	rng := NewRNG(seed)
	sum := 0.0
	for i := 0; i < samples; i++ {
		x := sampler(rng)
		w := weight(x)
		if w != 0 {
			sum += f(x) / w
		}
	}
	return sum / float64(samples)
}

func IntegrationErrorEstimate(f Function, a, b float64, n int) float64 {
	trap := TrapezoidalRule(f, a, b, n)
	trap2 := TrapezoidalRule(f, a, b, 2*n)
	return math.Abs(trap2-trap) / 3
}

func SimpsonErrorEstimate(f Function, a, b float64, n int) float64 {
	s1 := SimpsonRule(f, a, b, n)
	s2 := SimpsonRule(f, a, b, 2*n)
	return math.Abs(s2-s1) / 15
}

func CumulativeIntegral(f Function, a, b float64, n int) ([]float64, []float64) {
	xs := make([]float64, n+1)
	vals := make([]float64, n+1)
	dx := (b - a) / float64(n)
	sum := 0.0
	prev := f(a)
	xs[0] = a
	vals[0] = 0
	for i := 1; i <= n; i++ {
		x := a + float64(i)*dx
		cur := f(x)
		sum += 0.5 * (cur + prev) * dx
		xs[i] = x
		vals[i] = sum
		prev = cur
	}
	return xs, vals
}

func IntegrationByParts(u, vPrime Function, a, b float64) float64 {
	v := func(x float64) float64 {
		return TrapezoidalRule(vPrime, a, x, 200)
	}
	term1 := u(b)*v(b) - u(a)*v(a)
	integrand := func(x float64) float64 {
		return Derivative(u, x) * v(x)
	}
	term2 := TrapezoidalRule(integrand, a, b, 200)
	return term1 - term2
}

func IntegralTransform(f Function, a, b float64, transform func(float64) float64, jacobian func(float64) float64, n int) float64 {
	g := func(t float64) float64 {
		x := transform(t)
		return f(x) * jacobian(t)
	}
	return TrapezoidalRule(g, a, b, n)
}

func DoubleIntegralRect(f Function, ax, bx, ay, by float64, nx, ny int) float64 {
	dx := (bx - ax) / float64(nx)
	dy := (by - ay) / float64(ny)
	sum := 0.0
	for i := 0; i <= nx; i++ {
		x := ax + float64(i)*dx
		for j := 0; j <= ny; j++ {
			y := ay + float64(j)*dy
			w := 1.0
			if i == 0 || i == nx {
				w *= 0.5
			}
			if j == 0 || j == ny {
				w *= 0.5
			}
			sum += w * f(x+y)
		}
	}
	return sum * dx * dy
}

func IntegralMeanValue(f Function, a, b float64) float64 {
	avg := TrapezoidalRule(f, a, b, 200) / (b - a)
	return avg
}

func AdaptiveTrapezoid(f Function, a, b float64, tol float64) float64 {
	n := 1
	prev := TrapezoidalRule(f, a, b, n)
	for i := 0; i < 12; i++ {
		n *= 2
		curr := TrapezoidalRule(f, a, b, n)
		if math.Abs(curr-prev) < tol {
			return curr
		}
		prev = curr
	}
	return prev
}

func GaussianLegendre5(f Function, a, b float64) float64 {
	m := (a + b) / 2
	h := (b - a) / 2
	x1 := 0.0
	x2 := 0.5384693101056831
	x3 := 0.9061798459386640
	w1 := 0.5688888888888889
	w2 := 0.4786286704993665
	w3 := 0.2369268850561891
	return h * (w1*f(m+h*x1) + w2*(f(m-h*x2)+f(m+h*x2)) + w3*(f(m-h*x3)+f(m+h*x3)))
}

func IntegrateAbsolute(f Function, a, b float64, n int) float64 {
	g := func(x float64) float64 {
		return math.Abs(f(x))
	}
	return TrapezoidalRule(g, a, b, n)
}

func ImproperIntegral(f Function, a, b float64, tol float64) float64 {
	limit := 1.0
	prev := TrapezoidalRule(f, a, b, 200)
	for i := 0; i < 10; i++ {
		limit *= 2
		curr := TrapezoidalRule(f, a, b, 400)
		if math.Abs(curr-prev) < tol {
			return curr
		}
		prev = curr
	}
	return prev
}

func PrincipalValueIntegral(f Function, a, b, c float64, n int) float64 {
	if c <= a || c >= b {
		return TrapezoidalRule(f, a, b, n)
	}
	left := TrapezoidalRule(f, a, c-1e-6, n/2)
	right := TrapezoidalRule(f, c+1e-6, b, n/2)
	return left + right
}

func ConvolutionIntegral(f, g Function, t float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i <= n; i++ {
		x := a + float64(i)*h
		sum += f(x) * g(t-x)
	}
	return sum * h
}

func CumulativeAverage(f Function, a, b float64, n int) []float64 {
	xs, vals := CumulativeIntegral(f, a, b, n)
	avg := make([]float64, len(vals))
	for i := range vals {
		length := xs[i] - a
		if length == 0 {
			avg[i] = 0
		} else {
			avg[i] = vals[i] / length
		}
	}
	return avg
}

func IntegratePiecewise(f Function, points []float64, n int) float64 {
	if len(points) < 2 {
		return 0
	}
	sum := 0.0
	for i := 0; i < len(points)-1; i++ {
		sum += TrapezoidalRule(f, points[i], points[i+1], n)
	}
	return sum
}

func IntegratePositivePart(f Function, a, b float64, n int) float64 {
	g := func(x float64) float64 {
		v := f(x)
		if v < 0 {
			return 0
		}
		return v
	}
	return TrapezoidalRule(g, a, b, n)
}

func IntegrateNegativePart(f Function, a, b float64, n int) float64 {
	g := func(x float64) float64 {
		v := f(x)
		if v > 0 {
			return 0
		}
		return -v
	}
	return TrapezoidalRule(g, a, b, n)
}

func IntegralConvergence(f Function, a float64, n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += f(a + float64(i))
	}
	return sum
}

func AdaptiveMidpoint(f Function, a, b float64, tol float64) float64 {
	n := 1
	prev := RiemannSumMidpoint(f, a, b, n)
	for i := 0; i < 12; i++ {
		n *= 2
		curr := RiemannSumMidpoint(f, a, b, n)
		if math.Abs(curr-prev) < tol {
			return curr
		}
		prev = curr
	}
	return prev
}

func IntegralToInfinity(f Function, a float64, tol float64) float64 {
	b := a + 1
	prev := TrapezoidalRule(f, a, b, 200)
	for i := 0; i < 10; i++ {
		b *= 2
		curr := TrapezoidalRule(f, a, b, 400)
		if math.Abs(curr-prev) < tol {
			return curr
		}
		prev = curr
	}
	return prev
}

func IntegralFromInfinity(f Function, b float64, tol float64) float64 {
	a := b - 1
	prev := TrapezoidalRule(f, a, b, 200)
	for i := 0; i < 10; i++ {
		a *= 2
		curr := TrapezoidalRule(f, a, b, 400)
		if math.Abs(curr-prev) < tol {
			return curr
		}
		prev = curr
	}
	return prev
}
