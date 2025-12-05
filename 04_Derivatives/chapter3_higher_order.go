package derivatives

func ThirdDerivative(f Function, x float64) float64 {
	h := 1e-4
	return (f(x+2*h) - 2*f(x+h) + 2*f(x-h) - f(x-2*h)) / (2 * h * h * h)
}

func FourthDerivative(f Function, x float64) float64 {
	h := 1e-3
	return (f(x+2*h) - 4*f(x+h) + 6*f(x) - 4*f(x-h) + f(x-2*h)) / (h * h * h * h)
}

func GeneralNthDerivative(f Function, x float64, n int, h float64) float64 {
	if n == 0 {
		return f(x)
	}
	coeffs := binomialCoeffs(n)
	sum := 0.0
	sign := 1.0
	if n%2 == 1 {
		sign = -1.0
	}
	for k := 0; k <= n; k++ {
		sum += sign * coeffs[k] * f(x+float64(n-2*k)*h/2)
		sign = -sign
	}
	return sum / powerD(h, float64(n))
}

func binomialCoeffs(n int) []float64 {
	coeffs := make([]float64, n+1)
	coeffs[0] = 1
	for i := 1; i <= n; i++ {
		coeffs[i] = coeffs[i-1] * float64(n-i+1) / float64(i)
	}
	return coeffs
}

func CurvatureAtPoint(f Function, x float64) float64 {
	df := Derivative(f, x)
	d2f := SecondDerivative(f, x)
	denom := powerD(1+df*df, 1.5)
	if absD(denom) < 1e-12 {
		return 0
	}
	return absD(d2f) / denom
}

func RadiusOfCurvature(f Function, x float64) float64 {
	k := CurvatureAtPoint(f, x)
	if absD(k) < 1e-12 {
		return 1e308
	}
	return 1.0 / k
}

func Concavity(f Function, x float64) string {
	d2f := SecondDerivative(f, x)
	if d2f > 1e-9 {
		return "concave up"
	}
	if d2f < -1e-9 {
		return "concave down"
	}
	return "inflection"
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
// Created by: MOUAAD IDOUFKIR
// << The universe runs on equations. We just translate them >>
