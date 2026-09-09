package main

import (
	"bufio"
	// "fmt"

	"os"
	"strconv"
	"strings"

	// "fmt"

	"fmt"

	"github.com/carlosnainggolan/koda-b9-git/internal/concurrency"
	"github.com/carlosnainggolan/koda-b9-git/internal/model"
	"github.com/carlosnainggolan/koda-b9-git/internal/sales"
	"github.com/carlosnainggolan/koda-b9-git/internal/service"
)

func main() {
	// Minitask 1
	// area, keliling := Gabungan(8,5)
	// fmt.Printf("Luas: %d\n", area)
	// fmt.Printf("Keliling: %d\n", keliling)

	// Minitask 2

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Hitung persegi panjang")
		fmt.Println("2. Membentuk jendela")
		fmt.Println("3. Inject to Slice")
		fmt.Println("4. Lihat Biodata")
		fmt.Println("5. Open File Path")
		fmt.Println("6. Setter Getter")
		fmt.Println("7. Payment")
		fmt.Println("8. Concurrency")
		fmt.Println("9. Morning activities")
		fmt.Println("10. Channel example")
		fmt.Println("11. Message")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu (0-9): ")

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

		case "5":
			fmt.Print("Masukkan file path")
			scanner.Scan()
			f := strings.TrimSpace(scanner.Text())
			err := service.ReadingFile(f)
			if err != nil {
				fmt.Println("Terjadi Error", err)
			}

		case "6":
			personData := model.Constructor("Carlos", "Jakarta", "088294649371")
			fmt.Println(personData.PrintData())
			fmt.Println(personData.Greet())
			personData.Setter("Rahel")
			fmt.Println(personData.Greet())
			personData.Setter("Dodi")
			fmt.Println(personData.Greet())

		case "7":
			list := []int{1000, 10000, 3000, 4000}
			bank := sales.Bank{}
			online := sales.Online{}
			fiktif := sales.Fictional{}

			sales.Result(&bank, list)
			sales.Result(&online, list)
			sales.Result(&fiktif, list)

		case "8":
			concurrency.RunConcurrent()

		case "9":
			concurrency.MorningActivities()

		case "10":
			concurrency.RunChannel()

		case "11":
			concurrency.Run()
			

		case "0":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Menu tidak valid, silakan pilih lagi.")
		}
	}
}

	




