package probability

func BinomialPMF(n, k int, p float64) float64 {
	return Combination(n, k) * powerP(p, float64(k)) * powerP(1-p, float64(n-k))
}

func BinomialCDF(n, k int, p float64) float64 {
	sum := 0.0
	for i := 0; i <= k; i++ {
		sum += BinomialPMF(n, i, p)
	}
	return sum
}

func BinomialMean(n int, p float64) float64 {
	return float64(n) * p
}

func BinomialVariance(n int, p float64) float64 {
	return float64(n) * p * (1 - p)
}

func PoissonPMF(k int, lambda float64) float64 {
	return powerP(lambda, float64(k)) * expP(-lambda) / Factorial(k)
}

func PoissonCDF(k int, lambda float64) float64 {
	sum := 0.0
	for i := 0; i <= k; i++ {
		sum += PoissonPMF(i, lambda)
	}
	return sum
}

func GeometricPMF(k int, p float64) float64 {
	if k < 1 {
		return 0
	}
	return p * powerP(1-p, float64(k-1))
}

func GeometricCDF(k int, p float64) float64 {
	return 1 - powerP(1-p, float64(k))
}

func NegativeBinomialPMF(k, r int, p float64) float64 {
	return Combination(k-1, r-1) * powerP(p, float64(r)) * powerP(1-p, float64(k-r))
}

func HypergeometricPMF(N, K, n, k int) float64 {
	return Combination(K, k) * Combination(N-K, n-k) / Combination(N, n)
}

func powerP(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	if base <= 0 {
		return 0
	}
	return expP(exp * lnP(base))
}

func expP(x float64) float64 {
	if x < 0 {
		return 1.0 / expP(-x)
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 50; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func lnP(x float64) float64 {
	if x <= 0 {
		return -1e308
	}
	y := x - 1.0
	for i := 0; i < 50; i++ {
		ey := expP(y)
		diff := ey - x
		if absP(diff) < 1e-12 {
			return y
		}
		y -= diff / ey
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
