package sequences

type Sequence func(int) float64

func IsConvergent(seq Sequence) (float64, bool) {
	epsilon := 1e-7
	prev := seq(1000)
	for _, n := range []int{10000, 100000, 1000000} {
		curr := seq(n)
		if absS(curr-prev) < epsilon {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func Limit(seq Sequence) float64 {
	lim, _ := IsConvergent(seq)
	return lim
}

func AitkenAcceleration(seq Sequence) Sequence {
	return func(n int) float64 {
		sn := seq(n)
		sn1 := seq(n + 1)
		sn2 := seq(n + 2)
		denom := sn2 - 2*sn1 + sn
		if absS(denom) < 1e-15 {
			return sn2
		}
		return sn - (sn1-sn)*(sn1-sn)/denom
	}
}

func RichardsonExtrapolation(f func(float64) float64, h float64, k int) float64 {
	factor := powerS(2, float64(k))
	return (factor*f(h/2) - f(h)) / (factor - 1)
}

func CesaroMean(seq Sequence, n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += seq(i)
	}
	return sum / float64(n)
}

func RatioTest(seq Sequence, n int) string {
	an := seq(n)
	an1 := seq(n + 1)
	if absS(an) < 1e-15 {
		return "inconclusive"
	}
	ratio := absS(an1 / an)
	if ratio < 0.99 {
		return "converges"
	}
	if ratio > 1.01 {
		return "diverges"
	}
	return "inconclusive"
}

func RootTest(seq Sequence, n int) string {
	an := absS(seq(n))
	if an == 0 {
		return "inconclusive"
	}
	root := powerS(an, 1.0/float64(n))
	if root < 0.99 {
		return "converges"
	}
	if root > 1.01 {
		return "diverges"
	}
	return "inconclusive"
}
