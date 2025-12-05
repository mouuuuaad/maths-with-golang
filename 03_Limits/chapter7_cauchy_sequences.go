package limits

func CauchyCriterion(seq Sequence, n int) bool {
	return CauchySequence(seq, n, 1e-7)
}

func SubsequenceLimit(seq Sequence, indices func(int) int) (float64, bool) {
	subseq := func(n int) float64 {
		return seq(indices(n))
	}
	return LimitSequence(subseq)
}

func ConvergentSubsequence(seq Sequence, n int) func(int) float64 {
	return func(k int) float64 {
		return seq(k * 2)
	}
}

func NestedIntervals(lower, upper Sequence, n int) (float64, bool) {
	for i := 1; i <= n; i++ {
		l := lower(i)
		u := upper(i)
		if l > u {
			return 0, false
		}
	}
	lLim, lOk := LimitSequence(lower)
	uLim, uOk := LimitSequence(upper)
	if lOk && uOk && absLim(lLim-uLim) < 1e-7 {
		return lLim, true
	}
	return 0, false
}

func DedekindsConstruction(lower, upper []float64) float64 {
	if len(lower) == 0 || len(upper) == 0 {
		return 0
	}
	maxLower := lower[0]
	for _, v := range lower {
		if v > maxLower {
			maxLower = v
		}
	}
	minUpper := upper[0]
	for _, v := range upper {
		if v < minUpper {
			minUpper = v
		}
	}
	return (maxLower + minUpper) / 2
}
