package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	datalist, _ := os.ReadFile("input.in")

	// Creating the heatmap and checked map
	var heatmap = [][]int{}
	curList := []int{}
	for _, str := range datalist {
		num, err := strconv.Atoi(string(str))
		if err == nil {
			curList = append(curList, num)
		} else {
			heatmap = append(heatmap, curList)
			curList = nil
		}
	}
	sizei := len(heatmap)
	sizej := len(heatmap[0])
	checkedmap := make([][]bool, sizei)
	for i := 0; i < sizei; i++ {
		checkedmap[i] = make([]bool, sizej)
	}

	//Part 1
	// Starting the traversal
	lowestPoints := []int{}
	for i, row := range heatmap {
		for j, num := range row {
			lowVal := num
			// Checking up
			if i-1 >= 0 {
				if heatmap[i-1][j] < num {
					lowVal = heatmap[i-1][j]
				}
			}
			// Checking left
			if j-1 >= 0 {
				if heatmap[i][j-1] < num {
					lowVal = heatmap[i][j-1]
				}
			}
			// Checking right
			if j+1 < sizej {
				if heatmap[i][j+1] < num {
					lowVal = heatmap[i][j+1]
				}
			}
			// Checking down
			if i+1 < sizei {
				if heatmap[i+1][j] < lowVal {
					lowVal = heatmap[i+1][j]
				}
			}
			// Adding it to the bag
			if lowVal == num && num != 9 {
				lowestPoints = append(lowestPoints, num)
			}
		}
	}

	// Getting the sum
	sum := 0
	for _, num := range lowestPoints {
		sum += num + 1
	}

	fmt.Printf("Total Sum: %d\n", sum)

	// Part 2

	basins := []int{}
	for i, row := range heatmap {
		for j, num := range row {
			if num != 9 && !checkedmap[i][j] {
				basinSize := startTraversal(i, j, heatmap, checkedmap)
				basins = append(basins, basinSize)
			}
		}
	}

	bigBasins := []int{}
	maxN := 0
	maxInd := -1
	for i := 0; i < 3; i++ {
		for i, num := range basins {
			if maxN < num {
				maxN = num
				maxInd = i
			}
		}
		bigBasins = append(bigBasins, maxN)
		basins[maxInd] = 0
		maxN = 0
		maxInd = -1
	}

	totalBasins := 1
	for _, basin := range bigBasins {
		totalBasins *= basin
	}

	fmt.Println(bigBasins)

	fmt.Printf("Total Basins: %d\n", totalBasins)

}

func startTraversal(i int, j int, heatmap [][]int, checkedmap [][]bool) int {
	sizei := len(heatmap)
	sizej := len(heatmap[0])
	checkedmap[i][j] = true

	sum := 1
	// Checking up
	if i-1 >= 0 {
		if heatmap[i-1][j] < 9 && !checkedmap[i-1][j] {
			sum += startTraversal(i-1, j, heatmap, checkedmap)
		}
	}
	// Checking left
	if j-1 >= 0 {
		if heatmap[i][j-1] < 9 && !checkedmap[i][j-1] {
			sum += startTraversal(i, j-1, heatmap, checkedmap)
		}
	}
	// Checking right
	if j+1 < sizej {
		if heatmap[i][j+1] < 9 && !checkedmap[i][j+1] {
			sum += startTraversal(i, j+1, heatmap, checkedmap)
		}
	}
	// Checking down
	if i+1 < sizei {
		if heatmap[i+1][j] < 9 && !checkedmap[i+1][j] {
			sum += startTraversal(i+1, j, heatmap, checkedmap)
		}
	}

	return sum
}
