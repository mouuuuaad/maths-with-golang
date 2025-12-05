package optimization

func ConjugateGradient(A [][]float64, b []float64, tol float64) []float64 {
	n := len(b)
	x := make([]float64, n)
	r := make([]float64, n)
	copy(r, b)
	p := make([]float64, n)
	copy(p, r)
	rsOld := dotProduct(r, r)
	for i := 0; i < n; i++ {
		Ap := matVecMult(A, p)
		alpha := rsOld / dotProduct(p, Ap)
		for j := 0; j < n; j++ {
			x[j] += alpha * p[j]
			r[j] -= alpha * Ap[j]
		}
		rsNew := dotProduct(r, r)
		if sqrtO(rsNew) < tol {
			break
		}
		beta := rsNew / rsOld
		for j := 0; j < n; j++ {
			p[j] = r[j] + beta*p[j]
		}
		rsOld = rsNew
	}
	return x
}

func BFGS(f ObjectiveFunc, grad func([]float64) []float64, x0 []float64, tol float64) []float64 {
	n := len(x0)
	x := make([]float64, n)
	copy(x, x0)
	H := identityMatrix(n)
	g := grad(x)
	for iter := 0; iter < 1000; iter++ {
		if norm(g) < tol {
			break
		}
		p := make([]float64, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				p[i] -= H[i][j] * g[j]
			}
		}
		alpha := lineSearch(f, x, p)
		s := make([]float64, n)
		for i := 0; i < n; i++ {
			s[i] = alpha * p[i]
			x[i] += s[i]
		}
		gNew := grad(x)
		y := make([]float64, n)
		for i := 0; i < n; i++ {
			y[i] = gNew[i] - g[i]
		}
		rho := 1.0 / dotProduct(y, s)
		H = updateBFGS(H, s, y, rho)
		g = gNew
	}
	return x
}

func identityMatrix(n int) [][]float64 {
	I := make([][]float64, n)
	for i := 0; i < n; i++ {
		I[i] = make([]float64, n)
		I[i][i] = 1
	}
	return I
}

func updateBFGS(H [][]float64, s, y []float64, rho float64) [][]float64 {
	n := len(s)
	I := identityMatrix(n)
	sy := outerProduct(s, y)
	_ = outerProduct(y, s)
	ss := outerProduct(s, s)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			I[i][j] -= rho * sy[i][j]
		}
	}
	H = matMult(I, matMult(H, transposeM(I)))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			H[i][j] += rho * ss[i][j]
		}
	}
	return H
}

func outerProduct(a, b []float64) [][]float64 {
	n := len(a)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i] * b[j]
		}
	}
	return result
}

func transposeM(m [][]float64) [][]float64 {
	n := len(m)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func matMult(a, b [][]float64) [][]float64 {
	n := len(a)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func lineSearch(f ObjectiveFunc, x, p []float64) float64 {
	alpha := 1.0
	for i := 0; i < 20; i++ {
		xNew := make([]float64, len(x))
		for j := range x {
			xNew[j] = x[j] + alpha*p[j]
		}
		if f(xNew) < f(x) {
			return alpha
		}
		alpha *= 0.5
	}
	return alpha
}

func dotProduct(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func matVecMult(m [][]float64, v []float64) []float64 {
	result := make([]float64, len(v))
	for i := range m {
		for j := range v {
			result[i] += m[i][j] * v[j]
		}
	}
	return result
}

func norm(v []float64) float64 {
	return sqrtO(dotProduct(v, v))
}

func sqrtO(x float64) float64 {
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
// Created by: MOUAAD IDOUFKIR
// << The universe runs on equations. We just translate them >>
