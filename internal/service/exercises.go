package service

import "fmt"

func Window (n int) error {
	if n < 3 {
		return fmt.Errorf("Harus lebih dari 3")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
				if j == 0 || j == n-1 || i == 0 || i == n-1 {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		fmt.Println()
	}
	return nil
}

func Slice () {
	a := []int{50, 75, 66, 20, 32, 90}
	for i, v := range a {
		if v == 66 {
			a = append(a[:i+1], append([]int{88}, a[i+1:]...)...)
			break
		}
	}
	for _, v := range a {
		fmt.Println(v)
	}
}
