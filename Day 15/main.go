// Advent of Code -- Day 15
package main

import (
	"fmt"
	"math"
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
	data, _ := os.ReadFile("input.in")

	input := strings.Fields(string(data))

	riskMap := [][]node{}
	newId := 0
	for _, line := range input {
		newline := []node{}
		for _, char := range line {
			newInt, _ := strconv.Atoi(string(char))
			newNode := node{
				id:        newId,
				val:       newInt,
				neighbors: nil,
			}
			newId += 1
			newline = append(newline, newNode)
		}
		riskMap = append(riskMap, newline)
	}
	// Updating the map with the neighbors
	riskMap = getNeighbors(riskMap)

	totalRisk := shortestPath(riskMap)
	fmt.Println("Total Risk:", totalRisk)

	// Part 2: Expanding the map
	riskMap = expandMap(riskMap)

	bigRisk := shortestPath(riskMap)
	fmt.Println("Jumbo Risk:", bigRisk)

}

func expandMap(riskMap [][]node) [][]node {
	sizeX := len(riskMap[0])
	sizeY := len(riskMap)

	// Splitting the variations into sectors
	sectors := [][][]node{}
	sectors = append(sectors, riskMap)

	// First, adjusting the values of each sector
	for k := 1; k < 9; k++ {
		newMap := [][]node{}
		for _, line := range riskMap {
			newline := []node{}
			for _, this := range line {
				newVal := this.val + k
				if newVal > 9 {
					newVal -= 9
				}
				newNode := node{
					id:        this.id,
					val:       newVal,
					neighbors: nil,
				}
				newline = append(newline, newNode)
			}
			newMap = append(newMap, newline)
		}
		sectors = append(sectors, newMap)
	}

	// Creating the big boah
	jumboMap := [][]node{}
	newId := 0
	for i:=0; i < sizeY * 5; i++ {
		jumboline := []node{}
		for j:=0; j < sizeX * 5; j++ {
			sect := (i/sizeY) + (j/sizeX)

			newVal := sectors[sect][i%sizeY][j%sizeX].val
			newNode := node {
				id: newId,
				val: newVal,
				neighbors: nil,
			}
			jumboline = append(jumboline, newNode)
			newId += 1
		}
		jumboMap = append(jumboMap, jumboline)
	}
	jumboMap = getNeighbors(jumboMap)


	return jumboMap
}

func shortestPath(riskMap [][]node) int {
	sizeX := len(riskMap[0])
	sizeY := len(riskMap)
	totalSize := sizeX * sizeY
	distance := []int{}
	parents := []int{}
	visited := []bool{}

	for i := 0; i < totalSize; i++ {
		distance = append(distance, int(math.Pow(2, 32)-1))
		parents = append(parents, i)
		visited = append(visited, false)
	}
	distance[0] = 1

	startTraversal(riskMap, 0, distance, parents, visited)

	t := (sizeX * sizeY) - 1

	return distance[t] - distance[0]
}

func startTraversal(riskMap [][]node, id int, distance []int, parents []int, visited []bool) {
	thislocation := getLocation(riskMap, id)
	thisX := thislocation[1]
	thisY := thislocation[0]
	visited[id] = true

	// Going through neighbors and updating distances
	neighbors := riskMap[thisY][thisX].neighbors
	for neighId, loc := range neighbors {
		neighVal := riskMap[loc[0]][loc[1]].val
		newDist := distance[id] + neighVal
		if newDist < distance[neighId] && !visited[neighId] {
			distance[neighId] = newDist
			parents[neighId] = id
		}
	}

	// Traversing through the minimum distanced node
	minDist := getMinDist(distance, visited)
	if minDist != -1 {
		startTraversal(riskMap, minDist, distance, parents, visited)
	}

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
				neighbors[this.id-sizeX] = []int{i - 1, j}
			}
			if j-1 >= 0 {
				neighbors[this.id-1] = []int{i, j - 1}
			}
			if i+1 < sizeY {
				neighbors[this.id+sizeX] = []int{i + 1, j}
			}
			if j+1 < sizeX {
				neighbors[this.id+1] = []int{i, j + 1}
			}
			newNode := node{
				id:        this.id,
				val:       this.val,
				neighbors: neighbors,
			}
			newline = append(newline, newNode)
		}
		newMap = append(newMap, newline)
	}
	return newMap
}

func getMinDist(distance []int, visited []bool) int {
	minId := -1
	minVal := int(math.Pow(2, 32) - 1)

	for id, val := range distance {
		if val < minVal && !visited[id] {
			minId = id
			minVal = val
		}
	}
	return minId
}

func getLocation(amap [][]node, id int) []int {
	for i, line := range amap {
		for j, node := range line {
			if node.id == id {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
