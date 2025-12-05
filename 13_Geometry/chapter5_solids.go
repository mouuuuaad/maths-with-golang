package geometry

type Sphere struct {
	Center Point3D
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * Pi * s.Radius * s.Radius * s.Radius
}

func (s Sphere) SurfaceArea() float64 {
	return 4 * Pi * s.Radius * s.Radius
}

func (s Sphere) ContainsPoint(p Point3D) bool {
	return Distance3D(s.Center, p) <= s.Radius
}

type Cylinder struct {
	Base   Point3D
	Radius float64
	Height float64
}

func (c Cylinder) Volume() float64 {
	return Pi * c.Radius * c.Radius * c.Height
}

func (c Cylinder) SurfaceArea() float64 {
	return 2*Pi*c.Radius*c.Height + 2*Pi*c.Radius*c.Radius
}

type Cone struct {
	Apex   Point3D
	Radius float64
	Height float64
}

func (c Cone) Volume() float64 {
	return Pi * c.Radius * c.Radius * c.Height / 3
}

func (c Cone) SurfaceArea() float64 {
	slant := sqrtG(c.Radius*c.Radius + c.Height*c.Height)
	return Pi * c.Radius * (c.Radius + slant)
}

type Box struct {
	Corner               Point3D
	Width, Height, Depth float64
}

func (b Box) Volume() float64 {
	return b.Width * b.Height * b.Depth
}

func (b Box) SurfaceArea() float64 {
	return 2 * (b.Width*b.Height + b.Height*b.Depth + b.Depth*b.Width)
}

func (b Box) Diagonal() float64 {
	return sqrtG(b.Width*b.Width + b.Height*b.Height + b.Depth*b.Depth)
}

type Torus struct {
	Center      Point3D
	MajorRadius float64
	MinorRadius float64
}

func (t Torus) Volume() float64 {
	return 2 * Pi * Pi * t.MajorRadius * t.MinorRadius * t.MinorRadius
}

func (t Torus) SurfaceArea() float64 {
	return 4 * Pi * Pi * t.MajorRadius * t.MinorRadius
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
