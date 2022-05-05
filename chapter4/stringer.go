package chapter4

import "fmt"

type Person2 struct {
	Person
	Age int
}

func (p Person2) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func Stringer() {
	mike := Person2{Person{"Mike"}, 22}
	fmt.Println(mike)
}
