package derivatives

func NewtonMethod(f Function, x0 float64, maxIter int) float64 {
	x := x0
	for i := 0; i < maxIter; i++ {
		fx := f(x)
		if absD(fx) < 1e-12 {
			return x
		}
		dfx := Derivative(f, x)
		if absD(dfx) < 1e-15 {
			break
		}
		x -= fx / dfx
	}
	return x
}

func GradientDescent(f MultiFunction, x0 []float64, alpha float64, maxIter int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	for i := 0; i < maxIter; i++ {
		grad := Gradient(f, x)
		norm := 0.0
		for _, g := range grad {
			norm += g * g
		}
		if sqrtD(norm) < 1e-8 {
			break
		}
		for j := range x {
			x[j] -= alpha * grad[j]
		}
	}
	return x
}

func GradientDescentWithMomentum(f MultiFunction, x0 []float64, alpha, beta float64, maxIter int) []float64 {
	x := make([]float64, len(x0))
	copy(x, x0)
	v := make([]float64, len(x0))
	for i := 0; i < maxIter; i++ {
		grad := Gradient(f, x)
		for j := range x {
			v[j] = beta*v[j] + alpha*grad[j]
			x[j] -= v[j]
		}
		norm := 0.0
		for _, g := range grad {
			norm += g * g
		}
		if sqrtD(norm) < 1e-8 {
			break
		}
	}
	return x
}

func Bisection(f Function, a, b float64, tol float64) float64 {
	fa := f(a)
	for i := 0; i < 100; i++ {
		c := (a + b) / 2
		fc := f(c)
		if absD(fc) < tol || (b-a)/2 < tol {
			return c
		}
		if fa*fc < 0 {
			b = c
		} else {
			a = c
			fa = fc
		}
	}
	return (a + b) / 2
}

func Secant(f Function, x0, x1 float64, maxIter int) float64 {
	for i := 0; i < maxIter; i++ {
		f0 := f(x0)
		f1 := f(x1)
		if absD(f1-f0) < 1e-15 {
			break
		}
		x2 := x1 - f1*(x1-x0)/(f1-f0)
		if absD(x2-x1) < 1e-12 {
			return x2
		}
		x0 = x1
		x1 = x2
	}
	return x1
}

func FindExtremum(f Function, start float64) float64 {
	df := func(x float64) float64 { return Derivative(f, x) }
	return NewtonMethod(df, start, 100)
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
