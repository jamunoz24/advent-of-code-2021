// Advent of Code - Day 7
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	datalist, _ := os.ReadFile("input.in")

	// First parsing the string to get rid of commas
	newList := strings.Split(string(datalist), ",")

	// Converting that string array to int
	var numList = []int{}
	for _, i := range newList {
		num, _ := strconv.Atoi(i)
		numList = append(numList, num)
	}

	var min int = int(math.Pow(2, 32) - 1)

	for i := range numList {
		var curMin int = 0
		for _, num := range numList {
			fuelCost := num - i
			if fuelCost < 0 {
				fuelCost = -fuelCost
			}
			curMin += fuelCost + fib(fuelCost)
		}
		if min > curMin {
			min = curMin
		}
	}

	fmt.Println(min)

}

// This isn't really fibbonacci
func fib(num int) int {
	var sum int = 0
	for i := 1; i < num; i++ {
		sum += i
	}
	return sum
}
