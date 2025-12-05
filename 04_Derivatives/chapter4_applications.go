package derivatives

func TangentLine(f Function, a float64) Function {
	fa := f(a)
	dfa := Derivative(f, a)
	return func(x float64) float64 {
		return fa + dfa*(x-a)
	}
}

func NormalLine(f Function, a float64) Function {
	fa := f(a)
	dfa := Derivative(f, a)
	if absD(dfa) < 1e-12 {
		return func(x float64) float64 { return fa }
	}
	return func(x float64) float64 {
		return fa - (x-a)/dfa
	}
}

func LinearApproximation(f Function, a, x float64) float64 {
	return f(a) + Derivative(f, a)*(x-a)
}

func QuadraticApproximation(f Function, a, x float64) float64 {
	h := x - a
	return f(a) + Derivative(f, a)*h + SecondDerivative(f, a)*h*h/2
}

func Differential(f Function, x, dx float64) float64 {
	return Derivative(f, x) * dx
}

func TangentSlope(f Function, a float64) float64 {
	return Derivative(f, a)
}

func TangentIntercept(f Function, a float64) float64 {
	return f(a) - Derivative(f, a)*a
}

func AngleBetweenCurves(f, g Function, x float64) float64 {
	df := Derivative(f, x)
	dg := Derivative(g, x)
	return absD(atanD(df) - atanD(dg))
}

func atanD(x float64) float64 {
	pi := 3.14159265358979323846
	if x < 0 {
		return -atanD(-x)
	}
	if x > 1 {
		return pi/2 - atanD(1/x)
	}
	sum := 0.0
	term := x
	x2 := x * x
	sign := 1.0
	for i := 0; i < 50; i++ {
		sum += sign * term / float64(2*i+1)
		term *= x2
		sign = -sign
	}
	return sum
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
