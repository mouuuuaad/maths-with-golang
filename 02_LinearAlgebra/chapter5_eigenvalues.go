package linearalgebra

func PowerIteration(A Matrix, maxIter int, tol float64) (float64, Vector) {
	n := A.Rows()
	v := make(Vector, n)
	for i := range v {
		v[i] = 1.0
	}
	v = v.Normalize()
	eigenvalue := 0.0
	for iter := 0; iter < maxIter; iter++ {
		vNew := A.MultiplyVector(v)
		newEigenvalue := vNew.Dot(v)
		vNew = vNew.Normalize()
		if absV(newEigenvalue-eigenvalue) < tol {
			return newEigenvalue, vNew
		}
		eigenvalue = newEigenvalue
		v = vNew
	}
	return eigenvalue, v
}

func InversePowerIteration(A Matrix, shift float64, maxIter int, tol float64) (float64, Vector) {
	n := A.Rows()
	shiftedA := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			shiftedA[i][j] = A[i][j]
			if i == j {
				shiftedA[i][j] -= shift
			}
		}
	}
	v := make(Vector, n)
	for i := range v {
		v[i] = 1.0
	}
	v = v.Normalize()
	eigenvalue := 0.0
	for iter := 0; iter < maxIter; iter++ {
		vNew := GaussElimination(shiftedA, v)
		if vNew == nil {
			break
		}
		newEigenvalue := vNew.Dot(v)
		vNew = vNew.Normalize()
		if absV(newEigenvalue-eigenvalue) < tol {
			return shift + 1.0/newEigenvalue, vNew
		}
		eigenvalue = newEigenvalue
		v = vNew
	}
	return shift + 1.0/eigenvalue, v
}

func RayleighQuotient(A Matrix, v Vector) float64 {
	Av := A.MultiplyVector(v)
	return Av.Dot(v) / v.Dot(v)
}

func CharacteristicPolynomial2x2(A Matrix) (float64, float64, float64) {
	if A.Rows() != 2 || A.Cols() != 2 {
		return 0, 0, 0
	}
	a := 1.0
	b := -(A[0][0] + A[1][1])
	c := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	return a, b, c
}

func Eigenvalues2x2(A Matrix) (float64, float64) {
	_, b, c := CharacteristicPolynomial2x2(A)
	disc := b*b - 4*c
	if disc < 0 {
		return 0, 0
	}
	sqrtD := sqrtV(disc)
	return (-b + sqrtD) / 2, (-b - sqrtD) / 2
}

func QRAlgorithm(A Matrix, maxIter int, tol float64) Vector {
	n := A.Rows()
	Ak := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		copy(Ak[i], A[i])
	}
	for iter := 0; iter < maxIter; iter++ {
		Q, R := QRDecomposition(Ak)
		Ak = R.Multiply(Q)
		offDiag := 0.0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					offDiag += absV(Ak[i][j])
				}
			}
		}
		if offDiag < tol {
			break
		}
	}
	eigenvalues := make(Vector, n)
	for i := 0; i < n; i++ {
		eigenvalues[i] = Ak[i][i]
	}
	return eigenvalues
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
