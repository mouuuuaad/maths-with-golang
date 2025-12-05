package geometry

func sinG(x float64) float64 {
	for x > Pi {
		x -= 2 * Pi
	}
	for x < -Pi {
		x += 2 * Pi
	}
	sum, term := x, x
	x2 := x * x
	for i := 1; i < 20; i++ {
		term *= -x2 / float64(2*i*(2*i+1))
		sum += term
	}
	return sum
}

func cosG(x float64) float64 {
	for x > Pi {
		x -= 2 * Pi
	}
	for x < -Pi {
		x += 2 * Pi
	}
	sum, term := 1.0, 1.0
	x2 := x * x
	for i := 1; i < 20; i++ {
		term *= -x2 / float64((2*i-1)*(2*i))
		sum += term
	}
	return sum
}

func RotatePoint2D(p Point2D, angle float64) Point2D {
	c, s := cosG(angle), sinG(angle)
	return Point2D{p.X*c - p.Y*s, p.X*s + p.Y*c}
}

func RotateAroundPoint(p, center Point2D, angle float64) Point2D {
	translated := Point2D{p.X - center.X, p.Y - center.Y}
	rotated := RotatePoint2D(translated, angle)
	return Point2D{rotated.X + center.X, rotated.Y + center.Y}
}

func ScalePoint(p Point2D, factor float64) Point2D {
	return Point2D{p.X * factor, p.Y * factor}
}

func TranslatePoint(p Point2D, dx, dy float64) Point2D {
	return Point2D{p.X + dx, p.Y + dy}
}

func ReflectOverX(p Point2D) Point2D {
	return Point2D{p.X, -p.Y}
}

func ReflectOverY(p Point2D) Point2D {
	return Point2D{-p.X, p.Y}
}

func ReflectOverLine(p Point2D, a, b, c float64) Point2D {
	d := a*a + b*b
	x := (b*b*p.X - a*a*p.X - 2*a*b*p.Y + 2*a*c) / d
	y := (a*a*p.Y - b*b*p.Y - 2*a*b*p.X + 2*b*c) / d
	return Point2D{p.X - x, p.Y - y}
}

type Matrix3x3 [3][3]float64

func (m Matrix3x3) MultiplyPoint(p Point2D) Point2D {
	x := m[0][0]*p.X + m[0][1]*p.Y + m[0][2]
	y := m[1][0]*p.X + m[1][1]*p.Y + m[1][2]
	return Point2D{x, y}
}

func TransformationMatrix(angle, sx, sy, tx, ty float64) Matrix3x3 {
	c, s := cosG(angle), sinG(angle)
	return Matrix3x3{
		{sx * c, -sy * s, tx},
		{sx * s, sy * c, ty},
		{0, 0, 1},
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
