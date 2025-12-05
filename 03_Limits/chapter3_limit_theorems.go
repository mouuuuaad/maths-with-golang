package limits

func SqueezeLemma(f, g, h Function, a float64) (float64, bool) {
	fLim, fOk := Limit(f, a)
	hLim, hOk := Limit(h, a)
	if !fOk || !hOk || absLim(fLim-hLim) > 1e-9 {
		return 0, false
	}
	return fLim, true
}

func LHopital(f, g Function, a float64) (float64, bool) {
	h := 1e-7
	df := func(x float64) float64 {
		return (f(x+h) - f(x-h)) / (2 * h)
	}
	dg := func(x float64) float64 {
		return (g(x+h) - g(x-h)) / (2 * h)
	}
	fa, ga := f(a+h), g(a+h)
	if absLim(fa) > 1e-6 || absLim(ga) > 1e-6 {
		return 0, false
	}
	ratio := func(x float64) float64 {
		dgx := dg(x)
		if absLim(dgx) < 1e-12 {
			return 0
		}
		return df(x) / dgx
	}
	return Limit(ratio, a)
}

func LimitProduct(f, g Function, a float64) (float64, bool) {
	fLim, fOk := Limit(f, a)
	gLim, gOk := Limit(g, a)
	if fOk && gOk {
		return fLim * gLim, true
	}
	return 0, false
}

func LimitSum(f, g Function, a float64) (float64, bool) {
	fLim, fOk := Limit(f, a)
	gLim, gOk := Limit(g, a)
	if fOk && gOk {
		return fLim + gLim, true
	}
	return 0, false
}

func LimitQuotient(f, g Function, a float64) (float64, bool) {
	fLim, fOk := Limit(f, a)
	gLim, gOk := Limit(g, a)
	if fOk && gOk && absLim(gLim) > 1e-12 {
		return fLim / gLim, true
	}
	return 0, false
}

func LimitPower(f Function, n float64, a float64) (float64, bool) {
	fLim, fOk := Limit(f, a)
	if !fOk {
		return 0, false
	}
	return powerLim(fLim, n), true
}

func powerLim(base, exp float64) float64 {
	if base <= 0 {
		return 0
	}
	return expLim(exp * lnLim(base))
}

func expLim(x float64) float64 {
	if x < 0 {
		return 1.0 / expLim(-x)
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 50; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func lnLim(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	for i := 0; i < 50; i++ {
		ey := expLim(y)
		diff := ey - x
		if absLim(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
	}
	return y
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
