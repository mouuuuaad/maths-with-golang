package diffeq

func SHOSolver(omega, y0, v0, h float64, steps int) ([]float64, []float64) {
	y := make([]float64, steps+1)
	v := make([]float64, steps+1)
	y[0], v[0] = y0, v0
	for i := 0; i < steps; i++ {
		k1y := v[i]
		k1v := -omega * omega * y[i]
		k2y := v[i] + h*k1v/2
		k2v := -omega * omega * (y[i] + h*k1y/2)
		k3y := v[i] + h*k2v/2
		k3v := -omega * omega * (y[i] + h*k2y/2)
		k4y := v[i] + h*k3v
		k4v := -omega * omega * (y[i] + h*k3y)
		y[i+1] = y[i] + h*(k1y+2*k2y+2*k3y+k4y)/6
		v[i+1] = v[i] + h*(k1v+2*k2v+2*k3v+k4v)/6
	}
	return y, v
}

func DampedOscillator(omega, gamma, y0, v0, h float64, steps int) ([]float64, []float64) {
	y := make([]float64, steps+1)
	v := make([]float64, steps+1)
	y[0], v[0] = y0, v0
	for i := 0; i < steps; i++ {
		k1y := v[i]
		k1v := -2*gamma*v[i] - omega*omega*y[i]
		k2y := v[i] + h*k1v/2
		k2v := -2*gamma*(v[i]+h*k1v/2) - omega*omega*(y[i]+h*k1y/2)
		k3y := v[i] + h*k2v/2
		k3v := -2*gamma*(v[i]+h*k2v/2) - omega*omega*(y[i]+h*k2y/2)
		k4y := v[i] + h*k3v
		k4v := -2*gamma*(v[i]+h*k3v) - omega*omega*(y[i]+h*k3y)
		y[i+1] = y[i] + h*(k1y+2*k2y+2*k3y+k4y)/6
		v[i+1] = v[i] + h*(k1v+2*k2v+2*k3v+k4v)/6
	}
	return y, v
}

func VanDerPol(mu, y0, v0, h float64, steps int) ([]float64, []float64) {
	y := make([]float64, steps+1)
	v := make([]float64, steps+1)
	y[0], v[0] = y0, v0
	for i := 0; i < steps; i++ {
		f := func(yy, vv float64) (float64, float64) {
			return vv, mu*(1-yy*yy)*vv - yy
		}
		k1y, k1v := f(y[i], v[i])
		k2y, k2v := f(y[i]+h*k1y/2, v[i]+h*k1v/2)
		k3y, k3v := f(y[i]+h*k2y/2, v[i]+h*k2v/2)
		k4y, k4v := f(y[i]+h*k3y, v[i]+h*k3v)
		y[i+1] = y[i] + h*(k1y+2*k2y+2*k3y+k4y)/6
		v[i+1] = v[i] + h*(k1v+2*k2v+2*k3v+k4v)/6
	}
	return y, v
}

func LorenzSystem(sigma, rho, beta float64, x0, y0, z0, h float64, steps int) ([]float64, []float64, []float64) {
	x := make([]float64, steps+1)
	y := make([]float64, steps+1)
	z := make([]float64, steps+1)
	x[0], y[0], z[0] = x0, y0, z0
	for i := 0; i < steps; i++ {
		dx := sigma * (y[i] - x[i])
		dy := x[i]*(rho-z[i]) - y[i]
		dz := x[i]*y[i] - beta*z[i]
		x[i+1] = x[i] + h*dx
		y[i+1] = y[i] + h*dy
		z[i+1] = z[i] + h*dz
	}
	return x, y, z
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
