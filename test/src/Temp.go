package main

import (
	"fmt"
	"time"
)

//var channel chan int = make(chan, int)
//channel := make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	go loop()
	loop()
	time.Sleep(time.Second)
}


