package linearalgebra

func GaussElimination(A Matrix, b Vector) Vector {
	n := A.Rows()
	if n != len(b) {
		return nil
	}
	aug := NewMatrix(n, n+1)
	for i := 0; i < n; i++ {
		copy(aug[i], A[i])
		aug[i][n] = b[i]
	}
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if absV(aug[j][i]) > absV(aug[pivot][i]) {
				pivot = j
			}
		}
		aug[i], aug[pivot] = aug[pivot], aug[i]
		if absV(aug[i][i]) < 1e-12 {
			return nil
		}
		for j := i + 1; j < n; j++ {
			factor := aug[j][i] / aug[i][i]
			for k := i; k <= n; k++ {
				aug[j][k] -= factor * aug[i][k]
			}
		}
	}
	x := make(Vector, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += aug[i][j] * x[j]
		}
		x[i] = (aug[i][n] - sum) / aug[i][i]
	}
	return x
}

func GaussJordan(A Matrix, b Vector) Vector {
	n := A.Rows()
	if n != len(b) {
		return nil
	}
	aug := NewMatrix(n, n+1)
	for i := 0; i < n; i++ {
		copy(aug[i], A[i])
		aug[i][n] = b[i]
	}
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if absV(aug[j][i]) > absV(aug[pivot][i]) {
				pivot = j
			}
		}
		aug[i], aug[pivot] = aug[pivot], aug[i]
		if absV(aug[i][i]) < 1e-12 {
			return nil
		}
		scale := aug[i][i]
		for k := i; k <= n; k++ {
			aug[i][k] /= scale
		}
		for j := 0; j < n; j++ {
			if j != i {
				factor := aug[j][i]
				for k := i; k <= n; k++ {
					aug[j][k] -= factor * aug[i][k]
				}
			}
		}
	}
	x := make(Vector, n)
	for i := 0; i < n; i++ {
		x[i] = aug[i][n]
	}
	return x
}

func CramersRule(A Matrix, b Vector) Vector {
	det := A.Determinant()
	if absV(det) < 1e-12 {
		return nil
	}
	n := A.Rows()
	x := make(Vector, n)
	for j := 0; j < n; j++ {
		Aj := NewMatrix(n, n)
		for i := 0; i < n; i++ {
			for k := 0; k < n; k++ {
				if k == j {
					Aj[i][k] = b[i]
				} else {
					Aj[i][k] = A[i][k]
				}
			}
		}
		x[j] = Aj.Determinant() / det
	}
	return x
}

func JacobiIteration(A Matrix, b Vector, x0 Vector, maxIter int, tol float64) Vector {
	n := A.Rows()
	x := make(Vector, n)
	copy(x, x0)
	xNew := make(Vector, n)
	for iter := 0; iter < maxIter; iter++ {
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += A[i][j] * x[j]
				}
			}
			xNew[i] = (b[i] - sum) / A[i][i]
		}
		diff := 0.0
		for i := 0; i < n; i++ {
			diff += absV(xNew[i] - x[i])
		}
		copy(x, xNew)
		if diff < tol {
			break
		}
	}
	return x
}

func GaussSeidel(A Matrix, b Vector, x0 Vector, maxIter int, tol float64) Vector {
	n := A.Rows()
	x := make(Vector, n)
	copy(x, x0)
	for iter := 0; iter < maxIter; iter++ {
		prevX := make(Vector, n)
		copy(prevX, x)
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += A[i][j] * x[j]
				}
			}
			x[i] = (b[i] - sum) / A[i][i]
		}
		diff := 0.0
		for i := 0; i < n; i++ {
			diff += absV(x[i] - prevX[i])
		}
		if diff < tol {
			break
		}
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
