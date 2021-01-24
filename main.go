package main

import "fmt"

// １文字目を大文字にすると他のファイルでも使用できる
const Pi = 3.14

// constは型を指定しない
const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
}
