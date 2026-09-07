package main

import "fmt"

func main() {
	fmt.Println(area(8,5))
	fmt.Println(keliling(8,5))
	fmt.Println(gabungan(8,5))
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

func area (p uint8, l uint8) (a uint8) {
	a = p*l
	return a
}

func keliling (p uint8, l uint8) (k uint8) {
	k = 2*(p+l)
	return k
}

func gabungan (p uint8, l uint8) (a uint8, k uint8) {
	a = p*l
	k = 2*(p+l)
	return a, k
}