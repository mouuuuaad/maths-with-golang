package discrete

func BoolAND(a, b bool) bool   { return a && b }
func BoolOR(a, b bool) bool    { return a || b }
func BoolNOT(a bool) bool      { return !a }
func BoolXOR(a, b bool) bool   { return a != b }
func BoolNAND(a, b bool) bool  { return !(a && b) }
func BoolNOR(a, b bool) bool   { return !(a || b) }
func BoolIMPLY(a, b bool) bool { return !a || b }
func BoolIFF(a, b bool) bool   { return a == b }

func TruthTable(vars int, f func([]bool) bool) [][]bool {
	rows := 1 << vars
	result := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]bool, vars+1)
		for j := 0; j < vars; j++ {
			result[i][j] = (i>>(vars-1-j))&1 == 1
		}
		inputs := result[i][:vars]
		result[i][vars] = f(inputs)
	}
	return result
}

func IsTautology(vars int, f func([]bool) bool) bool {
	table := TruthTable(vars, f)
	for _, row := range table {
		if !row[vars] {
			return false
		}
	}
	return true
}

func IsContradiction(vars int, f func([]bool) bool) bool {
	table := TruthTable(vars, f)
	for _, row := range table {
		if row[vars] {
			return false
		}
	}
	return true
}

func IsSatisfiable(vars int, f func([]bool) bool) bool {
	table := TruthTable(vars, f)
	for _, row := range table {
		if row[vars] {
			return true
		}
	}
	return false
}

func AreEquivalent(vars int, f, g func([]bool) bool) bool {
	rows := 1 << vars
	for i := 0; i < rows; i++ {
		inputs := make([]bool, vars)
		for j := 0; j < vars; j++ {
			inputs[j] = (i>>(vars-1-j))&1 == 1
		}
		if f(inputs) != g(inputs) {
			return false
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
