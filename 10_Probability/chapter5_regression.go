package probability

func LinearRegression(x, y []float64) (slope, intercept float64) {
	n := float64(len(x))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}
	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept = (sumY - slope*sumX) / n
	return
}

func RSquared(x, y []float64) float64 {
	slope, intercept := LinearRegression(x, y)
	yMean := Mean(y)
	ssRes, ssTot := 0.0, 0.0
	for i := range x {
		predicted := slope*x[i] + intercept
		ssRes += (y[i] - predicted) * (y[i] - predicted)
		ssTot += (y[i] - yMean) * (y[i] - yMean)
	}
	if ssTot == 0 {
		return 0
	}
	return 1 - ssRes/ssTot
}

func PolynomialRegression(x, y []float64, degree int) []float64 {
	n := len(x)
	m := degree + 1
	X := make([][]float64, n)
	for i := 0; i < n; i++ {
		X[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			X[i][j] = powerP(x[i], float64(j))
		}
	}
	Xt := transpose(X)
	XtX := multiply(Xt, X)
	Xty := multiplyVector(Xt, y)
	return solveSystem(XtX, Xty)
}

func transpose(m [][]float64) [][]float64 {
	if len(m) == 0 {
		return nil
	}
	rows, cols := len(m), len(m[0])
	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
		for j := range result[i] {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func multiply(a, b [][]float64) [][]float64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	rows, cols, inner := len(a), len(b[0]), len(b)
	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
		for j := range result[i] {
			for k := 0; k < inner; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func multiplyVector(m [][]float64, v []float64) []float64 {
	result := make([]float64, len(m))
	for i := range m {
		for j := range v {
			result[i] += m[i][j] * v[j]
		}
	}
	return result
}

func solveSystem(A [][]float64, b []float64) []float64 {
	n := len(A)
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, n+1)
		copy(aug[i], A[i])
		aug[i][n] = b[i]
	}
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if absP(aug[j][i]) > absP(aug[pivot][i]) {
				pivot = j
			}
		}
		aug[i], aug[pivot] = aug[pivot], aug[i]
		for j := i + 1; j < n; j++ {
			factor := aug[j][i] / aug[i][i]
			for k := i; k <= n; k++ {
				aug[j][k] -= factor * aug[i][k]
			}
		}
	}
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += aug[i][j] * x[j]
		}
		x[i] = (aug[i][n] - sum) / aug[i][i]
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
