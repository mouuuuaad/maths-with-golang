package limits

func EpsilonDelta(f Function, a, L, epsilon float64) float64 {
	for delta := 0.1; delta > 1e-10; delta /= 10 {
		ok := true
		for x := a - delta; x <= a+delta; x += delta / 10 {
			if absLim(x-a) < 1e-12 {
				continue
			}
			if absLim(f(x)-L) >= epsilon {
				ok = false
				break
			}
		}
		if ok {
			return delta
		}
	}
	return 0
}

func LimitNM(seq Sequence, N, epsilon float64) bool {
	for n := int(N); n < int(N)+100; n++ {
		for m := n; m < n+100; m++ {
			if absLim(seq(n)-seq(m)) >= epsilon {
				return false
			}
		}
	}
	return true
}

func VerifyLimit(f Function, a, L, epsilon, delta float64) bool {
	for x := a - delta; x <= a+delta; x += delta / 100 {
		if absLim(x-a) < 1e-12 {
			continue
		}
		if absLim(f(x)-L) >= epsilon {
			return false
		}
	}
	return true
}

func CauchySequence(seq Sequence, n int, epsilon float64) bool {
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if absLim(seq(i)-seq(j)) >= epsilon {
				return false
			}
		}
	}
	return true
}

func CompletionR(seq Sequence, n int) float64 {
	if CauchySequence(seq, n, 1e-7) {
		lim, _ := LimitSequence(seq)
		return lim
	}
	return 0
}
