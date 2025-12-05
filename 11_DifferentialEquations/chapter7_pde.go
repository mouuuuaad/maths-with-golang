package diffeq

type PDE2D func(x, y, u, ux, uy float64) float64

func HeatEquation1D(u0 []float64, alpha, dx, dt float64, steps int) [][]float64 {
	n := len(u0)
	result := make([][]float64, steps+1)
	result[0] = make([]float64, n)
	copy(result[0], u0)
	r := alpha * dt / (dx * dx)
	for t := 0; t < steps; t++ {
		result[t+1] = make([]float64, n)
		result[t+1][0] = result[t][0]
		result[t+1][n-1] = result[t][n-1]
		for i := 1; i < n-1; i++ {
			result[t+1][i] = result[t][i] + r*(result[t][i+1]-2*result[t][i]+result[t][i-1])
		}
	}
	return result
}

func WaveEquation1D(u0, v0 []float64, c, dx, dt float64, steps int) [][]float64 {
	n := len(u0)
	result := make([][]float64, steps+1)
	result[0] = make([]float64, n)
	copy(result[0], u0)
	r := c * c * dt * dt / (dx * dx)
	result[1] = make([]float64, n)
	for i := 1; i < n-1; i++ {
		result[1][i] = result[0][i] + dt*v0[i] + 0.5*r*(result[0][i+1]-2*result[0][i]+result[0][i-1])
	}
	for t := 1; t < steps; t++ {
		result[t+1] = make([]float64, n)
		for i := 1; i < n-1; i++ {
			result[t+1][i] = 2*result[t][i] - result[t-1][i] + r*(result[t][i+1]-2*result[t][i]+result[t][i-1])
		}
	}
	return result
}

func LaplaceEquation2D(nx, ny int, boundary func(i, j int) (float64, bool), tol float64) [][]float64 {
	u := make([][]float64, ny)
	for j := 0; j < ny; j++ {
		u[j] = make([]float64, nx)
		for i := 0; i < nx; i++ {
			if val, isBoundary := boundary(i, j); isBoundary {
				u[j][i] = val
			}
		}
	}
	for iter := 0; iter < 10000; iter++ {
		maxDiff := 0.0
		for j := 1; j < ny-1; j++ {
			for i := 1; i < nx-1; i++ {
				if _, isBoundary := boundary(i, j); !isBoundary {
					newVal := 0.25 * (u[j][i+1] + u[j][i-1] + u[j+1][i] + u[j-1][i])
					diff := absD(newVal - u[j][i])
					if diff > maxDiff {
						maxDiff = diff
					}
					u[j][i] = newVal
				}
			}
		}
		if maxDiff < tol {
			break
		}
	}
	return u
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
