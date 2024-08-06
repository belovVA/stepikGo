package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	splitter := scanner.Text()
	_ = scanner.Scan()
	s1 := scanner.Text()
	_ = scanner.Scan()
	s2 := scanner.Text()
	_ = scanner.Scan()
	s3 := scanner.Text()
	fmt.Print(s1)
	fmt.Print(splitter)
	fmt.Print(s2)
	fmt.Print(splitter)
	fmt.Print(s3)

}
