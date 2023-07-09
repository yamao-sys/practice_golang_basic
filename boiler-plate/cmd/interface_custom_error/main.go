package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// fmt.Println(err.Message)
	e, ok := err.(*MyError) // 型アサーション
	if ok {
		fmt.Println(e.ErrCode)
	}
}
