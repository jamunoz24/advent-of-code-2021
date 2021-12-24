package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type cave struct {
	Neighbors []string //`default:{}`
	Small     bool
	Used      bool
}

func main() {
	datalist, _ := os.ReadFile("input.in")

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
	travSys := copyMap(System)
	paths := 0
	curPath := []string{"start"}
	passedTwice := false

	findPaths(travSys, "start", &paths, curPath, passedTwice)

	fmt.Printf("Paths found: %d\n", paths)
}

func findPaths(System map[string]cave, cur string, paths *int, curPath []string, passedTwice bool) {
	if cur == "end" {
		(*paths) += 1
		return
	}
	// Make Used if its a small cave
	if System[cur].Small {
		System[cur] = makeUsed(System[cur])
	}

	// Traversing through unused caves
	for _, node := range System[cur].Neighbors {
		if node != "start" &&
		(!System[node].Used || (System[node].Used && !passedTwice)) {
			nextPath := []string{}
			nextPath = append(nextPath, curPath...)
			nextPath = append(nextPath, node)
			travSys := copyMap(System)
			// if node == "end" {
			//	  fmt.Println(nextPath)
			// }
			if System[node].Used && !passedTwice {
			findPaths(travSys, node, paths, nextPath, true)
			} else {
				findPaths(travSys, node, paths, nextPath, passedTwice)
			}

		}
	}

}

func makeUsed(oldCave cave) cave {
	return cave{
		Neighbors: oldCave.Neighbors,
		Small:     oldCave.Small,
		Used:      true,
	}
}

func copyMap(amap map[string]cave) map[string]cave {
	newMap := make(map[string]cave)
	for key, val := range amap {
		newMap[key] = val
	}
	return newMap
}

func printSystem(System map[string]cave) {
	for key, val := range System {
		fmt.Printf("%s: %v\n", key, val)
	}
}
