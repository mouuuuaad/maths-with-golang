package linearalgebra

func LeastSquares(A Matrix, b Vector) Vector {
	At := A.Transpose()
	AtA := At.Multiply(A)
	Atb := At.MultiplyVector(b)
	return GaussElimination(AtA, Atb)
}

func PolynomialFit(x, y Vector, degree int) Vector {
	n := len(x)
	A := NewMatrix(n, degree+1)
	for i := 0; i < n; i++ {
		val := 1.0
		for j := 0; j <= degree; j++ {
			A[i][j] = val
			val *= x[i]
		}
	}
	return LeastSquares(A, y)
}

func LinearRegression(x, y Vector) (float64, float64) {
	n := float64(len(x))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}
	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n
	return slope, intercept
}

func PageRank(adjacency Matrix, damping float64, maxIter int) Vector {
	n := adjacency.Rows()
	outDegree := make([]float64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			outDegree[i] += adjacency[i][j]
		}
	}
	M := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if outDegree[j] > 0 {
				M[i][j] = adjacency[j][i] / outDegree[j]
			}
		}
	}
	pr := make(Vector, n)
	for i := range pr {
		pr[i] = 1.0 / float64(n)
	}
	for iter := 0; iter < maxIter; iter++ {
		newPr := make(Vector, n)
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				sum += M[i][j] * pr[j]
			}
			newPr[i] = (1-damping)/float64(n) + damping*sum
		}
		pr = newPr
	}
	return pr
}

func MatrixExponential(A Matrix, terms int) Matrix {
	n := A.Rows()
	result := Identity(n)
	term := Identity(n)
	for k := 1; k <= terms; k++ {
		term = term.Multiply(A).Scale(1.0 / float64(k))
		result = result.Add(term)
	}
	return result
}

func KroneckerProduct(A, B Matrix) Matrix {
	m1, n1 := A.Rows(), A.Cols()
	m2, n2 := B.Rows(), B.Cols()
	result := NewMatrix(m1*m2, n1*n2)
	for i := 0; i < m1; i++ {
		for j := 0; j < n1; j++ {
			for k := 0; k < m2; k++ {
				for l := 0; l < n2; l++ {
					result[i*m2+k][j*n2+l] = A[i][j] * B[k][l]
				}
			}
		}
	}
	return result
}

func HadamardProduct(A, B Matrix) Matrix {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return nil
	}
	result := NewMatrix(A.Rows(), A.Cols())
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			result[i][j] = A[i][j] * B[i][j]
		}
	}
	return result
}

func FrobeniusNorm(A Matrix) float64 {
	sum := 0.0
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			sum += A[i][j] * A[i][j]
		}
	}
	return sqrtV(sum)
}

func ConditionNumber(A Matrix) float64 {
	normA := FrobeniusNorm(A)
	inv := A.Inverse()
	if inv == nil {
		return 1e308
	}
	normInv := FrobeniusNorm(inv)
	return normA * normInv
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
