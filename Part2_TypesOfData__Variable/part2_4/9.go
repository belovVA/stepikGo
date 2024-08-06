package main

import "fmt"

func main() {
	var a, b, n, v int
	_, _ = fmt.Scan(&a, &b, &n)
	v = a*100 + b
	fmt.Println(v*n/100, v*n%100)
}
