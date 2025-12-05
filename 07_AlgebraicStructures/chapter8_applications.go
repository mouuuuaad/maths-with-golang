package algebra

type RSA struct {
	N int
	E int
	D int
}

func GenerateRSAKeys(p, q, e int) RSA {
	n := p * q
	phi := (p - 1) * (q - 1)
	d := modInverseA(e, phi)
	return RSA{N: n, E: e, D: d}
}

func (r RSA) Encrypt(m int) int {
	return modPowA(m, r.E, r.N)
}

func (r RSA) Decrypt(c int) int {
	return modPowA(c, r.D, r.N)
}

func modInverseA(a, m int) int {
	g, x, _ := extendedGCDA(a, m)
	if g != 1 {
		return 0
	}
	return ((x % m) + m) % m
}

func extendedGCDA(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extendedGCDA(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

type EllipticCurve struct {
	A int
	B int
	P int
}

type ECPoint struct {
	X        int
	Y        int
	Infinity bool
}

func (ec EllipticCurve) Add(p1, p2 ECPoint) ECPoint {
	if p1.Infinity {
		return p2
	}
	if p2.Infinity {
		return p1
	}
	if p1.X == p2.X && p1.Y != p2.Y {
		return ECPoint{Infinity: true}
	}
	var m int
	if p1.X == p2.X && p1.Y == p2.Y {
		num := (3*p1.X*p1.X + ec.A) % ec.P
		den := (2 * p1.Y) % ec.P
		m = (num * modInverseA(den, ec.P)) % ec.P
	} else {
		num := (p2.Y - p1.Y + ec.P) % ec.P
		den := (p2.X - p1.X + ec.P) % ec.P
		m = (num * modInverseA(den, ec.P)) % ec.P
	}
	x3 := (m*m - p1.X - p2.X + 2*ec.P) % ec.P
	y3 := (m*(p1.X-x3+ec.P) - p1.Y + ec.P) % ec.P
	return ECPoint{X: x3, Y: y3}
}

func (ec EllipticCurve) ScalarMult(k int, p ECPoint) ECPoint {
	result := ECPoint{Infinity: true}
	current := p
	for k > 0 {
		if k%2 == 1 {
			result = ec.Add(result, current)
		}
		current = ec.Add(current, current)
		k /= 2
	}
	return result
}

func (ec EllipticCurve) IsOnCurve(p ECPoint) bool {
	if p.Infinity {
		return true
	}
	lhs := (p.Y * p.Y) % ec.P
	rhs := (p.X*p.X*p.X + ec.A*p.X + ec.B) % ec.P
	return lhs == rhs
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
