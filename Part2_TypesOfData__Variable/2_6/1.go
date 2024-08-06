package main

import "fmt"

func main() {
	pi := 3.14
	var r float64
	_, _ = fmt.Scan(&r)
	fmt.Println(pi * r * r)
}
