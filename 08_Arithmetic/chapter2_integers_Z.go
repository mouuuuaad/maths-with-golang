package arithmetic

type Integer int64

func AbsI(n Integer) Integer {
	if n < 0 {
		return -n
	}
	return n
}

func SignI(n Integer) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}

func IsDivisible(n, d Integer) bool {
	if d == 0 {
		return false
	}
	return n%d == 0
}

func EuclideanDivision(n, d Integer) (Integer, Integer) {
	if d == 0 {
		return 0, 0
	}
	q := n / d
	r := n % d
	if r < 0 {
		if d > 0 {
			q--
			r += d
		} else {
			q++
			r -= d
		}
	}
	return q, r
}

func GCD(a, b Integer) Integer {
	a = AbsI(a)
	b = AbsI(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b Integer) Integer {
	if a == 0 || b == 0 {
		return 0
	}
	return AbsI(a*b) / GCD(a, b)
}

func ExtendedGCD(a, b Integer) (Integer, Integer, Integer) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := ExtendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

func ModInverse(a, m Integer) Integer {
	g, x, _ := ExtendedGCD(a, m)
	if g != 1 {
		return 0
	}
	return ((x % m) + m) % m
}

func ChineseRemainderTheorem(remainders, moduli []Integer) Integer {
	if len(remainders) != len(moduli) {
		return 0
	}
	M := Integer(1)
	for _, m := range moduli {
		M *= m
	}
	result := Integer(0)
	for i := range remainders {
		Mi := M / moduli[i]
		yi := ModInverse(Mi, moduli[i])
		result += remainders[i] * Mi * yi
	}
	return ((result % M) + M) % M
}

func BinaryGCD(a, b Integer) Integer {
	a = AbsI(a)
	b = AbsI(b)
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	shift := 0
	for (a|b)&1 == 0 {
		a >>= 1
		b >>= 1
		shift++
	}
	for a&1 == 0 {
		a >>= 1
	}
	for {
		for b&1 == 0 {
			b >>= 1
		}
		if a > b {
			a, b = b, a
		}
		b -= a
		if b == 0 {
			break
		}
	}
	return a << shift
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
