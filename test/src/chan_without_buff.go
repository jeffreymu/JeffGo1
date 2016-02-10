package main

import "fmt"

var a string
var c = make(chan int)

func f2() {
	a = "hello golang"
	<-c
}
func main() {
	go f2()
	c <- 0
	fmt.Println(a)
}
