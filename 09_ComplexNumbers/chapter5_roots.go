package complexnums

func (c ComplexNumber) NthRoots(n int) []ComplexNumber {
	if n <= 0 {
		return nil
	}
	p := c.ToPolar()
	r := powerC(p.Radius, 1.0/float64(n))
	theta := p.Theta
	roots := make([]ComplexNumber, n)
	for k := 0; k < n; k++ {
		angle := (theta + 2*PiC*float64(k)) / float64(n)
		roots[k] = ComplexNumber{
			R: r * cosC(angle),
			I: r * sinC(angle),
		}
	}
	return roots
}

func RootsOfUnity(n int) []ComplexNumber {
	return New(1, 0).NthRoots(n)
}

func QuadraticRoots(a, b, c ComplexNumber) (ComplexNumber, ComplexNumber) {
	discriminant := b.Multiply(b).Subtract(New(4, 0).Multiply(a).Multiply(c))
	sqrtD := discriminant.Sqrt()
	twoA := New(2, 0).Multiply(a)
	negB := b.Negate()
	r1 := negB.Add(sqrtD).Divide(twoA)
	r2 := negB.Subtract(sqrtD).Divide(twoA)
	return r1, r2
}

func CubicRoots(a, b, c, d ComplexNumber) []ComplexNumber {
	if a.Abs() < 1e-9 {
		r1, r2 := QuadraticRoots(b, c, d)
		return []ComplexNumber{r1, r2}
	}
	p := c.Divide(a).Subtract(b.Multiply(b).Divide(a.Multiply(a).Scale(3)))
	q := d.Divide(a).Add(b.PowerN(3).Divide(a.PowerN(3).Scale(13.5))).Subtract(b.Multiply(c).Divide(a.Multiply(a).Scale(3)))
	discriminant := q.Multiply(q).Scale(0.25).Add(p.PowerN(3).Scale(1.0 / 27.0))
	sqrtD := discriminant.Sqrt()
	u := q.Scale(-0.5).Add(sqrtD).NthRoots(3)[0]
	v := q.Scale(-0.5).Subtract(sqrtD).NthRoots(3)[0]
	offset := b.Divide(a.Scale(3)).Negate()
	return []ComplexNumber{u.Add(v).Add(offset)}
}

func powerC(base, exp float64) float64 {
	if base <= 0 {
		return 0
	}
	return expC(exp * lnC(base))
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
