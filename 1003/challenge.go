package main

import (
  "fmt"
)

func sum(a int, b int) int {
	return int(a + b)
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)


	fmt.Println(fmt.Sprintf("SOMA = %v", sum(a, b)))
}
	