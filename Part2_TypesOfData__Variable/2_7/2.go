package main

import "fmt"

func main() {
	var x int64
	_, _ = fmt.Scan(&x)
	fmt.Println(x % 1000 / 100)
}
