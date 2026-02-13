// 2026 Update: Linear Programming
package optimization

import "math"

type SimplexSettings struct {
	MaxIter int
	Tol     float64
	Bland   bool
}

func DefaultSimplexSettings() SimplexSettings {
	return SimplexSettings{
		MaxIter: 1000,
		Tol:     1e-9,
		Bland:   true,
	}
}

type SimplexResult struct {
	X          []float64
	Objective  float64
	Iterations int
	Status     string
}

func Simplex(c []float64, A [][]float64, b []float64) ([]float64, float64) {
	settings := DefaultSimplexSettings()
	res := SimplexWithSettings(c, A, b, settings)
	return res.X, res.Objective
}

func SimplexWithSettings(c []float64, A [][]float64, b []float64, settings SimplexSettings) SimplexResult {
	m, n := len(A), len(c)
	if m == 0 || n == 0 {
		return SimplexResult{X: nil, Objective: 0, Iterations: 0, Status: "empty"}
	}
	tableau := buildTableau(c, A, b)
	rows := m + 1
	cols := n + m + 1
	basis := make([]int, m)
	for i := 0; i < m; i++ {
		basis[i] = n + i
	}
	iter := 0
	for iter < settings.MaxIter {
		enter := chooseEntering(tableau[rows-1], settings.Tol, settings.Bland)
		if enter < 0 {
			return SimplexResult{X: extractSolution(tableau, basis, n), Objective: tableau[rows-1][cols-1], Iterations: iter, Status: "optimal"}
		}
		leave := chooseLeaving(tableau, enter, settings.Tol)
		if leave < 0 {
			return SimplexResult{X: nil, Objective: math.Inf(1), Iterations: iter, Status: "unbounded"}
		}
		pivot(tableau, leave, enter)
		basis[leave] = enter
		iter++
	}
	return SimplexResult{X: extractSolution(tableau, basis, n), Objective: tableau[rows-1][cols-1], Iterations: iter, Status: "max_iter"}
}

func TwoPhaseSimplex(c []float64, A [][]float64, b []float64, settings SimplexSettings) SimplexResult {
	m, n := len(A), len(c)
	if m == 0 || n == 0 {
		return SimplexResult{X: nil, Objective: 0, Iterations: 0, Status: "empty"}
	}
	A2 := make([][]float64, m)
	b2 := make([]float64, m)
	for i := 0; i < m; i++ {
		b2[i] = b[i]
		A2[i] = make([]float64, n)
		copy(A2[i], A[i])
		if b2[i] < 0 {
			b2[i] = -b2[i]
			for j := 0; j < n; j++ {
				A2[i][j] = -A2[i][j]
			}
		}
	}
	artificial := make([]float64, m)
	for i := 0; i < m; i++ {
		artificial[i] = 1
	}
	cPhase1 := make([]float64, n+m)
	for i := 0; i < m; i++ {
		cPhase1[n+i] = 1
	}
	tab := buildTableauExtended(cPhase1, A2, b2)
	basis := make([]int, m)
	for i := 0; i < m; i++ {
		basis[i] = n + i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n+m+1; j++ {
			tab[m][j] += tab[i][j]
		}
	}
	res1 := simplexIterate(tab, basis, settings)
	if res1.Status != "optimal" || absO(res1.Objective) > settings.Tol {
		res1.Status = "infeasible"
		return res1
	}
	newC := make([]float64, n+m)
	copy(newC, c)
	tab2 := buildTableauExtended(newC, A2, b2)
	for i := 0; i < m; i++ {
		row := basis[i]
		if row < n {
			cost := newC[row]
			for j := 0; j < n+m+1; j++ {
				tab2[m][j] -= cost * tab2[i][j]
			}
		}
	}
	res2 := simplexIterate(tab2, basis, settings)
	res2.X = extractSolution(tab2, basis, n)
	return res2
}

func simplexIterate(tableau [][]float64, basis []int, settings SimplexSettings) SimplexResult {
	rows := len(tableau)
	cols := len(tableau[0])
	iter := 0
	for iter < settings.MaxIter {
		enter := chooseEntering(tableau[rows-1], settings.Tol, settings.Bland)
		if enter < 0 {
			return SimplexResult{X: nil, Objective: tableau[rows-1][cols-1], Iterations: iter, Status: "optimal"}
		}
		leave := chooseLeaving(tableau, enter, settings.Tol)
		if leave < 0 {
			return SimplexResult{X: nil, Objective: math.Inf(1), Iterations: iter, Status: "unbounded"}
		}
		pivot(tableau, leave, enter)
		basis[leave] = enter
		iter++
	}
	return SimplexResult{X: nil, Objective: tableau[rows-1][cols-1], Iterations: iter, Status: "max_iter"}
}

func buildTableau(c []float64, A [][]float64, b []float64) [][]float64 {
	m, n := len(A), len(c)
	cols := n + m + 1
	tableau := make([][]float64, m+1)
	for i := 0; i <= m; i++ {
		tableau[i] = make([]float64, cols)
	}
	for i := 0; i < m; i++ {
		copy(tableau[i][:n], A[i])
		tableau[i][n+i] = 1
		tableau[i][cols-1] = b[i]
	}
	for j := 0; j < n; j++ {
		tableau[m][j] = -c[j]
	}
	return tableau
}

func buildTableauExtended(c []float64, A [][]float64, b []float64) [][]float64 {
	m := len(A)
	cols := len(c) + 1
	tableau := make([][]float64, m+1)
	for i := 0; i <= m; i++ {
		tableau[i] = make([]float64, cols)
	}
	for i := 0; i < m; i++ {
		copy(tableau[i][:len(c)], A[i])
		tableau[i][cols-1] = b[i]
	}
	for j := 0; j < len(c); j++ {
		tableau[m][j] = -c[j]
	}
	return tableau
}

func chooseEntering(costRow []float64, tol float64, bland bool) int {
	cols := len(costRow) - 1
	if bland {
		for j := 0; j < cols; j++ {
			if costRow[j] < -tol {
				return j
			}
		}
		return -1
	}
	idx := -1
	minVal := -tol
	for j := 0; j < cols; j++ {
		if costRow[j] < minVal {
			minVal = costRow[j]
			idx = j
		}
	}
	return idx
}

func chooseLeaving(tableau [][]float64, enter int, tol float64) int {
	rows := len(tableau) - 1
	cols := len(tableau[0])
	row := -1
	minRatio := math.Inf(1)
	for i := 0; i < rows; i++ {
		coeff := tableau[i][enter]
		if coeff > tol {
			ratio := tableau[i][cols-1] / coeff
			if ratio < minRatio {
				minRatio = ratio
				row = i
			}
		}
	}
	return row
}

func pivot(tableau [][]float64, row, col int) {
	cols := len(tableau[0])
	pivotVal := tableau[row][col]
	if pivotVal == 0 {
		return
	}
	inv := 1.0 / pivotVal
	for j := 0; j < cols; j++ {
		tableau[row][j] *= inv
	}
	for i := 0; i < len(tableau); i++ {
		if i == row {
			continue
		}
		factor := tableau[i][col]
		if factor == 0 {
			continue
		}
		for j := 0; j < cols; j++ {
			tableau[i][j] -= factor * tableau[row][j]
		}
	}
}

func extractSolution(tableau [][]float64, basis []int, n int) []float64 {
	x := make([]float64, n)
	cols := len(tableau[0])
	for i := 0; i < len(basis); i++ {
		if basis[i] < n {
			x[basis[i]] = tableau[i][cols-1]
		}
	}
	return x
}

type InteriorPointSettings struct {
	Iterations int
	Mu         float64
	Step       float64
	Tol        float64
}

func DefaultInteriorPointSettings() InteriorPointSettings {
	return InteriorPointSettings{
		Iterations: 200,
		Mu:         1.0,
		Step:       0.01,
		Tol:        1e-6,
	}
}

func InteriorPointSolve(c []float64, A [][]float64, b []float64, settings InteriorPointSettings) []float64 {
	n := len(c)
	m := len(A)
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = 1
	}
	for iter := 0; iter < settings.Iterations; iter++ {
		grad := make([]float64, n)
		for i := 0; i < n; i++ {
			grad[i] = c[i]
			if x[i] != 0 {
				grad[i] -= settings.Mu / x[i]
			}
		}
		for i := 0; i < m; i++ {
			ai := A[i]
			res := dotLP(ai, x) - b[i]
			for j := 0; j < n; j++ {
				grad[j] += 2 * res * ai[j]
			}
		}
		step := settings.Step
		for i := 0; i < n; i++ {
			x[i] -= step * grad[i]
			if x[i] <= 0 {
				x[i] = settings.Tol
			}
		}
		if vecNormLP(grad) < settings.Tol {
			break
		}
		settings.Mu *= 0.95
		if settings.Mu < settings.Tol {
			settings.Mu = settings.Tol
		}
	}
	return x
}

func dotLP(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func vecNormLP(v []float64) float64 {
	sum := 0.0
	for i := range v {
		sum += v[i] * v[i]
	}
	return math.Sqrt(sum)
}

func ObjectiveValueLP(c, x []float64) float64 {
	if len(c) != len(x) {
		return 0
	}
	sum := 0.0
	for i := range c {
		sum += c[i] * x[i]
	}
	return sum
}

func ConstraintResiduals(A [][]float64, b []float64, x []float64) []float64 {
	res := make([]float64, len(A))
	for i := range A {
		res[i] = dotLP(A[i], x) - b[i]
	}
	return res
}

func SlackValues(A [][]float64, b []float64, x []float64) []float64 {
	res := make([]float64, len(A))
	for i := range A {
		res[i] = b[i] - dotLP(A[i], x)
	}
	return res
}

func FeasibleLP(A [][]float64, b []float64, x []float64, tol float64) bool {
	if len(A) != len(b) {
		return false
	}
	for i := range A {
		if dotLP(A[i], x)-b[i] > tol {
			return false
		}
	}
	return true
}

func ScaleLP(A [][]float64, b []float64, c []float64) ([][]float64, []float64, []float64) {
	m := len(A)
	n := len(c)
	A2 := make([][]float64, m)
	b2 := make([]float64, m)
	c2 := make([]float64, n)
	colScale := make([]float64, n)
	rowScale := make([]float64, m)
	for i := 0; i < n; i++ {
		colScale[i] = 1
	}
	for i := 0; i < m; i++ {
		maxv := 0.0
		for j := 0; j < n; j++ {
			v := absO(A[i][j])
			if v > maxv {
				maxv = v
			}
		}
		if maxv == 0 {
			rowScale[i] = 1
		} else {
			rowScale[i] = 1 / maxv
		}
	}
	for j := 0; j < n; j++ {
		maxv := 0.0
		for i := 0; i < m; i++ {
			v := absO(A[i][j])
			if v > maxv {
				maxv = v
			}
		}
		if maxv == 0 {
			colScale[j] = 1
		} else {
			colScale[j] = 1 / maxv
		}
	}
	for i := 0; i < m; i++ {
		A2[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			A2[i][j] = A[i][j] * rowScale[i] * colScale[j]
		}
		b2[i] = b[i] * rowScale[i]
	}
	for j := 0; j < n; j++ {
		c2[j] = c[j] * colScale[j]
	}
	return A2, b2, c2
}
