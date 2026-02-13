// 2026 Update: Metaheuristics
package optimization

import "math"

type RNG struct {
	state uint64
}

func NewRNG(seed uint64) *RNG {
	if seed == 0 {
		seed = 1
	}
	return &RNG{state: seed}
}

func (r *RNG) Next() uint64 {
	r.state = r.state*6364136223846793005 + 1442695040888963407
	return r.state
}

func (r *RNG) Float64() float64 {
	return float64(r.Next()>>11) / float64(1<<53)
}

func (r *RNG) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.Next() % uint64(n))
}

func (r *RNG) Range(low, high float64) float64 {
	return low + r.Float64()*(high-low)
}

func clampValue(x, low, high float64) float64 {
	if x < low {
		return low
	}
	if x > high {
		return high
	}
	return x
}

func clampVector(x []float64, bounds [][2]float64) []float64 {
	if bounds == nil || len(bounds) != len(x) {
		return cloneVector(x)
	}
	out := make([]float64, len(x))
	for i := range x {
		low := bounds[i][0]
		high := bounds[i][1]
		if low > high {
			low, high = high, low
		}
		out[i] = clampValue(x[i], low, high)
	}
	return out
}

func randomVector(rng *RNG, dim int, bounds [][2]float64) []float64 {
	v := make([]float64, dim)
	for i := 0; i < dim; i++ {
		v[i] = rng.Range(bounds[i][0], bounds[i][1])
	}
	return v
}

func vectorCopy(dst, src []float64) {
	for i := range src {
		dst[i] = src[i]
	}
}

type Particle struct {
	Position     []float64
	Velocity     []float64
	BestPosition []float64
	BestValue    float64
}

type PSOSettings struct {
	Iterations     int
	Inertia        float64
	Cognitive      float64
	Social         float64
	VelocityClamp  float64
	PositionClamp  bool
	Neighborhood   int
	WarmStart      [][]float64
	InitialSeed    uint64
	MinImprovement float64
}

func DefaultPSOSettings() PSOSettings {
	return PSOSettings{
		Iterations:     300,
		Inertia:        0.72,
		Cognitive:      1.49,
		Social:         1.49,
		VelocityClamp:  0.0,
		PositionClamp:  true,
		Neighborhood:   0,
		WarmStart:      nil,
		InitialSeed:    42,
		MinImprovement: 1e-12,
	}
}

func PSO(f ObjectiveFunc, dim, particles, iterations int, bounds [][2]float64) []float64 {
	settings := DefaultPSOSettings()
	settings.Iterations = iterations
	return PSOWithSettings(f, dim, particles, bounds, settings)
}

func PSOWithSettings(f ObjectiveFunc, dim, particles int, bounds [][2]float64, settings PSOSettings) []float64 {
	rng := NewRNG(settings.InitialSeed)
	swarm := make([]Particle, particles)
	globalBest := make([]float64, dim)
	globalBestValue := math.Inf(1)
	for i := 0; i < particles; i++ {
		pos := randomVector(rng, dim, bounds)
		if settings.WarmStart != nil && i < len(settings.WarmStart) {
			pos = clampVector(settings.WarmStart[i], bounds)
		}
		vel := make([]float64, dim)
		for d := 0; d < dim; d++ {
			span := bounds[d][1] - bounds[d][0]
			vel[d] = (rng.Float64()*2 - 1) * span * 0.1
		}
		swarm[i] = Particle{
			Position:     pos,
			Velocity:     vel,
			BestPosition: cloneVector(pos),
			BestValue:    f(pos),
		}
		if swarm[i].BestValue < globalBestValue {
			globalBestValue = swarm[i].BestValue
			vectorCopy(globalBest, pos)
		}
	}
	bestPrev := globalBestValue
	stall := 0
	for iter := 0; iter < settings.Iterations; iter++ {
		for i := 0; i < particles; i++ {
			localBest := globalBest
			if settings.Neighborhood > 0 {
				bestIdx := i
				bestVal := swarm[i].BestValue
				start := i - settings.Neighborhood
				end := i + settings.Neighborhood
				for j := start; j <= end; j++ {
					idx := j
					if idx < 0 {
						idx += particles
					}
					if idx >= particles {
						idx -= particles
					}
					if swarm[idx].BestValue < bestVal {
						bestVal = swarm[idx].BestValue
						bestIdx = idx
					}
				}
				localBest = swarm[bestIdx].BestPosition
			}
			for d := 0; d < dim; d++ {
				r1 := rng.Float64()
				r2 := rng.Float64()
				vel := settings.Inertia*swarm[i].Velocity[d] + settings.Cognitive*r1*(swarm[i].BestPosition[d]-swarm[i].Position[d]) + settings.Social*r2*(localBest[d]-swarm[i].Position[d])
				if settings.VelocityClamp > 0 {
					vel = clampValue(vel, -settings.VelocityClamp, settings.VelocityClamp)
				}
				swarm[i].Velocity[d] = vel
				swarm[i].Position[d] += vel
				if settings.PositionClamp {
					swarm[i].Position[d] = clampValue(swarm[i].Position[d], bounds[d][0], bounds[d][1])
				}
			}
			val := f(swarm[i].Position)
			if val < swarm[i].BestValue {
				swarm[i].BestValue = val
				vectorCopy(swarm[i].BestPosition, swarm[i].Position)
				if val < globalBestValue {
					globalBestValue = val
					vectorCopy(globalBest, swarm[i].Position)
				}
			}
		}
		if bestPrev-globalBestValue < settings.MinImprovement {
			stall++
		} else {
			stall = 0
			bestPrev = globalBestValue
		}
		if stall > 50 {
			break
		}
	}
	return globalBest
}

type DEStrategy int

const (
	DERand1Bin DEStrategy = iota
	DEBest1Bin
	DERand2Bin
)

type DESettings struct {
	Population  int
	Generations int
	F           float64
	CR          float64
	Strategy    DEStrategy
	Seed        uint64
}

func DefaultDESettings(population, generations int) DESettings {
	return DESettings{
		Population:  population,
		Generations: generations,
		F:           0.8,
		CR:          0.9,
		Strategy:    DERand1Bin,
		Seed:        42,
	}
}

func DifferentialEvolution(f ObjectiveFunc, dim, population, generations int, bounds [][2]float64) []float64 {
	settings := DefaultDESettings(population, generations)
	return DifferentialEvolutionWithSettings(f, dim, bounds, settings)
}

func DifferentialEvolutionWithSettings(f ObjectiveFunc, dim int, bounds [][2]float64, settings DESettings) []float64 {
	rng := NewRNG(settings.Seed)
	pop := make([][]float64, settings.Population)
	fitness := make([]float64, settings.Population)
	bestIdx := 0
	for i := 0; i < settings.Population; i++ {
		pop[i] = randomVector(rng, dim, bounds)
		fitness[i] = f(pop[i])
		if fitness[i] < fitness[bestIdx] {
			bestIdx = i
		}
	}
	for gen := 0; gen < settings.Generations; gen++ {
		for i := 0; i < settings.Population; i++ {
			a := rng.Intn(settings.Population)
			b := rng.Intn(settings.Population)
			c := rng.Intn(settings.Population)
			d := rng.Intn(settings.Population)
			e := rng.Intn(settings.Population)
			for a == i {
				a = rng.Intn(settings.Population)
			}
			for b == i || b == a {
				b = rng.Intn(settings.Population)
			}
			for c == i || c == a || c == b {
				c = rng.Intn(settings.Population)
			}
			for d == i || d == a || d == b || d == c {
				d = rng.Intn(settings.Population)
			}
			for e == i || e == a || e == b || e == c || e == d {
				e = rng.Intn(settings.Population)
			}
			trial := make([]float64, dim)
			jRand := rng.Intn(dim)
			for j := 0; j < dim; j++ {
				if rng.Float64() < settings.CR || j == jRand {
					switch settings.Strategy {
					case DERand1Bin:
						trial[j] = pop[a][j] + settings.F*(pop[b][j]-pop[c][j])
					case DEBest1Bin:
						trial[j] = pop[bestIdx][j] + settings.F*(pop[b][j]-pop[c][j])
					case DERand2Bin:
						trial[j] = pop[a][j] + settings.F*(pop[b][j]-pop[c][j]+pop[d][j]-pop[e][j])
					}
					trial[j] = clampValue(trial[j], bounds[j][0], bounds[j][1])
				} else {
					trial[j] = pop[i][j]
				}
			}
			trialFit := f(trial)
			if trialFit < fitness[i] {
				pop[i] = trial
				fitness[i] = trialFit
				if trialFit < fitness[bestIdx] {
					bestIdx = i
				}
			}
		}
	}
	return pop[bestIdx]
}

type GASettings struct {
	Population    int
	Generations   int
	MutationRate  float64
	CrossoverRate float64
	Tournament    int
	Elitism       int
	Seed          uint64
}

func DefaultGASettings(population, generations int) GASettings {
	return GASettings{
		Population:    population,
		Generations:   generations,
		MutationRate:  0.08,
		CrossoverRate: 0.9,
		Tournament:    3,
		Elitism:       1,
		Seed:          42,
	}
}

func GeneticAlgorithm(f ObjectiveFunc, dim, popSize, generations int, bounds [][2]float64) []float64 {
	settings := DefaultGASettings(popSize, generations)
	return GeneticAlgorithmWithSettings(f, dim, bounds, settings)
}

func GeneticAlgorithmWithSettings(f ObjectiveFunc, dim int, bounds [][2]float64, settings GASettings) []float64 {
	rng := NewRNG(settings.Seed)
	pop := make([][]float64, settings.Population)
	fitness := make([]float64, settings.Population)
	for i := range pop {
		pop[i] = randomVector(rng, dim, bounds)
		fitness[i] = f(pop[i])
	}
	bestIdx := 0
	for gen := 0; gen < settings.Generations; gen++ {
		newPop := make([][]float64, 0, settings.Population)
		for e := 0; e < settings.Elitism; e++ {
			bestIdx = argmin(fitness)
			elite := cloneVector(pop[bestIdx])
			newPop = append(newPop, elite)
			fitness[bestIdx] = math.Inf(1)
		}
		for len(newPop) < settings.Population {
			p1 := tournamentSelect(rng, pop, fitness, settings.Tournament)
			p2 := tournamentSelect(rng, pop, fitness, settings.Tournament)
			child1 := cloneVector(p1)
			child2 := cloneVector(p2)
			if rng.Float64() < settings.CrossoverRate {
				c1, c2 := uniformCrossover(rng, p1, p2)
				child1, child2 = c1, c2
			}
			mutateVector(rng, child1, bounds, settings.MutationRate)
			mutateVector(rng, child2, bounds, settings.MutationRate)
			newPop = append(newPop, child1)
			if len(newPop) < settings.Population {
				newPop = append(newPop, child2)
			}
		}
		pop = newPop
		fitness = make([]float64, settings.Population)
		for i := range pop {
			fitness[i] = f(pop[i])
		}
	}
	best := 0
	for i := 1; i < len(pop); i++ {
		if fitness[i] < fitness[best] {
			best = i
		}
	}
	return pop[best]
}

func tournamentSelect(rng *RNG, pop [][]float64, fitness []float64, k int) []float64 {
	best := -1
	bestVal := math.Inf(1)
	for i := 0; i < k; i++ {
		idx := rng.Intn(len(pop))
		val := fitness[idx]
		if val < bestVal {
			bestVal = val
			best = idx
		}
	}
	return pop[best]
}

func uniformCrossover(rng *RNG, a, b []float64) ([]float64, []float64) {
	c1 := make([]float64, len(a))
	c2 := make([]float64, len(a))
	for i := range a {
		if rng.Float64() < 0.5 {
			c1[i] = a[i]
			c2[i] = b[i]
		} else {
			c1[i] = b[i]
			c2[i] = a[i]
		}
	}
	return c1, c2
}

func mutateVector(rng *RNG, x []float64, bounds [][2]float64, rate float64) {
	for i := range x {
		if rng.Float64() < rate {
			span := bounds[i][1] - bounds[i][0]
			shift := (rng.Float64()*2 - 1) * 0.2 * span
			x[i] = clampValue(x[i]+shift, bounds[i][0], bounds[i][1])
		}
	}
}

type AnnealSettings struct {
	InitialTemp float64
	MinTemp     float64
	Alpha       float64
	Iterations  int
	StepScale   float64
	Seed        uint64
}

func DefaultAnnealSettings() AnnealSettings {
	return AnnealSettings{
		InitialTemp: 1.0,
		MinTemp:     1e-4,
		Alpha:       0.95,
		Iterations:  500,
		StepScale:   0.1,
		Seed:        42,
	}
}

func SimulatedAnnealing(f ObjectiveFunc, x0 []float64, temp, cooling float64, iterations int) []float64 {
	settings := DefaultAnnealSettings()
	settings.InitialTemp = temp
	settings.Alpha = cooling
	settings.Iterations = iterations
	return SimulatedAnnealingWithSettings(f, x0, nil, settings)
}

func SimulatedAnnealingWithSettings(f ObjectiveFunc, x0 []float64, bounds [][2]float64, settings AnnealSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := clampVector(x0, bounds)
	best := cloneVector(x)
	bestVal := f(best)
	currentVal := bestVal
	temp := settings.InitialTemp
	for temp > settings.MinTemp {
		for i := 0; i < settings.Iterations; i++ {
			cand := neighborVector(rng, x, bounds, settings.StepScale)
			val := f(cand)
			delta := val - currentVal
			if delta < 0 || rng.Float64() < math.Exp(-delta/temp) {
				x = cand
				currentVal = val
				if val < bestVal {
					bestVal = val
					best = cloneVector(cand)
				}
			}
		}
		temp *= settings.Alpha
	}
	return best
}

func neighborVector(rng *RNG, x []float64, bounds [][2]float64, step float64) []float64 {
	out := cloneVector(x)
	idx := rng.Intn(len(out))
	if bounds == nil || len(bounds) != len(out) {
		scale := 1 + absO(out[idx])
		shift := (rng.Float64()*2 - 1) * step * scale
		out[idx] = out[idx] + shift
		return out
	}
	span := bounds[idx][1] - bounds[idx][0]
	shift := (rng.Float64()*2 - 1) * step * span
	out[idx] = clampValue(out[idx]+shift, bounds[idx][0], bounds[idx][1])
	return out
}

type TabuSettings struct {
	Iterations int
	TabuSize   int
	StepScale  float64
	Seed       uint64
}

func DefaultTabuSettings() TabuSettings {
	return TabuSettings{
		Iterations: 400,
		TabuSize:   20,
		StepScale:  0.15,
		Seed:       42,
	}
}

func TabuSearch(f ObjectiveFunc, x0 []float64, tabuSize, iterations int) []float64 {
	settings := DefaultTabuSettings()
	settings.TabuSize = tabuSize
	settings.Iterations = iterations
	return TabuSearchWithSettings(f, x0, nil, settings)
}

func TabuSearchWithSettings(f ObjectiveFunc, x0 []float64, bounds [][2]float64, settings TabuSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := clampVector(x0, bounds)
	best := cloneVector(x)
	bestVal := f(best)
	mem := make([][]float64, 0, settings.TabuSize)
	for iter := 0; iter < settings.Iterations; iter++ {
		candidates := make([][]float64, 0, 8)
		for i := 0; i < 8; i++ {
			cand := neighborVector(rng, x, bounds, settings.StepScale)
			candidates = append(candidates, cand)
		}
		selected := candidates[0]
		selectedVal := f(selected)
		for i := 1; i < len(candidates); i++ {
			val := f(candidates[i])
			if val < selectedVal {
				selectedVal = val
				selected = candidates[i]
			}
		}
		if !isTabu(selected, mem) {
			x = selected
			if selectedVal < bestVal {
				bestVal = selectedVal
				best = cloneVector(selected)
			}
			mem = append(mem, cloneVector(selected))
			if len(mem) > settings.TabuSize {
				mem = mem[1:]
			}
		}
	}
	return best
}

func isTabu(x []float64, tabu [][]float64) bool {
	for _, t := range tabu {
		if vectorNormInf(subVectors(x, t)) < 1e-9 {
			return true
		}
	}
	return false
}

type HarmonySettings struct {
	MemorySize   int
	Iterations   int
	HMCR         float64
	PAR          float64
	Bandwidth    float64
	Seed         uint64
	ImproveLimit float64
}

func DefaultHarmonySettings() HarmonySettings {
	return HarmonySettings{
		MemorySize:   20,
		Iterations:   300,
		HMCR:         0.9,
		PAR:          0.3,
		Bandwidth:    0.01,
		Seed:         42,
		ImproveLimit: 1e-12,
	}
}

func HarmonySearch(f ObjectiveFunc, dim int, bounds [][2]float64, settings HarmonySettings) []float64 {
	rng := NewRNG(settings.Seed)
	memory := make([][]float64, settings.MemorySize)
	values := make([]float64, settings.MemorySize)
	best := 0
	for i := 0; i < settings.MemorySize; i++ {
		memory[i] = randomVector(rng, dim, bounds)
		values[i] = f(memory[i])
		if values[i] < values[best] {
			best = i
		}
	}
	bestVal := values[best]
	stall := 0
	for iter := 0; iter < settings.Iterations; iter++ {
		newHarmony := make([]float64, dim)
		for j := 0; j < dim; j++ {
			if rng.Float64() < settings.HMCR {
				idx := rng.Intn(settings.MemorySize)
				newHarmony[j] = memory[idx][j]
				if rng.Float64() < settings.PAR {
					newHarmony[j] += (rng.Float64()*2 - 1) * settings.Bandwidth
				}
			} else {
				newHarmony[j] = rng.Range(bounds[j][0], bounds[j][1])
			}
			newHarmony[j] = clampValue(newHarmony[j], bounds[j][0], bounds[j][1])
		}
		val := f(newHarmony)
		worst := argmax(values)
		if val < values[worst] {
			memory[worst] = newHarmony
			values[worst] = val
			if val < bestVal {
				bestVal = val
				best = worst
			}
		}
		if absO(bestVal-values[best]) < settings.ImproveLimit {
			stall++
		} else {
			stall = 0
			bestVal = values[best]
		}
		if stall > 60 {
			break
		}
	}
	return memory[best]
}
