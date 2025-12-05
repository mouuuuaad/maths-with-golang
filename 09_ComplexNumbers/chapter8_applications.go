package complexnums

func MandelbrotIteration(c ComplexNumber, maxIter int) int {
	z := New(0, 0)
	for i := 0; i < maxIter; i++ {
		if z.Abs() > 2.0 {
			return i
		}
		z = z.Multiply(z).Add(c)
	}
	return maxIter
}

func JuliaIteration(z, c ComplexNumber, maxIter int) int {
	for i := 0; i < maxIter; i++ {
		if z.Abs() > 2.0 {
			return i
		}
		z = z.Multiply(z).Add(c)
	}
	return maxIter
}

func DFT(input []ComplexNumber) []ComplexNumber {
	N := len(input)
	output := make([]ComplexNumber, N)
	for k := 0; k < N; k++ {
		sum := New(0, 0)
		for n := 0; n < N; n++ {
			angle := -2.0 * PiC * float64(k) * float64(n) / float64(N)
			w := EulerFormula(angle)
			term := input[n].Multiply(w)
			sum = sum.Add(term)
		}
		output[k] = sum
	}
	return output
}

func IDFT(input []ComplexNumber) []ComplexNumber {
	N := len(input)
	output := make([]ComplexNumber, N)
	for n := 0; n < N; n++ {
		sum := New(0, 0)
		for k := 0; k < N; k++ {
			angle := 2.0 * PiC * float64(k) * float64(n) / float64(N)
			w := EulerFormula(angle)
			term := input[k].Multiply(w)
			sum = sum.Add(term)
		}
		output[n] = sum.Scale(1.0 / float64(N))
	}
	return output
}

func FFT(input []ComplexNumber) []ComplexNumber {
	n := len(input)
	if n <= 1 {
		return input
	}
	if n&(n-1) != 0 {
		return DFT(input)
	}
	even := make([]ComplexNumber, n/2)
	odd := make([]ComplexNumber, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = input[2*i]
		odd[i] = input[2*i+1]
	}
	evenFFT := FFT(even)
	oddFFT := FFT(odd)
	result := make([]ComplexNumber, n)
	for k := 0; k < n/2; k++ {
		angle := -2.0 * PiC * float64(k) / float64(n)
		w := EulerFormula(angle)
		t := w.Multiply(oddFFT[k])
		result[k] = evenFFT[k].Add(t)
		result[k+n/2] = evenFFT[k].Subtract(t)
	}
	return result
}

func IFFT(input []ComplexNumber) []ComplexNumber {
	n := len(input)
	conj := make([]ComplexNumber, n)
	for i, c := range input {
		conj[i] = c.Conjugate()
	}
	result := FFT(conj)
	for i := range result {
		result[i] = result[i].Conjugate().Scale(1.0 / float64(n))
	}
	return result
}

func Convolution(a, b []ComplexNumber) []ComplexNumber {
	n := len(a) + len(b) - 1
	size := 1
	for size < n {
		size *= 2
	}
	paddedA := make([]ComplexNumber, size)
	paddedB := make([]ComplexNumber, size)
	copy(paddedA, a)
	copy(paddedB, b)
	fftA := FFT(paddedA)
	fftB := FFT(paddedB)
	product := make([]ComplexNumber, size)
	for i := range product {
		product[i] = fftA[i].Multiply(fftB[i])
	}
	result := IFFT(product)
	return result[:n]
}

func NewtonFractal(z, c ComplexNumber, f, df func(ComplexNumber) ComplexNumber, maxIter int) ComplexNumber {
	for i := 0; i < maxIter; i++ {
		fz := f(z)
		dfz := df(z)
		if dfz.Abs() < 1e-10 {
			break
		}
		z = z.Subtract(fz.Divide(dfz))
	}
	return z
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
