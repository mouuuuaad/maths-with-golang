package arithmetic

type Rational struct {
	Num Integer
	Den Integer
}

func NewRational(num, den Integer) Rational {
	if den == 0 {
		return Rational{0, 1}
	}
	if den < 0 {
		num = -num
		den = -den
	}
	g := GCD(AbsI(num), AbsI(den))
	return Rational{num / g, den / g}
}

func (r Rational) Simplify() Rational {
	return NewRational(r.Num, r.Den)
}

func (r Rational) Add(other Rational) Rational {
	num := r.Num*other.Den + other.Num*r.Den
	den := r.Den * other.Den
	return NewRational(num, den)
}

func (r Rational) Subtract(other Rational) Rational {
	num := r.Num*other.Den - other.Num*r.Den
	den := r.Den * other.Den
	return NewRational(num, den)
}

func (r Rational) Multiply(other Rational) Rational {
	return NewRational(r.Num*other.Num, r.Den*other.Den)
}

func (r Rational) Divide(other Rational) Rational {
	if other.Num == 0 {
		return Rational{0, 1}
	}
	return NewRational(r.Num*other.Den, r.Den*other.Num)
}

func (r Rational) Inverse() Rational {
	if r.Num == 0 {
		return Rational{0, 1}
	}
	return NewRational(r.Den, r.Num)
}

func (r Rational) ToFloat64() float64 {
	return float64(r.Num) / float64(r.Den)
}

func (r Rational) Negate() Rational {
	return Rational{-r.Num, r.Den}
}

func (r Rational) Abs() Rational {
	return Rational{AbsI(r.Num), r.Den}
}

func (r Rational) Compare(other Rational) int {
	diff := r.Num*other.Den - other.Num*r.Den
	return SignI(diff)
}

func (r Rational) IsZero() bool {
	return r.Num == 0
}

func (r Rational) IsPositive() bool {
	return r.Num > 0
}

func (r Rational) IsNegative() bool {
	return r.Num < 0
}

func ContinuedFraction(r Rational, maxTerms int) []Integer {
	result := []Integer{}
	for i := 0; i < maxTerms && r.Den != 0; i++ {
		q := r.Num / r.Den
		result = append(result, q)
		r = Rational{r.Den, r.Num - q*r.Den}
	}
	return result
}

func FromContinuedFraction(cf []Integer) Rational {
	if len(cf) == 0 {
		return Rational{0, 1}
	}
	r := Rational{cf[len(cf)-1], 1}
	for i := len(cf) - 2; i >= 0; i-- {
		r = Rational{cf[i], 1}.Add(r.Inverse())
	}
	return r
}
