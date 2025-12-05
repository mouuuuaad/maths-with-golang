package sequences

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

func IsStrictlyIncreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) <= seq(i) {
			return false
		}
	}
	return true
}

func IsStrictlyDecreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) >= seq(i) {
			return false
		}
	}
	return true
}

func EventuallyMonotone(seq Sequence, n, start int) bool {
	subseq := func(k int) float64 {
		return seq(k + start)
	}
	return IsMonotone(subseq, n-start)
}

func MonotoneConvergenceTheorem(seq Sequence, n int) (float64, bool) {
	if !IsMonotone(seq, n) {
		return 0, false
	}
	bounded, _, _ := IsBounded(seq, n)
	if !bounded {
		return 0, false
	}
	return IsConvergent(seq)
}
