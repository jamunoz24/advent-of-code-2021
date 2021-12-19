// Advent of Code -- Day 6: Go edition
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	datalist, _ := os.ReadFile("input.in")

	// First parsing the string to get rid of commas
	newList := strings.Split(string(datalist), ",")

	// Converting that string array to int
	var fishList = []int{}
	for _, i := range newList {
		num, _ := strconv.Atoi(i)
		fishList = append(fishList, num)
	}

	// Find the number of lanternfish after 80 days
	var days int = 256

	fishCounts := make([]int, 9)

	for _, num := range fishList {
		fishCounts[num] += 1
	}

	for day := 0; day < days; day++ {
		newFishies := fishCounts[0]
		for i := 0; i < len(fishCounts)-1; i++ {
			fishCounts[i] = fishCounts[i+1]
		}
		fishCounts[6] += newFishies
		fishCounts[8] = newFishies
	}

	// Getting sum
	var sum uint = 0
	for i := range fishCounts {
		sum += uint(fishCounts[i])
	}
	fmt.Println(sum)

}
