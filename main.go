package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("Hello foo")
}

func main() {
	/*
			foo()
			defer fmt.Println("world")

			fmt.Println("Hello")


		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("seccess")
	*/

	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
