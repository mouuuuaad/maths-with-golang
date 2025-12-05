package algebra

type GroupHomomorphism struct {
	Domain   Group
	Codomain Group
	Map      func(interface{}) interface{}
}

func (h GroupHomomorphism) IsHomomorphism(elements []interface{}) bool {
	for _, a := range elements {
		for _, b := range elements {
			ab := h.Domain.Operate(a, b)
			faFb := h.Codomain.Operate(h.Map(a), h.Map(b))
			if h.Map(ab) != faFb {
				return false
			}
		}
	}
	return true
}

func (h GroupHomomorphism) Kernel(elements []interface{}) []interface{} {
	result := []interface{}{}
	identity := h.Codomain.Identity()
	for _, a := range elements {
		if h.Map(a) == identity {
			result = append(result, a)
		}
	}
	return result
}

func (h GroupHomomorphism) Image(elements []interface{}) []interface{} {
	result := []interface{}{}
	seen := make(map[interface{}]bool)
	for _, a := range elements {
		fa := h.Map(a)
		if !seen[fa] {
			seen[fa] = true
			result = append(result, fa)
		}
	}
	return result
}

func (h GroupHomomorphism) IsInjective(elements []interface{}) bool {
	kernel := h.Kernel(elements)
	return len(kernel) == 1 && kernel[0] == h.Domain.Identity()
}

func (h GroupHomomorphism) IsSurjective(domainElements, codomainElements []interface{}) bool {
	image := h.Image(domainElements)
	return len(image) == len(codomainElements)
}

func (h GroupHomomorphism) IsIsomorphism(domainElements, codomainElements []interface{}) bool {
	return h.IsInjective(domainElements) && h.IsSurjective(domainElements, codomainElements)
}

func ComposeMorphisms(f, g GroupHomomorphism) GroupHomomorphism {
	return GroupHomomorphism{
		Domain:   f.Domain,
		Codomain: g.Codomain,
		Map: func(x interface{}) interface{} {
			return g.Map(f.Map(x))
		},
	}
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
