package main

import "fmt"

func main() {
	var v int
	s := 1
	for i := 0; i < 3; i++ {
		_, _ = fmt.Scan(&v)
		s *= v

	}
	fmt.Println(s)
}
