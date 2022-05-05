package chapter5

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	defer c.mux.Unlock()
	c.mux.Lock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	defer c.mux.Unlock()
	c.mux.Lock()
	return c.v[key]
}

func Mutex_pra() {
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("key"))
}
