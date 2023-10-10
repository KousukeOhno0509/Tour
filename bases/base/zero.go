//go:build ignore

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var a uint64
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, a)
}

//go run確認済み
