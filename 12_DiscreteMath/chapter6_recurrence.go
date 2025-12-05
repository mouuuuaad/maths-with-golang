package discrete

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func LinearRecurrence(coeffs []int, initial []int, n int) int {
	k := len(coeffs)
	if n < k {
		return initial[n]
	}
	vals := make([]int, n+1)
	copy(vals, initial)
	for i := k; i <= n; i++ {
		for j := 0; j < k; j++ {
			vals[i] += coeffs[j] * vals[i-1-j]
		}
	}
	return vals[n]
}

func SolveRecurrence(a, b, c int) (r1, r2 float64) {
	discriminant := float64(b*b - 4*a*c)
	if discriminant >= 0 {
		sqrtD := sqrtDM(discriminant)
		r1 = (-float64(b) + sqrtD) / (2 * float64(a))
		r2 = (-float64(b) - sqrtD) / (2 * float64(a))
	}
	return
}

func sqrtDM(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}

func GeneratingFunction(coeffs []int, x float64, terms int) float64 {
	result := 0.0
	xPow := 1.0
	for i := 0; i < terms && i < len(coeffs); i++ {
		result += float64(coeffs[i]) * xPow
		xPow *= x
	}
	return result
}

func PartitionNumber(n int) int {
	p := make([]int, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			p[j] += p[j-i]
		}
	}
	return p[n]
}

func CatalanNumber(n int) int {
	if n <= 1 {
		return 1
	}
	cat := make([]int, n+1)
	cat[0], cat[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			cat[i] += cat[j] * cat[i-1-j]
		}
	}
	return cat[n]
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
