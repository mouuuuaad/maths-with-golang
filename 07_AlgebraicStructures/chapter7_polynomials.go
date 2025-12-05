package algebra

type Polynomial struct {
	Coeffs []float64
}

func NewPolynomial(coeffs ...float64) Polynomial {
	return Polynomial{Coeffs: coeffs}
}

func (p Polynomial) Degree() int {
	for i := len(p.Coeffs) - 1; i >= 0; i-- {
		if p.Coeffs[i] != 0 {
			return i
		}
	}
	return 0
}

func (p Polynomial) Evaluate(x float64) float64 {
	result := 0.0
	xPow := 1.0
	for _, c := range p.Coeffs {
		result += c * xPow
		xPow *= x
	}
	return result
}

func (p Polynomial) Add(q Polynomial) Polynomial {
	maxLen := len(p.Coeffs)
	if len(q.Coeffs) > maxLen {
		maxLen = len(q.Coeffs)
	}
	result := make([]float64, maxLen)
	for i := range result {
		if i < len(p.Coeffs) {
			result[i] += p.Coeffs[i]
		}
		if i < len(q.Coeffs) {
			result[i] += q.Coeffs[i]
		}
	}
	return Polynomial{Coeffs: result}
}

func (p Polynomial) Multiply(q Polynomial) Polynomial {
	if len(p.Coeffs) == 0 || len(q.Coeffs) == 0 {
		return Polynomial{Coeffs: []float64{0}}
	}
	result := make([]float64, len(p.Coeffs)+len(q.Coeffs)-1)
	for i, pc := range p.Coeffs {
		for j, qc := range q.Coeffs {
			result[i+j] += pc * qc
		}
	}
	return Polynomial{Coeffs: result}
}

func (p Polynomial) Derivative() Polynomial {
	if len(p.Coeffs) <= 1 {
		return Polynomial{Coeffs: []float64{0}}
	}
	result := make([]float64, len(p.Coeffs)-1)
	for i := 1; i < len(p.Coeffs); i++ {
		result[i-1] = p.Coeffs[i] * float64(i)
	}
	return Polynomial{Coeffs: result}
}

func (p Polynomial) Integral() Polynomial {
	result := make([]float64, len(p.Coeffs)+1)
	result[0] = 0
	for i, c := range p.Coeffs {
		result[i+1] = c / float64(i+1)
	}
	return Polynomial{Coeffs: result}
}

func (p Polynomial) Scale(c float64) Polynomial {
	result := make([]float64, len(p.Coeffs))
	for i, coeff := range p.Coeffs {
		result[i] = coeff * c
	}
	return Polynomial{Coeffs: result}
}

func LagrangeInterpolation(x, y []float64) Polynomial {
	n := len(x)
	result := Polynomial{Coeffs: []float64{0}}
	for i := 0; i < n; i++ {
		li := Polynomial{Coeffs: []float64{1}}
		for j := 0; j < n; j++ {
			if i != j {
				term := Polynomial{Coeffs: []float64{-x[j] / (x[i] - x[j]), 1 / (x[i] - x[j])}}
				li = li.Multiply(term)
			}
		}
		result = result.Add(li.Scale(y[i]))
	}
	return result
}
