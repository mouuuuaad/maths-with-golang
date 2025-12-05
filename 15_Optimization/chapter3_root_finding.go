package optimization

func BisectionMethod(f func(float64) float64, a, b, tol float64) float64 {
	for absO(b-a) > tol {
		c := (a + b) / 2
		if f(c) == 0 {
			return c
		}
		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return (a + b) / 2
}

func SecantMethod(f func(float64) float64, x0, x1, tol float64) float64 {
	for i := 0; i < 100; i++ {
		f0, f1 := f(x0), f(x1)
		if absO(f1) < tol {
			return x1
		}
		x2 := x1 - f1*(x1-x0)/(f1-f0)
		x0, x1 = x1, x2
	}
	return x1
}

func BrentMethod(f func(float64) float64, a, b, tol float64) float64 {
	fa, fb := f(a), f(b)
	if fa*fb > 0 {
		return (a + b) / 2
	}
	if absO(fa) < absO(fb) {
		a, b = b, a
		fa, fb = fb, fa
	}
	c, fc := a, fa
	mflag := true
	var s, d float64
	for absO(b-a) > tol && fb != 0 {
		if fa != fc && fb != fc {
			s = a*fb*fc/((fa-fb)*(fa-fc)) + b*fa*fc/((fb-fa)*(fb-fc)) + c*fa*fb/((fc-fa)*(fc-fb))
		} else {
			s = b - fb*(b-a)/(fb-fa)
		}
		cond1 := (s < (3*a+b)/4 || s > b) && (s > (3*a+b)/4 || s < b)
		cond2 := mflag && absO(s-b) >= absO(b-c)/2
		cond3 := !mflag && absO(s-b) >= absO(c-d)/2
		cond4 := mflag && absO(b-c) < tol
		cond5 := !mflag && absO(c-d) < tol
		if cond1 || cond2 || cond3 || cond4 || cond5 {
			s = (a + b) / 2
			mflag = true
		} else {
			mflag = false
		}
		fs := f(s)
		d, c, fc = c, b, fb
		if fa*fs < 0 {
			b, fb = s, fs
		} else {
			a, fa = s, fs
		}
		if absO(fa) < absO(fb) {
			a, b = b, a
			fa, fb = fb, fa
		}
	}
	return b
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
