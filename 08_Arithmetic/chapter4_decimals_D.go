package arithmetic

type Decimal struct {
	Mantissa int64
	Exponent int
}

func NewDecimal(mantissa int64, exponent int) Decimal {
	return Decimal{Mantissa: mantissa, Exponent: exponent}
}

func NewDecimalFromFloat(f float64, precision int) Decimal {
	scale := int64(1)
	for i := 0; i < precision; i++ {
		scale *= 10
	}
	mantissa := int64(f * float64(scale))
	return Decimal{Mantissa: mantissa, Exponent: -precision}
}

func (d Decimal) normalize() Decimal {
	for d.Mantissa != 0 && d.Mantissa%10 == 0 {
		d.Mantissa /= 10
		d.Exponent++
	}
	return d
}

func (d Decimal) alignExponent(other Decimal) (Decimal, Decimal) {
	if d.Exponent == other.Exponent {
		return d, other
	}
	if d.Exponent > other.Exponent {
		diff := d.Exponent - other.Exponent
		for i := 0; i < diff; i++ {
			d.Mantissa *= 10
		}
		d.Exponent = other.Exponent
	} else {
		diff := other.Exponent - d.Exponent
		for i := 0; i < diff; i++ {
			other.Mantissa *= 10
		}
		other.Exponent = d.Exponent
	}
	return d, other
}

func (d Decimal) Add(other Decimal) Decimal {
	d, other = d.alignExponent(other)
	return Decimal{Mantissa: d.Mantissa + other.Mantissa, Exponent: d.Exponent}.normalize()
}

func (d Decimal) Subtract(other Decimal) Decimal {
	d, other = d.alignExponent(other)
	return Decimal{Mantissa: d.Mantissa - other.Mantissa, Exponent: d.Exponent}.normalize()
}

func (d Decimal) Multiply(other Decimal) Decimal {
	return Decimal{
		Mantissa: d.Mantissa * other.Mantissa,
		Exponent: d.Exponent + other.Exponent,
	}.normalize()
}

func (d Decimal) Divide(other Decimal, precision int) Decimal {
	if other.Mantissa == 0 {
		return Decimal{0, 0}
	}
	scale := int64(1)
	for i := 0; i < precision; i++ {
		scale *= 10
	}
	mantissa := (d.Mantissa * scale) / other.Mantissa
	return Decimal{
		Mantissa: mantissa,
		Exponent: d.Exponent - other.Exponent - precision,
	}.normalize()
}

func (d Decimal) ToFloat64() float64 {
	result := float64(d.Mantissa)
	if d.Exponent > 0 {
		for i := 0; i < d.Exponent; i++ {
			result *= 10
		}
	} else {
		for i := 0; i < -d.Exponent; i++ {
			result /= 10
		}
	}
	return result
}

func (d Decimal) Compare(other Decimal) int {
	d, other = d.alignExponent(other)
	if d.Mantissa > other.Mantissa {
		return 1
	}
	if d.Mantissa < other.Mantissa {
		return -1
	}
	return 0
}

func (d Decimal) Abs() Decimal {
	if d.Mantissa < 0 {
		d.Mantissa = -d.Mantissa
	}
	return d
}

func (d Decimal) Negate() Decimal {
	return Decimal{-d.Mantissa, d.Exponent}
}
