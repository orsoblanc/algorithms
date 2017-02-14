package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {

	if len(os.Args) == 3 {
		start, _ := strconv.ParseInt(os.Args[1], 0, 32)
		end, _ := strconv.ParseInt(os.Args[2], 0, 32)

		var max int = 0

		for i := start ; i <= end ; i++ {
			cycle := find3Plus1Cycle(i)

			if max < cycle {
				max = cycle
			}
		}

		fmt.Printf("%d %d %d", start, end, max)
	}

}


func find3Plus1Cycle(num int64) int {
	var cycle = 1

	for num > 1 {
		if num % 2 == 0 {
			num = num / 2
		} else {
			num = num * 3 + 1
		}
		cycle ++
	}

	return cycle
}