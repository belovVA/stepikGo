package main

import "fmt"

func main() {
	var v int
	_, _ = fmt.Scan(&v)
	for i := 0; i < 3; i++ {
		fmt.Println(v)
		v++
	}
}
