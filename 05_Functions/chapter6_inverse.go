package functions

func Sqrt(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x == 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		next := 0.5 * (z + x/z)
		if Abs(next-z) < 1e-15 {
			return next
		}
		z = next
	}
	return z
}

func Atan(x float64) float64 {
	if x < 0 {
		return -Atan(-x)
	}
	if x > 1 {
		return Pi/2 - Atan(1/x)
	}
	sum := 0.0
	term := x
	x2 := x * x
	sign := 1.0
	for i := 0; i < 200; i++ {
		sum += sign * term / float64(2*i+1)
		term *= x2
		sign = -sign
		if Abs(term) < 1e-15 {
			break
		}
	}
	return sum
}

func Asin(x float64) float64 {
	if x < -1 || x > 1 {
		return 0
	}
	if Abs(x) == 1 {
		return x * Pi / 2
	}
	return Atan(x / Sqrt(1-x*x))
}

func Acos(x float64) float64 {
	return Pi/2 - Asin(x)
}

func Atan2(y, x float64) float64 {
	if x > 0 {
		return Atan(y / x)
	}
	if x < 0 {
		if y >= 0 {
			return Atan(y/x) + Pi
		}
		return Atan(y/x) - Pi
	}
	if y > 0 {
		return Pi / 2
	}
	if y < 0 {
		return -Pi / 2
	}
	return 0
}

func Asinh(x float64) float64 {
	return Ln(x + Sqrt(x*x+1))
}

func Acosh(x float64) float64 {
	if x < 1 {
		return 0
	}
	return Ln(x + Sqrt(x*x-1))
}

func Atanh(x float64) float64 {
	if Abs(x) >= 1 {
		return 0
	}
	return 0.5 * Ln((1+x)/(1-x))
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
