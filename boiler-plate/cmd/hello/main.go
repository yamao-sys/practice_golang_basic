package main // mailパッケージを使用する
// パッケージ宣言は一つのみ！

import "fmt" // formatパッケージ. printやprintlnを使用可能
// goは複数パッケージを部品のように分けてインポートして使用
// import "time"
// import (
// 	"fmt"
// 	"strconv"
// )
// import "os"
import "time"

// 関数外では暗黙的な変数定義NG
// i5 := 500

// 型指定した変数であれば、関数外であってもOK
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

const Pi = 3.14

const (
	URL = "http://example.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x /y
	r := x % y
	return q, r
}

// func Double(price int) (result int) {
// 	result = price * 2
// 	return
// }

func NoReturn() {
	fmt.Println("No Return")
	return
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("return func")
	}
}

func CallFunc(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func Integer() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

// func anything(a interface{}) {
// 	fmt.Println(a)
// }

// 引数を型アサーション
func anything(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println(v + 10000)
	case string:
		fmt.Println(v + "iiii")
	default:
		fmt.Println("I don't know")
	}
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	// deferは最後に登録したものから評価される
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func sub() {
	for {
		fmt.Println("sub routin")
		time.Sleep(100 + time.Millisecond)
	}
}

// func init() {
// 	fmt.Println("init")
// }

// func init() {
// 	fmt.Println("init2")
// }

// スライスによる可変長引数
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// func Receiver(c chan int) {
// 	for {
// 		i := <-c
// 		fmt.Println(i)
// 	}
// }

func Receiver(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "END")
}

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() { // main()関数 エントリポイント
	// fmt.Println("Hello World")
	// fmt.Println(time.Now())

	// var i int = 100
	// fmt.Println(i)

	// var s string = "Hello Go"
	// fmt.Println(s)

	// var t, f bool = true, false
	// fmt.Println(t, f)

	// var (
	// 	i2 int = 200
	// 	s2 string = "Golang"
	// )
	// fmt.Println(i2, s2)

	// var i3 int
	// var s3 string
	// // 変数宣言のみの場合はその型のデフォルト値が格納される
	// fmt.Println(i3, s3)

	// i3 = 300
	// s3 = "Go"
	// fmt.Println(i3, s3)

	// // 変数の更新
	// i = 150
	// fmt.Println(i)

	// // 暗黙的な変数定義
	// i4 := 400
	// fmt.Println(i4)

	// i4 = 450
	// fmt.Println(i4)

	// // 異なる型の値の代入
	// // i4 = "Hello"
	// // fmt.Println(i4)

	// fmt.Println(i5)

	// outer()

	// var s5 string = "Not used"

	// var i int = 100

	// var i2 int64 = 150

	// fmt.Println(i + 50)

	// 整数型同士であってもbit数が違えば演算NG
	// fmt.Println(i + i2)

	// 書式指定子で変数の型を表示
	// fmt.Printf("%T\n", i2)

	// 型変換
	// fmt.Printf("%T\n", int32(i2))

	// 型変換で型を合わせれば演算OK
	// fmt.Println(i + int(i2))

	// 浮動小数点型
	// var flt64 float64 = 2.4
	// fmt.Println(flt64);

	// flt := 3.2
	// fmt.Println(flt64 + flt)
	// fmt.Printf("%T, %T\n", flt64, flt)

	// var flt32 float32 = 1.2
	// fmt.Printf("%T\n", flt32)

	// fmt.Printf("%T\n", float64(flt32))

	// zero := 0.0
	// pinf := 1.0 / zero
	// fmt.Println(pinf)

	// ninf := -1.0 / zero
	// fmt.Println(ninf)

	// nan := zero / zero
	// fmt.Println(nan)

	// 論理値型
	// var t, f bool = true, false
	// fmt.Println(t, f)

	// 文字列型
	// var s string = "Hello Go"
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)

	// var si string = "300"
	// fmt.Println(si)
	// fmt.Printf("%T\n", si)

	// fmt.Println(`test
	// test
	// 	test`)

	// fmt.Println("\"")
	// fmt.Println(`"`)

	// 文字列はバイト配列の集まり(s[0]はHだが、バイト型だと72. Hで出力するにはstring型に変換してあげる)
	// fmt.Println(string(s[0]))

	// byte(uint8)型
	// byteA := []byte{72, 73}
	// fmt.Println(byteA)
	// fmt.Println(string(byteA))

	// c := []byte("HI")
	// fmt.Println(c)
	// fmt.Println(string(c))

	// 配列型
	// var arr1 [3]int
	// fmt.Println(arr1)
	// fmt.Printf("%T\n", arr1)

	// var arr2 [3]string = [3]string{"A", "B"}
	// fmt.Println(arr2)

	// arr3 := [3]int{1, 2, 3}
	// fmt.Println(arr3)

	// arr4 := [...]string{"C", "D"}
	// fmt.Println(arr4)
	// fmt.Printf("%T\n", arr4)

	// fmt.Println(arr2[0])
	// fmt.Println(arr2[1])
	// fmt.Println(arr2[2])
	// fmt.Println(arr2[2-1])

	// arr2[2] = "C"
	// fmt.Println(arr2)

	// 要素数が異なる配列型への代入はNG
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 配列の要素数を返すlen()
	// fmt.Println(len(arr1))

	// interface型
	// var x interface{}
	// fmt.Println(x)

	// interface型は互換性がある
	// x = 1
	// fmt.Println(x)
	// x = 3.14
	// fmt.Println(x)
	// x = "A"
	// fmt.Println(x)
	// x = [3]int{1, 2, 3}
	// fmt.Println(x)

	// interface型とint型の加算はNG
	// x = 2
	// fmt.Println(x + 1)

	// 型変換
	// var i int = 1
	// flt64 := float64(i)
	// fmt.Println(flt64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("flt64 = %T\n", flt64)

	// i2 := int(flt64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Println(i3)
	// fmt.Printf("i3 = %T\n", i3)

	// 文字列型→数値型
	// var s string = "100"
	// fmt.Printf("s = %T\n", s)

	// 「_」を使用することで戻り値を破棄することができる
	// i, _ := strconv.Atoi(s)
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	// i, err := strconv.Atoi("A")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	// 数値型→文字列型
	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	// var h string = "Hello World"
	// b := []byte(h)
	// fmt.Println(b)

	// h2 := string(b)
	// fmt.Println(h2)

	// 定数
	// fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)

	// fmt.Println(URL, SiteName)

	// fmt.Println(A, B, C, D, E, F)

	// fmt.Println(c0, c1, c2)

	// 算術演算子
	// fmt.Println(1 + 2)
	// fmt.Println("ABC" + "DE")

	// n := 0
	// n += 4
	// fmt.Println(n)
	// n++
	// fmt.Println(n)
	// n--
	// fmt.Println(n)

	// s := "ABC"
	// s += "DEF"

	// fmt.Println(s)

	// 比較演算子
	// fmt.Println(1 == 1)
	// fmt.Println(1 == 2)
	// fmt.Println(1 <= 2)
	// fmt.Println(1 >= 2)
	// fmt.Println(1 < 2)
	// fmt.Println(1 > 2)

	// fmt.Println(true == false)
	// fmt.Println(true != false)

	// 論理演算子
	// fmt.Println(true && false == true)
	// fmt.Println(true && true == true)

	// fmt.Println(true || false == true)
	// fmt.Println(false || false == true)

	// fmt.Println(!true)
	// fmt.Println(!false)

	// i := Plus(1, 2)
	// fmt.Println(i)

	// fmt.Println(Div(9, 3))
	// fmt.Println(Div(9, 4))

	// i2, _ := Div(9, 4)
	// fmt.Println(i2)

	// i4 := Double(10000)
	// fmt.Println(i4)

	// NoReturn()

	// 無名関数
	// f := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(f(1, 2))

	// i2 := func(x, y int) int {
	// 	return x + y
	// }(2, 3)
	// fmt.Println(i2)

	// 戻り値に関数を持つ関数
	// f := ReturnFunc()
	// f()

	// 引数に関数をとる関数
	// CallFunc(func() {
	// 	fmt.Println("call func")
	// })

	// クロージャ
	// f := Later()
	// fmt.Println(f("Hello"))
	// fmt.Println(f("My"))
	// fmt.Println(f("name"))
	// fmt.Println(f("is"))
	// fmt.Println(f("Golang"))

	// ジェネレータ
	// ints := Integer()
	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())

	// otherints := Integer()
	// fmt.Println(otherints())
	// fmt.Println(otherints())
	// fmt.Println(otherints())

	// 条件分岐
	// a := 1
	// if a == 2 {
	// 	fmt.Println("two")
	// } else if a == 1 {
	// 	fmt.Println("one")
	// } else {
	// 	fmt.Println("I don't know")
	// }

	// if b := 100; b == 100 {
	// 	fmt.Println("one hundred")
	// }
	
	// x := 0
	// if x := 2; true {
	//  // 分岐時の簡易分で定義した変数のスコープは分岐内になる
	// 	fmt.Println(x)
	// }
	// このxは「x := 0」
	// fmt.Println(x)

	// エラーハンドリング
	// var s string = "A"

	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("i = %T\n", i)
	// i := 0
	// for {
	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// i2 := 0
	// for i2 < 10 {
	// 	fmt.Println(i2)
	// 	i2++
	// }

	// for i3 := 0; i3 < 10; i3++ {
	// 	if i3 == 3 {
	// 		continue
	// 	}
	// 	if i3 == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i3)
	// }

	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1, 2, 3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
	// for i, _ := range arr {
	// 	fmt.Println(i)
	// }
	// for i := range arr {
	// 	fmt.Println(i)
	// }

	// sl := []string{"Python", "Ruby", "PHP"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 200}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// n := 1
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }
	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }

	// n := 6
	// switch {
	// case n > 0 && n <= 6:
	// 	fmt.Println("0 < n <= 6")
	// case n > 3 && n < 7:
	// 	fmt.Println("3 < n < 7")
	// }

	// var x interface{} = 3
	// i := x.(int) // 型アサーション
	// fmt.Println(i + 2)
	// fmt.Println(x + 2) // interface型なので演算NG

	// interface型で定義されているxはintの値なので、float64で型アサーションしようとするとエラー 
	// f := x.(float64)
	// fmt.Println(f)

	// interface型で定義されているxをfloat64で型アサーションしようとしているが、2つ目の返り値に型アサーションできるかのbool値を代入
	// → 型アサーション失敗のruntime errorを回避できる
	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	// if x == nil {
	// 	fmt.Println("None")
	// } else if i, isInt := x.(int); isInt {
	// 	fmt.Println(i, "x is int")
	// } else if s, isString := x.(string); isString {
	// 	fmt.Println(s, "x is string")
	// } else {
	// 	fmt.Println("I don't know")
	// }

	// switch文による型アサーション
	// switch x.(type) {
	// case int:
	// 	fmt.Println("x is int")
	// case string:
	// 	fmt.Println("x is string")
	// default:
	// 	fmt.Println("I don't know")
	// }

	// switch v := x.(type) {
	// case int:
	// 	fmt.Println(v, "x is int")
	// case string:
	// 	fmt.Println(v, "x is string")
	// default:
	// 	fmt.Println("I don't know")
	// }

	// anything("aaa")
	// anything(1)

	// defer
	// TestDefer()

	// defer func() {
	// 	fmt.Println(1)
	// 	fmt.Println(2)
	// 	fmt.Println(3)
	// }()

	// RunDefer()

	// deferはクリーンアップ処理が最も使用のしどころ
	// file, err := os.Create("./test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// file.Write([]byte("Hello"))

	// panicとrecover
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error")
	// fmt.Println("Hello")

	// goroutinを使用した並列処理
	// go sub()

	// for {
	// 	fmt.Println("main routin")
	// 	time.Sleep(200 + time.Millisecond)
	// }

	// fmt.Println("main")

	// スライス
	// var s1 []int
	// fmt.Println(s1)

	// var s2 []int = []int{100, 200}
	// fmt.Println(s2)

	// s3 := []string{"A", "B"}
	// fmt.Println(s3)

	// s4 := make([]int, 5)
	// fmt.Println(s4)

	// s2[0] = 1000
	// fmt.Println(s2)

	// s5 := []int{1, 2, 3, 4, 5}
	// fmt.Println(s5[0])

	// fmt.Println(s5[2:4])
	// fmt.Println(s5[:2])
	// fmt.Println(s5[2:])
	// fmt.Println(s5[:]) // 全要素
	// fmt.Println(s5[len(s5)-1]) // 最後の要素
	// fmt.Println(s5[1:len(s5)-1])

	// sl := []int{100, 200}
	// fmt.Println(sl)

	// sl = append(sl, 300)
	// fmt.Println(sl)

	// sl = append(sl, 400, 500, 600)
	// fmt.Println(sl)

	// sl2 := make([]int, 5)
	// fmt.Println(sl2)

	// fmt.Println(len(sl2))
	// fmt.Println(cap(sl2))

	// // makeの第3引数はcapacity: メモリの確保サイズ
	// sl3 := make([]int, 5, 10)
	// fmt.Println(len(sl3))
	// fmt.Println(cap(sl3))

	// sl3 = append(sl3, 1, 2, 3, 4, 5, 6)
	// fmt.Println(len(sl3))
	// fmt.Println(cap(sl3)) // capacityは現在の容量を超えると倍増

	// スライスは参照渡し
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000

	// fmt.Println(sl)

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2)

	// スライスのコピー
	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// fmt.Println(sl2)

	// n := copy(sl2, sl)
	// fmt.Println(n, sl2)

	// スライス for
	// sl := []string{"A", "B", "C"}
	// fmt.Println(sl)

	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }
	// for i := range sl {
	// 	fmt.Println(i, sl[i])
	// }
	// for _, v := range sl {
	// 	fmt.Println(v)
	// }
	// for i := 0; i < len(sl); i++ {
	// 	fmt.Println(i, sl[i])
	// }

	// スライスによる可変長引数
	// fmt.Println(Sum(1, 2, 3))

	// fmt.Println(Sum())

	// sl := []int{1, 2, 3}
	// fmt.Println(Sum(sl...))

	// マップ
	// var m = map[string]int{"A": 100, "B": 200}
	// fmt.Println(m)

	// m2 := map[string]int{"A": 100, "B": 200}
	// fmt.Println(m2)

	// m3 := map[int]string{
	// 	1: "A",
	// 	2: "B",
	// }
	// fmt.Println(m3)

	// m4 := make(map[int]string)
	// fmt.Println(m4)

	// m4[1] = "JAPAN"
	// m4[2] = "USA"
	// fmt.Println(m4)

	// fmt.Println(m2["A"])
	// fmt.Println(m4[2])
	// fmt.Println(m4[3]) // Goでは存在しないキーはnilではなく、その型の初期値が返る.

	// 存在しないキーのエラーハンドリング
	// s1 := m4[1]
	// fmt.Println(s1)

	// s2 := m4[3]
	// fmt.Println(s2)

	// s11, ok := m4[1]
	// fmt.Println(s11, ok)

	// s21, ok := m4[3]
	// if !ok {
	// 	fmt.Println("error!")
	// }
	// fmt.Println(s21)

	// _, ok := m4[3]
	// if !ok {
	// 	fmt.Println("error!!")
	// }

	// m4[2] = "US"
	// fmt.Println(m4)

	// m4[3] = "CHINA"
	// fmt.Println(m4)

	// delete(m4, 3)
	// fmt.Println(m4)

	// fmt.Println(len(m4))

	// マップ for
	// m := map[string]int{
	// 	"Apple": 100,
	// 	"Banana": 200,
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// for k := range m {
	// 	fmt.Println(k)
	// }

	// チャネル
	// var ch1 chan int

	// // 受信専用
	// var ch2 <-chan int

	// // 送信専用
	// var ch3 chan<- int

	// ch1 = make(chan int)

	// ch2 := make(chan int)

	// fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	// ch3 := make(chan int, 5)
	// fmt.Println(cap(ch3))

	// ch3 <- 1
	// fmt.Println(len(ch3))

	// ch3 <- 2
	// ch3 <- 3
	// fmt.Println("len", len(ch3))

	// i := <-ch3
	// fmt.Println(i)
	// fmt.Println("len", len(ch3))

	// i2 := <-ch3
	// fmt.Println(i2)
	// fmt.Println("len", len(ch3))

	// fmt.Println(<-ch3)
	// fmt.Println("len", len(ch3))

	// バッファサイズオーバーはdeadlock
	// ch3 <- 1
	// ch3 <- 2
	// ch3 <- 3
	// ch3 <- 4
	// ch3 <- 5
	// ch3 <- 6

	// ch3 <- 1
	// fmt.Println("len", len(ch3))
	// fmt.Println(<-ch3)
	// fmt.Println("len", len(ch3))
	// ch3 <- 2
	// ch3 <- 3
	// ch3 <- 4
	// ch3 <- 5
	// ch3 <- 6

	// チャネルとゴルーチン
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go Receiver(ch1)
	// go Receiver(ch2)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	i++
	// }

	// チャネルclose
	// ch1 := make(chan int, 2)

	// close(ch1)

	// // ch1 <- 1 // closeしたチャネルはデータ受信NG

	// fmt.Println(<-ch1) // closeしたチャネルはデータ送信は可能
	// i, ok := <-ch1 // okにはch1のcloseかつバッファサイズが0の場合にfalse
	// fmt.Println(i, ok)

	// ch1 := make(chan int, 2)

	// ch1 <- 1
	// i, ok := <-ch1
	// fmt.Println(i, ok)

	// ch1 := make(chan int, 2)

	// ch1 <- 1

	// close(ch1)

	// i, ok := <-ch1
	// fmt.Println(i, ok) // ch1はcloseされているけど、バッファにデータが残っているためokはtrue

	// i2, ok := <-ch1
	// fmt.Println(i2, ok) // ch1はcloseされ、バッファにデータなしのためokはfalse
	
	// ch1 := make(chan int, 2)

	// go Receiver("1.goroutin", ch1)
	// go Receiver("2.goroutin", ch1)
	// go Receiver("3.goroutin", ch1)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	i++
	// }
	// close(ch1)
	// time.Sleep(3 * time.Second)

	// ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// close(ch1) // channelのループは空になった状態でdeadlockが発生するのを防ぐため、closeして使うのが基本
	// for i := range ch1 {
	// 	fmt.Println(i)
	// }

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"

	// v1 := <-ch1 // まだ値が未格納のchannelからの送信はNG
	// v2 := <-ch2

	// fmt.Println(v1)
	// fmt.Println(v2)

	// selectの中のcaseはランダム
	// selectの中はchannelに対する処理である必要がある
	// どのcase式が実行されるかはランダム. 後続のcase式が実行されなくなってしまう. 後続のcase式で他のchannelのcase式が実行されなくなるためランダム
	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	default:
		fmt.Println("どちらでもない")
	}

	// // select活用例
	// ch3 := make(chan int)
	// ch4 := make(chan int)
	// ch5 := make(chan int)

	// // receiver
	// go func() {
	// 	for {
	// 		i := <-ch3
	// 		ch4 <- i * 2
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		i2 := <-ch4
	// 		ch5 <- i2 - 1
	// 	}
	// }()

	// n := 0
	// for {
	// 	select {
	// 	case ch3 <- n:
	// 		n++
	// 	case i3 := <-ch5:
	// 		fmt.Println("received", i3)
	// 	}
	// 	if n > 100 {
	// 		break
	// 	}
	// }

	var n int = 100
	fmt.Println(n)

	// fmt.Println(&n)

	// // 値渡しのためnの値は変わらず
	// Double(n)
	// fmt.Println(n)

	var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// *p = 300
	// fmt.Println(n)

	// n = 200
	// fmt.Println(*p)

	// 参照渡しのためnの値も変わる
	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(n)

	// スライスは参照渡しのデータ構造のため、そのまま渡しても変更される
	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)
}
