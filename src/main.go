package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func allocate(size int) [][]bool {
	table := make([][]bool, size)

	for i := 0; i < size; i++ {

		table[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			table[i][j] = false
		}
	}
	return table
}

func is_safe(table [][]bool, i int, j int, size int) bool {
	for k := 0; k < size; k++ {
		if table[k][j] && k != i {
			return false
		}
	}
	for k := 0; k < size; k++ {
		if table[i][k] && k != j {
			return false
		}
	}
	for k := 1; k < size; k++ {
		if i - k >= 0 && j - k >= 0 && table[i - k][j - k] {
			return false
		}
		if i + k < size && j - k >= 0 && table[i + k][j - k] {
			return false
		}
		if i + k < size && j + k < size && table[i + k][j + k] {
			return false
		}
		if i - k >= 0 && j + k < size && table[i - k][j + k] {
			return false
		}
	}
	return true
}

func nqueen(size int, table [][]bool, step int, wgprec *sync.WaitGroup, possibilities *int, col *int) int {
    if step == size {
        if possibilities != nil {
            *possibilities++
        }
        return 1
    }
    var total int
    if step == 0 {
        var wg sync.WaitGroup
        table = allocate(size)
        for i := 0; i < size; i++ {
            newTable := allocate(size)
            newTable[0][i] = true
            if is_safe(newTable, 0, i, size) {
                wg.Add(1)
                go func(col int) {
                    defer wg.Done()
                    localPoss := 0
                    nqueen(size, newTable, 1, nil, &localPoss, &col)
                    total += localPoss
                }(i)
            }
        }
        wg.Wait()
        return total
    }
    for i := 0; i < size; i++ {
        table[step][i] = true
        if is_safe(table, step, i, size) {
            total += nqueen(size, table, step+1, wgprec, possibilities, col)
        }
        table[step][i] = false
    }

    return total
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
	solution := nqueen(nb, nil, 0, nil, nil, nil)
	fmt.Println("Solutions :", solution)
	return
}
