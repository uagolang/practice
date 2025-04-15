package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse("hello"))
	fmt.Println(reverse("wør¬∂"))
}

func reverse(s string) string {
	runes := make([]rune, len(s))
	var n int
	for _, val := range s {
		runes[n] = val
		n++
	}

	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
