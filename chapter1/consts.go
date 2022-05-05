package chapter1

import "fmt"

const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

const big = 9223372036854775807 + 1

func Consts() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
