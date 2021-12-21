// Advent of Code - Day 8
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	datalist, _ := os.ReadFile("input.in")

	// First parsing the string to get rid of commas
	// newList := strings.Split(string(datalist), " ")
	newList := strings.FieldsFunc(string(datalist), Split)

	// Splitting the data into two
	clueList := [][]string{}
	outputList := [][]string{}
	tempL1 := []string{}
	for i := 0; i < len(newList); i++ {
		if newList[i] == "|" {
			i += 1
			tempL := []string{}
			for j := i; j < i+4; j++ {
				tempL = append(tempL, newList[j])
			}
			clueList = append(clueList, tempL1)
			tempL1 = nil
			outputList = append(outputList, tempL)
			i += 3
		} else {
			tempL1 = append(tempL1, newList[i])
		}
	}

	// Part 1: First filtering 1, 4, 7, 8
	var digitSum uint = 0
	for _, line := range outputList {
		for _, str := range line {
			if checkOneFourSevenEight(str) {
				digitSum += 1
			}
		}
	}

	// Part 2

	totalSum := 0

	for line := range clueList {
		// First we store our known numbers (1,4,7,8)
		var knownNums [10]string
		for _, num := range clueList[line] {
			if len(num) == 2 {
				knownNums[1] = num
			} else if len(num) == 4 {
				knownNums[4] = num
			} else if len(num) == 3 {
				knownNums[7] = num
			} else if len(num) == 7 {
				knownNums[8] = num
			}
		}

		one := knownNums[1]
		four := knownNums[4]
		seven := knownNums[7]
		fourMinusOne := fourMinusOne(four, one)

		for _, clue := range clueList[line] {
			//0,6,9 or 2,3,5
			if len(clue) == 5 {
				if numInNum(seven, clue) {
					knownNums[3] = clue
				} else if numInNum(fourMinusOne, clue) {
					knownNums[5] = clue
				} else {
					knownNums[2] = clue
				}

			} else if len(clue) == 6 {
				if numInNum(four, clue) && numInNum(one, clue) {
					knownNums[9] = clue
				} else if numInNum(one, clue) {
					knownNums[0] = clue
				} else {
					knownNums[6] = clue
				}
			}
		}

		// Getting output values
		curSum := 0
		multiplier := 1000
		for _, output := range outputList[line] {
			for i, num := range knownNums {
				if len(output) == len(num) && numInNum(output, num) {
					curSum += i * multiplier
					multiplier /= 10
					break
				}
			}
		}
		// fmt.Printf("%v | %v\n", clueList[line], outputList[line])
		// fmt.Println(knownNums)
		// fmt.Println(curSum)
		totalSum += curSum
	}

	fmt.Printf("Total Sum: %d\n", totalSum)

}

func checkOneFourSevenEight(str string) bool {
	len := len(str)
	return len == 2 || len == 4 || len == 3 || len == 7
}

func Split(char rune) bool {
	return char == ' ' || char == '\n'
}

func fourMinusOne(four string, one string) string {
	var str string
	for i := 0; i < len(four); i++ {
		if four[i] != one[0] && four[i] != one[1] {
			str += string(four[i])
		}
	}
	return str
}

func numInNum(smallNum string, bigNum string) bool {
	var charChecks []bool
	for i := 0; i < len(smallNum); i++ {
		charChecks = append(charChecks, false)
	}
	for i := 0; i < len(bigNum); i++ {
		for j := 0; j < len(smallNum); j++ {
			// found, check the next number
			if bigNum[i] == smallNum[j] {
				charChecks[j] = true
				break
			}
		}
	}
	for i := 0; i < len(charChecks); i++ {
		if !charChecks[i] {
			return false
		}
	}
	return true
}
