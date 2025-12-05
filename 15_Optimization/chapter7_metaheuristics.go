package optimization

type Particle struct {
	Position     []float64
	Velocity     []float64
	BestPosition []float64
	BestValue    float64
}

func PSO(f ObjectiveFunc, dim, particles, iterations int, bounds [][2]float64) []float64 {
	swarm := make([]Particle, particles)
	globalBest := make([]float64, dim)
	globalBestValue := 1e18
	w, c1, c2 := 0.7, 1.5, 1.5
	seed := uint64(42)
	randFloat := func() float64 {
		seed = seed*1103515245 + 12345
		return float64(seed%1000000) / 1000000.0
	}
	for i := 0; i < particles; i++ {
		swarm[i].Position = make([]float64, dim)
		swarm[i].Velocity = make([]float64, dim)
		swarm[i].BestPosition = make([]float64, dim)
		for d := 0; d < dim; d++ {
			swarm[i].Position[d] = bounds[d][0] + randFloat()*(bounds[d][1]-bounds[d][0])
			swarm[i].Velocity[d] = (randFloat() - 0.5) * (bounds[d][1] - bounds[d][0]) * 0.1
			swarm[i].BestPosition[d] = swarm[i].Position[d]
		}
		swarm[i].BestValue = f(swarm[i].Position)
		if swarm[i].BestValue < globalBestValue {
			globalBestValue = swarm[i].BestValue
			copy(globalBest, swarm[i].Position)
		}
	}
	for iter := 0; iter < iterations; iter++ {
		for i := 0; i < particles; i++ {
			for d := 0; d < dim; d++ {
				r1, r2 := randFloat(), randFloat()
				swarm[i].Velocity[d] = w*swarm[i].Velocity[d] + c1*r1*(swarm[i].BestPosition[d]-swarm[i].Position[d]) + c2*r2*(globalBest[d]-swarm[i].Position[d])
				swarm[i].Position[d] += swarm[i].Velocity[d]
				if swarm[i].Position[d] < bounds[d][0] {
					swarm[i].Position[d] = bounds[d][0]
				}
				if swarm[i].Position[d] > bounds[d][1] {
					swarm[i].Position[d] = bounds[d][1]
				}
			}
			val := f(swarm[i].Position)
			if val < swarm[i].BestValue {
				swarm[i].BestValue = val
				copy(swarm[i].BestPosition, swarm[i].Position)
				if val < globalBestValue {
					globalBestValue = val
					copy(globalBest, swarm[i].Position)
				}
			}
		}
	}
	return globalBest
}

func DifferentialEvolution(f ObjectiveFunc, dim, population, generations int, bounds [][2]float64) []float64 {
	pop := make([][]float64, population)
	fitness := make([]float64, population)
	seed := uint64(42)
	randFloat := func() float64 {
		seed = seed*1103515245 + 12345
		return float64(seed%1000000) / 1000000.0
	}
	randInt := func(n int) int {
		seed = seed*1103515245 + 12345
		return int(seed % uint64(n))
	}
	for i := 0; i < population; i++ {
		pop[i] = make([]float64, dim)
		for d := 0; d < dim; d++ {
			pop[i][d] = bounds[d][0] + randFloat()*(bounds[d][1]-bounds[d][0])
		}
		fitness[i] = f(pop[i])
	}
	F, CR := 0.8, 0.9
	for gen := 0; gen < generations; gen++ {
		for i := 0; i < population; i++ {
			a, b, c := i, i, i
			for a == i {
				a = randInt(population)
			}
			for b == i || b == a {
				b = randInt(population)
			}
			for c == i || c == a || c == b {
				c = randInt(population)
			}
			trial := make([]float64, dim)
			jRand := randInt(dim)
			for d := 0; d < dim; d++ {
				if d == jRand || randFloat() < CR {
					trial[d] = pop[a][d] + F*(pop[b][d]-pop[c][d])
					if trial[d] < bounds[d][0] {
						trial[d] = bounds[d][0]
					}
					if trial[d] > bounds[d][1] {
						trial[d] = bounds[d][1]
					}
				} else {
					trial[d] = pop[i][d]
				}
			}
			trialFit := f(trial)
			if trialFit < fitness[i] {
				pop[i] = trial
				fitness[i] = trialFit
			}
		}
	}
	best := 0
	for i := 1; i < population; i++ {
		if fitness[i] < fitness[best] {
			best = i
		}
	}
	return pop[best]
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
