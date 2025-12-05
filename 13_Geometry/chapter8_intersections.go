package geometry

func LineSegmentsIntersect(p1, p2, p3, p4 Point2D) bool {
	d1 := direction(p3, p4, p1)
	d2 := direction(p3, p4, p2)
	d3 := direction(p1, p2, p3)
	d4 := direction(p1, p2, p4)
	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) && ((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}
	if d1 == 0 && onSegment(p3, p4, p1) {
		return true
	}
	if d2 == 0 && onSegment(p3, p4, p2) {
		return true
	}
	if d3 == 0 && onSegment(p1, p2, p3) {
		return true
	}
	if d4 == 0 && onSegment(p1, p2, p4) {
		return true
	}
	return false
}

func direction(pi, pj, pk Point2D) float64 {
	return (pk.X-pi.X)*(pj.Y-pi.Y) - (pj.X-pi.X)*(pk.Y-pi.Y)
}

func onSegment(pi, pj, pk Point2D) bool {
	minX, maxX := pi.X, pj.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := pi.Y, pj.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	return pk.X >= minX && pk.X <= maxX && pk.Y >= minY && pk.Y <= maxY
}

func LineIntersection(p1, p2, p3, p4 Point2D) (Point2D, bool) {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	x3, y3, x4, y4 := p3.X, p3.Y, p4.X, p4.Y
	denom := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if absG(denom) < 1e-10 {
		return Point2D{}, false
	}
	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / denom
	return Point2D{x1 + t*(x2-x1), y1 + t*(y2-y1)}, true
}

func CirclesIntersect(c1, c2 Circle) bool {
	d := Distance2D(c1.Center, c2.Center)
	return d <= c1.Radius+c2.Radius && d >= absG(c1.Radius-c2.Radius)
}

func LineCircleIntersection(p1, p2 Point2D, c Circle) []Point2D {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	fx, fy := p1.X-c.Center.X, p1.Y-c.Center.Y
	a := dx*dx + dy*dy
	b := 2 * (fx*dx + fy*dy)
	cc := fx*fx + fy*fy - c.Radius*c.Radius
	disc := b*b - 4*a*cc
	if disc < 0 {
		return nil
	}
	sqrtDisc := sqrtG(disc)
	t1 := (-b - sqrtDisc) / (2 * a)
	t2 := (-b + sqrtDisc) / (2 * a)
	result := []Point2D{}
	if t1 >= 0 && t1 <= 1 {
		result = append(result, Point2D{p1.X + t1*dx, p1.Y + t1*dy})
	}
	if t2 >= 0 && t2 <= 1 && absG(t2-t1) > 1e-10 {
		result = append(result, Point2D{p1.X + t2*dx, p1.Y + t2*dy})
	}
	return result
}

func PointInPolygon(p Point2D, poly Polygon) bool {
	n := len(poly.Vertices)
	inside := false
	j := n - 1
	for i := 0; i < n; i++ {
		if ((poly.Vertices[i].Y > p.Y) != (poly.Vertices[j].Y > p.Y)) &&
			(p.X < (poly.Vertices[j].X-poly.Vertices[i].X)*(p.Y-poly.Vertices[i].Y)/(poly.Vertices[j].Y-poly.Vertices[i].Y)+poly.Vertices[i].X) {
			inside = !inside
		}
		j = i
	}
	return inside
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
