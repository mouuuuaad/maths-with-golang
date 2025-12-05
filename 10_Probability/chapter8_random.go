package probability

type LCG struct {
	state uint64
	a     uint64
	c     uint64
	m     uint64
}

func NewLCG(seed uint64) *LCG {
	return &LCG{state: seed, a: 1103515245, c: 12345, m: 1 << 31}
}

func (l *LCG) Next() uint64 {
	l.state = (l.a*l.state + l.c) % l.m
	return l.state
}

func (l *LCG) Float64() float64 {
	return float64(l.Next()) / float64(l.m)
}

func (l *LCG) Uniform(a, b float64) float64 {
	return a + (b-a)*l.Float64()
}

func (l *LCG) NormalSample(mu, sigma float64) float64 {
	u1 := l.Float64()
	u2 := l.Float64()
	z := sqrtP(-2*lnP(u1)) * cosP(2*Pi*u2)
	return mu + sigma*z
}

func (l *LCG) ExponentialSample(lambda float64) float64 {
	return -lnP(l.Float64()) / lambda
}

func (l *LCG) BernoulliSample(p float64) int {
	if l.Float64() < p {
		return 1
	}
	return 0
}

func (l *LCG) BinomialSample(n int, p float64) int {
	count := 0
	for i := 0; i < n; i++ {
		count += l.BernoulliSample(p)
	}
	return count
}

func cosP(x float64) float64 {
	for x > Pi {
		x -= 2 * Pi
	}
	for x < -Pi {
		x += 2 * Pi
	}
	sum := 1.0
	term := 1.0
	x2 := x * x
	for i := 1; i < 30; i++ {
		term *= -x2 / float64((2*i-1)*(2*i))
		sum += term
	}
	return sum
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
