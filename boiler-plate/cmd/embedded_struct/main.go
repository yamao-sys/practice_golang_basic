package main // mailパッケージを使用する

import "fmt"

type T struct {
	User User
	// User // 型を書くだけでもOK(field名省略可)
}

type User struct {
	Name string
	Age int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	// fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t.User)
}
