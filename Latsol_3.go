package main

import (
	"errors"
	"fmt"
)

// Fungsi untuk menyaring bilangan genap dari slice
func FilterEven(numbers []int) ([]int, error) {
	var evenNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}

	if len(evenNumbers) == 0 {
		return nil, errors.New("tidak ada bilangan genap yang ditemukan")
	}

	return evenNumbers, nil
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

	evenNumbers, err := FilterEven(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Bilangan genap yang ditemukan:", evenNumbers)
	}
}
