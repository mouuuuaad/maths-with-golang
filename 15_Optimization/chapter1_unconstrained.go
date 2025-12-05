package optimization

func absO(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

type ObjectiveFunc func(x []float64) float64

func GoldenSection(f func(float64) float64, a, b, tol float64) float64 {
	phi := (1 + 2.2360679774997896) / 2
	x1 := b - (b-a)/phi
	x2 := a + (b-a)/phi
	f1, f2 := f(x1), f(x2)
	for absO(b-a) > tol {
		if f1 < f2 {
			b = x2
			x2 = x1
			f2 = f1
			x1 = b - (b-a)/phi
			f1 = f(x1)
		} else {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + (b-a)/phi
			f2 = f(x2)
		}
	}
	return (a + b) / 2
}

func GradientDescent(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, lr float64, iters int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for i := 0; i < iters; i++ {
		g := grad(x)
		for j := range x {
			x[j] -= lr * g[j]
		}
	}
	return x
}

func NewtonMethod1D(f, df func(float64) float64, x0, tol float64) float64 {
	x := x0
	for i := 0; i < 100; i++ {
		fx := f(x)
		if absO(fx) < tol {
			return x
		}
		x = x - fx/df(x)
	}
	return x
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
