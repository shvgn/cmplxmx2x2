package main

import (
	"fmt"
)

// Mat is a type of 2×2 matrix of complex numbers
//  0:m11   1:m12
//  2:m21   3:m22
type Mat [4]complex128

func (m *Mat) String() string {
	s := fmt.Sprintf("[ %v\t%v  \n  %v\t%v ]", m[0], m[1], m[2], m[3])
	return s
}

// Elem returns matrix element chosen with (row, col) notation
func (m *Mat) Elem(r, c int) complex128 {
	return m[r*2+c]
}

// Mul multiplies two Mat's
func Mul(m1, m2 Mat) Mat {
	var m Mat
	m[0] = m1[0]*m2[0] + m1[1]*m2[2]
	m[1] = m1[0]*m2[1] + m1[1]*m2[3]
	m[2] = m1[2]*m2[0] + m1[3]*m2[2]
	m[3] = m1[2]*m2[1] + m1[3]*m2[3]
	return m
}

func main() {

	// This all should be in tests
	m1 := Mat{1i, 2, 3i, 4}
	m2 := Mat{3, 1, 4, 2}

	fmt.Println("M1 =\n" + m1.String() + "\n")
	fmt.Println("M2 =\n" + m2.String() + "\n")
	m := Mul(m1, m2)
	fmt.Println("M = M1 × M2 =\n" + m.String() + "\n")

	fmt.Printf("M(%d, %d) = %v\n", 0, 0, m.Elem(0, 0))
	fmt.Printf("M(%d, %d) = %v\n", 0, 1, m.Elem(0, 1))
	fmt.Printf("M(%d, %d) = %v\n", 1, 0, m.Elem(1, 0))
	fmt.Printf("M(%d, %d) = %v\n", 1, 1, m.Elem(1, 1))

}
