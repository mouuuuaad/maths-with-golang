package limits

func PointwiseContinuity(f Function, a float64) bool {
	lim, ok := Limit(f, a)
	if !ok {
		return false
	}
	return absLim(lim-f(a)) < 1e-9
}

func SequentialContinuity(f Function, seq Sequence, a float64, n int) bool {
	for i := 1; i <= n; i++ {
		if absLim(f(seq(i))-f(a)) > 1e-6 {
			return false
		}
	}
	return true
}

func UniformContinuity(f Function, a, b, epsilon float64) float64 {
	for delta := 0.1; delta > 1e-10; delta /= 2 {
		ok := true
		for x := a; x <= b; x += delta / 10 {
			for y := x; y <= x+delta && y <= b; y += delta / 100 {
				if absLim(f(x)-f(y)) >= epsilon {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			return delta
		}
	}
	return 0
}

func LipschitzConstant(f Function, a, b float64, n int) float64 {
	maxRatio := 0.0
	step := (b - a) / float64(n)
	for x := a; x <= b; x += step {
		for y := x + step; y <= b; y += step {
			if x == y {
				continue
			}
			ratio := absLim(f(x)-f(y)) / absLim(x-y)
			if ratio > maxRatio {
				maxRatio = ratio
			}
		}
	}
	return maxRatio
}

func IsLipschitz(f Function, a, b float64, L float64, n int) bool {
	return LipschitzConstant(f, a, b, n) <= L
}

func IntermediateValue(f Function, a, b, val float64) (float64, bool) {
	fa := f(a)
	fb := f(b)
	if (fa < val && fb < val) || (fa > val && fb > val) {
		return 0, false
	}
	for i := 0; i < 100; i++ {
		mid := (a + b) / 2
		fmid := f(mid)
		if absLim(fmid-val) < 1e-9 {
			return mid, true
		}
		if (fa < val && fmid < val) || (fa > val && fmid > val) {
			a = mid
			fa = fmid
		} else {
			b = mid
			fb = fmid
		}
	}
	return (a + b) / 2, true
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
