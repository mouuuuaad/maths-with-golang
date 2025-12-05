package complexnums

func MobiusTransformation(z, a, b, c, d ComplexNumber) ComplexNumber {
	numerator := a.Multiply(z).Add(b)
	denominator := c.Multiply(z).Add(d)
	if denominator.Abs() == 0 {
		return New(1e308, 0)
	}
	return numerator.Divide(denominator)
}

func JoukowskyTransform(z ComplexNumber) ComplexNumber {
	if z.Abs() == 0 {
		return New(1e308, 0)
	}
	return z.Add(z.Inverse())
}

func Inversion(z ComplexNumber) ComplexNumber {
	if z.Abs() == 0 {
		return New(1e308, 0)
	}
	return z.Inverse()
}

func Translation(z, w ComplexNumber) ComplexNumber {
	return z.Add(w)
}

func Rotation(z ComplexNumber, angle float64) ComplexNumber {
	return z.Multiply(EulerFormula(angle))
}

func Scaling(z ComplexNumber, factor float64) ComplexNumber {
	return z.Scale(factor)
}

func Reflection(z ComplexNumber) ComplexNumber {
	return z.Conjugate()
}

func ZhukovskyAirfoil(z ComplexNumber, a float64) ComplexNumber {
	aComplex := New(a, 0)
	return z.Add(aComplex.Multiply(aComplex).Divide(z))
}

func SchwarzChristoffel(z ComplexNumber, vertices []ComplexNumber, angles []float64) ComplexNumber {
	if len(vertices) == 0 || len(angles) == 0 {
		return z
	}
	result := New(1, 0)
	for i := range vertices {
		diff := z.Subtract(vertices[i])
		result = result.Multiply(Pow(diff, New(angles[i]-1, 0)))
	}
	return result
}

func ConformalMap(z ComplexNumber, f func(ComplexNumber) ComplexNumber) ComplexNumber {
	return f(z)
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
