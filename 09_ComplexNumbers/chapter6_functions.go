package complexnums

func (c ComplexNumber) Sin() ComplexNumber {
	return ComplexNumber{
		R: sinC(c.R) * coshC(c.I),
		I: cosC(c.R) * sinhC(c.I),
	}
}

func (c ComplexNumber) Cos() ComplexNumber {
	return ComplexNumber{
		R: cosC(c.R) * coshC(c.I),
		I: -sinC(c.R) * sinhC(c.I),
	}
}

func (c ComplexNumber) Tan() ComplexNumber {
	return c.Sin().Divide(c.Cos())
}

func (c ComplexNumber) Sinh() ComplexNumber {
	ez := c.Exp()
	emz := c.Negate().Exp()
	return ez.Subtract(emz).Scale(0.5)
}

func (c ComplexNumber) Cosh() ComplexNumber {
	ez := c.Exp()
	emz := c.Negate().Exp()
	return ez.Add(emz).Scale(0.5)
}

func (c ComplexNumber) Tanh() ComplexNumber {
	return c.Sinh().Divide(c.Cosh())
}

func (c ComplexNumber) Gamma() ComplexNumber {
	p := []float64{
		76.18009172947146,
		-86.50532032941677,
		24.01409824083091,
		-1.231739572450155,
		0.1208650973866179e-2,
		-0.5395239384953e-5,
	}
	if c.R < 0.5 {
		pi := New(PiC, 0)
		sinPiZ := c.Multiply(pi).Sin()
		oneMinusZ := New(1, 0).Subtract(c)
		return pi.Divide(sinPiZ.Multiply(oneMinusZ.Gamma()))
	}
	z := c.Subtract(New(1, 0))
	x := z.Add(New(5.5, 0))
	term1 := z.Add(New(0.5, 0)).Multiply(x.Log()).Subtract(x)
	ser := New(1.000000000190015, 0)
	for i := 0; i < len(p); i++ {
		num := New(p[i], 0)
		den := z.Add(New(float64(i)+1, 0))
		ser = ser.Add(num.Divide(den))
	}
	sqrt2Pi := New(2.5066282746310005, 0)
	return term1.Exp().Multiply(sqrt2Pi).Multiply(ser)
}

func (c ComplexNumber) Zeta(terms int) ComplexNumber {
	sum := New(0, 0)
	for n := 1; n <= terms; n++ {
		nC := New(float64(n), 0)
		term := Pow(nC, c).Inverse()
		sum = sum.Add(term)
	}
	return sum
}

func sinhC(x float64) float64 {
	ex := expC(x)
	emx := expC(-x)
	return (ex - emx) / 2
}

func coshC(x float64) float64 {
	ex := expC(x)
	emx := expC(-x)
	return (ex + emx) / 2
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
