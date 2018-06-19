package main

import "fmt"

var test1,test2, test3 = "v1", "v2", "v3"

func main() {
	fmt.Println("hello world!")

	//complex
	var c complex64 = 5+5i
	fmt.Printf("value is : %G",c)

	//string
	var s string = "hello"
	b := []byte(s)
	b[0] = 'c'
	s2 := string(b)
	fmt.Println(s2)
}
