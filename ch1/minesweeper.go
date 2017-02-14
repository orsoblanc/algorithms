package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"math"
	"strconv"
)

type Minesweeper struct {
	n int
	m int
	answer [][]int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int = -1, -1
	minesweepers := []Minesweeper{}

	n,m = readNM(reader)

	for n != 0 && m != 0 {
		minesweeper := Minesweeper{n: n, m: m}
		minesweeper.answer = make([][]int, n)

		for i := range minesweeper.answer {
			minesweeper.answer[i] = make([]int, m)
		}

		for i := 0; i < n; i++ {
			bytes, _ := reader.ReadBytes('\n')
			resolve(i, bytes, &minesweeper)
		}

		minesweepers = append(minesweepers, minesweeper)
		n,m = readNM(reader)

	}

	// print answer
	printResult(minesweepers)
}

func printResult(minesweepers []Minesweeper) {
	for i := 0 ; i < len(minesweepers); i++ {
		ms := minesweepers[i]
		fmt.Printf("Field #%d:\n" + strconv.Itoa(i+1))

		for i := 0 ; i < ms.n; i ++ {
			for j:= 0; j < ms.m ; j++ {
				if ms.answer[i][j] == math.MinInt8{
					fmt.Print("*")
				} else {
					fmt.Print(ms.answer[i][j])
				}
			}
			fmt.Print("\n")
		}
	}

}

func readNM(reader *bufio.Reader) (int, int) {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	dims := strings.Split(text, " ")

	n, _ := strconv.Atoi(dims[0])
	m, _ := strconv.Atoi(dims[1])

	return n, m
}

func resolve(currRow int, bytes []byte, m *Minesweeper) {

	for i:= 0 ; i < len(bytes); i++ {
		byte := bytes[i]

		if string(byte) == "*" {
			// inc 1, around
			// where am i (row, i)
			m.answer[currRow][i] = math.MinInt8

			for r := currRow - 1 ; r <= currRow + 1; r++ {
				for c := i - 1; c <= i + 1; c ++ {
					// boundary check
					if !(r == currRow && c == i) && (r >= 0 && r < m.n) && (c >= 0 && c < m.m) && (m.answer[r][c] != math.MinInt8) {
						m.answer[r][c] ++
					}
				}
			}
		}
	}

}