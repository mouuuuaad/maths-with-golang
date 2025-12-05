package linearalgebra

type Vector []float64

func (v Vector) Add(u Vector) Vector {
	if len(v) != len(u) {
		return nil
	}
	res := make(Vector, len(v))
	for i := range v {
		res[i] = v[i] + u[i]
	}
	return res
}

func (v Vector) Subtract(u Vector) Vector {
	if len(v) != len(u) {
		return nil
	}
	res := make(Vector, len(v))
	for i := range v {
		res[i] = v[i] - u[i]
	}
	return res
}

func (v Vector) Scale(s float64) Vector {
	res := make(Vector, len(v))
	for i := range v {
		res[i] = v[i] * s
	}
	return res
}

func (v Vector) Dot(u Vector) float64 {
	if len(v) != len(u) {
		return 0
	}
	sum := 0.0
	for i := range v {
		sum += v[i] * u[i]
	}
	return sum
}

func (v Vector) Norm() float64 {
	return sqrtV(v.Dot(v))
}

func (v Vector) Normalize() Vector {
	n := v.Norm()
	if n == 0 {
		return make(Vector, len(v))
	}
	return v.Scale(1.0 / n)
}

func (v Vector) Cross(u Vector) Vector {
	if len(v) != 3 || len(u) != 3 {
		return nil
	}
	return Vector{
		v[1]*u[2] - v[2]*u[1],
		v[2]*u[0] - v[0]*u[2],
		v[0]*u[1] - v[1]*u[0],
	}
}

func (v Vector) Angle(u Vector) float64 {
	dot := v.Dot(u)
	norms := v.Norm() * u.Norm()
	if norms == 0 {
		return 0
	}
	return acosV(dot / norms)
}

func (v Vector) ProjectOnto(u Vector) Vector {
	uNorm := u.Dot(u)
	if uNorm == 0 {
		return make(Vector, len(v))
	}
	scalar := v.Dot(u) / uNorm
	return u.Scale(scalar)
}

func (v Vector) IsOrthogonal(u Vector) bool {
	return absV(v.Dot(u)) < 1e-9
}

func sqrtV(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := x
	for i := 0; i < 50; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}

func absV(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func acosV(x float64) float64 {
	pi := 3.14159265358979323846
	if x >= 1 {
		return 0
	}
	if x <= -1 {
		return pi
	}
	return pi/2 - asinV(x)
}

func asinV(x float64) float64 {
	if x < -1 {
		x = -1
	}
	if x > 1 {
		x = 1
	}
	return atanV(x / sqrtV(1-x*x))
}

func atanV(x float64) float64 {
	pi := 3.14159265358979323846
	if x < 0 {
		return -atanV(-x)
	}
	if x > 1 {
		return pi/2 - atanV(1/x)
	}
	sum := 0.0
	term := x
	x2 := x * x
	sign := 1.0
	for i := 0; i < 50; i++ {
		sum += sign * term / float64(2*i+1)
		term *= x2
		sign = -sign
	}
	return sum
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
