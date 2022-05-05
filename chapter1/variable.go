package chapter1

import "fmt"

// これは外でも宣言できる
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func Variable() {
	fmt.Println(i, f64, s, t, f)
	fmt.Printf("%T\n", f64)
	fmt.Printf("%T\n", i)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}
