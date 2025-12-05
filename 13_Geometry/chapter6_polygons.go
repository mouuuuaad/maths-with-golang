package geometry

type Polygon struct {
	Vertices []Point2D
}

func (p Polygon) Area() float64 {
	n := len(p.Vertices)
	if n < 3 {
		return 0
	}
	area := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += p.Vertices[i].X * p.Vertices[j].Y
		area -= p.Vertices[j].X * p.Vertices[i].Y
	}
	if area < 0 {
		area = -area
	}
	return area / 2
}

func (p Polygon) Perimeter() float64 {
	n := len(p.Vertices)
	perimeter := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		perimeter += Distance2D(p.Vertices[i], p.Vertices[j])
	}
	return perimeter
}

func (p Polygon) Centroid() Point2D {
	n := len(p.Vertices)
	cx, cy := 0.0, 0.0
	for _, v := range p.Vertices {
		cx += v.X
		cy += v.Y
	}
	return Point2D{cx / float64(n), cy / float64(n)}
}

func (p Polygon) IsConvex() bool {
	n := len(p.Vertices)
	if n < 3 {
		return false
	}
	sign := 0
	for i := 0; i < n; i++ {
		dx1 := p.Vertices[(i+1)%n].X - p.Vertices[i].X
		dy1 := p.Vertices[(i+1)%n].Y - p.Vertices[i].Y
		dx2 := p.Vertices[(i+2)%n].X - p.Vertices[(i+1)%n].X
		dy2 := p.Vertices[(i+2)%n].Y - p.Vertices[(i+1)%n].Y
		cross := dx1*dy2 - dy1*dx2
		if cross != 0 {
			if sign == 0 {
				if cross > 0 {
					sign = 1
				} else {
					sign = -1
				}
			} else if (sign == 1 && cross < 0) || (sign == -1 && cross > 0) {
				return false
			}
		}
	}
	return true
}

func ConvexHull(points []Point2D) []Point2D {
	n := len(points)
	if n < 3 {
		return points
	}
	sorted := make([]Point2D, n)
	copy(sorted, points)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j].X > sorted[j+1].X || (sorted[j].X == sorted[j+1].X && sorted[j].Y > sorted[j+1].Y) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	cross := func(o, a, b Point2D) float64 {
		return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
	}
	lower := []Point2D{}
	for _, p := range sorted {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}
	upper := []Point2D{}
	for i := n - 1; i >= 0; i-- {
		p := sorted[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}
	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
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
