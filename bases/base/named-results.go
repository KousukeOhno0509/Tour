//go:build ignore

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 3 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

//go run確認済み
