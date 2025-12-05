package algebra

type Ring interface {
	Zero() interface{}
	One() interface{}
	Add(a, b interface{}) interface{}
	Multiply(a, b interface{}) interface{}
	Negate(a interface{}) interface{}
}

type IntegerRing struct{}

func (r IntegerRing) Zero() interface{} {
	return 0
}

func (r IntegerRing) One() interface{} {
	return 1
}

func (r IntegerRing) Add(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func (r IntegerRing) Multiply(a, b interface{}) interface{} {
	return a.(int) * b.(int)
}

func (r IntegerRing) Negate(a interface{}) interface{} {
	return -a.(int)
}

type ModularRing struct {
	N int
}

func (r ModularRing) Zero() interface{} {
	return 0
}

func (r ModularRing) One() interface{} {
	return 1
}

func (r ModularRing) Add(a, b interface{}) interface{} {
	return (a.(int) + b.(int)) % r.N
}

func (r ModularRing) Multiply(a, b interface{}) interface{} {
	return (a.(int) * b.(int)) % r.N
}

func (r ModularRing) Negate(a interface{}) interface{} {
	return (r.N - a.(int)) % r.N
}

func IsCommutativeRing(R Ring, elements []interface{}) bool {
	for _, a := range elements {
		for _, b := range elements {
			if R.Multiply(a, b) != R.Multiply(b, a) {
				return false
			}
		}
	}
	return true
}

func IsIntegralDomain(R Ring, elements []interface{}) bool {
	for _, a := range elements {
		if a == R.Zero() {
			continue
		}
		for _, b := range elements {
			if b == R.Zero() {
				continue
			}
			if R.Multiply(a, b) == R.Zero() {
				return false
			}
		}
	}
	return true
}

func Units(R ModularRing) []int {
	result := []int{}
	for i := 1; i < R.N; i++ {
		if gcdA(i, R.N) == 1 {
			result = append(result, i)
		}
	}
	return result
}

func gcdA(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
