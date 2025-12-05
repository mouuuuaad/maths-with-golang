package diffeq

func BVPShooting(f ODE, a, b, ya, yb float64, n int) []float64 {
	s0, s1 := 0.0, 1.0
	y0 := solveIVP(f, a, ya, s0, b, n)
	y1 := solveIVP(f, a, ya, s1, b, n)
	for i := 0; i < 50; i++ {
		if absD(y1[n]-yb) < 1e-10 {
			return y1
		}
		s2 := s1 - (y1[n]-yb)*(s1-s0)/(y1[n]-y0[n])
		s0, s1 = s1, s2
		y0, y1 = y1, solveIVP(f, a, ya, s2, b, n)
	}
	return y1
}

func solveIVP(f ODE, t0, y0, slope, tEnd float64, n int) []float64 {
	h := (tEnd - t0) / float64(n)
	y := make([]float64, n+1)
	y[0] = y0
	t := t0
	for i := 0; i < n; i++ {
		k1 := slope + f(t, y[i])
		k2 := slope + f(t+h/2, y[i]+h*k1/2)
		k3 := slope + f(t+h/2, y[i]+h*k2/2)
		k4 := slope + f(t+h, y[i]+h*k3)
		y[i+1] = y[i] + h*(k1+2*k2+2*k3+k4)/6
		t += h
	}
	return y
}

func FiniteDifference(f ODE, a, b, ya, yb float64, n int) []float64 {
	h := (b - a) / float64(n)
	y := make([]float64, n+1)
	y[0] = ya
	y[n] = yb
	for k := 0; k < 100; k++ {
		for i := 1; i < n; i++ {
			t := a + float64(i)*h
			y[i] = (y[i-1] + y[i+1] + h*h*f(t, y[i])) / 2
		}
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
