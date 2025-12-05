package graphs

func EulerianPath(g *Graph) []int {
	n := g.Vertices
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		edges[i] = make([]int, len(g.Edges[i]))
		copy(edges[i], g.Edges[i])
	}
	start := 0
	oddCount := 0
	for v := 0; v < n; v++ {
		if len(edges[v])%2 == 1 {
			oddCount++
			start = v
		}
	}
	if oddCount != 0 && oddCount != 2 {
		return nil
	}
	stack := []int{start}
	path := []int{}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		if len(edges[v]) > 0 {
			u := edges[v][len(edges[v])-1]
			edges[v] = edges[v][:len(edges[v])-1]
			for i, w := range edges[u] {
				if w == v {
					edges[u] = append(edges[u][:i], edges[u][i+1:]...)
					break
				}
			}
			stack = append(stack, u)
		} else {
			path = append(path, v)
			stack = stack[:len(stack)-1]
		}
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func HamiltonianPath(g *Graph) []int {
	n := g.Vertices
	path := make([]int, n)
	visited := make([]bool, n)
	var solve func(pos int) bool
	solve = func(pos int) bool {
		if pos == n {
			return true
		}
		for v := 0; v < n; v++ {
			if !visited[v] {
				canAdd := pos == 0
				if !canAdd {
					for _, u := range g.Edges[path[pos-1]] {
						if u == v {
							canAdd = true
							break
						}
					}
				}
				if canAdd {
					path[pos] = v
					visited[v] = true
					if solve(pos + 1) {
						return true
					}
					visited[v] = false
				}
			}
		}
		return false
	}
	if solve(0) {
		return path
	}
	return nil
}

func GraphColoring(g *Graph, numColors int) []int {
	n := g.Vertices
	colors := make([]int, n)
	for i := range colors {
		colors[i] = -1
	}
	var canColor func(v, c int) bool
	canColor = func(v, c int) bool {
		for _, u := range g.Edges[v] {
			if colors[u] == c {
				return false
			}
		}
		return true
	}
	var solve func(v int) bool
	solve = func(v int) bool {
		if v == n {
			return true
		}
		for c := 0; c < numColors; c++ {
			if canColor(v, c) {
				colors[v] = c
				if solve(v + 1) {
					return true
				}
				colors[v] = -1
			}
		}
		return false
	}
	if solve(0) {
		return colors
	}
	return nil
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
