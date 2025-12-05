package algebra

type Field interface {
	Ring
	Inverse(a interface{}) interface{}
}

type FiniteField struct {
	P int
}

func (f FiniteField) Zero() interface{} {
	return 0
}

func (f FiniteField) One() interface{} {
	return 1
}

func (f FiniteField) Add(a, b interface{}) interface{} {
	return (a.(int) + b.(int)) % f.P
}

func (f FiniteField) Multiply(a, b interface{}) interface{} {
	return (a.(int) * b.(int)) % f.P
}

func (f FiniteField) Negate(a interface{}) interface{} {
	return (f.P - a.(int)) % f.P
}

func (f FiniteField) Inverse(a interface{}) interface{} {
	ai := a.(int)
	if ai == 0 {
		return 0
	}
	return modPowA(ai, f.P-2, f.P)
}

func modPowA(base, exp, mod int) int {
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}
	return result
}

type RationalField struct{}

func (f RationalField) Zero() interface{} {
	return [2]int{0, 1}
}

func (f RationalField) One() interface{} {
	return [2]int{1, 1}
}

func (f RationalField) Add(a, b interface{}) interface{} {
	r1 := a.([2]int)
	r2 := b.([2]int)
	num := r1[0]*r2[1] + r2[0]*r1[1]
	den := r1[1] * r2[1]
	g := gcdA(absA(num), absA(den))
	return [2]int{num / g, den / g}
}

func (f RationalField) Multiply(a, b interface{}) interface{} {
	r1 := a.([2]int)
	r2 := b.([2]int)
	num := r1[0] * r2[0]
	den := r1[1] * r2[1]
	g := gcdA(absA(num), absA(den))
	return [2]int{num / g, den / g}
}

func (f RationalField) Negate(a interface{}) interface{} {
	r := a.([2]int)
	return [2]int{-r[0], r[1]}
}

func (f RationalField) Inverse(a interface{}) interface{} {
	r := a.([2]int)
	if r[0] == 0 {
		return [2]int{0, 1}
	}
	if r[0] < 0 {
		return [2]int{-r[1], -r[0]}
	}
	return [2]int{r[1], r[0]}
}

func absA(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
