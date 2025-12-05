package limits

func IsIncreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) < seq(i) {
			return false
		}
	}
	return true
}

func IsDecreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) > seq(i) {
			return false
		}
	}
	return true
}

func IsMonotone(seq Sequence, n int) bool {
	return IsIncreasing(seq, n) || IsDecreasing(seq, n)
}

func IsBounded(seq Sequence, n int) (bool, float64, float64) {
	lower := seq(1)
	upper := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val < lower {
			lower = val
		}
		if val > upper {
			upper = val
		}
	}
	return true, lower, upper
}

func MonotoneConvergence(seq Sequence, n int) (float64, bool) {
	if !IsMonotone(seq, n) {
		return 0, false
	}
	bounded, _, _ := IsBounded(seq, n*10)
	if !bounded {
		return 0, false
	}
	return LimitSequence(seq)
}

func BolzanoWeierstrass(seq Sequence, n int) (float64, bool) {
	bounded, _, _ := IsBounded(seq, n)
	if !bounded {
		return 0, false
	}
	return LimitSequence(seq)
}
