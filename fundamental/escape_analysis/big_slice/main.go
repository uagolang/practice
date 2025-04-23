package main

import "fmt"

func main() {
	// 8 byte - розмір int для 64-bit процесорів

	// маленький слайс залишиться на стеку
	smallSlice := make([]int, 10) // 10 * 8 byte (для 64-bit) = 80 byte
	fmt.Println("Small slice len:", len(smallSlice))

	// дуже великий слайс, ймовірно, втече на купу
	// поріг залежить від версії Go, але зазвичай > 64KB
	largeSlice := make([]int, 10000) // 10000 * 8 byte = 80000 byte (~78KB)
	fmt.Println("Large slice len:", len(largeSlice))
}

// команда компіляції:
// go build -gcflags="-m" ./fundamental/escape_analysis/big_slice/main.go
