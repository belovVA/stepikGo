package main

import "fmt"

func main() {
	s := 0.0
	m := 1.0
	var val float64
	for i := 0; i < 3; i++ {
		_, _ = fmt.Scan(&val)
		s += val
		m *= val
	}
	fmt.Println(s, m)
}
