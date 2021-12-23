package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type cave struct {
	Neighbors []string `default:{}`
	Small     bool
	Used      bool
}

func main() {
	datalist, _ := os.ReadFile("example.in")

	input := strings.Fields(string(datalist))

	System := make(map[string]cave)
	for _, path := range input {
		startend := strings.Split(path, "-")
		// Checking if it exists
		if _, exists := System[startend[0]]; !exists {
			System[startend[0]] = cave{
				Small: unicode.IsLower(rune(startend[0][0])),
				Used:  startend[0] == "start",
			}
		}
		if _, exists := System[startend[1]]; !exists {
			System[startend[1]] = cave{
				Small: unicode.IsLower(rune(startend[1][0])),
				Used:  startend[1] == "start",
			}
		}
		// Adding to System
		newNeighbors := System[startend[0]].Neighbors
		newNeighbors = append(newNeighbors, startend[1])
		System[startend[0]] = cave{
			Neighbors: newNeighbors,
			Small:     unicode.IsLower(rune(startend[0][0])),
			Used:      startend[0] == "start",
		}
		newNeighbors = System[startend[1]].Neighbors
		newNeighbors = append(newNeighbors, startend[0])
		System[startend[1]] = cave{
			Neighbors: newNeighbors,
			Small:     unicode.IsLower(rune(startend[1][0])),
			Used:      startend[1] == "start",
		}
	}

	printSystem(System)


}

func findPaths(System map[string]cave, cur string, paths [][]string) {

}

func printSystem(System map[string]cave) {
	for key, val := range System {
			fmt.Printf("%s: %v\n", key, val)
	}
}
