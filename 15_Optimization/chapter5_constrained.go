package optimization

func LagrangeMultiplier(f, g ObjectiveFunc, gradF, gradG func([]float64) []float64, x0 []float64, lambda0, tol float64) ([]float64, float64) {
	n := len(x0)
	x := make([]float64, n)
	copy(x, x0)
	lambda := lambda0
	for iter := 0; iter < 1000; iter++ {
		gf := gradF(x)
		gg := gradG(x)
		gVal := g(x)
		maxGrad := 0.0
		for i := 0; i < n; i++ {
			grad := gf[i] - lambda*gg[i]
			if absO(grad) > maxGrad {
				maxGrad = absO(grad)
			}
		}
		if maxGrad < tol && absO(gVal) < tol {
			break
		}
		step := 0.01
		for i := 0; i < n; i++ {
			x[i] -= step * (gf[i] - lambda*gg[i])
		}
		lambda += step * gVal
	}
	return x, lambda
}

func PenaltyMethod(f, g ObjectiveFunc, x0 []float64, rho, tol float64) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for k := 0; k < 20; k++ {
		penalty := func(xx []float64) float64 {
			gVal := g(xx)
			return f(xx) + rho*gVal*gVal
		}
		x = NelderMead(penalty, x, tol)
		if absO(g(x)) < tol {
			break
		}
		rho *= 10
	}
	return x
}

func BarrierMethod(f ObjectiveFunc, inequalities []ObjectiveFunc, x0 []float64, mu, tol float64) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for k := 0; k < 20; k++ {
		barrier := func(xx []float64) float64 {
			val := f(xx)
			for _, g := range inequalities {
				gVal := g(xx)
				if gVal <= 0 {
					return 1e18
				}
				val -= mu * logO(gVal)
			}
			return val
		}
		x = NelderMead(barrier, x, tol)
		mu /= 10
		if mu < tol {
			break
		}
	}
	return x
}

func logO(x float64) float64 {
	if x <= 0 {
		return -1e18
	}
	y := 0.0
	for x > 2 {
		x /= 2.718281828
		y++
	}
	for x < 0.5 {
		x *= 2.718281828
		y--
	}
	z := (x - 1) / (x + 1)
	z2 := z * z
	term := z
	sum := z
	for i := 1; i < 50; i++ {
		term *= z2
		sum += term / float64(2*i+1)
	}
	return y + 2*sum
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
