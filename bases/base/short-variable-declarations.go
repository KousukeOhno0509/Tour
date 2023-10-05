//go:build ignore

package main

import "fmt"

func main() {
	//関数内でのver宣言のかわりの:=を使用し暗黙的な方宣言を行う
	i, j := 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

//go run確認済み
