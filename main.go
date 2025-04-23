package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("====== Kalkulator CLI ======")
		fmt.Println("1. Tambah (+)")
		fmt.Println("2. Kurang (-)")
		fmt.Println("3. Kali (*)")
		fmt.Println("4. Bagi (/)")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih operasi (1-5): ")

		var choice int
		fmt.Scan(&choice)

		if choice == 5 {
			fmt.Println("Terima kasih! 👋")
			os.Exit(0)
		}

		var num1, num2 float64
		fmt.Print("Masukkan angka pertama: ")
		fmt.Scan(&num1)
		fmt.Print("Masukkan angka kedua: ")
		fmt.Scan(&num2)

		switch choice {
		case 1:
			fmt.Printf("Hasil: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)
		case 2:
			fmt.Printf("Hasil: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)
		case 3:
			fmt.Printf("Hasil: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)
		case 4:
			if num2 == 0 {
				fmt.Println("Error: Tidak bisa dibagi nol")
			} else {
				fmt.Printf("Hasil: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)
			}
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println() // Biar jaraknya rapi
	}
}
