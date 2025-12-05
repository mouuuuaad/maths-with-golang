package algebra

type QuotientGroup struct {
	G      Group
	N      []interface{}
	Cosets [][]interface{}
}

func NewQuotientGroup(G Group, N []interface{}, elements []interface{}) QuotientGroup {
	cosets := [][]interface{}{}
	used := make(map[interface{}]bool)
	for _, g := range elements {
		if used[g] {
			continue
		}
		coset := LeftCoset(G, N, g)
		for _, c := range coset {
			used[c] = true
		}
		cosets = append(cosets, coset)
	}
	return QuotientGroup{G: G, N: N, Cosets: cosets}
}

func (q QuotientGroup) FindCoset(a interface{}) int {
	for i, coset := range q.Cosets {
		for _, c := range coset {
			if c == a {
				return i
			}
		}
	}
	return -1
}

func (q QuotientGroup) Operate(aCoset, bCoset interface{}) interface{} {
	a := aCoset
	b := bCoset
	return q.G.Operate(a, b)
}

func (q QuotientGroup) Order() int {
	return len(q.Cosets)
}

func FirstIsomorphismTheorem(h GroupHomomorphism, elements []interface{}) bool {
	kernel := h.Kernel(elements)
	image := h.Image(elements)
	return len(elements)/len(kernel) == len(image)
}

func LagrangesTheorem(groupOrder, subgroupOrder int) bool {
	if subgroupOrder == 0 {
		return false
	}
	return groupOrder%subgroupOrder == 0
}
