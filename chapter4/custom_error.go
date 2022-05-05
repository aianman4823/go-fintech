package chapter4

import "fmt"

type UserNotFound struct {
	Username string
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	} else {
		return &UserNotFound{Username: "mike"}
	}
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v",e.Username)
}

func CustomError() {
	err := myFunc()
	if err != nil {
		fmt.Println(err)
	}
}
