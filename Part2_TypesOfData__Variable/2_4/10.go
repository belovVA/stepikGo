package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println(n, "мин - это", n/60%24, "час", n%60, "минут.")
}
