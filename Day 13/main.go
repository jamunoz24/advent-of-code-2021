package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type fold struct {
	Axis  string
	Value int
}

func main() {
	datalist, _ := os.ReadFile("example.in")

	input := strings.Fields(string(datalist))

	coords := []coord{}
	folds := []fold{}
	sizeX := 0
	sizeY := 0
	for _, line := range input {
		if line != "fold" && line != "along" {
			if line[0] != 'x' && line[0] != 'y' {
				newline := strings.Split(line, ",")
				newx, _ := strconv.Atoi(newline[0])
				newy, _ := strconv.Atoi(newline[1])
				newCoord := coord{
					x: newx,
					y: newy,
				}
				coords = append(coords, newCoord)
				if newx > sizeX {
					sizeX = newx
				}
				if newy > sizeY {
					sizeY = newy
				}
			} else {
				newline := strings.Split(line, "=")
				newCoord, _ := strconv.Atoi(newline[1])
				newFold := fold {
					Axis: newline[0],
					Value: newCoord,
				}
				folds = append(folds, newFold)
			}
		}
	}

	// Creating the matrix
	sizeX += 1
	sizeY += 1
	matrix := [][]string{}
	for i:=0; i < sizeY; i++ {
		newline := []string{}
		for j:=0; j < sizeX; j++ {
			newline = append(newline, ".")
		}
		matrix = append(matrix, newline)
	}
	// Putting in the stamps
	for _, coord := range coords {
		matrix[coord.y][coord.x] = "#"
	}
	
	// Applying the folds
	for _, fold := range folds {
		sizeY = len(matrix)
		sizeX = len(matrix[0])
		// Drawing the folding line
		cut := fold.Value
		if fold.Axis == "y" {
			for i:=0; i < sizeX; i++ {
				matrix[cut][i] = "-"
			}
			// Splitting the Matrix
			firstHalf := matrix[:cut]
			secondHalf := matrix[cut+1:]
			// Inversing the second matrix up
			sizeY -= cut+1
			for j:=0; j < sizeX; j++ {
				k := sizeY-1
				for i:=0; i < sizeY/2; i++ {
					secondHalf[i][j], secondHalf[k][j] = secondHalf[k][j], secondHalf[i][j]
					k -= 1
				}
			}
			// Intersecting the Matrices
			for i:=0; i < sizeY; i++ {
				for j:=0; j < sizeX; j++ {
					if secondHalf[i][j] == "#" {
						firstHalf[i][j] = "#"
					}
				}
			}
			matrix = firstHalf
		} else {
			for i:=0; i < sizeY; i++ {
				matrix[i][cut] = "|"
			}
			// Splitting the Matrix
			firstHalf := [][]string{}
			secondHalf := [][]string{}
			for i:=0; i < sizeY; i++ {
				newline := []string{}
				for j:=0; j < cut; j++ {
					newline = append(newline, matrix[i][j])
				}
				firstHalf = append(firstHalf, newline)
			}
			for i:=0; i < sizeY; i++ {
				newline := []string{}
				for j:=cut+1; j < sizeX; j++ {
					newline = append(newline, matrix[i][j])
				}
				secondHalf = append(secondHalf, newline)
			}
			// Inversing the second matrix left
			printMatrix(firstHalf)
			printMatrix(secondHalf)
			sizeX -= cut+1
			for i:=0; i < sizeY; i++ {
				k := sizeX-1
				for j:=0; j < sizeX/2; j++ {
					secondHalf[i][j], secondHalf[i][k] = secondHalf[i][k], secondHalf[i][j]
					k -= 1
				}
			}
			printMatrix(secondHalf)
			// Intersecting the Matrices
			for i:=0; i < sizeY; i++ {
				for j:=0; j < sizeX; j++ {
					if secondHalf[i][j] == "#" {
						firstHalf[i][j] = "#"
					}
				}
			}
			matrix = firstHalf
		}
	}

	// Counting the dots
	dots := 0
	for _, line := range matrix {
		for _, char := range line {
			if char == "#" {
				dots += 1
			}
		}
	}

	printMatrix(matrix)
	fmt.Printf("Dots: %d\n", dots)
	fmt.Printf("Folds: %v\n", folds)

}

func printMatrix(matrix [][]string) {
	fmt.Println("Matrix:")
	for _, line := range matrix {
		for _, ind := range line {
			if ind == "." {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}