package chapter4

import "fmt"

func (v Vertex) Plus() int {
	return v.x + v.y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.x, v.y)
}

func Exercise() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())

	fmt.Println(v)
}
