package main

import "fmt"

func main() {
	x := 42
	y := &x         // беремо адресу x
	fmt.Println(*y) // використовуємо всередині main
}

// команда компіляції:
// go build -gcflags="-m" ./fundamental/escape_analysis/interface_args/main.go
