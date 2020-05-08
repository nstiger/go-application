package main

import (
	"fmt"
	"os"
)

//ReadData is a function to read data from data.in
func ReadData(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(f, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(f, "%d", &maze[i][j])
		}
	}
	return maze 
}

func main() {
	maze :=ReadData("data.in") 
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}
}