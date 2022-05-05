package chapter2

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func Panic_recover() {
	save()
	fmt.Println("OK?")
}
