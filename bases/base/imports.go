//go:build ignore

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(time.Now())
}

//go run確認済み
