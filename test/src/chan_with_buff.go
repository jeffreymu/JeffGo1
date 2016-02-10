package main

import (
	"fmt"
)

var a string
var c = make(chan int, 10)

func f1() {
	a = "hello golang"
	c <-0
}

func main() {
	go f1()
	<-c
	fmt.Println(a)
}

