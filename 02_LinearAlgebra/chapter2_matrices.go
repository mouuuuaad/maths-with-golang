package linearalgebra

type Matrix [][]float64

func NewMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]float64, cols)
	}
	return m
}

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix) Add(other Matrix) Matrix {
	if m.Rows() != other.Rows() || m.Cols() != other.Cols() {
		return nil
	}
	result := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[i][j] = m[i][j] + other[i][j]
		}
	}
	return result
}

func (m Matrix) Subtract(other Matrix) Matrix {
	if m.Rows() != other.Rows() || m.Cols() != other.Cols() {
		return nil
	}
	result := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[i][j] = m[i][j] - other[i][j]
		}
	}
	return result
}

func (m Matrix) Scale(s float64) Matrix {
	result := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[i][j] = m[i][j] * s
		}
	}
	return result
}

func (m Matrix) Multiply(other Matrix) Matrix {
	if m.Cols() != other.Rows() {
		return nil
	}
	result := NewMatrix(m.Rows(), other.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < other.Cols(); j++ {
			sum := 0.0
			for k := 0; k < m.Cols(); k++ {
				sum += m[i][k] * other[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func (m Matrix) Transpose() Matrix {
	result := NewMatrix(m.Cols(), m.Rows())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[j][i] = m[i][j]
		}
	}
	return result
}

func Identity(n int) Matrix {
	m := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}

func (m Matrix) MultiplyVector(v Vector) Vector {
	if m.Cols() != len(v) {
		return nil
	}
	result := make(Vector, m.Rows())
	for i := 0; i < m.Rows(); i++ {
		sum := 0.0
		for j := 0; j < m.Cols(); j++ {
			sum += m[i][j] * v[j]
		}
		result[i] = sum
	}
	return result
}

func (m Matrix) Trace() float64 {
	min := m.Rows()
	if m.Cols() < min {
		min = m.Cols()
	}
	sum := 0.0
	for i := 0; i < min; i++ {
		sum += m[i][i]
	}
	return sum
}

func (m Matrix) IsSquare() bool {
	return m.Rows() == m.Cols()
}

func (m Matrix) IsSymmetric() bool {
	if !m.IsSquare() {
		return false
	}
	for i := 0; i < m.Rows(); i++ {
		for j := i + 1; j < m.Cols(); j++ {
			if absV(m[i][j]-m[j][i]) > 1e-9 {
				return false
			}
		}
	}
	return true
}
