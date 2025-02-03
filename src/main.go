package main

import (
	"fmt"
	"os"
	"strconv"
)

func allocate(size int) [][]bool {
	plate := make([][]bool, size)

	for i := 0; i < size; i++ {
		plate[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			plate[i][j] = false
		}
	}
	return plate
}

func is_safe(plate [][]bool, i int, j int, size int) bool {
	for k := 0; k < size; k++ {
		if plate[k][j] && k != i {
			return false
		}
	}
	for k := 0; k < size; k++ {
		if plate[i][k] && k != j {
			return false
		}
	}
	// Diagonals check :
	return true
}

func nqueen(size int, plate [][]bool, step int) int {
	possibilities := 0

	if step == size {
		return 1
	}
	for i := 0; i < size; i++ {
		plate[step][i] = true
		if is_safe(plate, step, i, size) {
			possibilities += nqueen(size, plate, step + 1)
		}
		plate[step][i] = false
	}
	return possibilities
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Invalid number of arguments !")
		os.Exit(84)
	}
	nb, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number !")
		os.Exit(84)
	}
	plate := allocate(nb)
	solution := nqueen(nb, plate, 0)
	fmt.Println("Solutions :", solution)
	return
}
