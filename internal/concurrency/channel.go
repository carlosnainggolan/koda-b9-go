package concurrency

import (
	"fmt"
	"time"
)

func RunChannel() {
	chn := make(chan int)
	go quadraticEq(chn)
	chn <- 4
	result := <-chn
	fmt.Println(result)
}

func quadraticEq(chn chan int) {
	num := <-chn
	result := num * num
	time.Sleep(500 * time.Millisecond)
	chn <- result
}