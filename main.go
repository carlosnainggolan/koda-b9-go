package main

import "fmt"

func main() {
	a, k := gabungan(8,5)
	fmt.Printf("Luas: %d\n", a)
	fmt.Printf("Keliling: %d\n", k)
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

func area (p uint8, l uint8) uint8 {
	return p * l
}

func keliling (p uint8, l uint8) uint8 {
	return 2 * (p + l)
}

func gabungan (p uint8, l uint8) (a uint8, k uint8) {
	a = area(p, l)
	k = keliling(p, l)
	return a, k
}