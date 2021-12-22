package main

import (
	"fmt"
	"os"
	"sort"
)

// Creating our own Struct
type Stack []byte

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(b byte) {
	*s = append((*s), b)
}

func (s *Stack) pop() (byte, bool) {
	if s.isEmpty() {
		fmt.Println("Pop error")
		return byte(0), false
	}
	last := len(*s) - 1
	item := (*s)[last]
	*s = (*s)[:last]
	return item, true
}

func main() {
	datalist, _ := os.ReadFile("input.in")

	chungus := [][]byte{}
	chunk := []byte{}
	for _, data := range datalist {
		if string(data) == "\n" {
			chungus = append(chungus, chunk)
			chunk = nil
		} else {
			chunk = append(chunk, data)
		}
	}

	// Part 1: Poppin'
	falseBits := []byte{}
	for _, chunk := range chungus {
		var stack Stack
		for _, bit := range chunk {
			if isOpen(bit) {
				stack.push(bit)
			} else {
				open, popped := stack.pop()
				if popped && !isCompatible(open, bit) {
					falseBits = append(falseBits, bit)
				}
			}
		}
	}

	// Counting points
	totalPoints := 0
	for _, bit := range falseBits {
		if bit == ')' {
			totalPoints += 3
		} else if bit == ']' {
			totalPoints += 57
		} else if bit == '}' {
			totalPoints += 1197
		} else if bit == '>' {
			totalPoints += 25137
		}
	}

	// Part 2:
	pointSlice := []uint{}
	for _, chunk := range chungus {
		var stack Stack
		createdBits := []byte{}
		contin := true
		for _, bit := range chunk {
			if isOpen(bit) {
				stack.push(bit)
				continue
			}
			open, _ := stack.pop()
			if !isCompatible(open, bit) {
				contin = false
				break
			}
		}
		if contin {
			// Emptying the stack
			for !stack.isEmpty() {
				open, _ := stack.pop()
				if open == '[' {
					createdBits = append(createdBits, ']')
				} else if open == '(' {
					createdBits = append(createdBits, ')')
				} else if open == '{' {
					createdBits = append(createdBits, '}')
				} else if open == '<' {
					createdBits = append(createdBits, '>')
				}
				// Counting the points
			}
			var createdPoints uint = 0
			for _, bit := range createdBits {
				createdPoints *= 5
				if bit == ')' {
					createdPoints += 1
				} else if bit == ']' {
					createdPoints += 2
				} else if bit == '}' {
					createdPoints += 3
				} else if bit == '>' {
					createdPoints += 4
				}
			}
			pointSlice = append(pointSlice, createdPoints)
		}
	}
	// Sorting the slice
	sort.Slice(pointSlice, func(i, j int) bool { return pointSlice[i] < pointSlice[j] })


	fmt.Printf("False Bits: ")
	for _, bit := range falseBits {
		fmt.Printf("%s ", string(bit))
	}
	fmt.Println()

	fmt.Printf("Total Points: %d\n", totalPoints)

	fmt.Printf("Collected Points: %v\n", pointSlice)
	fmt.Printf("Middle Slice: %d\n", pointSlice[len(pointSlice)/2])
}

func isOpen(b byte) bool {
	return b == 91 || b == 40 || b == 123 || b == 60
}

func isCompatible(b byte, c byte) bool {
	return (b == '[' && c == ']') || (b == '(' && c == ')') ||
		(b == '{' && c == '}') || (b == '<' && c == '>')
}
