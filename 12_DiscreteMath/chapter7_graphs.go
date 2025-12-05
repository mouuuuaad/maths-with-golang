package discrete

type SimpleGraph struct {
	V     int
	Edges [][]bool
}

func NewSimpleGraph(n int) *SimpleGraph {
	edges := make([][]bool, n)
	for i := range edges {
		edges[i] = make([]bool, n)
	}
	return &SimpleGraph{V: n, Edges: edges}
}

func (g *SimpleGraph) AddEdge(u, v int) {
	g.Edges[u][v] = true
	g.Edges[v][u] = true
}

func (g *SimpleGraph) Degree(v int) int {
	count := 0
	for _, e := range g.Edges[v] {
		if e {
			count++
		}
	}
	return count
}

func (g *SimpleGraph) IsConnected() bool {
	if g.V == 0 {
		return true
	}
	visited := make([]bool, g.V)
	queue := []int{0}
	visited[0] = true
	count := 1
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for u := 0; u < g.V; u++ {
			if g.Edges[v][u] && !visited[u] {
				visited[u] = true
				count++
				queue = append(queue, u)
			}
		}
	}
	return count == g.V
}

func (g *SimpleGraph) IsEulerian() bool {
	if !g.IsConnected() {
		return false
	}
	for v := 0; v < g.V; v++ {
		if g.Degree(v)%2 != 0 {
			return false
		}
	}
	return true
}

func (g *SimpleGraph) IsBipartite() bool {
	color := make([]int, g.V)
	for i := range color {
		color[i] = -1
	}
	for start := 0; start < g.V; start++ {
		if color[start] == -1 {
			queue := []int{start}
			color[start] = 0
			for len(queue) > 0 {
				v := queue[0]
				queue = queue[1:]
				for u := 0; u < g.V; u++ {
					if g.Edges[v][u] {
						if color[u] == -1 {
							color[u] = 1 - color[v]
							queue = append(queue, u)
						} else if color[u] == color[v] {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func (g *SimpleGraph) ChromaticNumber() int {
	for k := 1; k <= g.V; k++ {
		if g.canColor(k) {
			return k
		}
	}
	return g.V
}

func (g *SimpleGraph) canColor(k int) bool {
	colors := make([]int, g.V)
	return g.colorHelper(0, k, colors)
}

func (g *SimpleGraph) colorHelper(v, k int, colors []int) bool {
	if v == g.V {
		return true
	}
	for c := 1; c <= k; c++ {
		valid := true
		for u := 0; u < v; u++ {
			if g.Edges[v][u] && colors[u] == c {
				valid = false
				break
			}
		}
		if valid {
			colors[v] = c
			if g.colorHelper(v+1, k, colors) {
				return true
			}
		}
	}
	return false
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
