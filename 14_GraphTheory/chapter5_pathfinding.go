package graphs

func BellmanFord(g *Graph, weights [][]int, start int) ([]float64, bool) {
	n := g.Vertices
	dist := make([]float64, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0
	for i := 0; i < n-1; i++ {
		for u := 0; u < n; u++ {
			for _, v := range g.Edges[u] {
				if dist[u]+float64(weights[u][v]) < dist[v] {
					dist[v] = dist[u] + float64(weights[u][v])
				}
			}
		}
	}
	for u := 0; u < n; u++ {
		for _, v := range g.Edges[u] {
			if dist[u]+float64(weights[u][v]) < dist[v] {
				return nil, false
			}
		}
	}
	return dist, true
}

func AStar(g *WeightedGraph, start, goal int, heuristic func(int) float64) []int {
	n := g.Vertices
	gScore := make([]float64, n)
	fScore := make([]float64, n)
	cameFrom := make([]int, n)
	inOpenSet := make([]bool, n)
	for i := 0; i < n; i++ {
		gScore[i] = INF
		fScore[i] = INF
		cameFrom[i] = -1
	}
	gScore[start] = 0
	fScore[start] = heuristic(start)
	inOpenSet[start] = true
	for {
		current := -1
		for v := 0; v < n; v++ {
			if inOpenSet[v] && (current == -1 || fScore[v] < fScore[current]) {
				current = v
			}
		}
		if current == -1 {
			return nil
		}
		if current == goal {
			path := []int{current}
			for cameFrom[current] != -1 {
				current = cameFrom[current]
				path = append([]int{current}, path...)
			}
			return path
		}
		inOpenSet[current] = false
		for _, e := range g.Edges[current] {
			neighbor, weight := e[0], e[1]
			tentative := gScore[current] + float64(weight)
			if tentative < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentative
				fScore[neighbor] = tentative + heuristic(neighbor)
				inOpenSet[neighbor] = true
			}
		}
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
// Created by: MOUAAD IDOUFKIR
// << The universe runs on equations. We just translate them >>
