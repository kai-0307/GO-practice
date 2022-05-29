package main

import (
	"fmt"
	"math"
)

func main() {

	// Printf / Sqrt はそれぞれのパッケージからエクスポートされている
	fmt.Printf("Sqrt(9) is %g\n", math.Sqrt(9))

	// 小文字で始まる名前は外部に公開されないので、エラーが発生する
	fmt.Println(math.Pi)

	// hello!!!!!!!!!!!
}

//関数

//関数はfuncで定義する。
//引数の型は変数の後ろに書く。

// ２つのintを受け取り、intを返す関数

func add(x int, y int) int {
	return x + y 
}