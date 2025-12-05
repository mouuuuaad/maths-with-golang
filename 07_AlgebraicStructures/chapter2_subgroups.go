package algebra

func IsSubgroup(G Group, elements []interface{}) bool {
	if len(elements) == 0 {
		return false
	}
	hasIdentity := false
	for _, e := range elements {
		if e == G.Identity() {
			hasIdentity = true
			break
		}
	}
	if !hasIdentity {
		return false
	}
	for _, a := range elements {
		invA := G.Inverse(a)
		found := false
		for _, e := range elements {
			if e == invA {
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

func LeftCoset(G Group, H []interface{}, g interface{}) []interface{} {
	result := make([]interface{}, len(H))
	for i, h := range H {
		result[i] = G.Operate(g, h)
	}
	return result
}

func RightCoset(G Group, H []interface{}, g interface{}) []interface{} {
	result := make([]interface{}, len(H))
	for i, h := range H {
		result[i] = G.Operate(h, g)
	}
	return result
}

func OrderOfElement(G CyclicGroup, a int) int {
	if a == 0 {
		return 1
	}
	current := a
	for order := 1; order <= G.Order; order++ {
		if current == 0 {
			return order
		}
		current = G.Operate(current, a).(int)
	}
	return G.Order
}

func Index(groupOrder, subgroupOrder int) int {
	if subgroupOrder == 0 {
		return 0
	}
	return groupOrder / subgroupOrder
}

func NormalSubgroup(G Group, H []interface{}, generators []interface{}) bool {
	for _, g := range generators {
		for _, h := range H {
			ghg1 := G.Operate(G.Operate(g, h), G.Inverse(g))
			found := false
			for _, e := range H {
				if e == ghg1 {
					found = true
					break
				}
			}
			if !found {
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
// Created by: MOUAAD
// MathsWithGolang - Pure Golang Mathematical Library
