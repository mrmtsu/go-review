package main

import (
	"fmt"
)

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	// :=は関数の中でしか使用できない
	xi := 1
	xi = 2
	xf64 := 1.2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xf32, xs, xt, xf)

	// Printfの"%T"で型を調べて出力することができる
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
