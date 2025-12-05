package limits

func CauchyCriterion(seq Sequence, n int) bool {
	return CauchySequence(seq, n, 1e-7)
}

func SubsequenceLimit(seq Sequence, indices func(int) int) (float64, bool) {
	subseq := func(n int) float64 {
		return seq(indices(n))
	}
	return LimitSequence(subseq)
}

func ConvergentSubsequence(seq Sequence, n int) func(int) float64 {
	return func(k int) float64 {
		return seq(k * 2)
	}
}

func NestedIntervals(lower, upper Sequence, n int) (float64, bool) {
	for i := 1; i <= n; i++ {
		l := lower(i)
		u := upper(i)
		if l > u {
			return 0, false
		}
	}
	lLim, lOk := LimitSequence(lower)
	uLim, uOk := LimitSequence(upper)
	if lOk && uOk && absLim(lLim-uLim) < 1e-7 {
		return lLim, true
	}
	return 0, false
}

func DedekindsConstruction(lower, upper []float64) float64 {
	if len(lower) == 0 || len(upper) == 0 {
		return 0
	}
	maxLower := lower[0]
	for _, v := range lower {
		if v > maxLower {
			maxLower = v
		}
	}
	minUpper := upper[0]
	for _, v := range upper {
		if v < minUpper {
			minUpper = v
		}
	}
	return (maxLower + minUpper) / 2
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
