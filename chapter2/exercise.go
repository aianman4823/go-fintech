package chapter2

import (
	"fmt"
	"math"
)

func Exercise() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	ans := int(math.Pow(10, 10))
	for _, v := range l {
		if ans > v {
			ans = v
		}
	}
	fmt.Println(ans)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum_value := 0
	for _, v := range m {
		sum_value += v
	}
	fmt.Println(sum_value)
}
