package calculus

type DifferentialEquation func(x, y float64) float64

func EulerMethod(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		y += step * f(x, y)
		x += step
	}
	return y
}

func HeunMethod(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		k1 := f(x, y)
		k2 := f(x+step, y+step*k1)
		y += step * (k1 + k2) / 2
		x += step
	}
	return y
}

func RungeKutta4(f DifferentialEquation, x0, y0, xEnd, step float64) float64 {
	x := x0
	y := y0
	for x < xEnd {
		k1 := step * f(x, y)
		k2 := step * f(x+step/2, y+k1/2)
		k3 := step * f(x+step/2, y+k2/2)
		k4 := step * f(x+step, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
		x += step
	}
	return y
}

func RungeKutta4Full(f DifferentialEquation, x0, y0, xEnd, step float64) ([]float64, []float64) {
	xs := []float64{x0}
	ys := []float64{y0}
	x := x0
	y := y0
	for x < xEnd {
		k1 := step * f(x, y)
		k2 := step * f(x+step/2, y+k1/2)
		k3 := step * f(x+step/2, y+k2/2)
		k4 := step * f(x+step, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
		x += step
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func AdaptiveRK45(f DifferentialEquation, x0, y0, xEnd, tol float64) float64 {
	x := x0
	y := y0
	h := (xEnd - x0) / 100
	for x < xEnd {
		k1 := h * f(x, y)
		k2 := h * f(x+h/4, y+k1/4)
		k3 := h * f(x+3*h/8, y+3*k1/32+9*k2/32)
		k4 := h * f(x+12*h/13, y+1932*k1/2197-7200*k2/2197+7296*k3/2197)
		k5 := h * f(x+h, y+439*k1/216-8*k2+3680*k3/513-845*k4/4104)
		k6 := h * f(x+h/2, y-8*k1/27+2*k2-3544*k3/2565+1859*k4/4104-11*k5/40)
		y4 := y + 25*k1/216 + 1408*k3/2565 + 2197*k4/4104 - k5/5
		y5 := y + 16*k1/135 + 6656*k3/12825 + 28561*k4/56430 - 9*k5/50 + 2*k6/55
		err := absL(y5 - y4)
		if err < tol {
			y = y5
			x += h
			if err > 0 {
				h *= 0.9 * powerS(tol/err, 0.2)
			}
		} else {
			h *= 0.9 * powerS(tol/err, 0.25)
		}
		if x+h > xEnd {
			h = xEnd - x
		}
	}
	return y
}

func NewtonRaphson(f, df Function, x0 float64, maxIter int) float64 {
	x := x0
	for i := 0; i < maxIter; i++ {
		fx := f(x)
		if absL(fx) < 1e-12 {
			return x
		}
		dfx := df(x)
		if absL(dfx) < 1e-15 {
			break
		}
		x -= fx / dfx
	}
	return x
}

func Bisection(f Function, a, b float64, tol float64) float64 {
	fa := f(a)
	for i := 0; i < 100; i++ {
		c := (a + b) / 2
		fc := f(c)
		if absL(fc) < tol || (b-a)/2 < tol {
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
		if absL(f1-f0) < 1e-15 {
			break
		}
		x2 := x1 - f1*(x1-x0)/(f1-f0)
		if absL(x2-x1) < 1e-12 {
			return x2
		}
		x0 = x1
		x1 = x2
	}
	return x1
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
		if sqrtS(norm) < 1e-8 {
			break
		}
		for j := range x {
			x[j] -= alpha * grad[j]
		}
	}
	return x
}

func sqrtS(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
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
// Created by: MOUAAD
// MathsWithGolang - Pure Golang Mathematical Library
