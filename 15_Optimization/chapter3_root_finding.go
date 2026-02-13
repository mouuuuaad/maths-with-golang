// 2026 Update: Root Finding
package optimization

func maxOF64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func minOF64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func safeMidpoint(a, b float64) float64 {
	return a + (b-a)/2
}

type RootSettings struct {
	MaxIter   int
	AbsTol    float64
	RelTol    float64
	DerivStep float64
	Grow      float64
}

func DefaultRootSettings() RootSettings {
	return RootSettings{
		MaxIter:   200,
		AbsTol:    1e-10,
		RelTol:    1e-10,
		DerivStep: 1e-6,
		Grow:      1.6,
	}
}

type RootResult struct {
	Root       float64
	Iterations int
	Converged  bool
	Residual   float64
	Bracketed  bool
}

func rootConverged(x, prev, fx, absTol, relTol float64) bool {
	if absO(fx) <= absTol {
		return true
	}
	step := absO(x - prev)
	scale := absTol + relTol*maxOF64(absO(x), absO(prev))
	return step <= scale
}

func RootBracket(f func(float64) float64, a, b, grow float64, maxSteps int) (float64, float64, bool) {
	fa := f(a)
	fb := f(b)
	if fa == 0 {
		return a, a, true
	}
	if fb == 0 {
		return b, b, true
	}
	if fa*fb < 0 {
		return a, b, true
	}
	left := a
	right := b
	step := absO(b-a) * grow
	if step == 0 {
		step = grow
	}
	for i := 0; i < maxSteps; i++ {
		left -= step
		right += step
		fa = f(left)
		fb = f(right)
		if fa*fb < 0 {
			return left, right, true
		}
		step *= grow
	}
	return a, b, false
}

func BisectionMethod(f func(float64) float64, a, b, tol float64) float64 {
	fa := f(a)
	fb := f(b)
	if fa == 0 {
		return a
	}
	if fb == 0 {
		return b
	}
	if fa*fb > 0 {
		la, lb, ok := RootBracket(f, a, b, 1.6, 16)
		if ok {
			a, b = la, lb
			fa, fb = f(a), f(b)
		}
	}
	for i := 0; i < 200; i++ {
		m := safeMidpoint(a, b)
		fm := f(m)
		if absO(fm) <= tol || absO(b-a) <= tol {
			return m
		}
		if fa*fm < 0 {
			b = m
			fb = fm
		} else {
			a = m
			fa = fm
		}
	}
	return safeMidpoint(a, b)
}

func SecantMethod(f func(float64) float64, x0, x1, tol float64) float64 {
	prev := x0
	x := x1
	f0 := f(prev)
	f1 := f(x)
	for i := 0; i < 200; i++ {
		if absO(f1) <= tol {
			return x
		}
		den := f1 - f0
		if den == 0 {
			return x
		}
		next := x - f1*(x-prev)/den
		if absO(next-x) <= tol {
			return next
		}
		prev, x = x, next
		f0, f1 = f1, f(x)
	}
	return x
}

func RegulaFalsiMethod(f func(float64) float64, a, b, tol float64) float64 {
	fa := f(a)
	fb := f(b)
	if fa*fb > 0 {
		return safeMidpoint(a, b)
	}
	x := a
	for i := 0; i < 200; i++ {
		x = (a*fb - b*fa) / (fb - fa)
		fx := f(x)
		if absO(fx) <= tol {
			return x
		}
		if fa*fx < 0 {
			b, fb = x, fx
		} else {
			a, fa = x, fx
		}
		if absO(b-a) <= tol {
			return x
		}
	}
	return x
}

func IllinoisMethod(f func(float64) float64, a, b, tol float64) float64 {
	fa := f(a)
	fb := f(b)
	if fa*fb > 0 {
		return safeMidpoint(a, b)
	}
	x := a
	wa := 1.0
	wb := 1.0
	for i := 0; i < 200; i++ {
		x = (a*fb*wb - b*fa*wa) / (fb*wb - fa*wa)
		fx := f(x)
		if absO(fx) <= tol {
			return x
		}
		if fa*fx < 0 {
			b, fb = x, fx
			wb = 1.0
			wa *= 0.5
		} else {
			a, fa = x, fx
			wa = 1.0
			wb *= 0.5
		}
		if absO(b-a) <= tol {
			return x
		}
	}
	return x
}

func RidderMethod(f func(float64) float64, a, b, tol float64) float64 {
	fa := f(a)
	fb := f(b)
	if fa*fb > 0 {
		return safeMidpoint(a, b)
	}
	x := a
	for i := 0; i < 200; i++ {
		m := safeMidpoint(a, b)
		fm := f(m)
		if absO(fm) <= tol {
			return m
		}
		s := fm*fm - fa*fb
		if s < 0 {
			s = 0
		}
		den := 1.0
		if s > 0 {
			den = 1.0 / s
		}
		q := 0.0
		if fm != 0 {
			q = (m - a) * fm * den
		}
		if fa-fb < 0 {
			q = -q
		}
		x = m + q
		fx := f(x)
		if absO(fx) <= tol {
			return x
		}
		if fm*fx < 0 {
			a, fa = m, fm
			b, fb = x, fx
		} else if fa*fx < 0 {
			b, fb = x, fx
		} else {
			a, fa = x, fx
		}
		if absO(b-a) <= tol {
			return x
		}
	}
	return x
}

func BrentMethod(f func(float64) float64, a, b, tol float64) float64 {
	fa, fb := f(a), f(b)
	if fa*fb > 0 {
		return safeMidpoint(a, b)
	}
	if absO(fa) < absO(fb) {
		a, b = b, a
		fa, fb = fb, fa
	}
	c, fc := a, fa
	d := b - a
	e := d
	for i := 0; i < 200; i++ {
		if fb == 0 {
			return b
		}
		if fa*fb > 0 {
			a, fa = c, fc
			d = b - a
			e = d
		}
		if absO(fa) < absO(fb) {
			c, fc = b, fb
			b, fb = a, fa
			a, fa = c, fc
		}
		m := 0.5 * (a - b)
		tol1 := 2*tol*maxOF64(1, absO(b)) + 0.5*tol
		if absO(m) <= tol1 {
			return b
		}
		if absO(e) >= tol1 && absO(fc) > absO(fb) {
			var p, q float64
			s := fb / fc
			if a == c {
				p = 2 * m * s
				q = 1 - s
			} else {
				q = fc / fa
				r := fb / fa
				p = s * (2*m*q*(q-r) - (b-c)*(r-1))
				q = (q - 1) * (r - 1) * (s - 1)
			}
			if p > 0 {
				q = -q
			} else {
				p = -p
			}
			cond1 := 2*p < 3*m*q-absO(tol1*q)
			cond2 := p < absO(0.5*e*q)
			if cond1 && cond2 {
				e = d
				d = p / q
			} else {
				e = d
				d = m
			}
		} else {
			e = d
			d = m
		}
		c, fc = b, fb
		if absO(d) > tol1 {
			b += d
		} else if m > 0 {
			b += tol1
		} else {
			b -= tol1
		}
		fb = f(b)
	}
	return b
}

func NewtonRaphson(f, df func(float64) float64, x0 float64, settings RootSettings) RootResult {
	x := x0
	prev := x0
	for i := 0; i < settings.MaxIter; i++ {
		fx := f(x)
		dfx := df(x)
		if dfx == 0 {
			return RootResult{Root: x, Iterations: i + 1, Converged: false, Residual: fx}
		}
		next := x - fx/dfx
		if rootConverged(next, x, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: next, Iterations: i + 1, Converged: true, Residual: f(next)}
		}
		prev = x
		x = next
		if rootConverged(x, prev, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: x, Iterations: i + 1, Converged: true, Residual: f(x)}
		}
	}
	return RootResult{Root: x, Iterations: settings.MaxIter, Converged: false, Residual: f(x)}
}

func AutoDerivative(f func(float64) float64, x, h float64) float64 {
	h = absO(h)
	if h == 0 {
		h = 1e-6
	}
	return (f(x+h) - f(x-h)) / (2 * h)
}

func AutoSecondDerivative(f func(float64) float64, x, h float64) float64 {
	h = absO(h)
	if h == 0 {
		h = 1e-4
	}
	return (f(x+h) - 2*f(x) + f(x-h)) / (h * h)
}

func NewtonRaphsonAuto(f func(float64) float64, x0 float64, settings RootSettings) RootResult {
	x := x0
	prev := x
	for i := 0; i < settings.MaxIter; i++ {
		fx := f(x)
		dfx := AutoDerivative(f, x, settings.DerivStep)
		if dfx == 0 {
			return RootResult{Root: x, Iterations: i + 1, Converged: false, Residual: fx}
		}
		next := x - fx/dfx
		if rootConverged(next, x, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: next, Iterations: i + 1, Converged: true, Residual: f(next)}
		}
		prev = x
		x = next
		if rootConverged(x, prev, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: x, Iterations: i + 1, Converged: true, Residual: f(x)}
		}
	}
	return RootResult{Root: x, Iterations: settings.MaxIter, Converged: false, Residual: f(x)}
}

func HalleyMethod(f, df, ddf func(float64) float64, x0 float64, settings RootSettings) RootResult {
	x := x0
	prev := x
	for i := 0; i < settings.MaxIter; i++ {
		fx := f(x)
		dfx := df(x)
		ddfx := ddf(x)
		den := 2*dfx*dfx - fx*ddfx
		if den == 0 {
			return RootResult{Root: x, Iterations: i + 1, Converged: false, Residual: fx}
		}
		next := x - (2*fx*dfx)/den
		if rootConverged(next, x, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: next, Iterations: i + 1, Converged: true, Residual: f(next)}
		}
		prev = x
		x = next
		if rootConverged(x, prev, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: x, Iterations: i + 1, Converged: true, Residual: f(x)}
		}
	}
	return RootResult{Root: x, Iterations: settings.MaxIter, Converged: false, Residual: f(x)}
}

func MullerMethod(f func(float64) float64, x0, x1, x2 float64, settings RootSettings) RootResult {
	a := x0
	b := x1
	c := x2
	for i := 0; i < settings.MaxIter; i++ {
		fa := f(a)
		fb := f(b)
		fc := f(c)
		h1 := b - a
		h2 := c - b
		if h1 == 0 || h2 == 0 {
			return RootResult{Root: c, Iterations: i + 1, Converged: false, Residual: fc}
		}
		d1 := (fb - fa) / h1
		d2 := (fc - fb) / h2
		d := (d2 - d1) / (h2 + h1)
		b2 := d2 + h2*d
		disc := b2*b2 - 4*fc*d
		if disc < 0 {
			disc = 0
		}
		den := b2 + absO(b2)
		if den == 0 {
			den = b2 - absO(b2)
		}
		if den == 0 {
			return RootResult{Root: c, Iterations: i + 1, Converged: false, Residual: fc}
		}
		step := -2 * fc / den
		next := c + step
		if rootConverged(next, c, fc, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: next, Iterations: i + 1, Converged: true, Residual: f(next)}
		}
		a, b, c = b, c, next
	}
	return RootResult{Root: c, Iterations: settings.MaxIter, Converged: false, Residual: f(c)}
}

func FixedPointIteration(g func(float64) float64, x0 float64, settings RootSettings) RootResult {
	x := x0
	prev := x
	for i := 0; i < settings.MaxIter; i++ {
		next := g(x)
		fx := next - x
		if rootConverged(next, x, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: next, Iterations: i + 1, Converged: true, Residual: fx}
		}
		prev = x
		x = next
		if rootConverged(x, prev, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: x, Iterations: i + 1, Converged: true, Residual: fx}
		}
	}
	return RootResult{Root: x, Iterations: settings.MaxIter, Converged: false, Residual: g(x) - x}
}

func SteffensenMethod(g func(float64) float64, x0 float64, settings RootSettings) RootResult {
	x := x0
	for i := 0; i < settings.MaxIter; i++ {
		y := g(x)
		z := g(y)
		den := z - 2*y + x
		if den == 0 {
			return RootResult{Root: x, Iterations: i + 1, Converged: false, Residual: y - x}
		}
		next := x - (y-x)*(y-x)/den
		fx := next - x
		if rootConverged(next, x, fx, settings.AbsTol, settings.RelTol) {
			return RootResult{Root: next, Iterations: i + 1, Converged: true, Residual: fx}
		}
		x = next
	}
	return RootResult{Root: x, Iterations: settings.MaxIter, Converged: false, Residual: g(x) - x}
}

func HybridBracketedSecant(f func(float64) float64, a, b float64, settings RootSettings) RootResult {
	fa := f(a)
	fb := f(b)
	if fa*fb > 0 {
		return RootResult{Root: safeMidpoint(a, b), Iterations: 0, Converged: false, Residual: f(safeMidpoint(a, b)), Bracketed: false}
	}
	x0 := a
	x1 := b
	f0 := fa
	f1 := fb
	for i := 0; i < settings.MaxIter; i++ {
		if f1 == f0 {
			m := safeMidpoint(x0, x1)
			fm := f(m)
			if f0*fm < 0 {
				x1, f1 = m, fm
			} else {
				x0, f0 = m, fm
			}
			continue
		}
		sec := x1 - f1*(x1-x0)/(f1-f0)
		if sec < minOF64(x0, x1) || sec > maxOF64(x0, x1) {
			sec = safeMidpoint(x0, x1)
		}
		fs := f(sec)
		if absO(fs) <= settings.AbsTol {
			return RootResult{Root: sec, Iterations: i + 1, Converged: true, Residual: fs, Bracketed: true}
		}
		if f0*fs < 0 {
			x1, f1 = sec, fs
		} else {
			x0, f0 = sec, fs
		}
		if absO(x1-x0) <= settings.AbsTol+settings.RelTol*maxOF64(absO(x0), absO(x1)) {
			m := safeMidpoint(x0, x1)
			return RootResult{Root: m, Iterations: i + 1, Converged: true, Residual: f(m), Bracketed: true}
		}
	}
	m := safeMidpoint(x0, x1)
	return RootResult{Root: m, Iterations: settings.MaxIter, Converged: false, Residual: f(m), Bracketed: true}
}

func RootScan(f func(float64) float64, start, end, step float64) []float64 {
	if step == 0 {
		return nil
	}
	if end < start {
		start, end = end, start
	}
	roots := make([]float64, 0)
	x := start
	fx := f(x)
	for x+step <= end {
		next := x + step
		fn := f(next)
		if fx == 0 {
			roots = append(roots, x)
		} else if fx*fn < 0 {
			root := BisectionMethod(f, x, next, 1e-8)
			roots = append(roots, root)
		}
		x = next
		fx = fn
	}
	if fx == 0 {
		roots = append(roots, x)
	}
	return roots
}

func RootRefineWithBrent(f func(float64) float64, guess, span float64, settings RootSettings) RootResult {
	a := guess - span
	b := guess + span
	la, lb, ok := RootBracket(f, a, b, settings.Grow, 20)
	if !ok {
		return RootResult{Root: guess, Iterations: 0, Converged: false, Residual: f(guess), Bracketed: false}
	}
	root := BrentMethod(f, la, lb, settings.AbsTol)
	return RootResult{Root: root, Iterations: settings.MaxIter, Converged: absO(f(root)) <= settings.AbsTol, Residual: f(root), Bracketed: true}
}
