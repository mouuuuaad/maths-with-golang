package probability

const Pi = 3.14159265358979323846

func NormalPDF(x, mu, sigma float64) float64 {
	z := (x - mu) / sigma
	return expP(-0.5*z*z) / (sigma * sqrtP(2*Pi))
}

func NormalCDF(x, mu, sigma float64) float64 {
	return 0.5 * (1 + erf((x-mu)/(sigma*sqrtP(2))))
}

func StandardNormalPDF(x float64) float64 {
	return NormalPDF(x, 0, 1)
}

func StandardNormalCDF(x float64) float64 {
	return NormalCDF(x, 0, 1)
}

func ExponentialPDF(x, lambda float64) float64 {
	if x < 0 {
		return 0
	}
	return lambda * expP(-lambda*x)
}

func ExponentialCDF(x, lambda float64) float64 {
	if x < 0 {
		return 0
	}
	return 1 - expP(-lambda*x)
}

func UniformPDF(x, a, b float64) float64 {
	if x < a || x > b {
		return 0
	}
	return 1 / (b - a)
}

func UniformCDF(x, a, b float64) float64 {
	if x < a {
		return 0
	}
	if x > b {
		return 1
	}
	return (x - a) / (b - a)
}

func GammaPDF(x, alpha, beta float64) float64 {
	if x <= 0 {
		return 0
	}
	return powerP(beta, alpha) * powerP(x, alpha-1) * expP(-beta*x) / gammaFunc(alpha)
}

func BetaPDF(x, alpha, beta float64) float64 {
	if x <= 0 || x >= 1 {
		return 0
	}
	return powerP(x, alpha-1) * powerP(1-x, beta-1) / betaFunc(alpha, beta)
}

func gammaFunc(z float64) float64 {
	g := 7
	c := []float64{
		0.99999999999980993,
		676.5203681218851,
		-1259.1392167224028,
		771.32342877765313,
		-176.61502916214059,
		12.507343278686905,
		-0.13857109526572012,
		9.9843695780195716e-6,
		1.5056327351493116e-7,
	}
	if z < 0.5 {
		return Pi / (sinP(Pi*z) * gammaFunc(1-z))
	}
	z -= 1
	x := c[0]
	for i := 1; i < g+2; i++ {
		x += c[i] / (z + float64(i))
	}
	t := z + float64(g) + 0.5
	return sqrtP(2*Pi) * powerP(t, z+0.5) * expP(-t) * x
}

func betaFunc(a, b float64) float64 {
	return gammaFunc(a) * gammaFunc(b) / gammaFunc(a+b)
}

func sinP(x float64) float64 {
	for x > Pi {
		x -= 2 * Pi
	}
	for x < -Pi {
		x += 2 * Pi
	}
	sum := 0.0
	term := x
	x2 := x * x
	for i := 1; i < 30; i++ {
		sum += term
		term *= -x2 / float64(2*i*(2*i+1))
	}
	return sum
}

func erf(x float64) float64 {
	a1 := 0.254829592
	a2 := -0.284496736
	a3 := 1.421413741
	a4 := -1.453152027
	a5 := 1.061405429
	p := 0.3275911
	sign := 1.0
	if x < 0 {
		sign = -1
		x = -x
	}
	t := 1.0 / (1.0 + p*x)
	y := 1.0 - (((((a5*t+a4)*t)+a3)*t+a2)*t+a1)*t*expP(-x*x)
	return sign * y
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
