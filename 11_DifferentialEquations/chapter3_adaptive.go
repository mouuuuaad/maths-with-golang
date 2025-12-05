package diffeq

func absD(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func AdaptiveRK45(f ODE, t0, y0, tEnd, tol float64) ([]float64, []float64) {
	ts := []float64{t0}
	ys := []float64{y0}
	t, y := t0, y0
	h := (tEnd - t0) / 100
	for t < tEnd {
		k1 := f(t, y)
		k2 := f(t+h/4, y+h*k1/4)
		k3 := f(t+3*h/8, y+3*h*k1/32+9*h*k2/32)
		k4 := f(t+12*h/13, y+1932*h*k1/2197-7200*h*k2/2197+7296*h*k3/2197)
		k5 := f(t+h, y+439*h*k1/216-8*h*k2+3680*h*k3/513-845*h*k4/4104)
		k6 := f(t+h/2, y-8*h*k1/27+2*h*k2-3544*h*k3/2565+1859*h*k4/4104-11*h*k5/40)
		y4 := y + h*(25*k1/216+1408*k3/2565+2197*k4/4104-k5/5)
		y5 := y + h*(16*k1/135+6656*k3/12825+28561*k4/56430-9*k5/50+2*k6/55)
		err := absD(y5 - y4)
		if err < tol {
			t += h
			y = y5
			ts = append(ts, t)
			ys = append(ys, y)
		}
		if err > 0 {
			h *= 0.84 * powD(tol/err, 0.25)
		}
		if t+h > tEnd {
			h = tEnd - t
		}
	}
	return ts, ys
}

func powD(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
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
