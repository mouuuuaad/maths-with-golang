package discrete

func ChineseRemainderTheorem(remainders, moduli []int) int {
	n := len(remainders)
	M := 1
	for _, m := range moduli {
		M *= m
	}
	result := 0
	for i := 0; i < n; i++ {
		Mi := M / moduli[i]
		yi := ModInverse(Mi, moduli[i])
		result += remainders[i] * Mi * yi
	}
	return ((result % M) + M) % M
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func SieveOfEratosthenes(n int) []int {
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false
	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}
	primes := []int{}
	for i, isPrime := range sieve {
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func EulerPhi(n int) int {
	result := n
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			for n%p == 0 {
				n /= p
			}
			result -= result / p
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func MobiusFunction(n int) int {
	if n == 1 {
		return 1
	}
	primeFactors := 0
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			n /= p
			primeFactors++
			if n%p == 0 {
				return 0
			}
		}
	}
	if n > 1 {
		primeFactors++
	}
	if primeFactors%2 == 0 {
		return 1
	}
	return -1
}

func DiscreteLog(base, target, mod int) int {
	m := 1
	for m*m < mod {
		m++
	}
	table := make(map[int]int)
	power := 1
	for j := 0; j < m; j++ {
		table[power] = j
		power = (power * base) % mod
	}
	factor := ModPow(base, mod-1-m, mod)
	gamma := target
	for i := 0; i < m; i++ {
		if j, ok := table[gamma]; ok {
			return i*m + j
		}
		gamma = (gamma * factor) % mod
	}
	return -1
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
