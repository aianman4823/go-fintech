package chapter1

import "fmt"

func Exercise() {
	f := 1.11
	fmt.Println(int(f))

	fmt.Println(5, 6)

	// m := make(map[string]int)
	// m["Mike"] = 20
	// m["Nancy"] = 24
	// m["Messi"] = 30
	// fmt.Printf("%T %v", m, m)
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v\n", m, m)

}
