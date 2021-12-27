// Advent of Code -- Day 15
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id        int
	val       int
	neighbors map[int][]int
}

func main() {
	data, _ := os.ReadFile("example.in")

	input := strings.Fields(string(data))

	riskMap := [][]node{}
	newId := 0
	for _, line := range input {
		newline := []node{}
		for _, char := range line {
			newInt, _ := strconv.Atoi(string(char))
			newNode := node{
				id: newId,
				val: newInt,
				neighbors: nil,
			}
			newId += 1
			newline = append(newline, newNode)
		}
		riskMap = append(riskMap, newline)
	}
	// Updating the map with the neighbors
	riskMap = getNeighbors(riskMap)

	shortestPath := shortestPath(riskMap)


	fmt.Println("Total Risk:", shortestPath)
}

func shortestPath(riskMap [][]node) int {



	return 0
}

func getNeighbors(amap [][]node) [][]node {
	newMap := [][]node{}
	sizeX := len(amap[0])
	sizeY := len(amap)
	
	for i, line := range amap {
		newline := []node{}
		for j, this := range line {
			neighbors := make(map[int][]int)
			if i-1 >= 0 {
				neighbors[this.id-sizeX] = []int{i-1,j}
			}
			if j-1 >= 0 {
				neighbors[this.id-1] = []int{i,j-1}
			}
			if i+1 < sizeY {
				neighbors[this.id+sizeX] = []int{i+1,j}
			}
			if j+1 < sizeX {
				neighbors[this.id+1] = []int{i,j+1}
			}
			newNode := node{
				id: this.id,
				val: this.val,
				neighbors: neighbors,
			}
			newline = append(newline, newNode)
		}
		newMap = append(newMap, newline)
	}
	return newMap
}

func getLocation(amap[][]node, id int) []int {
	for i, line := range amap {
		for j, node := range line {
			if node.id == id {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func printMap(amap [][]node) {
	for _, line := range amap {
		for _, node := range line {
			fmt.Print(node.val)
		}
		fmt.Println()
	}
}
