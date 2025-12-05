package graphs

const INF = 1e18

type WeightedGraph struct {
	Vertices int
	Edges    [][][2]int
}

func NewWeightedGraph(n int) *WeightedGraph {
	edges := make([][][2]int, n)
	for i := range edges {
		edges[i] = [][2]int{}
	}
	return &WeightedGraph{Vertices: n, Edges: edges}
}

func (g *WeightedGraph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], [2]int{v, w})
	g.Edges[v] = append(g.Edges[v], [2]int{u, w})
}

func (g *WeightedGraph) Dijkstra(start int) []float64 {
	dist := make([]float64, g.Vertices)
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0
	visited := make([]bool, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		u := -1
		for v := 0; v < g.Vertices; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}
		if dist[u] == INF {
			break
		}
		visited[u] = true
		for _, e := range g.Edges[u] {
			v, w := e[0], e[1]
			if dist[u]+float64(w) < dist[v] {
				dist[v] = dist[u] + float64(w)
			}
		}
	}
	return dist
}

func FloydWarshall(adjMatrix [][]float64) [][]float64 {
	n := len(adjMatrix)
	dist := make([][]float64, n)
	for i := range dist {
		dist[i] = make([]float64, n)
		copy(dist[i], adjMatrix[i])
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
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
