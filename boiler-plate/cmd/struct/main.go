package main // mailパッケージを使用する

import "fmt"

type User struct {
	Name string
	Age int
	// X, Y int // 2つ以上のfieldをまとめて定義することもできる
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// user5 := User{30, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User) // newはポインタ型を返す
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	// structも値渡し. → UpdateUser(user1)を呼び出してもuser1は変わらず
	// 参照渡しにしたければuser8のようにポインタ型を渡してあげる.
	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}
