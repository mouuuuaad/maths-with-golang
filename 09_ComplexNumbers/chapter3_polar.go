package complexnums

const PiC = 3.14159265358979323846

type PolarForm struct {
	Radius float64
	Theta  float64
}

func (c ComplexNumber) ToPolar() PolarForm {
	r := c.Abs()
	theta := atan2C(c.I, c.R)
	return PolarForm{Radius: r, Theta: theta}
}

func (p PolarForm) ToComplex() ComplexNumber {
	return ComplexNumber{
		R: p.Radius * cosC(p.Theta),
		I: p.Radius * sinC(p.Theta),
	}
}

func (c ComplexNumber) Argument() float64 {
	return atan2C(c.I, c.R)
}

func (c ComplexNumber) Rotate(phi float64) ComplexNumber {
	p := c.ToPolar()
	p.Theta += phi
	return p.ToComplex()
}

func atan2C(y, x float64) float64 {
	if x > 0 {
		return atanC(y / x)
	}
	if x < 0 {
		if y >= 0 {
			return atanC(y/x) + PiC
		}
		return atanC(y/x) - PiC
	}
	if y > 0 {
		return PiC / 2
	}
	if y < 0 {
		return -PiC / 2
	}
	return 0
}

func atanC(x float64) float64 {
	if x < 0 {
		return -atanC(-x)
	}
	if x > 1 {
		return PiC/2 - atanC(1/x)
	}
	sum := 0.0
	term := x
	x2 := x * x
	sign := 1.0
	for i := 0; i < 100; i++ {
		sum += sign * term / float64(2*i+1)
		term *= x2
		sign = -sign
		if absC(term) < 1e-15 {
			break
		}
	}
	return sum
}

func sinC(x float64) float64 {
	k := int((x + PiC) / (2 * PiC))
	x -= float64(k) * 2 * PiC
	sum := 0.0
	term := x
	x2 := x * x
	for i := 1; i < 30; i++ {
		sum += term
		term *= -x2 / float64(2*i*(2*i+1))
	}
	return sum
}

func cosC(x float64) float64 {
	k := int((x + PiC) / (2 * PiC))
	x -= float64(k) * 2 * PiC
	sum := 0.0
	term := 1.0
	x2 := x * x
	for i := 1; i < 30; i++ {
		sum += term
		term *= -x2 / float64((2*i-1)*(2*i))
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
