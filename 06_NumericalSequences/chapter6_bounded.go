package sequences

func IsBounded(seq Sequence, n int) (bool, float64, float64) {
	lower := seq(1)
	upper := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val < lower {
			lower = val
		}
		if val > upper {
			upper = val
		}
	}
	return true, lower, upper
}

func IsBoundedAbove(seq Sequence, n int) (bool, float64) {
	upper := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val > upper {
			upper = val
		}
	}
	return true, upper
}

func IsBoundedBelow(seq Sequence, n int) (bool, float64) {
	lower := seq(1)
	for i := 2; i <= n; i++ {
		val := seq(i)
		if val < lower {
			lower = val
		}
	}
	return true, lower
}

func Supremum(seq Sequence, n int) float64 {
	_, upper := IsBoundedAbove(seq, n)
	return upper
}

func Infimum(seq Sequence, n int) float64 {
	_, lower := IsBoundedBelow(seq, n)
	return lower
}

func LimSup(seq Sequence, n int) float64 {
	maxVals := make([]float64, n)
	for i := 1; i <= n; i++ {
		max := seq(i)
		for j := i; j <= n; j++ {
			if seq(j) > max {
				max = seq(j)
			}
		}
		maxVals[i-1] = max
	}
	return maxVals[n-1]
}

func LimInf(seq Sequence, n int) float64 {
	minVals := make([]float64, n)
	for i := 1; i <= n; i++ {
		min := seq(i)
		for j := i; j <= n; j++ {
			if seq(j) < min {
				min = seq(j)
			}
		}
		minVals[i-1] = min
	}
	return minVals[n-1]
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
