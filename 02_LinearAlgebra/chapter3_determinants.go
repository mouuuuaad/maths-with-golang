package linearalgebra

func (m Matrix) Determinant() float64 {
	if !m.IsSquare() {
		return 0
	}
	n := m.Rows()
	if n == 1 {
		return m[0][0]
	}
	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	temp := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		copy(temp[i], m[i])
	}
	det := 1.0
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if absV(temp[j][i]) > absV(temp[pivot][i]) {
				pivot = j
			}
		}
		if pivot != i {
			temp[i], temp[pivot] = temp[pivot], temp[i]
			det = -det
		}
		if absV(temp[i][i]) < 1e-12 {
			return 0
		}
		det *= temp[i][i]
		for j := i + 1; j < n; j++ {
			factor := temp[j][i] / temp[i][i]
			for k := i; k < n; k++ {
				temp[j][k] -= factor * temp[i][k]
			}
		}
	}
	return det
}

func (m Matrix) Minor(row, col int) Matrix {
	n := m.Rows()
	result := NewMatrix(n-1, n-1)
	r := 0
	for i := 0; i < n; i++ {
		if i == row {
			continue
		}
		c := 0
		for j := 0; j < n; j++ {
			if j == col {
				continue
			}
			result[r][c] = m[i][j]
			c++
		}
		r++
	}
	return result
}

func (m Matrix) Cofactor(row, col int) float64 {
	sign := 1.0
	if (row+col)%2 != 0 {
		sign = -1.0
	}
	return sign * m.Minor(row, col).Determinant()
}

func (m Matrix) CofactorMatrix() Matrix {
	n := m.Rows()
	result := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = m.Cofactor(i, j)
		}
	}
	return result
}

func (m Matrix) Adjugate() Matrix {
	return m.CofactorMatrix().Transpose()
}

func (m Matrix) Inverse() Matrix {
	det := m.Determinant()
	if absV(det) < 1e-12 {
		return nil
	}
	return m.Adjugate().Scale(1.0 / det)
}

func (m Matrix) Rank() int {
	n := m.Rows()
	c := m.Cols()
	temp := NewMatrix(n, c)
	for i := 0; i < n; i++ {
		copy(temp[i], m[i])
	}
	rank := 0
	for col := 0; col < c && rank < n; col++ {
		pivot := -1
		for row := rank; row < n; row++ {
			if absV(temp[row][col]) > 1e-9 {
				pivot = row
				break
			}
		}
		if pivot == -1 {
			continue
		}
		temp[rank], temp[pivot] = temp[pivot], temp[rank]
		for row := rank + 1; row < n; row++ {
			factor := temp[row][col] / temp[rank][col]
			for k := col; k < c; k++ {
				temp[row][k] -= factor * temp[rank][k]
			}
		}
		rank++
	}
	return rank
}

func (m Matrix) Nullity() int {
	return m.Cols() - m.Rank()
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
