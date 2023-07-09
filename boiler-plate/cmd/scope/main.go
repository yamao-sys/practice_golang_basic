package main

import (
	"fmt"
	f "fmt"
	. "fmt"
	"github.com/boiler-plate/cmd/scope/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = "s"
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)// 外部パッケージでは大文字スタートのものはpublic
	// fmt.Println(foo.min) // 外部パッケージでは小文字スタートのものはprivate

	f.Println(foo.ReturnMin()) // 大文字スタートなのでpublic
	Println(foo.ReturnMin())

	fmt.Println(appName())

	fmt.Println(Do("AAA"))
}
