package main

import (
	"fmt"
	"math"
)

func main() {
	var value float64
	_, _ = fmt.Scan(&value)
	fmt.Println(value / math.Pow(2.0, 13))

}
