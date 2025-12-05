package sequences

func SeriesSum(seq Sequence, n int) float64 {
	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += seq(i)
	}
	return sum
}

func PartialSums(seq Sequence, n int) []float64 {
	result := make([]float64, n+1)
	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += seq(i)
		result[i] = sum
	}
	return result
}

func AlternatingSeries(seq Sequence, n int) float64 {
	sum := 0.0
	sign := 1.0
	for i := 0; i <= n; i++ {
		sum += sign * absS(seq(i))
		sign = -sign
	}
	return sum
}

func PowerSeries(coeffs []float64, x float64, n int) float64 {
	sum := 0.0
	xPow := 1.0
	for i := 0; i <= n && i < len(coeffs); i++ {
		sum += coeffs[i] * xPow
		xPow *= x
	}
	return sum
}

func TelescipingSeries(seq Sequence, n int) float64 {
	return seq(0) - seq(n+1)
}

func HarmonicSeries(n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	return sum
}

func GeometricSeries(a, r float64, n int) float64 {
	if absS(r) >= 1 {
		sum := 0.0
		term := a
		for i := 0; i <= n; i++ {
			sum += term
			term *= r
		}
		return sum
	}
	return a * (1 - powerS(r, float64(n+1))) / (1 - r)
}

func SeriesConverges(seq Sequence) bool {
	sums := PartialSums(seq, 1000)
	epsilon := 1e-7
	for i := 900; i < 1000; i++ {
		if absS(sums[i]-sums[i-1]) > epsilon {
			return false
		}
	}
	return true
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
