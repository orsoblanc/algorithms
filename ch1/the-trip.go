package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for n != 0 {

		var total float64 = 0.0
		var avg float64 = 0.0
		var diff float64 = 0.0

		money := make([]float64, n)

		for i := 0; i < n; i ++ {
			fmt.Scanf("%f", &money[i])
			total += money[i]
		}

		avg = math.Ceil((total / float64(n)) * 100) / 100	// precision 2

		for i := 0; i < n; i++ {
			if money[i] > avg {
				diff += money[i] - avg
			}
		}

		fmt.Printf("$%.2f \n", diff)
		fmt.Scanf("%d", &n)

	}

}