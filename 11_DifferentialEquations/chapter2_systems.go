package diffeq

type SystemODE func(t float64, y []float64) []float64

func RK4System(f SystemODE, t0 float64, y0 []float64, h float64, steps int) [][]float64 {
	n := len(y0)
	result := make([][]float64, steps+1)
	result[0] = make([]float64, n)
	copy(result[0], y0)
	t := t0
	for i := 0; i < steps; i++ {
		y := result[i]
		k1 := f(t, y)
		y2 := make([]float64, n)
		for j := 0; j < n; j++ {
			y2[j] = y[j] + h*k1[j]/2
		}
		k2 := f(t+h/2, y2)
		y3 := make([]float64, n)
		for j := 0; j < n; j++ {
			y3[j] = y[j] + h*k2[j]/2
		}
		k3 := f(t+h/2, y3)
		y4 := make([]float64, n)
		for j := 0; j < n; j++ {
			y4[j] = y[j] + h*k3[j]
		}
		k4 := f(t+h, y4)
		result[i+1] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i+1][j] = y[j] + h*(k1[j]+2*k2[j]+2*k3[j]+k4[j])/6
		}
		t += h
	}
	return result
}

func EulerSystem(f SystemODE, t0 float64, y0 []float64, h float64, steps int) [][]float64 {
	n := len(y0)
	result := make([][]float64, steps+1)
	result[0] = make([]float64, n)
	copy(result[0], y0)
	t := t0
	for i := 0; i < steps; i++ {
		dy := f(t, result[i])
		result[i+1] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i+1][j] = result[i][j] + h*dy[j]
		}
		t += h
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
