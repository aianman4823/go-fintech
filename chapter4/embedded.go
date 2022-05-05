package chapter4

import "fmt"

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func Embedded() {
	v := New3D(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}
