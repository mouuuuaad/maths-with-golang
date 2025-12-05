package algebra

type Ring interface {
	Zero() interface{}
	One() interface{}
	Add(a, b interface{}) interface{}
	Multiply(a, b interface{}) interface{}
	Negate(a interface{}) interface{}
}

type IntegerRing struct{}

func (r IntegerRing) Zero() interface{} {
	return 0
}

func (r IntegerRing) One() interface{} {
	return 1
}

func (r IntegerRing) Add(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func (r IntegerRing) Multiply(a, b interface{}) interface{} {
	return a.(int) * b.(int)
}

func (r IntegerRing) Negate(a interface{}) interface{} {
	return -a.(int)
}

type ModularRing struct {
	N int
}

func (r ModularRing) Zero() interface{} {
	return 0
}

func (r ModularRing) One() interface{} {
	return 1
}

func (r ModularRing) Add(a, b interface{}) interface{} {
	return (a.(int) + b.(int)) % r.N
}

func (r ModularRing) Multiply(a, b interface{}) interface{} {
	return (a.(int) * b.(int)) % r.N
}

func (r ModularRing) Negate(a interface{}) interface{} {
	return (r.N - a.(int)) % r.N
}

func IsCommutativeRing(R Ring, elements []interface{}) bool {
	for _, a := range elements {
		for _, b := range elements {
			if R.Multiply(a, b) != R.Multiply(b, a) {
				return false
			}
		}
	}
	return true
}

func IsIntegralDomain(R Ring, elements []interface{}) bool {
	for _, a := range elements {
		if a == R.Zero() {
			continue
		}
		for _, b := range elements {
			if b == R.Zero() {
				continue
			}
			if R.Multiply(a, b) == R.Zero() {
				return false
			}
		}
	}
	return true
}

func Units(R ModularRing) []int {
	result := []int{}
	for i := 1; i < R.N; i++ {
		if gcdA(i, R.N) == 1 {
			result = append(result, i)
		}
	}
	return result
}

func gcdA(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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
