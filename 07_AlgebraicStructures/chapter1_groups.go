package algebra

type Group interface {
	Identity() interface{}
	Operate(a, b interface{}) interface{}
	Inverse(a interface{}) interface{}
}

type IntegerAddGroup struct{}

func (g IntegerAddGroup) Identity() interface{} {
	return 0
}

func (g IntegerAddGroup) Operate(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func (g IntegerAddGroup) Inverse(a interface{}) interface{} {
	return -a.(int)
}

type IntegerMulGroup struct{}

func (g IntegerMulGroup) Identity() interface{} {
	return 1
}

func (g IntegerMulGroup) Operate(a, b interface{}) interface{} {
	return a.(int) * b.(int)
}

func (g IntegerMulGroup) Inverse(a interface{}) interface{} {
	return 0
}

type CyclicGroup struct {
	Order int
}

func (g CyclicGroup) Identity() interface{} {
	return 0
}

func (g CyclicGroup) Operate(a, b interface{}) interface{} {
	return (a.(int) + b.(int)) % g.Order
}

func (g CyclicGroup) Inverse(a interface{}) interface{} {
	return (g.Order - a.(int)) % g.Order
}

func (g CyclicGroup) Generator() int {
	return 1
}

func (g CyclicGroup) Generate(n int) int {
	return n % g.Order
}

type SymmetricGroup struct {
	N int
}

func (g SymmetricGroup) Identity() []int {
	perm := make([]int, g.N)
	for i := range perm {
		perm[i] = i
	}
	return perm
}

func (g SymmetricGroup) Compose(a, b []int) []int {
	result := make([]int, g.N)
	for i := 0; i < g.N; i++ {
		result[i] = a[b[i]]
	}
	return result
}

func (g SymmetricGroup) Inverse(a []int) []int {
	result := make([]int, g.N)
	for i := 0; i < g.N; i++ {
		result[a[i]] = i
	}
	return result
}
