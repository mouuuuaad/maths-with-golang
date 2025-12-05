package calculus

func Derivative(f Function, x float64) float64 {
	h := 1e-7
	return (f(x+h) - f(x-h)) / (2 * h)
}

func DerivativeForward(f Function, x float64) float64 {
	h := 1e-7
	return (f(x+h) - f(x)) / h
}

func DerivativeBackward(f Function, x float64) float64 {
	h := 1e-7
	return (f(x) - f(x-h)) / h
}

func SecondDerivative(f Function, x float64) float64 {
	h := 1e-5
	return (f(x+h) - 2*f(x) + f(x-h)) / (h * h)
}

func NthDerivative(f Function, x float64, n int) float64 {
	if n == 0 {
		return f(x)
	}
	if n == 1 {
		return Derivative(f, x)
	}
	df := func(t float64) float64 {
		return Derivative(f, t)
	}
	return NthDerivative(df, x, n-1)
}

func TangentLine(f Function, a float64) Function {
	fa := f(a)
	dfa := Derivative(f, a)
	return func(x float64) float64 {
		return fa + dfa*(x-a)
	}
}

func LinearApproximation(f Function, a, x float64) float64 {
	return f(a) + Derivative(f, a)*(x-a)
}

func QuadraticApproximation(f Function, a, x float64) float64 {
	h := x - a
	return f(a) + Derivative(f, a)*h + SecondDerivative(f, a)*h*h/2
}

func IsDifferentiableAt(f Function, a float64) bool {
	left := DerivativeBackward(f, a)
	right := DerivativeForward(f, a)
	return absL(left-right) < 1e-5
}

func CriticalPoints(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	for x := a; x <= b; x += step {
		d := Derivative(f, x)
		if absL(d) < 1e-6 {
			result = append(result, x)
		}
	}
	return result
}

func InflectionPoints(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	prev := SecondDerivative(f, a)
	for x := a + step; x <= b; x += step {
		curr := SecondDerivative(f, x)
		if prev*curr < 0 {
			result = append(result, x)
		}
		prev = curr
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
// Created by: MOUAAD IDOUFKIR
// << The universe runs on equations. We just translate them >>
