package main

import "fmt"

func main() {
	var a, b float64
	_, _ = fmt.Scan(&a, &b)
	fmt.Println((a + b) / 2.0)
}
