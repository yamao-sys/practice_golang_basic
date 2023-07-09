package main // mailパッケージを使用する

import "fmt"

type User struct {
	Name string
	Age int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	// 値渡しのメソッドなのでuser1のNameは不変
	user1.SetName("A")
	user1.SayName()

	// 参照渡しのメソッドなのでuser1のName変更
	user1.SetName2("A")
	user1.SayName()

	// user2がポインタ型なので参照渡しのメソッドに渡すと値が更新される
	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()
}
