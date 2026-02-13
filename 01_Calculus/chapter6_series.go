// 2026 Update: Series
package calculus

import "math"

type Sequence func(int) float64

func GeometricSeriesSum(a, r float64) float64 {
	if absL(r) >= 1 {
		return 1e308
	}
	return a / (1 - r)
}

func PartialSum(seq Sequence, n int) float64 {
	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += seq(i)
	}
	return sum
}

func RatioTest(seq Sequence, n int) string {
	an := seq(n)
	an1 := seq(n + 1)
	if absL(an) < 1e-15 {
		return "Inconclusive"
	}
	ratio := absL(an1 / an)
	if ratio < 0.99 {
		return "Converges"
	}
	if ratio > 1.01 {
		return "Diverges"
	}
	return "Inconclusive"
}

func RootTest(seq Sequence, n int) string {
	an := absL(seq(n))
	if an == 0 {
		return "Inconclusive"
	}
	root := powerS(an, 1.0/float64(n))
	if root < 0.99 {
		return "Converges"
	}
	if root > 1.01 {
		return "Diverges"
	}
	return "Inconclusive"
}

func IntegralTest(f Function, a float64, n int) float64 {
	return TrapezoidalRule(f, a, float64(n), n*10)
}

func ComparisonTest(seq1, seq2 Sequence, n int) bool {
	for i := 1; i <= n; i++ {
		if absL(seq1(i)) > absL(seq2(i)) {
			return false
		}
	}
	return true
}

func LimitComparisonTest(seq1, seq2 Sequence, n int) float64 {
	a := seq1(n)
	b := seq2(n)
	if absL(b) < 1e-15 {
		return 0
	}
	return a / b
}

func AlternatingSeries(seq Sequence, n int) float64 {
	sum := 0.0
	sign := 1.0
	for i := 0; i <= n; i++ {
		sum += sign * absL(seq(i))
		sign = -sign
	}
	return sum
}

func AlternatingTest(seq Sequence, n int) bool {
	for i := 1; i <= n; i++ {
		if seq(i) > seq(i-1) {
			return false
		}
	}
	return true
}

func TelescopingSum(seq Sequence, n int) float64 {
	return PartialSum(seq, n)
}

func PowerSeries(coeffs []float64) Function {
	return func(x float64) float64 {
		sum := 0.0
		pow := 1.0
		for i := 0; i < len(coeffs); i++ {
			sum += coeffs[i] * pow
			pow *= x
		}
		return sum
	}
}

func TaylorSeries(f Function, a float64, n int) Function {
	coeffs := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		coeffs[i] = NthDerivative(f, a, i) / factorialS(i)
	}
	return func(x float64) float64 {
		sum := 0.0
		pow := 1.0
		h := x - a
		for i := 0; i <= n; i++ {
			sum += coeffs[i] * pow
			pow *= h
		}
		return sum
	}
}

func MaclaurinSeries(f Function, n int) Function {
	return TaylorSeries(f, 0, n)
}

func SeriesExp(x float64, n int) float64 {
	sum := 1.0
	term := 1.0
	for i := 1; i <= n; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func SeriesSin(x float64, n int) float64 {
	sum := 0.0
	sign := 1.0
	term := x
	for i := 1; i <= 2*n-1; i += 2 {
		sum += sign * term
		term *= x * x / float64((i+1)*(i+2))
		sign = -sign
	}
	return sum
}

func SeriesCos(x float64, n int) float64 {
	sum := 1.0
	sign := -1.0
	term := x * x
	for i := 2; i <= 2*n; i += 2 {
		sum += sign * term
		term *= x * x / float64((i+1)*(i+2))
		sign = -sign
	}
	return sum
}

func SeriesLog1p(x float64, n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		term := powerS(-1, float64(i+1)) * powerS(x, float64(i)) / float64(i)
		sum += term
	}
	return sum
}

func SeriesArctan(x float64, n int) float64 {
	sum := 0.0
	sign := 1.0
	pow := x
	for i := 1; i <= 2*n-1; i += 2 {
		sum += sign * pow / float64(i)
		pow *= x * x
		sign = -sign
	}
	return sum
}

func SeriesSinh(x float64, n int) float64 {
	sum := 0.0
	term := x
	for i := 1; i <= 2*n-1; i += 2 {
		sum += term
		term *= x * x / float64((i+1)*(i+2))
	}
	return sum
}

func SeriesCosh(x float64, n int) float64 {
	sum := 1.0
	term := x * x
	for i := 2; i <= 2*n; i += 2 {
		sum += term
		term *= x * x / float64((i+1)*(i+2))
	}
	return sum
}

func SeriesBinomial(alpha, x float64, n int) float64 {
	sum := 1.0
	term := 1.0
	for i := 1; i <= n; i++ {
		term *= (alpha - float64(i-1)) * x / float64(i)
		sum += term
	}
	return sum
}

func SeriesConverges(seq Sequence, tol float64, n int) bool {
	prev := seq(1)
	for i := 2; i <= n; i++ {
		curr := seq(i)
		if absL(curr-prev) < tol {
			return true
		}
		prev = curr
	}
	return false
}

func SeriesLimit(seq Sequence, n int) float64 {
	prev := seq(1)
	for i := 2; i <= n; i++ {
		curr := seq(i)
		if absL(curr-prev) < 1e-9 {
			return curr
		}
		prev = curr
	}
	return prev
}

func SequenceLimit(seq Sequence, n int) float64 {
	return SeriesLimit(seq, n)
}

func SequenceDifference(seq Sequence, n int) Sequence {
	return func(k int) float64 {
		return seq(k+1) - seq(k)
	}
}

func SequenceRatio(seq Sequence, n int) Sequence {
	return func(k int) float64 {
		den := seq(k)
		if den == 0 {
			return 0
		}
		return seq(k+1) / den
	}
}

func SequencePartialSums(seq Sequence, n int) []float64 {
	sums := make([]float64, n+1)
	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += seq(i)
		sums[i] = sum
	}
	return sums
}

func CauchyCond(seq Sequence, n int, tol float64) bool {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := 0.0
			for k := i; k <= j; k++ {
				sum += seq(k)
			}
			if absL(sum) > tol {
				return false
			}
		}
	}
	return true
}

func AbelTest(a, b Sequence, n int) bool {
	if !SequenceIsMonotone(b, n) {
		return false
	}
	if absL(SequenceLimit(b, n)) > 1e3 {
		return false
	}
	if !CauchyCond(a, n, 1e-3) {
		return false
	}
	return true
}

func DirichletTest(a, b Sequence, n int) bool {
	if !SequenceIsMonotone(b, n) {
		return false
	}
	if !SequenceTendsToZero(b, n) {
		return false
	}
	if !CauchyCond(a, n, 1e-3) {
		return false
	}
	return true
}

func SequenceIsMonotone(seq Sequence, n int) bool {
	inc := true
	dec := true
	prev := seq(0)
	for i := 1; i <= n; i++ {
		curr := seq(i)
		if curr < prev {
			inc = false
		}
		if curr > prev {
			dec = false
		}
		prev = curr
	}
	return inc || dec
}

func SequenceTendsToZero(seq Sequence, n int) bool {
	prev := absL(seq(0))
	for i := 1; i <= n; i++ {
		curr := absL(seq(i))
		if curr > prev && curr > 1e-3 {
			return false
		}
		prev = curr
	}
	return prev < 1e-6
}

func SeriesAccelerationShanks(s0, s1, s2 float64) float64 {
	num := s2*s0 - s1*s1
	den := s2 - 2*s1 + s0
	if den == 0 {
		return s2
	}
	return num / den
}

func AitkenDeltaSquared(seq Sequence, n int) float64 {
	s0 := seq(n)
	s1 := seq(n + 1)
	s2 := seq(n + 2)
	return SeriesAccelerationShanks(s0, s1, s2)
}

func CesaroMean(seq Sequence, n int) float64 {
	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += seq(i)
	}
	return sum / float64(n+1)
}

func SeriesConvolution(a, b Sequence, n int) float64 {
	sum := 0.0
	for k := 0; k <= n; k++ {
		sum += a(k) * b(n-k)
	}
	return sum
}

func powerS(base, exp float64) float64 {
	return math.Pow(base, exp)
}

func factorialS(n int) float64 {
	if n <= 1 {
		return 1
	}
	res := 1.0
	for i := 2; i <= n; i++ {
		res *= float64(i)
	}
	return res
}

func SeriesFromCoeffs(coeffs []float64) Sequence {
	return func(n int) float64 {
		if n < 0 || n >= len(coeffs) {
			return 0
		}
		return coeffs[n]
	}
}

func BinomialCoefficients(n int) []float64 {
	coeffs := make([]float64, n+1)
	coeffs[0] = 1
	for k := 1; k <= n; k++ {
		coeffs[k] = coeffs[k-1] * float64(n-k+1) / float64(k)
	}
	return coeffs
}

func SeriesEvaluate(coeffs []float64, x float64, n int) float64 {
	sum := 0.0
	pow := 1.0
	limit := n
	if limit > len(coeffs)-1 {
		limit = len(coeffs) - 1
	}
	for i := 0; i <= limit; i++ {
		sum += coeffs[i] * pow
		pow *= x
	}
	return sum
}

func SeriesShift(coeffs []float64, shift float64) []float64 {
	out := make([]float64, len(coeffs))
	for k := 0; k < len(coeffs); k++ {
		for j := k; j < len(coeffs); j++ {
			out[j] += coeffs[k] * BinomialCoefficients(j)[k] * powerS(shift, float64(j-k))
		}
	}
	return out
}

func AbelTransform(a, b Sequence, n int) float64 {
	sum := 0.0
	partial := 0.0
	for k := 0; k <= n; k++ {
		partial += a(k)
		sum += partial * (b(k) - b(k+1))
	}
	return sum
}

func CesaroSequence(seq Sequence, n int) Sequence {
	return func(k int) float64 {
		sum := 0.0
		for i := 0; i <= k; i++ {
			sum += seq(i)
		}
		return sum / float64(k+1)
	}
}

func SeriesDifference(seq Sequence, n int) []float64 {
	vals := make([]float64, n)
	for i := 0; i < n; i++ {
		vals[i] = seq(i+1) - seq(i)
	}
	return vals
}
