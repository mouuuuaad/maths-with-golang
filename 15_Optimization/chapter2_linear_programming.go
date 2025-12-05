package optimization

func Simplex(c []float64, A [][]float64, b []float64) ([]float64, float64) {
	m, n := len(A), len(c)
	tableau := make([][]float64, m+1)
	for i := 0; i <= m; i++ {
		tableau[i] = make([]float64, n+m+1)
	}
	for i := 0; i < m; i++ {
		copy(tableau[i][:n], A[i])
		tableau[i][n+i] = 1
		tableau[i][n+m] = b[i]
	}
	for j := 0; j < n; j++ {
		tableau[m][j] = -c[j]
	}
	for {
		pivot := -1
		for j := 0; j < n+m; j++ {
			if tableau[m][j] < -1e-10 {
				pivot = j
				break
			}
		}
		if pivot == -1 {
			break
		}
		row := -1
		minRatio := 1e18
		for i := 0; i < m; i++ {
			if tableau[i][pivot] > 1e-10 {
				ratio := tableau[i][n+m] / tableau[i][pivot]
				if ratio < minRatio {
					minRatio = ratio
					row = i
				}
			}
		}
		if row == -1 {
			return nil, -1e18
		}
		pivotVal := tableau[row][pivot]
		for j := range tableau[row] {
			tableau[row][j] /= pivotVal
		}
		for i := 0; i <= m; i++ {
			if i != row {
				factor := tableau[i][pivot]
				for j := range tableau[i] {
					tableau[i][j] -= factor * tableau[row][j]
				}
			}
		}
	}
	x := make([]float64, n)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if absO(tableau[i][j]-1) < 1e-10 {
				x[j] = tableau[i][n+m]
				break
			}
		}
	}
	return x, tableau[m][n+m]
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
