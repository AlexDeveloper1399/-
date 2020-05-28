package main

import (
	"fmt"
	"math"
)

func Two() {
	var a float64
	fmt.Scan(&a)
	var c float64

	b := math.Ceil(math.Log10(math.Abs(a) + 0.5))

	c = a / math.Pow(10, b-1)

	fmt.Println(int(c))

}
