//go:build ignore

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func stringAdd(a string, b string) string {
	return a + b
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(stringAdd("42", "13"))
}

//go run確認済み
