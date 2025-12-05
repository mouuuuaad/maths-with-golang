package functions

func Log10(x float64) float64 {
	return Ln(x) / Ln(10)
}

func Log2(x float64) float64 {
	return Ln(x) / Ln(2)
}

func LogBase(x, base float64) float64 {
	return Ln(x) / Ln(base)
}

func Log10Fast(x float64) float64 {
	ln10 := 2.302585092994046
	return Ln(x) / ln10
}

func Log2Fast(x float64) float64 {
	ln2 := 0.6931471805599453
	return Ln(x) / ln2
}
