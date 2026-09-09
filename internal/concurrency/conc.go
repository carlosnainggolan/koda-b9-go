package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RunConcurrent() {
	var wg sync.WaitGroup

	fmt.Println("Start")

	// wg.Add(2)
	wg.Go(getCharacter)
	wg.Go(getDigit)
	// go getChars(&wg)
	// go getDigits(&wg)
	wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}

func getChars(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Get charts done")
	for _,char := range "hello" {
		time.Sleep(30 * time.Millisecond)
		fmt.Println(string(char))
	}
}

func getCharacter() {
	// defer wg.Done()
	defer fmt.Println("Get charts done")
	for _,char := range "hello" {
		time.Sleep(30 * time.Millisecond)
		fmt.Println(string(char))
	}
}

func getDigits(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Get digits done")
	for i:=1; i <=5; i++ {
		time.Sleep(time.Duration(10*1)*time.Millisecond)
		fmt.Println(i)
	}
}

func getDigit() {
	// defer wg.Done()
	defer fmt.Println("Get digits done")
	for i:=1; i <=5; i++ {
		time.Sleep(time.Duration(10*1)*time.Millisecond)
		fmt.Println(i)
	}
}