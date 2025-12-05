package probability

func Covariance(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}
	mx := Mean(x)
	my := Mean(y)
	sum := 0.0
	for i := range x {
		sum += (x[i] - mx) * (y[i] - my)
	}
	return sum / float64(len(x))
}

func Correlation(x, y []float64) float64 {
	cov := Covariance(x, y)
	sx := StandardDeviation(x)
	sy := StandardDeviation(y)
	if sx == 0 || sy == 0 {
		return 0
	}
	return cov / (sx * sy)
}

func SampleCovariance(x, y []float64) float64 {
	if len(x) != len(y) || len(x) <= 1 {
		return 0
	}
	mx := Mean(x)
	my := Mean(y)
	sum := 0.0
	for i := range x {
		sum += (x[i] - mx) * (y[i] - my)
	}
	return sum / float64(len(x)-1)
}

func SampleVariance(data []float64) float64 {
	if len(data) <= 1 {
		return 0
	}
	m := Mean(data)
	sum := 0.0
	for _, v := range data {
		diff := v - m
		sum += diff * diff
	}
	return sum / float64(len(data)-1)
}

func SampleStandardDeviation(data []float64) float64 {
	return sqrtP(SampleVariance(data))
}

func CovarianceMatrix(data [][]float64) [][]float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
		for j := range result[i] {
			result[i][j] = Covariance(data[i], data[j])
		}
	}
	return result
}

func CorrelationMatrix(data [][]float64) [][]float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
		for j := range result[i] {
			result[i][j] = Correlation(data[i], data[j])
		}
	}
	return result
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
