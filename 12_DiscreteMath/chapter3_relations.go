package discrete

type Relation struct {
	Domain   Set
	Codomain Set
	Pairs    [][2]int
}

func (r Relation) IsReflexive() bool {
	for e := range r.Domain {
		found := false
		for _, p := range r.Pairs {
			if p[0] == e && p[1] == e {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (r Relation) IsSymmetric() bool {
	for _, p := range r.Pairs {
		found := false
		for _, q := range r.Pairs {
			if q[0] == p[1] && q[1] == p[0] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (r Relation) IsTransitive() bool {
	for _, p := range r.Pairs {
		for _, q := range r.Pairs {
			if p[1] == q[0] {
				found := false
				for _, s := range r.Pairs {
					if s[0] == p[0] && s[1] == q[1] {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
		}
	}
	return true
}

func (r Relation) IsEquivalence() bool {
	return r.IsReflexive() && r.IsSymmetric() && r.IsTransitive()
}

func (r Relation) IsPartialOrder() bool {
	return r.IsReflexive() && r.IsAntisymmetric() && r.IsTransitive()
}

func (r Relation) IsAntisymmetric() bool {
	for _, p := range r.Pairs {
		for _, q := range r.Pairs {
			if p[0] == q[1] && p[1] == q[0] && p[0] != p[1] {
				return false
			}
		}
	}
	return true
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
