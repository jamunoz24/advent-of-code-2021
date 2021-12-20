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
	firstList := [][]string{}
	secondList := [][]string{}
	tempL1 := []string{}
	for i := 0; i < len(newList); i++ {
		if  newList[i] == "|" {
			i += 1
			tempL := []string{}
			for j := i; j < i+4; j++ {
				tempL = append(tempL, newList[j])
			}
			firstList = append(firstList, tempL1)
			tempL1 = nil
			secondList = append(secondList, tempL)
			i += 3
		} else {
			tempL1 = append(tempL1, newList[i])
		}
	}

	// Part 1: First filtering 1, 4, 7, 8
	var digitSum uint = 0
	for _, line := range secondList {
		for _, str := range line {
			if checkOneFourSevenEight(str) {
				digitSum += 1
			}
		}
	}

	// Part 2
	var str1 string = "ABC"
	for i := 0; i < len(str1); i++ {
		fmt.Println(string(str1[i]))
	}

}

func checkOneFourSevenEight(str string) bool {
	len := len(str)
	return len == 2 || len == 4 || len == 3 || len == 7
}

func Split(char rune) bool {
	return char == ' ' || char == '\n'
}


