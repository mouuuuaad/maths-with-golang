package diffeq

func StiffSolver(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		yNext := y[i]
		for iter := 0; iter < 20; iter++ {
			k1 := f(t, y[i])
			k2 := f(t+h/2, y[i]+h*k1/2)
			k3 := f(t+h/2, yNext-h*k2/2+h*k1)
			k4 := f(t+h, yNext)
			yNew := y[i] + h*(k1+2*k2+2*k3+k4)/6
			if absD(yNew-yNext) < 1e-12 {
				yNext = yNew
				break
			}
			yNext = yNew
		}
		y[i+1] = yNext
		t += h
	}
	return y
}

func BDFMethod2(f ODE, t0, y0, y1, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	if steps > 0 {
		y[1] = y1
	}
	t := t0 + h
	for i := 1; i < steps; i++ {
		yNext := (4*y[i] - y[i-1]) / 3
		for iter := 0; iter < 20; iter++ {
			fNext := f(t+h, yNext)
			yNew := (4*y[i] - y[i-1] + 2*h*fNext) / 3
			if absD(yNew-yNext) < 1e-12 {
				break
			}
			yNext = yNew
		}
		y[i+1] = yNext
		t += h
	}
	return y
}

func CrankNicolson(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		fn := f(t, y[i])
		yNext := y[i] + h*fn
		for iter := 0; iter < 20; iter++ {
			fnp1 := f(t+h, yNext)
			yNew := y[i] + h*(fn+fnp1)/2
			if absD(yNew-yNext) < 1e-12 {
				break
			}
			yNext = yNew
		}
		y[i+1] = yNext
		t += h
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
