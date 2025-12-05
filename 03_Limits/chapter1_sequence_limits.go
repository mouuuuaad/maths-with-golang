package limits

type Sequence func(int) float64

func absLim(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func LimitSequence(seq Sequence) (float64, bool) {
	epsilon := 1e-9
	prev := seq(1000)
	for _, n := range []int{10000, 100000, 1000000} {
		curr := seq(n)
		if absLim(curr-prev) < epsilon {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func IsCauchy(seq Sequence) bool {
	epsilon := 1e-7
	n := 10000
	m := 20000
	return absLim(seq(n)-seq(m)) < epsilon
}

func SeqLimit(seq Sequence, epsilon float64, maxN int) (float64, bool) {
	prev := seq(100)
	for n := 1000; n <= maxN; n *= 10 {
		curr := seq(n)
		if absLim(curr-prev) < epsilon {
			return curr, true
		}
		prev = curr
	}
	return prev, false
}

func SequenceSupremum(seq Sequence, n int) float64 {
	sup := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val > sup {
			sup = val
		}
	}
	return sup
}

func SequenceInfimum(seq Sequence, n int) float64 {
	inf := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val < inf {
			inf = val
		}
	}
	return inf
}

func LimSup(seq Sequence, n int) float64 {
	maxVals := []float64{}
	for i := 1; i <= n; i++ {
		max := seq(i)
		for j := i; j <= n; j++ {
			if seq(j) > max {
				max = seq(j)
			}
		}
		maxVals = append(maxVals, max)
	}
	return maxVals[len(maxVals)-1]
}

func LimInf(seq Sequence, n int) float64 {
	minVals := []float64{}
	for i := 1; i <= n; i++ {
		min := seq(i)
		for j := i; j <= n; j++ {
			if seq(j) < min {
				min = seq(j)
			}
		}
		minVals = append(minVals, min)
	}
	return minVals[len(minVals)-1]
}
