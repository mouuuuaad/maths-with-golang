package diffeq

func VerletMethod(f ODE, t0, y0, v0, h float64, steps int) ([]float64, []float64) {
	y := make([]float64, steps+1)
	v := make([]float64, steps+1)
	y[0], v[0] = y0, v0
	t := t0
	for i := 0; i < steps; i++ {
		a := f(t, y[i])
		y[i+1] = y[i] + v[i]*h + 0.5*a*h*h
		aNew := f(t+h, y[i+1])
		v[i+1] = v[i] + 0.5*(a+aNew)*h
		t += h
	}
	return y, v
}

func Leapfrog(f ODE, t0, y0, v0, h float64, steps int) ([]float64, []float64) {
	y := make([]float64, steps+1)
	v := make([]float64, steps+1)
	y[0], v[0] = y0, v0
	vHalf := v0 + 0.5*h*f(t0, y0)
	t := t0
	for i := 0; i < steps; i++ {
		y[i+1] = y[i] + h*vHalf
		a := f(t+h, y[i+1])
		vHalf = vHalf + h*a
		v[i+1] = vHalf - 0.5*h*a
		t += h
	}
	return y, v
}

func BackwardEuler(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		yNext := y[i] + h*f(t+h, y[i])
		for j := 0; j < 10; j++ {
			yNext = y[i] + h*f(t+h, yNext)
		}
		y[i+1] = yNext
		t += h
	}
	return y
}

func TrapezoidalMethod(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		k1 := f(t, y[i])
		yPredict := y[i] + h*k1
		for j := 0; j < 10; j++ {
			k2 := f(t+h, yPredict)
			yPredict = y[i] + h*(k1+k2)/2
		}
		y[i+1] = yPredict
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
