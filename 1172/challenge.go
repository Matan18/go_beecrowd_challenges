package main

import (
	"fmt"
)

func main() {
	var scanned int
	var arr []int
	for i := 0; i < 10; i++ {
		fmt.Scan(&scanned)
		if scanned <= 0 {
			scanned = 1
		}
		arr = append(arr, scanned)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("X[%v] = %v\n", i, arr[i])
	}
}
