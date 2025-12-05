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
