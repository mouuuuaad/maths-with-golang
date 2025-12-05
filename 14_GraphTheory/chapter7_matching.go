package graphs

func BipartiteMatching(adj [][]int) int {
	n := len(adj)
	match := make([]int, n)
	for i := range match {
		match[i] = -1
	}
	maxMatching := 0
	for u := 0; u < n; u++ {
		visited := make([]bool, n)
		if augment(adj, u, match, visited) {
			maxMatching++
		}
	}
	return maxMatching
}

func augment(adj [][]int, u int, match []int, visited []bool) bool {
	for _, v := range adj[u] {
		if !visited[v] {
			visited[v] = true
			if match[v] == -1 || augment(adj, match[v], match, visited) {
				match[v] = u
				return true
			}
		}
	}
	return false
}

func ArticulationPoints(g *Graph) []int {
	n := g.Vertices
	visited := make([]bool, n)
	disc := make([]int, n)
	low := make([]int, n)
	parent := make([]int, n)
	ap := make([]bool, n)
	timer := 0
	for i := range parent {
		parent[i] = -1
	}
	var dfs func(u int)
	dfs = func(u int) {
		children := 0
		visited[u] = true
		disc[u] = timer
		low[u] = timer
		timer++
		for _, v := range g.Edges[u] {
			if !visited[v] {
				children++
				parent[v] = u
				dfs(v)
				if low[v] < low[u] {
					low[u] = low[v]
				}
				if parent[u] == -1 && children > 1 {
					ap[u] = true
				}
				if parent[u] != -1 && low[v] >= disc[u] {
					ap[u] = true
				}
			} else if v != parent[u] {
				if disc[v] < low[u] {
					low[u] = disc[v]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	result := []int{}
	for i, isAP := range ap {
		if isAP {
			result = append(result, i)
		}
	}
	return result
}

func Bridges(g *Graph) [][2]int {
	n := g.Vertices
	visited := make([]bool, n)
	disc := make([]int, n)
	low := make([]int, n)
	parent := make([]int, n)
	bridges := [][2]int{}
	timer := 0
	for i := range parent {
		parent[i] = -1
	}
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		disc[u] = timer
		low[u] = timer
		timer++
		for _, v := range g.Edges[u] {
			if !visited[v] {
				parent[v] = u
				dfs(v)
				if low[v] < low[u] {
					low[u] = low[v]
				}
				if low[v] > disc[u] {
					bridges = append(bridges, [2]int{u, v})
				}
			} else if v != parent[u] {
				if disc[v] < low[u] {
					low[u] = disc[v]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	return bridges
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
