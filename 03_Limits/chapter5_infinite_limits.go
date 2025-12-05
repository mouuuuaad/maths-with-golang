package limits

func PInfinity(f Function, a float64) bool {
	left, _ := LimitLeft(f, a)
	right, _ := LimitRight(f, a)
	return left > 1e10 || right > 1e10
}

func NInfinity(f Function, a float64) bool {
	left, _ := LimitLeft(f, a)
	right, _ := LimitRight(f, a)
	return left < -1e10 || right < -1e10
}

func DetectAsymptote(f Function, a float64) string {
	left, lOk := LimitLeft(f, a)
	right, rOk := LimitRight(f, a)
	if !lOk && !rOk {
		return "undefined"
	}
	if absLim(left) > 1e10 || absLim(right) > 1e10 {
		return "vertical"
	}
	return "none"
}

func HorizontalAsymptote(f Function) (float64, bool) {
	return LimitPosInfinity(f)
}

func ObliqueAsymptote(f Function) (float64, float64, bool) {
	m := func(x float64) float64 {
		return f(x) / x
	}
	slope, mOk := LimitPosInfinity(m)
	if !mOk {
		return 0, 0, false
	}
	c := func(x float64) float64 {
		return f(x) - slope*x
	}
	intercept, cOk := LimitPosInfinity(c)
	if !cOk {
		return 0, 0, false
	}
	return slope, intercept, true
}

func FindVerticalAsymptotes(f Function, start, end float64, n int) []float64 {
	result := []float64{}
	step := (end - start) / float64(n)
	for x := start; x <= end; x += step {
		val := f(x)
		if val != val || absLim(val) > 1e10 {
			result = append(result, x)
		}
	}
	return result
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
