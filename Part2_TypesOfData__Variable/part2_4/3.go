package main

import "fmt"

func main() {
	var dig int
	_, _ = fmt.Scan(&dig)
	digDouble := dig * dig
	fmt.Println((digDouble * digDouble * digDouble))
}
