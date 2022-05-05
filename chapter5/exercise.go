package chapter5

import "fmt"

func goroutine_ex(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func Exercise() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine_ex(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
