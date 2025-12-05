package calculus

type Sequence func(int) float64

func GeometricSeriesSum(a, r float64) float64 {
	if absL(r) >= 1 {
		return 1e308
	}
	return a / (1 - r)
}

func PartialSum(seq Sequence, n int) float64 {
	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += seq(i)
	}
	return sum
}

func RatioTest(seq Sequence, n int) string {
	an := seq(n)
	an1 := seq(n + 1)
	if absL(an) < 1e-15 {
		return "Inconclusive"
	}
	ratio := absL(an1 / an)
	if ratio < 0.99 {
		return "Converges"
	}
	if ratio > 1.01 {
		return "Diverges"
	}
	return "Inconclusive"
}

func RootTest(seq Sequence, n int) string {
	an := absL(seq(n))
	if an == 0 {
		return "Inconclusive"
	}
	root := powerS(an, 1.0/float64(n))
	if root < 0.99 {
		return "Converges"
	}
	if root > 1.01 {
		return "Diverges"
	}
	return "Inconclusive"
}

func IntegralTest(f Function, a float64, n int) float64 {
	return TrapezoidalRule(f, a, float64(n), n*10)
}

func ComparisonTest(seq1, seq2 Sequence, n int) bool {
	for i := 1; i <= n; i++ {
		if absL(seq1(i)) > absL(seq2(i)) {
			return false
		}
	}
	return true
}

func LimitComparisonTest(seq1, seq2 Sequence, n int) float64 {
	a := seq1(n)
	b := seq2(n)
	if absL(b) < 1e-15 {
		return 0
	}
	return a / b
}

func AlternatingSeries(seq Sequence, n int) float64 {
	sum := 0.0
	sign := 1.0
	for i := 0; i <= n; i++ {
		sum += sign * absL(seq(i))
		sign = -sign
	}
	return sum
}

func TaylorSeries(f Function, a float64, n int) func(float64) float64 {
	coeffs := make([]float64, n+1)
	coeffs[0] = f(a)
	df := f
	for i := 1; i <= n; i++ {
		dfNew := func(g Function) Function {
			return func(x float64) float64 {
				return Derivative(g, x)
			}
		}(df)
		df = dfNew
		coeffs[i] = df(a) / factorialS(i)
	}
	return func(x float64) float64 {
		sum := 0.0
		h := 1.0
		for i := 0; i <= n; i++ {
			sum += coeffs[i] * h
			h *= (x - a)
		}
		return sum
	}
}

func powerS(base, exp float64) float64 {
	if base <= 0 {
		return 0
	}
	return expS(exp * lnS(base))
}

func expS(x float64) float64 {
	if x < 0 {
		return 1.0 / expS(-x)
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 50; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func lnS(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	for i := 0; i < 50; i++ {
		ey := expS(y)
		diff := ey - x
		if absL(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
	}
	return y
}

func factorialS(n int) float64 {
	if n <= 1 {
		return 1
	}
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return result
}

//MMMMMMMM               MMMMMMMM     OOOOOOOOO     UUUUUUUU     UUUUUUUU           AAA                              AAA               DDDDDDDDDDDDD        
//M:::::::M             M:::::::M   OO:::::::::OO   U::::::U     U::::::U          A:::A                            A:::A              D::::::::::::DDD     
//M::::::::M           M::::::::M OO:::::::::::::OO U::::::U     U::::::U         A:::::A                          A:::::A             D:::::::::::::::DD   
//M:::::::::M         M:::::::::MO:::::::OOO:::::::OUU:::::U     U:::::UU        A:::::::A                        A:::::::A            DDD:::::DDDDD:::::D  
//M::::::::::M       M::::::::::MO::::::O   O::::::O U:::::U     U:::::U        A:::::::::A                      A:::::::::A             D:::::D    D:::::D 
//M:::::::::::M     M:::::::::::MO:::::O     O:::::O U:::::D     D:::::U       A:::::A:::::A                    A:::::A:::::A            D:::::D     D:::::D
//M:::::::M::::M   M::::M:::::::MO:::::O     O:::::O U:::::D     D:::::U      A:::::A A:::::A                  A:::::A A:::::A           D:::::D     D:::::D
//M::::::M M::::M M::::M M::::::MO:::::O     O:::::O U:::::D     D:::::U     A:::::A   A:::::A                A:::::A   A:::::A          D:::::D     D:::::D
//M::::::M  M::::M::::M  M::::::MO:::::O     O:::::O U:::::D     D:::::U    A:::::A     A:::::A              A:::::A     A:::::A         D:::::D     D:::::D
//M::::::M   M:::::::M   M::::::MO:::::O     O:::::O U:::::D     D:::::U   A:::::AAAAAAAAA:::::A            A:::::AAAAAAAAA:::::A        D:::::D     D:::::D
//M::::::M    M:::::M    M::::::MO:::::O     O:::::O U:::::D     D:::::U  A:::::::::::::::::::::A          A:::::::::::::::::::::A       D:::::D     D:::::D
//M::::::M     MMMMM     M::::::MO::::::O   O::::::O U::::::U   U::::::U A:::::AAAAAAAAAAAAA:::::A        A:::::AAAAAAAAAAAAA:::::A      D:::::D    D:::::D 
//M::::::M               M::::::MO:::::::OOO:::::::O U:::::::UUU:::::::UA:::::A             A:::::A      A:::::A             A:::::A   DDD:::::DDDDD:::::D  
//M::::::M               M::::::M OO:::::::::::::OO   UU:::::::::::::UUA:::::A               A:::::A    A:::::A               A:::::A  D:::::::::::::::DD   
//M::::::M               M::::::M   OO:::::::::OO       UU:::::::::UU A:::::A                 A:::::A  A:::::A                 A:::::A D::::::::::::DDD     
//MMMMMMMM               MMMMMMMM     OOOOOOOOO           UUUUUUUUU  AAAAAAA                   AAAAAAAAAAAAAA                   AAAAAAADDDDDDDDDDDDD        
// Created by: MOUAAD
// MathsWithGolang - Pure Golang Mathematical Library
