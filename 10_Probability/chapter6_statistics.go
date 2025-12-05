package probability

func Median(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, data)
	quickSort(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

func Mode(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	counts := make(map[float64]int)
	for _, v := range data {
		counts[v]++
	}
	var mode float64
	maxCount := 0
	for v, c := range counts {
		if c > maxCount {
			maxCount = c
			mode = v
		}
	}
	return mode
}

func Percentile(data []float64, p float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, data)
	quickSort(sorted)
	index := p * float64(n-1)
	lower := int(index)
	upper := lower + 1
	if upper >= n {
		return sorted[n-1]
	}
	frac := index - float64(lower)
	return sorted[lower]*(1-frac) + sorted[upper]*frac
}

func Quartiles(data []float64) (q1, q2, q3 float64) {
	q1 = Percentile(data, 0.25)
	q2 = Percentile(data, 0.50)
	q3 = Percentile(data, 0.75)
	return
}

func IQR(data []float64) float64 {
	q1, _, q3 := Quartiles(data)
	return q3 - q1
}

func Skewness(data []float64) float64 {
	n := float64(len(data))
	if n < 3 {
		return 0
	}
	m := Mean(data)
	s := StandardDeviation(data)
	if s == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range data {
		sum += powerP((v-m)/s, 3)
	}
	return sum / n
}

func Kurtosis(data []float64) float64 {
	n := float64(len(data))
	if n < 4 {
		return 0
	}
	m := Mean(data)
	s := StandardDeviation(data)
	if s == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range data {
		sum += powerP((v-m)/s, 4)
	}
	return sum/n - 3
}

func quickSort(arr []float64) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
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
