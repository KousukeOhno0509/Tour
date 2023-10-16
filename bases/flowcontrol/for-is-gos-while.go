//go:build ignore

package main

import "fmt"

func main() {
	sum := 1
	for sum < 50 {
		sum += sum
	}
	fmt.Println(sum)
}

// go run確認済み
