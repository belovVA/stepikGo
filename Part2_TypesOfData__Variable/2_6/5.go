package main

import "fmt"

func main() {
	var numb float64
	_, _ = fmt.Scan(&numb)
	fmt.Println(numb - (float64(int64(numb))))
}
