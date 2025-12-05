package optimization

func NelderMead(f ObjectiveFunc, x0 []float64, tol float64) []float64 {
	n := len(x0)
	simplex := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		simplex[i] = make([]float64, n)
		copy(simplex[i], x0)
		if i > 0 {
			simplex[i][i-1] += 0.5
		}
	}
	values := make([]float64, n+1)
	for i := range simplex {
		values[i] = f(simplex[i])
	}
	alpha, gamma, rho, sigma := 1.0, 2.0, 0.5, 0.5
	for iter := 0; iter < 1000; iter++ {
		best, worst, second := 0, 0, 0
		for i := range values {
			if values[i] < values[best] {
				best = i
			}
			if values[i] > values[worst] {
				worst = i
			}
		}
		for i := range values {
			if i != worst && values[i] > values[second] {
				second = i
			}
		}
		if values[worst]-values[best] < tol {
			return simplex[best]
		}
		centroid := make([]float64, n)
		for i := range simplex {
			if i != worst {
				for j := 0; j < n; j++ {
					centroid[j] += simplex[i][j] / float64(n)
				}
			}
		}
		reflected := make([]float64, n)
		for j := 0; j < n; j++ {
			reflected[j] = centroid[j] + alpha*(centroid[j]-simplex[worst][j])
		}
		fr := f(reflected)
		if fr < values[second] && fr >= values[best] {
			copy(simplex[worst], reflected)
			values[worst] = fr
			continue
		}
		if fr < values[best] {
			expanded := make([]float64, n)
			for j := 0; j < n; j++ {
				expanded[j] = centroid[j] + gamma*(reflected[j]-centroid[j])
			}
			fe := f(expanded)
			if fe < fr {
				copy(simplex[worst], expanded)
				values[worst] = fe
			} else {
				copy(simplex[worst], reflected)
				values[worst] = fr
			}
			continue
		}
		contracted := make([]float64, n)
		for j := 0; j < n; j++ {
			contracted[j] = centroid[j] + rho*(simplex[worst][j]-centroid[j])
		}
		fc := f(contracted)
		if fc < values[worst] {
			copy(simplex[worst], contracted)
			values[worst] = fc
			continue
		}
		for i := range simplex {
			if i != best {
				for j := 0; j < n; j++ {
					simplex[i][j] = simplex[best][j] + sigma*(simplex[i][j]-simplex[best][j])
				}
				values[i] = f(simplex[i])
			}
		}
	}
	best := 0
	for i := range values {
		if values[i] < values[best] {
			best = i
		}
	}
	return simplex[best]
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
