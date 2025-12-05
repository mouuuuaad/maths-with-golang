package sequences

type ArithmeticSequence struct {
	Start float64
	Diff  float64
}

func (s ArithmeticSequence) NthTerm(n int) float64 {
	return s.Start + float64(n)*s.Diff
}

func (s ArithmeticSequence) SumN(n int) float64 {
	return float64(n+1) * (s.Start + s.NthTerm(n)) / 2
}

func (s ArithmeticSequence) Generate(n int) []float64 {
	result := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = s.NthTerm(i)
	}
	return result
}

func (s ArithmeticSequence) FindN(term float64) int {
	if s.Diff == 0 {
		if term == s.Start {
			return 0
		}
		return -1
	}
	n := (term - s.Start) / s.Diff
	if n < 0 || n != float64(int(n)) {
		return -1
	}
	return int(n)
}

func ArithmeticMean(a, b float64) float64 {
	return (a + b) / 2
}

func InsertArithmeticMeans(a, b float64, n int) []float64 {
	d := (b - a) / float64(n+1)
	result := make([]float64, n+2)
	result[0] = a
	for i := 1; i <= n; i++ {
		result[i] = a + float64(i)*d
	}
	result[n+1] = b
	return result
}
