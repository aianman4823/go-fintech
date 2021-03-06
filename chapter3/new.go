package chapter3

import "fmt"

func New() {
	// var p *int = new(int)
	// fmt.Println(*p)
	// *p++
	// fmt.Println(*p)

	// var p2 *int
	// fmt.Println(p2)
	// *p2++
	// fmt.Println(p2)

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	// ポインタが返ってくるものに関してはnew()を利用する。
	// それ以外にはmake()を使う
	var p *int = new(int)
	fmt.Printf("%T\n", p)
}
