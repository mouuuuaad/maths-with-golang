package discrete

type Set map[int]bool

func NewSet(elements ...int) Set {
	s := make(Set)
	for _, e := range elements {
		s[e] = true
	}
	return s
}

func (s Set) Add(e int) { s[e] = true }

func (s Set) Remove(e int) { delete(s, e) }

func (s Set) Contains(e int) bool { return s[e] }

func (s Set) Size() int { return len(s) }

func (s Set) Union(other Set) Set {
	result := NewSet()
	for e := range s {
		result[e] = true
	}
	for e := range other {
		result[e] = true
	}
	return result
}

func (s Set) Intersection(other Set) Set {
	result := NewSet()
	for e := range s {
		if other[e] {
			result[e] = true
		}
	}
	return result
}

func (s Set) Difference(other Set) Set {
	result := NewSet()
	for e := range s {
		if !other[e] {
			result[e] = true
		}
	}
	return result
}

func (s Set) SymmetricDifference(other Set) Set {
	return s.Difference(other).Union(other.Difference(s))
}

func (s Set) IsSubset(other Set) bool {
	for e := range s {
		if !other[e] {
			return false
		}
	}
	return true
}

func (s Set) PowerSet() []Set {
	elements := make([]int, 0, len(s))
	for e := range s {
		elements = append(elements, e)
	}
	n := 1 << len(elements)
	result := make([]Set, n)
	for i := 0; i < n; i++ {
		result[i] = NewSet()
		for j := 0; j < len(elements); j++ {
			if i&(1<<j) != 0 {
				result[i][elements[j]] = true
			}
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
