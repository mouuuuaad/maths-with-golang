package probability

func ChiSquarePDF(x float64, k int) float64 {
	if x <= 0 {
		return 0
	}
	kHalf := float64(k) / 2
	return powerP(x, kHalf-1) * expP(-x/2) / (powerP(2, kHalf) * gammaFunc(kHalf))
}

func TDistPDF(t float64, df int) float64 {
	v := float64(df)
	coef := gammaFunc((v+1)/2) / (sqrtP(v*Pi) * gammaFunc(v/2))
	return coef * powerP(1+t*t/v, -(v+1)/2)
}

func ZTest(sampleMean, popMean, popStd float64, n int) float64 {
	return (sampleMean - popMean) / (popStd / sqrtP(float64(n)))
}

func TTest(sampleMean, popMean, sampleStd float64, n int) float64 {
	return (sampleMean - popMean) / (sampleStd / sqrtP(float64(n)))
}

func ChiSquareTest(observed, expected []float64) float64 {
	chi2 := 0.0
	for i := range observed {
		if expected[i] != 0 {
			diff := observed[i] - expected[i]
			chi2 += diff * diff / expected[i]
		}
	}
	return chi2
}

func ANOVA(groups ...[]float64) float64 {
	k := len(groups)
	if k < 2 {
		return 0
	}
	allData := []float64{}
	for _, g := range groups {
		allData = append(allData, g...)
	}
	grandMean := Mean(allData)
	ssb, ssw := 0.0, 0.0
	for _, g := range groups {
		gMean := Mean(g)
		ssb += float64(len(g)) * (gMean - grandMean) * (gMean - grandMean)
		for _, v := range g {
			ssw += (v - gMean) * (v - gMean)
		}
	}
	n := float64(len(allData))
	return (ssb / float64(k-1)) / (ssw / (n - float64(k)))
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
