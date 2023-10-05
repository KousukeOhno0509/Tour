//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//0から10までのランダムな数値を表示する
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
