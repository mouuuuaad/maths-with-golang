# MathsWithGolang

A pure Go mathematics library implementing a wide range of algorithms across 15 modules. The code is structured as chapter-style packages and focuses on clarity, correctness, and breadth. No external dependencies.

## Highlights

- Pure Go implementations only
- 15 modules covering core and applied mathematics
- Chapter-structured packages for easy navigation
- Clean, consistent APIs

## Modules

1. 01_Calculus
2. 02_LinearAlgebra
3. 03_Limits
4. 04_Derivatives
5. 05_Functions
6. 06_NumericalSequences
7. 07_AlgebraicStructures
8. 08_Arithmetic
9. 09_ComplexNumbers
10. 10_Probability
11. 11_DifferentialEquations
12. 12_DiscreteMath
13. 13_Geometry
14. 14_GraphTheory
15. 15_Optimization

## Requirements

- Go 1.24.10 or newer

## Build

```bash
go build ./...
```

## Run

```bash
go run main.go
```

## Tests

```bash
go test ./...
```

## Layout

Each module is a Go package in its own folder, with chapter files that group related topics. The root `main.go` demonstrates usage across modules.

## License

MIT License. See `LICENSE`.
