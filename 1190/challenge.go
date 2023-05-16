package main

import (
	"fmt"
)

func calc(op string, matriz [12][12]float32) float32 {
	var final float32 = 0
	var count float32 = 0
	col_count := len(matriz)
	row_count := len(matriz[0])

	for i := 0; i < col_count; i++ {
		for j := 0; j < row_count; j++ {
			if j > i && i+j >= col_count {
				final += matriz[i][j]
				count++
			}
		}
	}

	if op == "M" {
		return final / count
	}

	return final
}

func main() {
	var op string
	var matriz [12][12]float32

	fmt.Scan(&op)
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}
	fmt.Printf("%.1f\n", calc(op, matriz))
}
