package arithmetic

func mulMod(a, b, m uint64) uint64 {
	var res uint64 = 0
	a %= m
	for b > 0 {
		if b%2 == 1 {
			res = (res + a) % m
		}
		a = (a * 2) % m
		b /= 2
	}
	return res
}

func ModularExponentiation(base, exp, mod uint64) uint64 {
	if mod == 1 {
		return 0
	}
	var result uint64 = 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = mulMod(result, base, mod)
		}
		exp = exp >> 1
		base = mulMod(base, base, mod)
	}
	return result
}

func IsPrimeMillerRabin(n uint64, k int) bool {
	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}
	witnesses := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, a := range witnesses {
		if a >= n {
			continue
		}
		x := ModularExponentiation(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for j := 0; j < r-1; j++ {
			x = mulMod(x, x, n)
			if x == n-1 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}

func IsPrime(n uint64) bool {
	return IsPrimeMillerRabin(n, 20)
}

func gcdU(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func PollardsRho(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return 2
	}
	x := uint64(2)
	y := uint64(2)
	d := uint64(1)
	c := uint64(1)
	f := func(v uint64) uint64 {
		return (mulMod(v, v, n) + c) % n
	}
	for d == 1 {
		x = f(x)
		y = f(f(y))
		val := x - y
		if x < y {
			val = y - x
		}
		d = gcdU(val, n)
		if d == n {
			return n
		}
	}
	return d
}

func PrimeFactors(n uint64) []uint64 {
	factors := []uint64{}
	if n <= 1 {
		return factors
	}
	for _, p := range []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29} {
		for n%p == 0 {
			factors = append(factors, p)
			n /= p
		}
	}
	if n == 1 {
		return factors
	}
	var factorize func(m uint64)
	factorize = func(m uint64) {
		if m == 1 {
			return
		}
		if IsPrime(m) {
			factors = append(factors, m)
			return
		}
		factor := PollardsRho(m)
		if factor == m {
			factors = append(factors, m)
			return
		}
		factorize(factor)
		factorize(m / factor)
	}
	factorize(n)
	return factors
}

func SieveOfEratosthenes(limit uint64) []uint64 {
	if limit < 2 {
		return []uint64{}
	}
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for p := uint64(2); p*p <= limit; p++ {
		if isPrime[p] {
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
	}
	primes := []uint64{}
	for p := uint64(2); p <= limit; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func EulerTotient(n uint64) uint64 {
	result := n
	p := uint64(2)
	for p*p <= n {
		if n%p == 0 {
			for n%p == 0 {
				n /= p
			}
			result -= result / p
		}
		p++
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func Mobius(n uint64) int {
	if n == 1 {
		return 1
	}
	primeFactors := 0
	p := uint64(2)
	for p*p <= n {
		if n%p == 0 {
			n /= p
			primeFactors++
			if n%p == 0 {
				return 0
			}
		}
		p++
	}
	if n > 1 {
		primeFactors++
	}
	if primeFactors%2 == 0 {
		return 1
	}
	return -1
}

func Jacobi(a, n int64) int {
	if n <= 0 || n%2 == 0 {
		return 0
	}
	if a < 0 {
		a = ((a % n) + n) % n
	}
	a = a % n
	result := 1
	for a != 0 {
		for a%2 == 0 {
			a /= 2
			if n%8 == 3 || n%8 == 5 {
				result = -result
			}
		}
		a, n = n, a
		if a%4 == 3 && n%4 == 3 {
			result = -result
		}
		a = a % n
	}
	if n == 1 {
		return result
	}
	return 0
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
