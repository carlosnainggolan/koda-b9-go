package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/carlosnainggolan/koda-b9-git/internal/model"
	"github.com/carlosnainggolan/koda-b9-git/internal/service"
)

func main() {
	// Minitask 1
	// area, keliling := gabungan(8,5)
	// fmt.Printf("Luas: %d\n", area)
	// fmt.Printf("Keliling: %d\n", keliling)

	// Minitask 2

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Hitung persegi panjang")
		fmt.Println("2. Membentuk jendela")
		fmt.Println("3. Inject to Slice")
		fmt.Println("4. Lihat Biodata")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu (0-4): ")

		if !scanner.Scan() {
			break
	}
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
		case "1":
			fmt.Print("Masukkan lebar")
			scanner.Scan()
			l, _ := strconv.Atoi(scanner.Text()) 

			fmt.Print("Masukkan panjang")
			scanner.Scan()
			p, _ := strconv.Atoi(scanner.Text()) 

			luas, keliling := service.Gabungan(uint8(p), uint8(l))
			fmt.Printf("Luas: %d\n Keliling: %d", luas, keliling)
		
		case "2":
			fmt.Print("Masukkan ukuran jendela")
			scanner.Scan()
			p,_ := strconv.Atoi(scanner.Text())

			err := service.Window(p)
			if err != nil {
				fmt.Println(err.Error())
			}

		case "3":
			fmt.Print("Insert angka ke slice")
			service.Slice()


		case "4":
			myBiodata := model.Biodata {
				Nama: "Carlos",
				Foto: "Img",
				Email: "carlosnainggolan@gmail.com",
				Umur: 25,
				NomorTelepon: "088294649371",
				StatusPernikahan: true,
				RiwayatPendidikan: []model.Pendidikan{
					{
						Nama: "UPN Veteran Yogyakarta",
						Jurusan: "Sistem Informasi",
					},
					{
						Nama: "SMA RK Budi Mulia",
						Jurusan: "Saintek",
					},
				},
			}
			fmt.Println(myBiodata)

		case "0":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Menu tidak valid, silakan pilih lagi.")
		}
	}
}

	// err := window(21)
	// if err != nil {
	// 	fmt.Println("Error:", err.Error())
	// } 

	// slice(88)

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




