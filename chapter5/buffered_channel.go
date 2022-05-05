package chapter5

import "fmt"

func Bufferd_channel() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))

	ch <- 200
	fmt.Println(len(ch))

	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func Bufferd_channel2() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine3(s, c)

	for i := range c {
		fmt.Println(i)
	}
}
