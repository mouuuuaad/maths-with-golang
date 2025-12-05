# MathsWithGolang

```
MMMMMMMM               MMMMMMMM     OOOOOOOOO     UUUUUUUU     UUUUUUUU           AAA                              AAA               DDDDDDDDDDDDD        
M:::::::M             M:::::::M   OO:::::::::OO   U::::::U     U::::::U          A:::A                            A:::A              D::::::::::::DDD     
M::::::::M           M::::::::M OO:::::::::::::OO U::::::U     U::::::U         A:::::A                          A:::::A             D:::::::::::::::DD   
M:::::::::M         M:::::::::MO:::::::OOO:::::::OUU:::::U     U:::::UU        A:::::::A                        A:::::::A            DDD:::::DDDDD:::::D  
M::::::::::M       M::::::::::MO::::::O   O::::::O U:::::U     U:::::U        A:::::::::A                      A:::::::::A             D:::::D    D:::::D 
M:::::::::::M     M:::::::::::MO:::::O     O:::::O U:::::D     D:::::U       A:::::A:::::A                    A:::::A:::::A            D:::::D     D:::::D
M:::::::M::::M   M::::M:::::::MO:::::O     O:::::O U:::::D     D:::::U      A:::::A A:::::A                  A:::::A A:::::A           D:::::D     D:::::D
M::::::M M::::M M::::M M::::::MO:::::O     O:::::O U:::::D     D:::::U     A:::::A   A:::::A                A:::::A   A:::::A          D:::::D     D:::::D
M::::::M  M::::M::::M  M::::::MO:::::O     O:::::O U:::::D     D:::::U    A:::::A     A:::::A              A:::::A     A:::::A         D:::::D     D:::::D
M::::::M   M:::::::M   M::::::MO:::::O     O:::::O U:::::D     D:::::U   A:::::AAAAAAAAA:::::A            A:::::AAAAAAAAA:::::A        D:::::D     D:::::D
M::::::M    M:::::M    M::::::MO:::::O     O:::::O U:::::D     D:::::U  A:::::::::::::::::::::A          A:::::::::::::::::::::A       D:::::D     D:::::D
M::::::M     MMMMM     M::::::MO::::::O   O::::::O U::::::U   U::::::U A:::::AAAAAAAAAAAAA:::::A        A:::::AAAAAAAAAAAAA:::::A      D:::::D    D:::::D 
M::::::M               M::::::MO:::::::OOO:::::::O U:::::::UUU:::::::UA:::::A             A:::::A      A:::::A             A:::::A   DDD:::::DDDDD:::::D  
M::::::M               M::::::M OO:::::::::::::OO   UU:::::::::::::UUA:::::A               A:::::A    A:::::A               A:::::A  D:::::::::::::::DD   
M::::::M               M::::::M   OO:::::::::OO       UU:::::::::UU A:::::A                 A:::::A  A:::::A                 A:::::A D::::::::::::DDD     
MMMMMMMM               MMMMMMMM     OOOOOOOOO           UUUUUUUUU  AAAAAAA                   AAAAAAAAAAAAAA                   AAAAAAADDDDDDDDDDDDD        
```

**<< The universe runs on equations. We just translate them >>**

---

A comprehensive **Pure Golang** mathematical library implementing deep algorithms across 15 modules with 124 chapters. No external dependencies - just pure mathematical implementations.

## 🌟 Features

- **100% Pure Golang** - No external packages or dependencies
- **15 Mathematical Modules** - From calculus to optimization
- **124 Deep Chapters** - Each with comprehensive implementations
- **Production Ready** - All code compiles and builds successfully

## 📚 Modules Overview

| Module | Chapters | Topics |
|--------|----------|--------|
| **01_Calculus** | 8 | Limits, Continuity, Differentiation, Integration, Series, Multivariable |
| **02_LinearAlgebra** | 8 | Vectors, Matrices, Determinants, Eigenvalues, Decompositions |
| **03_Limits** | 8 | Sequences, Functions, L'Hôpital, Squeeze Theorem, Continuity |
| **04_Derivatives** | 8 | Rules, Higher-Order, Implicit, Partial, Extrema, Applications |
| **05_Functions** | 8 | Exponential, Logarithmic, Trigonometric, Hyperbolic, Special |
| **06_NumericalSequences** | 8 | Arithmetic, Geometric, Recursive, Convergence, Series |
| **07_AlgebraicStructures** | 8 | Groups, Rings, Fields, Polynomials, Abstract Algebra |
| **08_Arithmetic** | 8 | ℕ, ℤ, ℚ, 𝔻, ℝ, ℂ, ℍ Number Systems |
| **09_ComplexNumbers** | 8 | Complex Operations, Polar Forms, Applications |
| **10_Probability** | 8 | Distributions, Statistics, Regression, Hypothesis Testing, ANOVA |
| **11_DifferentialEquations** | 8 | ODE Solvers, PDEs, Stiff Systems, Lorenz, Van der Pol |
| **12_DiscreteMath** | 8 | Number Theory, Sets, Logic, Combinatorics, Primes, CRT |
| **13_Geometry** | 8 | 2D/3D Geometry, Transformations, Polygons, Convex Hull |
| **14_GraphTheory** | 8 | BFS, DFS, Dijkstra, A*, MST, Max Flow, Eulerian/Hamiltonian |
| **15_Optimization** | 8 | Gradient Descent, BFGS, Simplex, GA, PSO, Simulated Annealing |

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

```bash
git clone https://github.com/mouuuuaad/maths-with-golang.git
cd maths-with-golang
```

### Build

```bash
go build ./...
```

### Usage

```bash
go run main.go
```

### Run Tests

```bash
go test ./...
```

## 📖 Module Highlights

### 🔢 Probability & Statistics (`10_Probability`)
- Binomial, Poisson, Normal distributions
- Linear & Polynomial Regression
- Hypothesis Testing (Z-test, T-test, Chi-Square, ANOVA)
- Covariance & Correlation matrices

### 📐 Differential Equations (`11_DifferentialEquations`)
- Euler, Heun, Midpoint, Runge-Kutta methods
- Adaptive RK45 with error control
- Stiff equation solvers (BDF, Crank-Nicolson)
- PDEs: Heat, Wave, Laplace equations
- Chaos: Lorenz attractor, Van der Pol oscillator

### 🧮 Discrete Mathematics (`12_DiscreteMath`)
- GCD, LCM, Extended Euclidean Algorithm
- Chinese Remainder Theorem
- Sieve of Eratosthenes, Euler's Totient
- Boolean Algebra, Truth Tables
- Permutations, Combinations, Stirling numbers

### 📊 Graph Theory (`14_GraphTheory`)
- BFS, DFS, Topological Sort
- Dijkstra, Bellman-Ford, Floyd-Warshall, A*
- Kruskal & Prim MST algorithms
- Max Flow (Ford-Fulkerson)
- Graph Coloring, Eulerian/Hamiltonian paths

### ⚡ Optimization (`15_Optimization`)
- Golden Section, Newton's Method, Brent's Method
- Gradient Descent, Conjugate Gradient, BFGS
- Simplex Method (Linear Programming)
- Nelder-Mead, Lagrange Multipliers
- Genetic Algorithm, PSO, Differential Evolution
- Simulated Annealing, Tabu Search

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

**Created by: MOUAAD IDOUFKIR**

*<< The universe runs on equations. We just translate them >>*
