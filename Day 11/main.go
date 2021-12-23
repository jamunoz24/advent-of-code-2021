package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	datalist, _ := os.ReadFile("input.in")

	energyLvlz := [][]int{}
	curLine := []int{}
	for _, bit := range datalist {
		if bit != '\n' {
			bitint, _ := strconv.Atoi(string(bit))
			curLine = append(curLine, bitint)
		} else {
			energyLvlz = append(energyLvlz, curLine)
			curLine = nil
		}
	}

	printMap := false
	if printMap {
		fmt.Println("Starting Step:")
		printMatrix(energyLvlz)
	}

	steps := 0
	flashes := 0
	allZeroes := false
	// for k := 0; k < steps; k++ { // Step 1
	for !allZeroes {
		for i := range energyLvlz {
			for j, num := range energyLvlz[i] {
				num += 1
				energyLvlz[i][j] = num
			}
		}
		splat(energyLvlz, &flashes)
		if printMap {
			fmt.Printf("Step %d:\n", steps+1)
			printMatrix(energyLvlz)
		}
		steps += 1
		allZeroes = allZeros(energyLvlz)
	}

	fmt.Printf("Flashes: %d\n", flashes)
	fmt.Printf("Steps 'til Zeroes: %d\n", steps)

}

// Part 1:
func splat(arr [][]int, flashes *int) {
	size := len(arr)
	convert := true
	for convert {
		convert = false
		for i := range arr {
			for j, num := range arr[i] {
				if num > 9 {
					convert = true
					(*flashes) += 1
					arr[i][j] = 0
					// Ups
					if i-1 >= 0 {
						nonZeroAdd(arr, i-1, j)
						if j-1 >= 0 {
							nonZeroAdd(arr, i-1, j-1)
						}
						if j+1 < size {
							nonZeroAdd(arr, i-1, j+1)
						}
					}
					// Left
					if j-1 >= 0 {
						nonZeroAdd(arr, i, j-1)
					}
					// Right
					if j+1 < size {
						nonZeroAdd(arr, i, j+1)
					}
					// Downs
					if i+1 < size {
						nonZeroAdd(arr, i+1, j)
						if j-1 >= 0 {
							nonZeroAdd(arr, i+1, j-1)
						}
						if j+1 < size {
							nonZeroAdd(arr, i+1, j+1)
						}
					}
				}
			}
		}
	}
}

func nonZeroAdd(arr [][]int, k int, l int) {
	if arr[k][l] != 0 {
		arr[k][l] += 1
	}
}

func allZeros(arr [][]int) bool {
	allZeros := true
	for _, row := range arr {
		for _, num := range row {
			if num != 0 {
				allZeros = false
				break
			}
		}
		if !allZeros {
			break
		}
	}
	return allZeros
}

func printMatrix(arr [][]int) {
	// Printing Matrix
	for _, line := range arr {
		for _, num := range line {
			print(num)
		}
		fmt.Println()
	}
}
