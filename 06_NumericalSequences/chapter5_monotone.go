package sequences

func IsIncreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) < seq(i) {
			return false
		}
	}
	return true
}

func IsDecreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) > seq(i) {
			return false
		}
	}
	return true
}

func IsMonotone(seq Sequence, n int) bool {
	return IsIncreasing(seq, n) || IsDecreasing(seq, n)
}

func IsStrictlyIncreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) <= seq(i) {
			return false
		}
	}
	return true
}

func IsStrictlyDecreasing(seq Sequence, n int) bool {
	for i := 1; i < n; i++ {
		if seq(i+1) >= seq(i) {
			return false
		}
	}
	return true
}

func EventuallyMonotone(seq Sequence, n, start int) bool {
	subseq := func(k int) float64 {
		return seq(k + start)
	}
	return IsMonotone(subseq, n-start)
}

func MonotoneConvergenceTheorem(seq Sequence, n int) (float64, bool) {
	if !IsMonotone(seq, n) {
		return 0, false
	}
	bounded, _, _ := IsBounded(seq, n)
	if !bounded {
		return 0, false
	}
	return IsConvergent(seq)
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
