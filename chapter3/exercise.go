package chapter3

import "fmt"

func Exercise() {

	// var i int = 10
	// var p *int
	// p = &i
	// var j int
	// j = *p
	// fmt.Println(j)

	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2
	p2 = p1
	j = *p2 + i
	fmt.Println(j)
}
