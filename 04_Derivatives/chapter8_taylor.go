package derivatives

func TaylorPolynomial(f Function, a float64, n int) func(float64) float64 {
	coeffs := make([]float64, n+1)
	factorial := 1.0
	for i := 0; i <= n; i++ {
		if i > 0 {
			factorial *= float64(i)
		}
		coeffs[i] = NthDerivative(f, a, i) / factorial
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

func MaclaurinPolynomial(f Function, n int) func(float64) float64 {
	return TaylorPolynomial(f, 0, n)
}

func TaylorError(f Function, a, x float64, n int) float64 {
	h := x - a
	factorial := 1.0
	for i := 1; i <= n+1; i++ {
		factorial *= float64(i)
	}
	M := 0.0
	step := absD(h) / 10
	for t := a; t <= x || t >= x; {
		val := absD(NthDerivative(f, t, n+1))
		if val > M {
			M = val
		}
		if x > a {
			t += step
			if t > x {
				break
			}
		} else {
			t -= step
			if t < x {
				break
			}
		}
	}
	return M * powerD(absD(h), float64(n+1)) / factorial
}

func TaylorSeries(f Function, a, x float64, n int) float64 {
	return TaylorPolynomial(f, a, n)(x)
}

func MaclaurinSeries(f Function, x float64, n int) float64 {
	return TaylorSeries(f, 0, x, n)
}

func PadePoly(f Function, m, n int) ([]float64, []float64) {
	numCoeffs := make([]float64, m+1)
	denCoeffs := make([]float64, n+1)
	numCoeffs[0] = f(0)
	denCoeffs[0] = 1
	return numCoeffs, denCoeffs
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
