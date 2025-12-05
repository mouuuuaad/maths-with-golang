package linearalgebra

func LUDecomposition(A Matrix) (Matrix, Matrix) {
	n := A.Rows()
	if n != A.Cols() {
		return nil, nil
	}
	L := NewMatrix(n, n)
	U := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for k := i; k < n; k++ {
			sum := 0.0
			for j := 0; j < i; j++ {
				sum += L[i][j] * U[j][k]
			}
			U[i][k] = A[i][k] - sum
		}
		for k := i; k < n; k++ {
			if i == k {
				L[i][i] = 1
			} else {
				sum := 0.0
				for j := 0; j < i; j++ {
					sum += L[k][j] * U[j][i]
				}
				if absV(U[i][i]) < 1e-12 {
					L[k][i] = 0
				} else {
					L[k][i] = (A[k][i] - sum) / U[i][i]
				}
			}
		}
	}
	return L, U
}

func QRDecomposition(A Matrix) (Matrix, Matrix) {
	m := A.Rows()
	n := A.Cols()
	Q := NewMatrix(m, m)
	R := NewMatrix(m, n)
	for i := 0; i < m; i++ {
		Q[i][i] = 1
	}
	for i := 0; i < m; i++ {
		copy(R[i], A[i])
	}
	for k := 0; k < n && k < m-1; k++ {
		x := make(Vector, m-k)
		normX := 0.0
		for i := k; i < m; i++ {
			x[i-k] = R[i][k]
			normX += x[i-k] * x[i-k]
		}
		normX = sqrtV(normX)
		if normX == 0 {
			continue
		}
		sign := 1.0
		if x[0] < 0 {
			sign = -1.0
		}
		x[0] += sign * normX
		normV := 0.0
		for _, val := range x {
			normV += val * val
		}
		normV = sqrtV(normV)
		for i := range x {
			x[i] /= normV
		}
		for j := k; j < n; j++ {
			dot := 0.0
			for i := k; i < m; i++ {
				dot += x[i-k] * R[i][j]
			}
			for i := k; i < m; i++ {
				R[i][j] -= 2 * x[i-k] * dot
			}
		}
		for i := 0; i < m; i++ {
			dot := 0.0
			for j := k; j < m; j++ {
				dot += Q[i][j] * x[j-k]
			}
			for j := k; j < m; j++ {
				Q[i][j] -= 2 * dot * x[j-k]
			}
		}
	}
	return Q, R
}

func CholeskyDecomposition(A Matrix) Matrix {
	n := A.Rows()
	if n != A.Cols() {
		return nil
	}
	L := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sum := 0.0
			for k := 0; k < j; k++ {
				sum += L[i][k] * L[j][k]
			}
			if i == j {
				val := A[i][i] - sum
				if val <= 0 {
					return nil
				}
				L[i][j] = sqrtV(val)
			} else {
				L[i][j] = (A[i][j] - sum) / L[j][j]
			}
		}
	}
	return L
}

func GramSchmidt(vectors []Vector) []Vector {
	n := len(vectors)
	if n == 0 {
		return nil
	}
	orthogonal := make([]Vector, n)
	for i := 0; i < n; i++ {
		orthogonal[i] = make(Vector, len(vectors[0]))
		copy(orthogonal[i], vectors[i])
		for j := 0; j < i; j++ {
			proj := orthogonal[i].ProjectOnto(orthogonal[j])
			orthogonal[i] = orthogonal[i].Subtract(proj)
		}
	}
	orthonormal := make([]Vector, n)
	for i := 0; i < n; i++ {
		orthonormal[i] = orthogonal[i].Normalize()
	}
	return orthonormal
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
