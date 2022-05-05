package chapter5

import (
	"fmt"
	"time"
)

func gorutine4(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func gorutine5(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)
	go gorutine4(c1)
	go gorutine5(c2)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func Select_default() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!!!")
			break OuterLoop
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("==============")
}
