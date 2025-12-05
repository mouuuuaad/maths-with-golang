package optimization

func SimulatedAnnealing(f ObjectiveFunc, x0 []float64, temp, cooling float64, iterations int) []float64 {
	n := len(x0)
	x := make([]float64, n)
	copy(x, x0)
	best := make([]float64, n)
	copy(best, x)
	fx := f(x)
	fBest := fx
	seed := uint64(42)
	randFloat := func() float64 {
		seed = seed*1103515245 + 12345
		return float64(seed%1000000) / 1000000.0
	}
	for iter := 0; iter < iterations; iter++ {
		xNew := make([]float64, n)
		for i := 0; i < n; i++ {
			xNew[i] = x[i] + (randFloat()-0.5)*temp
		}
		fNew := f(xNew)
		delta := fNew - fx
		if delta < 0 || randFloat() < expO(-delta/temp) {
			copy(x, xNew)
			fx = fNew
			if fx < fBest {
				fBest = fx
				copy(best, x)
			}
		}
		temp *= cooling
	}
	return best
}

func expO(x float64) float64 {
	if x < -20 {
		return 0
	}
	if x > 20 {
		return 1e9
	}
	sum := 1.0
	term := 1.0
	for i := 1; i < 30; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func GeneticAlgorithm(f ObjectiveFunc, dim, popSize, generations int, bounds [][2]float64) []float64 {
	seed := uint64(42)
	randFloat := func() float64 {
		seed = seed*1103515245 + 12345
		return float64(seed%1000000) / 1000000.0
	}
	randInt := func(n int) int {
		seed = seed*1103515245 + 12345
		return int(seed % uint64(n))
	}
	pop := make([][]float64, popSize)
	fitness := make([]float64, popSize)
	for i := 0; i < popSize; i++ {
		pop[i] = make([]float64, dim)
		for d := 0; d < dim; d++ {
			pop[i][d] = bounds[d][0] + randFloat()*(bounds[d][1]-bounds[d][0])
		}
		fitness[i] = f(pop[i])
	}
	mutationRate := 0.1
	crossoverRate := 0.8
	for gen := 0; gen < generations; gen++ {
		newPop := make([][]float64, popSize)
		for i := 0; i < popSize; i++ {
			p1 := tournamentSelect(fitness, randInt)
			p2 := tournamentSelect(fitness, randInt)
			child := make([]float64, dim)
			if randFloat() < crossoverRate {
				point := randInt(dim)
				for d := 0; d < dim; d++ {
					if d < point {
						child[d] = pop[p1][d]
					} else {
						child[d] = pop[p2][d]
					}
				}
			} else {
				copy(child, pop[p1])
			}
			for d := 0; d < dim; d++ {
				if randFloat() < mutationRate {
					child[d] += (randFloat() - 0.5) * (bounds[d][1] - bounds[d][0]) * 0.1
					if child[d] < bounds[d][0] {
						child[d] = bounds[d][0]
					}
					if child[d] > bounds[d][1] {
						child[d] = bounds[d][1]
					}
				}
			}
			newPop[i] = child
		}
		pop = newPop
		for i := 0; i < popSize; i++ {
			fitness[i] = f(pop[i])
		}
	}
	best := 0
	for i := 1; i < popSize; i++ {
		if fitness[i] < fitness[best] {
			best = i
		}
	}
	return pop[best]
}

func tournamentSelect(fitness []float64, randInt func(int) int) int {
	n := len(fitness)
	a, b := randInt(n), randInt(n)
	if fitness[a] < fitness[b] {
		return a
	}
	return b
}

func TabuSearch(f ObjectiveFunc, x0 []float64, tabuSize, iterations int) []float64 {
	n := len(x0)
	x := make([]float64, n)
	copy(x, x0)
	best := make([]float64, n)
	copy(best, x)
	fBest := f(best)
	tabuList := make([][]float64, 0, tabuSize)
	seed := uint64(42)
	randFloat := func() float64 {
		seed = seed*1103515245 + 12345
		return float64(seed%1000000) / 1000000.0
	}
	for iter := 0; iter < iterations; iter++ {
		bestNeighbor := make([]float64, n)
		bestNeighborF := 1e18
		for k := 0; k < 20; k++ {
			neighbor := make([]float64, n)
			for i := 0; i < n; i++ {
				neighbor[i] = x[i] + (randFloat()-0.5)*0.5
			}
			if !isTabu(neighbor, tabuList) {
				fNeighbor := f(neighbor)
				if fNeighbor < bestNeighborF {
					bestNeighborF = fNeighbor
					copy(bestNeighbor, neighbor)
				}
			}
		}
		copy(x, bestNeighbor)
		tabuList = append(tabuList, x)
		if len(tabuList) > tabuSize {
			tabuList = tabuList[1:]
		}
		if bestNeighborF < fBest {
			fBest = bestNeighborF
			copy(best, x)
		}
	}
	return best
}

func isTabu(x []float64, tabuList [][]float64) bool {
	for _, t := range tabuList {
		same := true
		for i := range x {
			if absO(x[i]-t[i]) > 0.01 {
				same = false
				break
			}
		}
		if same {
			return true
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
