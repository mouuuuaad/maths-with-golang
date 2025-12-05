package graphs

func MaxFlow(capacity [][]int, source, sink int) int {
	n := len(capacity)
	residual := make([][]int, n)
	for i := range residual {
		residual[i] = make([]int, n)
		copy(residual[i], capacity[i])
	}
	parent := make([]int, n)
	maxFlow := 0
	for {
		for i := range parent {
			parent[i] = -1
		}
		parent[source] = source
		queue := []int{source}
		for len(queue) > 0 && parent[sink] == -1 {
			u := queue[0]
			queue = queue[1:]
			for v := 0; v < n; v++ {
				if parent[v] == -1 && residual[u][v] > 0 {
					parent[v] = u
					queue = append(queue, v)
				}
			}
		}
		if parent[sink] == -1 {
			break
		}
		pathFlow := residual[parent[sink]][sink]
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			if residual[u][v] < pathFlow {
				pathFlow = residual[u][v]
			}
		}
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			residual[u][v] -= pathFlow
			residual[v][u] += pathFlow
		}
		maxFlow += pathFlow
	}
	return maxFlow
}

func MinCut(capacity [][]int, source, sink int) [][]int {
	n := len(capacity)
	MaxFlow(capacity, source, sink)
	visited := make([]bool, n)
	queue := []int{source}
	visited[source] = true
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for v := 0; v < n; v++ {
			if !visited[v] && capacity[u][v] > 0 {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	cut := [][]int{}
	for u := 0; u < n; u++ {
		if visited[u] {
			for v := 0; v < n; v++ {
				if !visited[v] && capacity[u][v] > 0 {
					cut = append(cut, []int{u, v})
				}
			}
		}
	}
	return cut
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
