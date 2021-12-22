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

}

// First attempt lol
func startTraversal(i int, j int, heatmap *[][]int, checkedmap *[][]bool) int {
	lowVal := (*heatmap)[i][j]
	potentialLow := []int{-1, -1}

	sizei := len((*heatmap))
	sizej := len((*heatmap)[0])
	// Checking up
	if i-1 >= 0 {
		if (*heatmap)[i-1][j] < lowVal {
			lowVal = (*heatmap)[i-1][j]
			potentialLow[0] = i - 1
			potentialLow[1] = j
		}
	}
	// Checking left
	if j-1 >= 0 {
		if (*heatmap)[i][j-1] < lowVal {
			lowVal = (*heatmap)[i][j-1]
			potentialLow[0] = i
			potentialLow[1] = j - 1
		}
	}
	// Checking right
	if j+1 < sizej {
		if (*heatmap)[i][j+1] < lowVal {
			lowVal = (*heatmap)[i][j+1]
			potentialLow[0] = i
			potentialLow[1] = j + 1
		}
	}
	// Checking down
	if i+1 < sizei {
		if (*heatmap)[i+1][j] < lowVal {
			lowVal = (*heatmap)[i+1][j]
			potentialLow[0] = i + 1
			potentialLow[1] = j
		}
	}

	// Marking our checks
	lowi := potentialLow[0]
	lowj := potentialLow[1]

	// Nothing was found; this is the lowest; return this val
	if lowi < 0 || lowj < 0 {
		markAll(i, j, checkedmap)
		return lowVal
	}

	// The low val location was already marked/checked
	if (*checkedmap)[lowi][lowj] {
		markAll(i, j, checkedmap)
		return -1
	}

	// fmt.Printf("%d ", lowVal)
	markAll(i, j, checkedmap)
	return startTraversal(lowi, lowj, heatmap, checkedmap)
}

func markAll(i int, j int, checkedmap *[][]bool) {
	sizei := len((*checkedmap))
	sizej := len((*checkedmap)[0])
	(*checkedmap)[i][j] = true
	if i-1 >= 0 {
		(*checkedmap)[i-1][j] = true
	}
	if j-1 >= 0 {
		(*checkedmap)[i][j-1] = true
	}
	if j+1 < sizej {
		(*checkedmap)[i][j+1] = true
	}
	if i+1 < sizei {
		(*checkedmap)[i+1][j] = true
	}
}
