package main

import (
	"fmt"
	"os"
)

//ReadData read data from datafile
func ReadData(file string) [][]int {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(f, "%d %d", &row, &col)

	var s [][]int
	// for _, m := range s {
	// 	for j, val := range m {
	// 	m[j], _ = fmt.Fscanf(f, "%d ", val)
	// 	}
	// }

	for i := 0; i < row-1; i++ {
		for j := 0; j < col-1; j++ {
			_, err = fmt.Fscanf(f, "%d ", s[i][j])
		}
	}
	fmt.Println(s)

	return s

}

func main() {
	var data [][]int
	data = ReadData("data.in")

	for _, m := range data {
		for _, val := range m {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
