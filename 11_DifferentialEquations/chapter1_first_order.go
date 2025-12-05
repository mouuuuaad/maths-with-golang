package diffeq

type ODE func(t, y float64) float64

func EulerMethod(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		y[i+1] = y[i] + h*f(t, y[i])
		t += h
	}
	return y
}

func HeunMethod(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		k1 := f(t, y[i])
		k2 := f(t+h, y[i]+h*k1)
		y[i+1] = y[i] + h*(k1+k2)/2
		t += h
	}
	return y
}

func MidpointMethod(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		k1 := f(t, y[i])
		k2 := f(t+h/2, y[i]+h*k1/2)
		y[i+1] = y[i] + h*k2
		t += h
	}
	return y
}

func RungeKutta4(f ODE, t0, y0, h float64, steps int) []float64 {
	y := make([]float64, steps+1)
	y[0] = y0
	t := t0
	for i := 0; i < steps; i++ {
		k1 := f(t, y[i])
		k2 := f(t+h/2, y[i]+h*k1/2)
		k3 := f(t+h/2, y[i]+h*k2/2)
		k4 := f(t+h, y[i]+h*k3)
		y[i+1] = y[i] + h*(k1+2*k2+2*k3+k4)/6
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
