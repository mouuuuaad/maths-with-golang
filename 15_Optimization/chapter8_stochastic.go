// 2026 Update: Stochastic Optimization
package optimization

import "math"

type StochasticSettings struct {
	Samples    int
	Iterations int
	Seed       uint64
}

func DefaultStochasticSettings() StochasticSettings {
	return StochasticSettings{
		Samples:    500,
		Iterations: 200,
		Seed:       42,
	}
}

func stochClone(x []float64) []float64 {
	out := make([]float64, len(x))
	copy(out, x)
	return out
}

func stochClamp(x []float64, bounds [][2]float64) []float64 {
	out := make([]float64, len(x))
	if bounds == nil || len(bounds) != len(x) {
		copy(out, x)
		return out
	}
	for i := range x {
		low := bounds[i][0]
		high := bounds[i][1]
		if low > high {
			low, high = high, low
		}
		v := x[i]
		if v < low {
			v = low
		}
		if v > high {
			v = high
		}
		out[i] = v
	}
	return out
}

func stochMean(samples [][]float64) []float64 {
	if len(samples) == 0 {
		return nil
	}
	dim := len(samples[0])
	mean := make([]float64, dim)
	for i := range samples {
		for j := 0; j < dim; j++ {
			mean[j] += samples[i][j]
		}
	}
	inv := 1.0 / float64(len(samples))
	for j := 0; j < dim; j++ {
		mean[j] *= inv
	}
	return mean
}

func stochStd(samples [][]float64, mean []float64) []float64 {
	if len(samples) == 0 {
		return nil
	}
	dim := len(samples[0])
	std := make([]float64, dim)
	for i := range samples {
		for j := 0; j < dim; j++ {
			d := samples[i][j] - mean[j]
			std[j] += d * d
		}
	}
	inv := 1.0 / float64(len(samples))
	for j := 0; j < dim; j++ {
		std[j] = math.Sqrt(std[j]*inv + 1e-12)
	}
	return std
}

func stochUniform(rng *RNG, dim int, bounds [][2]float64) []float64 {
	out := make([]float64, dim)
	for i := 0; i < dim; i++ {
		low := bounds[i][0]
		high := bounds[i][1]
		if low > high {
			low, high = high, low
		}
		out[i] = rng.Range(low, high)
	}
	return out
}

func RandomSearch(f ObjectiveFunc, dim int, bounds [][2]float64, settings StochasticSettings) []float64 {
	rng := NewRNG(settings.Seed)
	best := stochUniform(rng, dim, bounds)
	bestVal := f(best)
	for i := 0; i < settings.Samples; i++ {
		cand := stochUniform(rng, dim, bounds)
		val := f(cand)
		if val < bestVal {
			bestVal = val
			best = cand
		}
	}
	return best
}

func RandomSearchRestart(f ObjectiveFunc, dim int, bounds [][2]float64, settings StochasticSettings) []float64 {
	best := RandomSearch(f, dim, bounds, settings)
	bestVal := f(best)
	for i := 0; i < settings.Iterations; i++ {
		cand := RandomSearch(f, dim, bounds, settings)
		val := f(cand)
		if val < bestVal {
			bestVal = val
			best = cand
		}
	}
	return best
}

type CEMSettings struct {
	Samples    int
	EliteRatio float64
	Iterations int
	Seed       uint64
	InitStd    float64
	MinStd     float64
	Bounds     [][2]float64
}

func DefaultCEMSettings(samples int, bounds [][2]float64) CEMSettings {
	return CEMSettings{
		Samples:    samples,
		EliteRatio: 0.2,
		Iterations: 100,
		Seed:       42,
		InitStd:    1.0,
		MinStd:     1e-4,
		Bounds:     bounds,
	}
}

func CrossEntropyMethod(f ObjectiveFunc, mean []float64, settings CEMSettings) []float64 {
	rng := NewRNG(settings.Seed)
	dim := len(mean)
	std := make([]float64, dim)
	for i := range std {
		std[i] = settings.InitStd
	}
	best := stochClone(mean)
	bestVal := f(best)
	for iter := 0; iter < settings.Iterations; iter++ {
		samples := make([][]float64, settings.Samples)
		values := make([]float64, settings.Samples)
		for i := 0; i < settings.Samples; i++ {
			cand := make([]float64, dim)
			for j := 0; j < dim; j++ {
				cand[j] = mean[j] + std[j]*normalSample(rng)
			}
			cand = stochClamp(cand, settings.Bounds)
			samples[i] = cand
			values[i] = f(cand)
			if values[i] < bestVal {
				bestVal = values[i]
				best = stochClone(cand)
			}
		}
		eliteCount := int(math.Max(1, math.Round(float64(settings.Samples)*settings.EliteRatio)))
		eliteIdx := argsort(values)
		elite := make([][]float64, eliteCount)
		for i := 0; i < eliteCount; i++ {
			elite[i] = samples[eliteIdx[i]]
		}
		mean = stochMean(elite)
		std = stochStd(elite, mean)
		for j := range std {
			if std[j] < settings.MinStd {
				std[j] = settings.MinStd
			}
		}
	}
	return best
}

func argsort(values []float64) []int {
	idx := make([]int, len(values))
	for i := range idx {
		idx[i] = i
	}
	for i := 0; i < len(idx); i++ {
		for j := i + 1; j < len(idx); j++ {
			if values[idx[j]] < values[idx[i]] {
				idx[i], idx[j] = idx[j], idx[i]
			}
		}
	}
	return idx
}

func normalSample(rng *RNG) float64 {
	u1 := rng.Float64()
	u2 := rng.Float64()
	if u1 < 1e-12 {
		u1 = 1e-12
	}
	r := math.Sqrt(-2 * math.Log(u1))
	t := 2 * math.Pi * u2
	return r * math.Cos(t)
}

type CMAESSettings struct {
	Population int
	Iterations int
	Seed       uint64
	Sigma      float64
	Bounds     [][2]float64
}

func DefaultCMAESSettings(population int, bounds [][2]float64) CMAESSettings {
	return CMAESSettings{
		Population: population,
		Iterations: 150,
		Seed:       42,
		Sigma:      0.5,
		Bounds:     bounds,
	}
}

func CMAESDiagonal(f ObjectiveFunc, mean []float64, settings CMAESSettings) []float64 {
	rng := NewRNG(settings.Seed)
	dim := len(mean)
	weights := make([]float64, settings.Population)
	for i := 0; i < settings.Population; i++ {
		weights[i] = math.Log(float64(settings.Population)+0.5) - math.Log(float64(i)+1)
		if weights[i] < 0 {
			weights[i] = 0
		}
	}
	wsum := 0.0
	for _, w := range weights {
		wsum += w
	}
	for i := range weights {
		weights[i] /= wsum
	}
	sigma := settings.Sigma
	diag := make([]float64, dim)
	for i := range diag {
		diag[i] = 1
	}
	best := stochClone(mean)
	bestVal := f(best)
	for iter := 0; iter < settings.Iterations; iter++ {
		pop := make([][]float64, settings.Population)
		vals := make([]float64, settings.Population)
		for i := 0; i < settings.Population; i++ {
			cand := make([]float64, dim)
			for j := 0; j < dim; j++ {
				cand[j] = mean[j] + sigma*diag[j]*normalSample(rng)
			}
			cand = stochClamp(cand, settings.Bounds)
			pop[i] = cand
			vals[i] = f(cand)
			if vals[i] < bestVal {
				bestVal = vals[i]
				best = stochClone(cand)
			}
		}
		order := argsort(vals)
		newMean := make([]float64, dim)
		for i := 0; i < settings.Population; i++ {
			w := weights[i]
			idx := order[i]
			for j := 0; j < dim; j++ {
				newMean[j] += w * pop[idx][j]
			}
		}
		mean = newMean
		for j := 0; j < dim; j++ {
			varsum := 0.0
			for i := 0; i < settings.Population; i++ {
				d := pop[order[i]][j] - mean[j]
				varsum += weights[i] * d * d
			}
			diag[j] = math.Sqrt(varsum + 1e-12)
			if diag[j] < 1e-6 {
				diag[j] = 1e-6
			}
		}
		sigma *= 0.99
		if sigma < 1e-4 {
			sigma = 1e-4
		}
	}
	return best
}

type SPSASettings struct {
	Iterations int
	A          float64
	C          float64
	Alpha      float64
	Gamma      float64
	Seed       uint64
}

func DefaultSPSASettings() SPSASettings {
	return SPSASettings{
		Iterations: 200,
		A:          0.2,
		C:          0.1,
		Alpha:      0.602,
		Gamma:      0.101,
		Seed:       42,
	}
}

func SPSA(f ObjectiveFunc, x0 []float64, bounds [][2]float64, settings SPSASettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := stochClone(x0)
	for k := 0; k < settings.Iterations; k++ {
		ak := settings.A / math.Pow(float64(k)+1, settings.Alpha)
		ck := settings.C / math.Pow(float64(k)+1, settings.Gamma)
		delta := make([]float64, len(x))
		for i := range delta {
			if rng.Float64() < 0.5 {
				delta[i] = -1
			} else {
				delta[i] = 1
			}
		}
		xPlus := make([]float64, len(x))
		xMinus := make([]float64, len(x))
		for i := range x {
			xPlus[i] = x[i] + ck*delta[i]
			xMinus[i] = x[i] - ck*delta[i]
		}
		xPlus = stochClamp(xPlus, bounds)
		xMinus = stochClamp(xMinus, bounds)
		fPlus := f(xPlus)
		fMinus := f(xMinus)
		for i := range x {
			g := (fPlus - fMinus) / (2 * ck * delta[i])
			x[i] = x[i] - ak*g
		}
		x = stochClamp(x, bounds)
	}
	return x
}

type SampleGrad func(x []float64, idx int) []float64

type SGDSettings struct {
	Steps        int
	LearningRate float64
	BatchSize    int
	DataSize     int
	Seed         uint64
}

func DefaultSGDSettings(dataSize int) SGDSettings {
	return SGDSettings{
		Steps:        500,
		LearningRate: 0.01,
		BatchSize:    16,
		DataSize:     dataSize,
		Seed:         42,
	}
}

func StochasticGradientDescent(grad SampleGrad, x0 []float64, settings SGDSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := stochClone(x0)
	if settings.DataSize <= 0 {
		return x
	}
	for step := 0; step < settings.Steps; step++ {
		g := make([]float64, len(x))
		for b := 0; b < settings.BatchSize; b++ {
			idx := rng.Intn(settings.DataSize)
			gi := grad(x, idx)
			for i := range x {
				g[i] += gi[i]
			}
		}
		inv := 1.0 / float64(settings.BatchSize)
		for i := range x {
			x[i] -= settings.LearningRate * g[i] * inv
		}
	}
	return x
}

type MomentumSettings struct {
	Steps        int
	LearningRate float64
	BatchSize    int
	DataSize     int
	Momentum     float64
	Seed         uint64
}

func DefaultMomentumSettings(dataSize int) MomentumSettings {
	return MomentumSettings{
		Steps:        600,
		LearningRate: 0.01,
		BatchSize:    16,
		DataSize:     dataSize,
		Momentum:     0.9,
		Seed:         42,
	}
}

func MomentumSGD(grad SampleGrad, x0 []float64, settings MomentumSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := stochClone(x0)
	v := make([]float64, len(x))
	if settings.DataSize <= 0 {
		return x
	}
	for step := 0; step < settings.Steps; step++ {
		g := make([]float64, len(x))
		for b := 0; b < settings.BatchSize; b++ {
			idx := rng.Intn(settings.DataSize)
			gi := grad(x, idx)
			for i := range x {
				g[i] += gi[i]
			}
		}
		inv := 1.0 / float64(settings.BatchSize)
		for i := range x {
			v[i] = settings.Momentum*v[i] + settings.LearningRate*g[i]*inv
			x[i] -= v[i]
		}
	}
	return x
}

type RMSPropSettings struct {
	Steps        int
	LearningRate float64
	BatchSize    int
	DataSize     int
	Decay        float64
	Epsilon      float64
	Seed         uint64
}

func DefaultRMSPropSettings(dataSize int) RMSPropSettings {
	return RMSPropSettings{
		Steps:        700,
		LearningRate: 0.001,
		BatchSize:    16,
		DataSize:     dataSize,
		Decay:        0.9,
		Epsilon:      1e-8,
		Seed:         42,
	}
}

func RMSProp(grad SampleGrad, x0 []float64, settings RMSPropSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := stochClone(x0)
	avg := make([]float64, len(x))
	if settings.DataSize <= 0 {
		return x
	}
	for step := 0; step < settings.Steps; step++ {
		g := make([]float64, len(x))
		for b := 0; b < settings.BatchSize; b++ {
			idx := rng.Intn(settings.DataSize)
			gi := grad(x, idx)
			for i := range x {
				g[i] += gi[i]
			}
		}
		inv := 1.0 / float64(settings.BatchSize)
		for i := range x {
			g[i] *= inv
			avg[i] = settings.Decay*avg[i] + (1-settings.Decay)*g[i]*g[i]
			adj := settings.LearningRate * g[i] / (math.Sqrt(avg[i]) + settings.Epsilon)
			x[i] -= adj
		}
	}
	return x
}

type AdamSettings struct {
	Steps        int
	LearningRate float64
	BatchSize    int
	DataSize     int
	Beta1        float64
	Beta2        float64
	Epsilon      float64
	Seed         uint64
}

func DefaultAdamSettings(dataSize int) AdamSettings {
	return AdamSettings{
		Steps:        800,
		LearningRate: 0.001,
		BatchSize:    16,
		DataSize:     dataSize,
		Beta1:        0.9,
		Beta2:        0.999,
		Epsilon:      1e-8,
		Seed:         42,
	}
}

func AdamOptimizer(grad SampleGrad, x0 []float64, settings AdamSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := stochClone(x0)
	m := make([]float64, len(x))
	v := make([]float64, len(x))
	if settings.DataSize <= 0 {
		return x
	}
	beta1 := settings.Beta1
	beta2 := settings.Beta2
	for step := 1; step <= settings.Steps; step++ {
		g := make([]float64, len(x))
		for b := 0; b < settings.BatchSize; b++ {
			idx := rng.Intn(settings.DataSize)
			gi := grad(x, idx)
			for i := range x {
				g[i] += gi[i]
			}
		}
		inv := 1.0 / float64(settings.BatchSize)
		for i := range x {
			g[i] *= inv
			m[i] = beta1*m[i] + (1-beta1)*g[i]
			v[i] = beta2*v[i] + (1-beta2)*g[i]*g[i]
			mhat := m[i] / (1 - math.Pow(beta1, float64(step)))
			vhat := v[i] / (1 - math.Pow(beta2, float64(step)))
			adj := settings.LearningRate * mhat / (math.Sqrt(vhat) + settings.Epsilon)
			x[i] -= adj
		}
	}
	return x
}

type HillSettings struct {
	Iterations int
	StepScale  float64
	Seed       uint64
}

func DefaultHillSettings() HillSettings {
	return HillSettings{
		Iterations: 500,
		StepScale:  0.2,
		Seed:       42,
	}
}

func StochasticHillClimb(f ObjectiveFunc, x0 []float64, bounds [][2]float64, settings HillSettings) []float64 {
	rng := NewRNG(settings.Seed)
	x := stochClamp(x0, bounds)
	best := stochClone(x)
	bestVal := f(best)
	for iter := 0; iter < settings.Iterations; iter++ {
		cand := stochClone(x)
		idx := rng.Intn(len(cand))
		span := bounds[idx][1] - bounds[idx][0]
		shift := (rng.Float64()*2 - 1) * settings.StepScale * span
		cand[idx] = clampValue(cand[idx]+shift, bounds[idx][0], bounds[idx][1])
		val := f(cand)
		if val < bestVal {
			bestVal = val
			best = stochClone(cand)
			x = cand
		}
	}
	return best
}
