package main

import (
	"fmt"

	calculus "github.com/mouaadid/MathsWithGolang/01_Calculus"
	linearalgebra "github.com/mouaadid/MathsWithGolang/02_LinearAlgebra"
	functions "github.com/mouaadid/MathsWithGolang/05_Functions"
	sequences "github.com/mouaadid/MathsWithGolang/06_NumericalSequences"
	algebra "github.com/mouaadid/MathsWithGolang/07_AlgebraicStructures"
	arithmetic "github.com/mouaadid/MathsWithGolang/08_Arithmetic"
	complexnums "github.com/mouaadid/MathsWithGolang/09_ComplexNumbers"
)

func main() {
	fmt.Println("=== MathsWithGolang - Pure Golang Deep Math Library ===")
	fmt.Println()

	fmt.Println("1. Functions Module (Pure Taylor Series)")
	fmt.Printf("   Exp(1) = %.10f (e)\n", functions.Exp(1))
	fmt.Printf("   Ln(e)  = %.10f\n", functions.Ln(functions.E))
	fmt.Printf("   Sin(π/2) = %.10f\n", functions.Sin(functions.Pi/2))
	fmt.Printf("   Cos(0) = %.10f\n", functions.Cos(0))
	fmt.Printf("   Gamma(5) = %.10f (should be 24)\n", functions.Gamma(5))
	fmt.Println()

	fmt.Println("2. Arithmetic Module (Number Theory)")
	fmt.Printf("   IsPrime(97) = %v\n", arithmetic.IsPrime(97))
	fmt.Printf("   IsPrime(100) = %v\n", arithmetic.IsPrime(100))
	fmt.Printf("   PrimeFactors(84) = %v\n", arithmetic.PrimeFactors(84))
	fmt.Printf("   EulerTotient(12) = %d\n", arithmetic.EulerTotient(12))
	fmt.Println()

	fmt.Println("3. Complex Numbers Module")
	z1 := complexnums.New(3, 4)
	z2 := complexnums.New(1, 2)
	fmt.Printf("   z1 = 3+4i, |z1| = %.4f\n", z1.Abs())
	fmt.Printf("   z1 * z2 = (%.4f, %.4f)\n", z1.Multiply(z2).R, z1.Multiply(z2).I)
	fmt.Printf("   exp(i*π) = (%.4f, %.4f)\n", complexnums.New(0, functions.Pi).Exp().R, complexnums.New(0, functions.Pi).Exp().I)
	fmt.Println()

	fmt.Println("4. Linear Algebra Module")
	A := linearalgebra.Matrix{{1, 2}, {3, 4}}
	fmt.Printf("   Matrix A = [[1,2],[3,4]]\n")
	fmt.Printf("   Det(A) = %.4f\n", A.Determinant())
	fmt.Printf("   Trace(A) = %.4f\n", A.Trace())
	v := linearalgebra.Vector{1, 2, 3}
	u := linearalgebra.Vector{4, 5, 6}
	fmt.Printf("   v·u = %.4f\n", v.Dot(u))
	fmt.Println()

	fmt.Println("5. Calculus Module")
	f := func(x float64) float64 { return x * x }
	fmt.Printf("   d/dx(x²) at x=3 = %.6f\n", calculus.Derivative(f, 3))
	fmt.Printf("   ∫x² dx from 0 to 1 = %.6f\n", calculus.SimpsonRule(f, 0, 1, 100))
	fmt.Println()

	fmt.Println("6. Numerical Sequences Module")
	fmt.Printf("   Fibonacci(10) = %d\n", sequences.Fibonacci(10))
	fmt.Printf("   Catalan(5) = %d\n", sequences.Catalan(5))
	arith := sequences.ArithmeticSequence{Start: 1, Diff: 2}
	fmt.Printf("   Arithmetic Sum(10 terms) = %.4f\n", arith.SumN(9))
	fmt.Println()

	fmt.Println("7. Algebraic Structures Module")
	F := algebra.FiniteField{P: 7}
	fmt.Printf("   Finite Field GF(7): 3 * 5 = %d (mod 7)\n", F.Multiply(3, 5))
	fmt.Printf("   Inverse of 3 in GF(7) = %d\n", F.Inverse(3))
	poly := algebra.NewPolynomial(1, 2, 1)
	fmt.Printf("   Polynomial (1 + 2x + x²) at x=2 = %.4f\n", poly.Evaluate(2))
	fmt.Println()

	fmt.Println("=== All modules use pure Golang - no external packages ===")
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
