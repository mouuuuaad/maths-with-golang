package graphs

func TopologicalSort(g *Graph) []int {
	indegree := make([]int, g.Vertices)
	for u := 0; u < g.Vertices; u++ {
		for _, v := range g.Edges[u] {
			indegree[v]++
		}
	}
	queue := []int{}
	for v := 0; v < g.Vertices; v++ {
		if indegree[v] == 0 {
			queue = append(queue, v)
		}
	}
	result := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, v := range g.Edges[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return result
}

func HasCycle(g *Graph) bool {
	visited := make([]int, g.Vertices)
	for v := 0; v < g.Vertices; v++ {
		if visited[v] == 0 {
			if hasCycleDFS(g, v, visited) {
				return true
			}
		}
	}
	return false
}

func hasCycleDFS(g *Graph, v int, visited []int) bool {
	visited[v] = 1
	for _, u := range g.Edges[v] {
		if visited[u] == 1 {
			return true
		}
		if visited[u] == 0 && hasCycleDFS(g, u, visited) {
			return true
		}
	}
	visited[v] = 2
	return false
}

func StronglyConnectedComponents(g *Graph) [][]int {
	n := g.Vertices
	order := []int{}
	visited := make([]bool, n)
	var dfs1 func(v int)
	dfs1 = func(v int) {
		visited[v] = true
		for _, u := range g.Edges[v] {
			if !visited[u] {
				dfs1(u)
			}
		}
		order = append(order, v)
	}
	for v := 0; v < n; v++ {
		if !visited[v] {
			dfs1(v)
		}
	}
	reverse := NewGraph(n)
	for u := 0; u < n; u++ {
		for _, v := range g.Edges[u] {
			reverse.Edges[v] = append(reverse.Edges[v], u)
		}
	}
	visited = make([]bool, n)
	sccs := [][]int{}
	var dfs2 func(v int, comp *[]int)
	dfs2 = func(v int, comp *[]int) {
		visited[v] = true
		*comp = append(*comp, v)
		for _, u := range reverse.Edges[v] {
			if !visited[u] {
				dfs2(u, comp)
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		v := order[i]
		if !visited[v] {
			comp := []int{}
			dfs2(v, &comp)
			sccs = append(sccs, comp)
		}
	}
	return sccs
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
