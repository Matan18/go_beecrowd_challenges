package main

import (
	"fmt"
)

func main() {
	var selectedCol int
	count := 0
	var final float32 = 0.0

	var op string
	fmt.Scan(&selectedCol)
	fmt.Scan(&op)

	for ok := true; ok; ok = count < 144 {
		var num float32
		fmt.Scan(&num)
		if count%12 == selectedCol {
			final += num
		}
		count++
	}
	if op == "M" {
		final = final / 12
	}
	fmt.Printf("%.1f\n", final)
}
