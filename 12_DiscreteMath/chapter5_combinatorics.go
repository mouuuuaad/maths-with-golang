package discrete

func GeneratePermutations(n int) [][]int {
	if n <= 0 {
		return nil
	}
	result := [][]int{}
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	var generate func(k int)
	generate = func(k int) {
		if k == 1 {
			tmp := make([]int, n)
			copy(tmp, perm)
			result = append(result, tmp)
			return
		}
		for i := 0; i < k; i++ {
			generate(k - 1)
			if k%2 == 0 {
				perm[i], perm[k-1] = perm[k-1], perm[i]
			} else {
				perm[0], perm[k-1] = perm[k-1], perm[0]
			}
		}
	}
	generate(n)
	return result
}

func GenerateCombinations(n, k int) [][]int {
	if k > n || k < 0 {
		return nil
	}
	result := [][]int{}
	comb := make([]int, k)
	var generate func(start, idx int)
	generate = func(start, idx int) {
		if idx == k {
			tmp := make([]int, k)
			copy(tmp, comb)
			result = append(result, tmp)
			return
		}
		for i := start; i <= n-(k-idx); i++ {
			comb[idx] = i
			generate(i+1, idx+1)
		}
	}
	generate(0, 0)
	return result
}

func Derangements(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	d := make([]int, n+1)
	d[0], d[1] = 1, 0
	for i := 2; i <= n; i++ {
		d[i] = (i - 1) * (d[i-1] + d[i-2])
	}
	return d[n]
}

func StirlingSecond(n, k int) int {
	if n == 0 && k == 0 {
		return 1
	}
	if n == 0 || k == 0 {
		return 0
	}
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, k+1)
	}
	s[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= minI(i, k); j++ {
			s[i][j] = j*s[i-1][j] + s[i-1][j-1]
		}
	}
	return s[n][k]
}

func BellNumber(n int) int {
	b := make([]int, n+1)
	b[0] = 1
	for i := 1; i <= n; i++ {
		for k := 0; k < i; k++ {
			b[i] += Combination(i-1, k) * b[k]
		}
	}
	return b[n]
}

func minI(a, b int) int {
	if a < b {
		return a
	}
	return b
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
