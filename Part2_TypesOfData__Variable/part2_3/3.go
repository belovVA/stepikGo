package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	s1 := scanner.Text()
	_ = scanner.Scan()
	s2 := scanner.Text()
	_ = scanner.Scan()
	s3 := scanner.Text()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
