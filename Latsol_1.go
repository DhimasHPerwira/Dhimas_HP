package main

import "fmt"

// Fungsi swap yang menggunakan pointer untuk menukar nilai dua string
func Swap(a, b *string) {
	*a, *b = *b, *a
}

func main() {
	str1 := "Hello"
	str2 := "World"

	fmt.Println("Sebelum swap:")
	fmt.Println("str1 =", str1)
	fmt.Println("str2 =", str2)

	Swap(&str1, &str2)

	fmt.Println("\nSetelah swap:")
	fmt.Println("str1 =", str1)
	fmt.Println("str2 =", str2)
}
