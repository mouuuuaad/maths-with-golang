package geometry

type Vector3D struct {
	X, Y, Z float64
}

func (v Vector3D) Add(u Vector3D) Vector3D {
	return Vector3D{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vector3D) Sub(u Vector3D) Vector3D {
	return Vector3D{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vector3D) Scale(s float64) Vector3D {
	return Vector3D{v.X * s, v.Y * s, v.Z * s}
}

func (v Vector3D) Dot(u Vector3D) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vector3D) Cross(u Vector3D) Vector3D {
	return Vector3D{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

func (v Vector3D) Magnitude() float64 {
	return sqrtG(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3D) Normalize() Vector3D {
	mag := v.Magnitude()
	if mag == 0 {
		return v
	}
	return v.Scale(1 / mag)
}

type Plane struct {
	A, B, C, D float64
}

func PlaneFromPoints(p1, p2, p3 Point3D) Plane {
	v1 := Vector3D{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}
	v2 := Vector3D{p3.X - p1.X, p3.Y - p1.Y, p3.Z - p1.Z}
	normal := v1.Cross(v2)
	d := -(normal.X*p1.X + normal.Y*p1.Y + normal.Z*p1.Z)
	return Plane{normal.X, normal.Y, normal.Z, d}
}

func (pl Plane) PointDistance(p Point3D) float64 {
	num := absG(pl.A*p.X + pl.B*p.Y + pl.C*p.Z + pl.D)
	denom := sqrtG(pl.A*pl.A + pl.B*pl.B + pl.C*pl.C)
	return num / denom
}

type Line3D struct {
	Point     Point3D
	Direction Vector3D
}

func LineFromPoints3D(p1, p2 Point3D) Line3D {
	dir := Vector3D{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}
	return Line3D{p1, dir}
}

func (l Line3D) PointAt(t float64) Point3D {
	return Point3D{
		l.Point.X + t*l.Direction.X,
		l.Point.Y + t*l.Direction.Y,
		l.Point.Z + t*l.Direction.Z,
	}
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
