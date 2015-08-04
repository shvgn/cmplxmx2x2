package cmpxmx2x2

import (
	"fmt"
)

// Matrix of complex numbers
//  0:m11   1:m12
//  2:m21   3:m22
type Mat [4]complex128

func (m *Mat) String() string {
	s := fmt.Sprintf("[ %v\t%v  \n  %v\t%v ]", m[0], m[1], m[2], m[3])
	return s
}

func Mul(m1, m2 Mat) Mat {
	var m Mat
	m[0] = m1[0]*m2[0] + m1[1]*m2[2]
	m[1] = m1[0]*m2[1] + m1[1]*m2[3]
	m[2] = m1[2]*m2[0] + m1[3]*m2[2]
	m[3] = m1[2]*m2[1] + m1[3]*m2[3]
	return m
}

// func main() {
// 	m1 := Mat{1, 2, 3, 4}
// 	m2 := Mat{3, 1, 4, 2}

// 	fmt.Println(m1.String())
// 	fmt.Println(m2.String())
// 	m := Mul(m1, m2)
// 	fmt.Println(m.String())

// }
