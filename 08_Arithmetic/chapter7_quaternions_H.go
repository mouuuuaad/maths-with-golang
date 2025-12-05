package arithmetic

type Quaternion struct {
	W float64
	X float64
	Y float64
	Z float64
}

func NewQuaternion(w, x, y, z float64) Quaternion {
	return Quaternion{W: w, X: x, Y: y, Z: z}
}

func (q Quaternion) Add(other Quaternion) Quaternion {
	return Quaternion{
		W: q.W + other.W,
		X: q.X + other.X,
		Y: q.Y + other.Y,
		Z: q.Z + other.Z,
	}
}

func (q Quaternion) Subtract(other Quaternion) Quaternion {
	return Quaternion{
		W: q.W - other.W,
		X: q.X - other.X,
		Y: q.Y - other.Y,
		Z: q.Z - other.Z,
	}
}

func (q Quaternion) Multiply(other Quaternion) Quaternion {
	return Quaternion{
		W: q.W*other.W - q.X*other.X - q.Y*other.Y - q.Z*other.Z,
		X: q.W*other.X + q.X*other.W + q.Y*other.Z - q.Z*other.Y,
		Y: q.W*other.Y - q.X*other.Z + q.Y*other.W + q.Z*other.X,
		Z: q.W*other.Z + q.X*other.Y - q.Y*other.X + q.Z*other.W,
	}
}

func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{W: q.W, X: -q.X, Y: -q.Y, Z: -q.Z}
}

func (q Quaternion) NormSquared() float64 {
	return q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z
}

func (q Quaternion) Norm() float64 {
	return sqrtF(q.NormSquared())
}

func (q Quaternion) Inverse() Quaternion {
	n2 := q.NormSquared()
	if n2 == 0 {
		return Quaternion{}
	}
	conj := q.Conjugate()
	return Quaternion{
		W: conj.W / n2,
		X: conj.X / n2,
		Y: conj.Y / n2,
		Z: conj.Z / n2,
	}
}

func (q Quaternion) Normalize() Quaternion {
	n := q.Norm()
	if n == 0 {
		return Quaternion{}
	}
	return Quaternion{q.W / n, q.X / n, q.Y / n, q.Z / n}
}

func (q Quaternion) Scale(s float64) Quaternion {
	return Quaternion{q.W * s, q.X * s, q.Y * s, q.Z * s}
}

func (q Quaternion) Dot(other Quaternion) float64 {
	return q.W*other.W + q.X*other.X + q.Y*other.Y + q.Z*other.Z
}

func QuaternionFromAxisAngle(x, y, z, angle float64) Quaternion {
	halfAngle := angle / 2
	s := sinF(halfAngle)
	return Quaternion{
		W: cosF(halfAngle),
		X: x * s,
		Y: y * s,
		Z: z * s,
	}
}

func (q Quaternion) ToAxisAngle() (float64, float64, float64, float64) {
	n := sqrtF(q.X*q.X + q.Y*q.Y + q.Z*q.Z)
	if n < 1e-10 {
		return 0, 0, 1, 0
	}
	return q.X / n, q.Y / n, q.Z / n, 2 * atanF(n/q.W)
}

func (q Quaternion) RotateVector(vx, vy, vz float64) (float64, float64, float64) {
	v := Quaternion{0, vx, vy, vz}
	result := q.Multiply(v).Multiply(q.Inverse())
	return result.X, result.Y, result.Z
}

func Slerp(q1, q2 Quaternion, t float64) Quaternion {
	dot := q1.Dot(q2)
	if dot < 0 {
		q2 = q2.Scale(-1)
		dot = -dot
	}
	if dot > 0.9995 {
		result := Quaternion{
			W: q1.W + t*(q2.W-q1.W),
			X: q1.X + t*(q2.X-q1.X),
			Y: q1.Y + t*(q2.Y-q1.Y),
			Z: q1.Z + t*(q2.Z-q1.Z),
		}
		return result.Normalize()
	}
	theta0 := acosF(dot)
	theta := theta0 * t
	s0 := cosF(theta) - dot*sinF(theta)/sinF(theta0)
	s1 := sinF(theta) / sinF(theta0)
	return Quaternion{
		W: s0*q1.W + s1*q2.W,
		X: s0*q1.X + s1*q2.X,
		Y: s0*q1.Y + s1*q2.Y,
		Z: s0*q1.Z + s1*q2.Z,
	}
}

func sinF(x float64) float64 {
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

func cosF(x float64) float64 {
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

func acosF(x float64) float64 {
	pi := 3.14159265358979323846
	return pi/2 - asinF(x)
}

func asinF(x float64) float64 {
	if x < -1 {
		x = -1
	}
	if x > 1 {
		x = 1
	}
	if absF(x) == 1 {
		pi := 3.14159265358979323846
		return x * pi / 2
	}
	return atanF(x / sqrtF(1-x*x))
}
