package main

import "fmt"

func main() {
	s := 0
	var v int
	for i := 0; i < 4; i++ {
		_, _ = fmt.Scan(&v)
		s += v

	}
	fmt.Println(s * 3)
}
