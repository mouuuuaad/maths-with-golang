package sequences

type RecurrenceRelation func(prev []float64) float64

type RecursiveSequence struct {
	InitialValues []float64
	Relation      RecurrenceRelation
}

func (s RecursiveSequence) NthTerm(n int) float64 {
	if n < len(s.InitialValues) {
		return s.InitialValues[n]
	}
	terms := make([]float64, n+1)
	copy(terms, s.InitialValues)
	for i := len(s.InitialValues); i <= n; i++ {
		terms[i] = s.Relation(terms[:i])
	}
	return terms[n]
}

func (s RecursiveSequence) Generate(n int) []float64 {
	result := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = s.NthTerm(i)
	}
	return result
}

func LinearRecurrence(coeffs []float64, initial []float64, n int) float64 {
	if n < len(initial) {
		return initial[n]
	}
	terms := make([]float64, n+1)
	copy(terms, initial)
	for i := len(initial); i <= n; i++ {
		sum := 0.0
		for j, c := range coeffs {
			sum += c * terms[i-j-1]
		}
		terms[i] = sum
	}
	return terms[n]
}

func FixedPointIteration(f func(float64) float64, x0 float64, maxIter int, tol float64) (float64, bool) {
	x := x0
	for i := 0; i < maxIter; i++ {
		xNew := f(x)
		if absS(xNew-x) < tol {
			return xNew, true
		}
		x = xNew
	}
	return x, false
}

func NewtonsRecurrence(f, df func(float64) float64, x0 float64, n int) float64 {
	x := x0
	for i := 0; i < n; i++ {
		fx := f(x)
		dfx := df(x)
		if absS(dfx) < 1e-12 {
			break
		}
		x -= fx / dfx
	}
	return x
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
