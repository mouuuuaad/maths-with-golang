package derivatives

func CriticalPoints(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	for x := a; x <= b; x += step {
		d := Derivative(f, x)
		if absD(d) < 1e-6 {
			result = append(result, x)
		}
	}
	return result
}

func InflectionPoints(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	step := (b - a) / float64(n)
	prev := SecondDerivative(f, a)
	for x := a + step; x <= b; x += step {
		curr := SecondDerivative(f, x)
		if prev*curr < 0 {
			result = append(result, x)
		}
		prev = curr
	}
	return result
}

func LocalMaxima(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	criticals := CriticalPoints(f, a, b, n)
	for _, x := range criticals {
		if SecondDerivative(f, x) < -1e-9 {
			result = append(result, x)
		}
	}
	return result
}

func LocalMinima(f Function, a, b float64, n int) []float64 {
	result := []float64{}
	criticals := CriticalPoints(f, a, b, n)
	for _, x := range criticals {
		if SecondDerivative(f, x) > 1e-9 {
			result = append(result, x)
		}
	}
	return result
}

func GlobalMaximum(f Function, a, b float64, n int) (float64, float64) {
	maxX := a
	maxY := f(a)
	step := (b - a) / float64(n)
	for x := a; x <= b; x += step {
		y := f(x)
		if y > maxY {
			maxY = y
			maxX = x
		}
	}
	return maxX, maxY
}

func GlobalMinimum(f Function, a, b float64, n int) (float64, float64) {
	minX := a
	minY := f(a)
	step := (b - a) / float64(n)
	for x := a; x <= b; x += step {
		y := f(x)
		if y < minY {
			minY = y
			minX = x
		}
	}
	return minX, minY
}

func IsIncreasingAt(f Function, x float64) bool {
	return Derivative(f, x) > 1e-9
}

func IsDecreasingAt(f Function, x float64) bool {
	return Derivative(f, x) < -1e-9
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
