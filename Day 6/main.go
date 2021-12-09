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
	var days int = 80
	var fishNum int = len(fishList)
	for i := 0; i < days; i++ {
		for fish, daysLeft := range fishList {
			daysLeft--
			// New fishy
			if daysLeft < 0 {
				daysLeft = 6
				fishList = append(fishList, 8)
				fishNum += 1
			}
			fishList[fish] = daysLeft
		}
	}


	fmt.Println(fishNum)

}
