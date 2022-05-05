package chapter1

import (
	"fmt"
	"os/user"
	"time"
)

func Import_func() {
	fmt.Println("Hello World!", time.Now())
	fmt.Println(user.Current())
}
