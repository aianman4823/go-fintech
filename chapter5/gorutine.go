package chapter5

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Gorutine() {
	// ゴルーチンの処理が終了しなくてもプログラムのコードが終了すれば終了する
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)

	normal("hello")
	wg.Wait()
}
