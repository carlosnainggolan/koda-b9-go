package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func MorningActivities() {
	var wg sync.WaitGroup

	fmt.Println("start")
	wg.Go(takeABath)
	wg.Go(makingCoffee)
	wg.Go(breakFast)
	wg.Go(cleaningBedroom)
	wg.Wait()
	fmt.Println("Go to Work")
}

func takeABath() {
	// defer fmt.Println("Finish")
	time.Sleep(100*time.Millisecond)
	fmt.Println("Start a bath")
	time.Sleep(200*time.Millisecond)
	fmt.Println("Take a bath")
}

func makingCoffee() {
	// defer fmt.Println("Finish")
	time.Sleep(300*time.Millisecond)
	fmt.Println("Making Coffee")
	time.Sleep(400*time.Millisecond)
	fmt.Println("Making Coffee")
}

func breakFast() {
	// defer fmt.Println("Finish")
	time.Sleep(500*time.Millisecond)
	fmt.Println("Take a breakfast")
	time.Sleep(600*time.Millisecond)
	fmt.Println("Take a breakfast")
}

func cleaningBedroom() {
	// defer fmt.Println("Finish")
	time.Sleep(700*time.Millisecond)
	fmt.Println("Cleaning Bedroom")
	time.Sleep(800*time.Millisecond)
	fmt.Println("Cleaning Bedroom")
}