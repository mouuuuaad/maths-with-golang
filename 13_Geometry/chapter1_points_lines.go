package geometry

func absG(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func sqrtG(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}

type Point2D struct {
	X, Y float64
}

type Point3D struct {
	X, Y, Z float64
}

func Distance2D(p1, p2 Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return sqrtG(dx*dx + dy*dy)
}

func Distance3D(p1, p2 Point3D) float64 {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	return sqrtG(dx*dx + dy*dy + dz*dz)
}

func Midpoint2D(p1, p2 Point2D) Point2D {
	return Point2D{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2}
}

func Midpoint3D(p1, p2 Point3D) Point3D {
	return Point3D{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2, (p1.Z + p2.Z) / 2}
}

func Slope(p1, p2 Point2D) float64 {
	if p2.X == p1.X {
		return 1e308
	}
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

func LineEquation(p1, p2 Point2D) (a, b, c float64) {
	a = p2.Y - p1.Y
	b = p1.X - p2.X
	c = a*p1.X + b*p1.Y
	return
}

func PointToLineDistance(p Point2D, a, b, c float64) float64 {
	return absG(a*p.X+b*p.Y-c) / sqrtG(a*a+b*b)
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
