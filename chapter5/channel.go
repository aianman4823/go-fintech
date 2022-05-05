package chapter5

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Channel() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // キューのようなイメージ

	go goroutine1(s, c)
	go goroutine2(s, c)
	// チャネルが次に進むのをブロックしている
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
