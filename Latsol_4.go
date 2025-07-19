package main

import (
	"errors"
	"fmt"
)

// Fungsi untuk menyaring bilangan ganjil dari slice
func FilterOdd(numbers []int) ([]int, error) {
	var oddNumbers []int
	for _, num := range numbers {
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}

	if len(oddNumbers) == 0 {
		return nil, errors.New("tidak ada bilangan ganjil yang ditemukan")
	}

	return oddNumbers, nil
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah bilangan harus lebih dari 0.")
		return
	}

	numbers := make([]int, n)

	fmt.Println("Masukkan bilangan satu per satu:")
	for i := 0; i < n; i++ {
		fmt.Printf("Bilangan ke-%d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	oddNumbers, err := FilterOdd(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Bilangan ganjil yang ditemukan:", oddNumbers)
	}
}
