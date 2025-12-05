package linearalgebra

func RotationMatrix2D(theta float64) Matrix {
	c := cosV(theta)
	s := sinV(theta)
	return Matrix{{c, -s}, {s, c}}
}

func RotationMatrix3DX(theta float64) Matrix {
	c := cosV(theta)
	s := sinV(theta)
	return Matrix{{1, 0, 0}, {0, c, -s}, {0, s, c}}
}

func RotationMatrix3DY(theta float64) Matrix {
	c := cosV(theta)
	s := sinV(theta)
	return Matrix{{c, 0, s}, {0, 1, 0}, {-s, 0, c}}
}

func RotationMatrix3DZ(theta float64) Matrix {
	c := cosV(theta)
	s := sinV(theta)
	return Matrix{{c, -s, 0}, {s, c, 0}, {0, 0, 1}}
}

func ScalingMatrix(scales []float64) Matrix {
	n := len(scales)
	m := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		m[i][i] = scales[i]
	}
	return m
}

func TranslationMatrix2D(tx, ty float64) Matrix {
	return Matrix{{1, 0, tx}, {0, 1, ty}, {0, 0, 1}}
}

func TranslationMatrix3D(tx, ty, tz float64) Matrix {
	return Matrix{{1, 0, 0, tx}, {0, 1, 0, ty}, {0, 0, 1, tz}, {0, 0, 0, 1}}
}

func ShearMatrix2D(shx, shy float64) Matrix {
	return Matrix{{1, shx}, {shy, 1}}
}

func ReflectionMatrix2D(axis int) Matrix {
	if axis == 0 {
		return Matrix{{1, 0}, {0, -1}}
	}
	return Matrix{{-1, 0}, {0, 1}}
}

func OrthogonalProjection2D(axis int) Matrix {
	if axis == 0 {
		return Matrix{{1, 0}, {0, 0}}
	}
	return Matrix{{0, 0}, {0, 1}}
}

func ApplyTransformation(M Matrix, v Vector) Vector {
	return M.MultiplyVector(v)
}

func ComposeTransformations(transforms ...Matrix) Matrix {
	if len(transforms) == 0 {
		return nil
	}
	result := transforms[0]
	for i := 1; i < len(transforms); i++ {
		result = transforms[i].Multiply(result)
	}
	return result
}

func cosV(x float64) float64 {
	pi := 3.14159265358979323846
	k := int((x + pi) / (2 * pi))
	x -= float64(k) * 2 * pi
	sum := 0.0
	term := 1.0
	x2 := x * x
	for i := 1; i < 30; i++ {
		sum += term
		term *= -x2 / float64((2*i-1)*(2*i))
	}
	return sum
}

func sinV(x float64) float64 {
	pi := 3.14159265358979323846
	k := int((x + pi) / (2 * pi))
	x -= float64(k) * 2 * pi
	sum := 0.0
	term := x
	x2 := x * x
	for i := 1; i < 30; i++ {
		sum += term
		term *= -x2 / float64(2*i*(2*i+1))
	}
	return sum
}
