package sequences

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

func IsBoundedAbove(seq Sequence, n int) (bool, float64) {
	upper := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val > upper {
			upper = val
		}
	}
	return true, upper
}

func IsBoundedBelow(seq Sequence, n int) (bool, float64) {
	lower := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val < lower {
			lower = val
		}
	}
	return true, lower
}

func Supremum(seq Sequence, n int) float64 {
	_, upper := IsBoundedAbove(seq, n)
	return upper
}

func Infimum(seq Sequence, n int) float64 {
	_, lower := IsBoundedBelow(seq, n)
	return lower
}

func LimSup(seq Sequence, n int) float64 {
	maxVals := make([]float64, n)
	for i := 1; i <= n; i++ {
		max := seq(i)
		for j := i; j <= n; j++ {
			if seq(j) > max {
				max = seq(j)
			}
		}
		maxVals[i-1] = max
	}
	return maxVals[n-1]
}

func LimInf(seq Sequence, n int) float64 {
	minVals := make([]float64, n)
	for i := 1; i <= n; i++ {
		min := seq(i)
		for j := i; j <= n; j++ {
			if seq(j) < min {
				min = seq(j)
			}
		}
		minVals[i-1] = min
	}
	return minVals[n-1]
}
