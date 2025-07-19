package main

import (
	"fmt"
)

// Fungsi untuk mencari bilangan terbesar, menghindari panic
func Max(numbers []int) (int, bool) {
	if len(numbers) == 0 {
		return 0, false
	}

	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max, true
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

	if maxValue, ok := Max(numbers); ok {
		fmt.Println("Bilangan terbesar adalah:", maxValue)
	} else {
		fmt.Println("Data kosong, tidak bisa mencari bilangan terbesar.")
	}
}
