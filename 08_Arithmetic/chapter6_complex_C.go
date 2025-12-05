package arithmetic

type Complex struct {
	Re float64
	Im float64
}

func NewComplex(re, im float64) Complex {
	return Complex{Re: re, Im: im}
}

func (c Complex) RealPart() float64 {
	return c.Re
}

func (c Complex) ImagPart() float64 {
	return c.Im
}

func (c Complex) Add(other Complex) Complex {
	return Complex{c.Re + other.Re, c.Im + other.Im}
}

func (c Complex) Subtract(other Complex) Complex {
	return Complex{c.Re - other.Re, c.Im - other.Im}
}

func (c Complex) Multiply(other Complex) Complex {
	return Complex{
		c.Re*other.Re - c.Im*other.Im,
		c.Re*other.Im + c.Im*other.Re,
	}
}

func (c Complex) Divide(other Complex) Complex {
	denom := other.Re*other.Re + other.Im*other.Im
	if denom == 0 {
		return Complex{0, 0}
	}
	return Complex{
		(c.Re*other.Re + c.Im*other.Im) / denom,
		(c.Im*other.Re - c.Re*other.Im) / denom,
	}
}

func (c Complex) Conjugate() Complex {
	return Complex{c.Re, -c.Im}
}

func (c Complex) Modulus() float64 {
	return sqrtF(c.Re*c.Re + c.Im*c.Im)
}

func (c Complex) ModulusSquared() float64 {
	return c.Re*c.Re + c.Im*c.Im
}

func (c Complex) Argument() float64 {
	return atan2F(c.Im, c.Re)
}

func (c Complex) Inverse() Complex {
	denom := c.Re*c.Re + c.Im*c.Im
	if denom == 0 {
		return Complex{0, 0}
	}
	return Complex{c.Re / denom, -c.Im / denom}
}

func (c Complex) Power(n int) Complex {
	if n == 0 {
		return Complex{1, 0}
	}
	if n < 0 {
		return c.Inverse().Power(-n)
	}
	result := Complex{1, 0}
	base := c
	for n > 0 {
		if n%2 == 1 {
			result = result.Multiply(base)
		}
		base = base.Multiply(base)
		n /= 2
	}
	return result
}

func (c Complex) Scale(s float64) Complex {
	return Complex{c.Re * s, c.Im * s}
}

func sqrtF(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}

func atan2F(y, x float64) float64 {
	pi := 3.14159265358979323846
	if x > 0 {
		return atanF(y / x)
	}
	if x < 0 {
		if y >= 0 {
			return atanF(y/x) + pi
		}
		return atanF(y/x) - pi
	}
	if y > 0 {
		return pi / 2
	}
	if y < 0 {
		return -pi / 2
	}
	return 0
}

func atanF(x float64) float64 {
	pi := 3.14159265358979323846
	if x < 0 {
		return -atanF(-x)
	}
	if x > 1 {
		return pi/2 - atanF(1/x)
	}
	sum := 0.0
	term := x
	x2 := x * x
	sign := 1.0
	for i := 0; i < 100; i++ {
		sum += sign * term / float64(2*i+1)
		term *= x2
		sign = -sign
		if absF(term) < 1e-15 {
			break
		}
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
// Created by: MOUAAD IDOUFKIR
// << The universe runs on equations. We just translate them >>
