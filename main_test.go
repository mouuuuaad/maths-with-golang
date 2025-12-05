package main

import (
	"testing"

	calculus "github.com/mouaadid/MathsWithGolang/01_Calculus"
	linearalgebra "github.com/mouaadid/MathsWithGolang/02_LinearAlgebra"
	limits "github.com/mouaadid/MathsWithGolang/03_Limits"
	derivatives "github.com/mouaadid/MathsWithGolang/04_Derivatives"
	functions "github.com/mouaadid/MathsWithGolang/05_Functions"
	sequences "github.com/mouaadid/MathsWithGolang/06_NumericalSequences"
	algebra "github.com/mouaadid/MathsWithGolang/07_AlgebraicStructures"
	arithmetic "github.com/mouaadid/MathsWithGolang/08_Arithmetic"
	complexnums "github.com/mouaadid/MathsWithGolang/09_ComplexNumbers"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func TestFunctionsModule(t *testing.T) {
	if abs(functions.Exp(0)-1) > 1e-9 {
		t.Error("Exp(0) should be 1")
	}
	if abs(functions.Exp(1)-functions.E) > 1e-6 {
		t.Error("Exp(1) should be e")
	}
	if abs(functions.Ln(functions.E)-1) > 1e-6 {
		t.Error("Ln(e) should be 1")
	}
	if abs(functions.Sin(0)) > 1e-9 {
		t.Error("Sin(0) should be 0")
	}
	if abs(functions.Sin(functions.Pi/2)-1) > 1e-6 {
		t.Error("Sin(π/2) should be 1")
	}
	if abs(functions.Cos(0)-1) > 1e-9 {
		t.Error("Cos(0) should be 1")
	}
	if abs(functions.Sqrt(4)-2) > 1e-9 {
		t.Error("Sqrt(4) should be 2")
	}
	if abs(functions.Gamma(5)-24) > 1e-6 {
		t.Error("Gamma(5) should be 24")
	}
}

func TestArithmeticModule(t *testing.T) {
	if !arithmetic.IsPrime(97) {
		t.Error("97 should be prime")
	}
	if arithmetic.IsPrime(100) {
		t.Error("100 should not be prime")
	}
	factors := arithmetic.PrimeFactors(84)
	if len(factors) == 0 {
		t.Error("84 should have prime factors")
	}
	if arithmetic.EulerTotient(12) != 4 {
		t.Error("EulerTotient(12) should be 4")
	}
	gcd := arithmetic.GCD(48, 18)
	if gcd != 6 {
		t.Errorf("GCD(48, 18) should be 6, got %d", gcd)
	}
}

func TestComplexModule(t *testing.T) {
	z := complexnums.New(3, 4)
	if abs(z.Abs()-5) > 1e-9 {
		t.Error("|3+4i| should be 5")
	}
	z1 := complexnums.New(1, 1)
	z2 := complexnums.New(2, 3)
	product := z1.Multiply(z2)
	if abs(product.R-(-1)) > 1e-9 || abs(product.I-5) > 1e-9 {
		t.Error("(1+i)*(2+3i) should be -1+5i")
	}
	euler := complexnums.New(0, functions.Pi).Exp()
	if abs(euler.R-(-1)) > 1e-6 || abs(euler.I) > 1e-6 {
		t.Error("e^(iπ) should be -1")
	}
}

func TestLinearAlgebraModule(t *testing.T) {
	A := linearalgebra.Matrix{{1, 2}, {3, 4}}
	det := A.Determinant()
	if abs(det-(-2)) > 1e-9 {
		t.Errorf("Det([[1,2],[3,4]]) should be -2, got %f", det)
	}
	I := linearalgebra.Identity(3)
	if abs(I.Determinant()-1) > 1e-9 {
		t.Error("Det(I3) should be 1")
	}
	v := linearalgebra.Vector{1, 0, 0}
	u := linearalgebra.Vector{0, 1, 0}
	cross := v.Cross(u)
	if abs(cross[2]-1) > 1e-9 {
		t.Error("i × j should be k")
	}
}

func TestCalculusModule(t *testing.T) {
	f := func(x float64) float64 { return x * x }
	deriv := calculus.Derivative(f, 3)
	if abs(deriv-6) > 1e-5 {
		t.Errorf("d/dx(x²) at x=3 should be 6, got %f", deriv)
	}
	integral := calculus.SimpsonRule(f, 0, 1, 100)
	if abs(integral-1.0/3.0) > 1e-6 {
		t.Errorf("∫x² dx from 0 to 1 should be 1/3, got %f", integral)
	}
}

func TestLimitsModule(t *testing.T) {
	f := func(x float64) float64 { return x * x }
	lim, ok := limits.Limit(f, 2)
	if !ok || abs(lim-4) > 1e-6 {
		t.Error("lim(x²) as x→2 should be 4")
	}
}

func TestDerivativesModule(t *testing.T) {
	f := func(x float64) float64 { return x * x * x }
	d := derivatives.Derivative(f, 2)
	if abs(d-12) > 1e-4 {
		t.Errorf("d/dx(x³) at x=2 should be 12, got %f", d)
	}
	d2 := derivatives.SecondDerivative(f, 2)
	if abs(d2-12) > 1e-3 {
		t.Errorf("d²/dx²(x³) at x=2 should be 12, got %f", d2)
	}
}

func TestSequencesModule(t *testing.T) {
	if sequences.Fibonacci(10) != 55 {
		t.Errorf("Fibonacci(10) should be 55, got %d", sequences.Fibonacci(10))
	}
	if sequences.Catalan(5) != 42 {
		t.Errorf("Catalan(5) should be 42, got %d", sequences.Catalan(5))
	}
	arith := sequences.ArithmeticSequence{Start: 1, Diff: 2}
	if abs(arith.NthTerm(5)-11) > 1e-9 {
		t.Error("Arithmetic sequence 1,3,5,... term 5 should be 11")
	}
}

func TestAlgebraModule(t *testing.T) {
	F := algebra.FiniteField{P: 7}
	product := F.Multiply(3, 5)
	if product != 1 {
		t.Errorf("3*5 mod 7 should be 1, got %d", product)
	}
	inv := F.Inverse(3)
	check := F.Multiply(3, inv)
	if check != 1 {
		t.Errorf("3 * inv(3) mod 7 should be 1, got %d", check)
	}
	poly := algebra.NewPolynomial(1, 2, 1)
	val := poly.Evaluate(2)
	if abs(val-9) > 1e-9 {
		t.Errorf("(1+2x+x²) at x=2 should be 9, got %f", val)
	}
}
