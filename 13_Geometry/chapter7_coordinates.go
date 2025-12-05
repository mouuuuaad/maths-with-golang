package geometry

func atanG(x float64) float64 {
	if x > 1 {
		return Pi/2 - atanG(1/x)
	}
	if x < -1 {
		return -Pi/2 - atanG(1/x)
	}
	sum := x
	term := x
	x2 := x * x
	for i := 1; i < 50; i++ {
		term *= -x2
		sum += term / float64(2*i+1)
	}
	return sum
}

func atan2G(y, x float64) float64 {
	if x > 0 {
		return atanG(y / x)
	}
	if x < 0 && y >= 0 {
		return atanG(y/x) + Pi
	}
	if x < 0 && y < 0 {
		return atanG(y/x) - Pi
	}
	if x == 0 && y > 0 {
		return Pi / 2
	}
	if x == 0 && y < 0 {
		return -Pi / 2
	}
	return 0
}

func CartesianToPolar(p Point2D) (r, theta float64) {
	r = sqrtG(p.X*p.X + p.Y*p.Y)
	theta = atan2G(p.Y, p.X)
	return
}

func PolarToCartesian(r, theta float64) Point2D {
	return Point2D{r * cosG(theta), r * sinG(theta)}
}

func CartesianToSpherical(p Point3D) (r, theta, phi float64) {
	r = sqrtG(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
	if r == 0 {
		return 0, 0, 0
	}
	theta = atan2G(p.Y, p.X)
	phi = acos(p.Z / r)
	return
}

func acos(x float64) float64 {
	return Pi/2 - asin(x)
}

func asin(x float64) float64 {
	if x < -1 {
		x = -1
	}
	if x > 1 {
		x = 1
	}
	return atanG(x / sqrtG(1-x*x))
}

func SphericalToCartesian(r, theta, phi float64) Point3D {
	return Point3D{
		r * sinG(phi) * cosG(theta),
		r * sinG(phi) * sinG(theta),
		r * cosG(phi),
	}
}

func CartesianToCylindrical(p Point3D) (r, theta, z float64) {
	r = sqrtG(p.X*p.X + p.Y*p.Y)
	theta = atan2G(p.Y, p.X)
	z = p.Z
	return
}

func CylindricalToCartesian(r, theta, z float64) Point3D {
	return Point3D{r * cosG(theta), r * sinG(theta), z}
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
