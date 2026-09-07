package main

import "fmt"

func main() {
	// Minitask 1
	// area, keliling := gabungan(8,5)
	// fmt.Printf("Luas: %d\n", area)
	// fmt.Printf("Keliling: %d\n", keliling)

	// Minitask 2
	err := window(21)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} 

	slice(88)
}

// func greet (name string) {
// 	fmt.Printf("Hello %s", name)
// }

// func prinTwoDecimalsFloat (num float32) string {
// 	result :=	fmt.Sprintf("\nHasil: %.2f\n", num)
// 	return result
// }

// //return 2
// func addAndSub (a int8, b int8)(resultAdd int16, resultSub int8) {
// 	resultAdd = int16(a) + int16(b)
// 	resultSub = a- b
// 	return resultAdd, resultSub
// }

// func area (p uint8, l uint8) uint8 {
// 	return p * l
// }

// func keliling (p uint8, l uint8) uint8 {
// 	return 2 * (p + l)
// }

// func gabungan (p uint8, l uint8) (a uint8, k uint8) {
// 	a = area(p, l)
// 	k = keliling(p, l)
// 	return a, k
// }

func window (n int) error {
	if n < 3 {
		return fmt.Errorf("Harus lebih dari 3")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
				if j == 0 || j == n-1 || i == 0 || i == n-1 || j == n/2 || i == n/2{
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		fmt.Println()
	}
	return nil
}

//menyisipkan angka setleah 66
func slice (num int8) {
	a := []int{50, 75, 66, 20, 32, 90}
	for i, v := range a {
		if v == 66 {
			a = append(a[:i+1], append([]int{int(num)}, a[i+1:]...)...)
			break
		}
	}
	for _, v := range a {
		fmt.Println(v)
	}
}