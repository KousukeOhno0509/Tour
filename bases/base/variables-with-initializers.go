//go:build ignore

package main

import "fmt"

//初期化子による型省略
var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//go run確認済み
