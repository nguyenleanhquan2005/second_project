package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var firstname string = "Thang"
	var lastname string = "Nguyen"
	fmt.Println("hello!I am", firstname, lastname)
	a := "Quan"
	b := "Nguyen"
	msg := getMessages(a, b)
	fmt.Println(msg)

}

func getMessages(firstname string, lastname string) string {
	return "Hello I am" + " " + firstname + " " + lastname

}
