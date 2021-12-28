package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func ParseNavigationSyntaxFile() {
	DisplayOceanVents()
	fmt.Println("We need to parse through the file to find the issues! Do you want to check for errors?(1) or include incomplete pairs?(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day10InputExample.txt")
	switch text {
	case "1":
		processSyntaxFile(inputData)
		//crabCalc(inputData)
	case "2":
		processSyntaxFileAutocomplete(inputData)
	}
}

func processSyntaxFile(inputData string) {
	lines := strings.Split(inputData, "\n")
	var openBraces Stack
	totalScore := 0
	for _, line := range lines {
		foundError := false
		for x := 0; x < len(line) && !foundError; x++ {
			if string(line[x]) == "(" || string(line[x]) == "[" || string(line[x]) == "{" || string(line[x]) == "<" {
				openBraces.Push(string(line[x]))
			} else {
				switch string(line[x]) {
				case ")":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "(" {
						openBraces.Pop()
					} else {
						totalScore += 3
						foundError = true
					}
				case "]":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "[" {
						openBraces.Pop()
					} else {
						totalScore += 57
						foundError = true
					}
				case "}":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "{" {
						openBraces.Pop()
					} else {
						totalScore += 1197
						foundError = true
					}
				case ">":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "<" {
						openBraces.Pop()
					} else {
						totalScore += 25137
						foundError = true
					}
				}
			}
		}
	}
	fmt.Printf("The total value of all errors is %d\n", totalScore)
}

func processSyntaxFileAutocomplete(inputData string) {
	lines := strings.Split(inputData, "\n")
	var openBraces Stack
	allScores := make([]big.Int, 0)
	for _, line := range lines {
		foundError := false
		for x := 0; x < len(line) && !foundError; x++ {
			if string(line[x]) == "(" || string(line[x]) == "[" || string(line[x]) == "{" || string(line[x]) == "<" {
				openBraces.Push(string(line[x]))
			} else {
				switch string(line[x]) {
				case ")":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "(" {
						openBraces.Pop()
					} else {
						foundError = true
					}
				case "]":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "[" {
						openBraces.Pop()
					} else {
						foundError = true
					}
				case "}":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "{" {
						openBraces.Pop()
					} else {
						foundError = true
					}
				case ">":
					expectedValue, _ := openBraces.Peek()
					if expectedValue == "<" {
						openBraces.Pop()
					} else {
						foundError = true
					}
				}
			}
		}
		if !foundError {
			score := big.NewInt(0)
			for !openBraces.IsEmpty() {
				value, _ := openBraces.Pop()

				score.Mul(score, big.NewInt(5))
				switch value {
				case "(":
					score.Add(score, big.NewInt(1))
				case "[":
					score.Add(score, big.NewInt(2))
				case "{":
					score.Add(score, big.NewInt(3))
				case "<":
					score.Add(score, big.NewInt(4))
				}
			}
			allScores = append(allScores, *score)
		}
	}
	fmt.Printf("The list of line values is %v\n", allScores)
	orderedScores := make([]big.Int, 0)
	for _, score := range allScores {
		orderedScores = addElementInOrder(orderedScores, score)
	}
	fmt.Printf("The list in order is %v\n", orderedScores)
	halfPosition := (len(orderedScores) / 2) + 1
	fmt.Printf("The halfway point of %d is %d and the value there is %v\n", len(orderedScores), halfPosition, orderedScores[halfPosition])

}

/*): 3 points.
]: 57 points.
}: 1197 points.
>: 25137 points.*/

func addElementInOrder(list []big.Int, newElement big.Int) []big.Int {
	added := false
	for x := 0; x < len(list) && !added; x++ {
		if newElement.Cmp(&(list)[x]) > 0 {
			fmt.Printf("The list is %v and we are adding at position %d\n", list, x)
			tempList := make([]big.Int, 0)
			endOfElements := list[x:]
			startOfElements := list[:x]
			fmt.Printf("The bit we took off the end is %v and the bit from the start is %v\n", endOfElements, startOfElements)
			tempList = append(startOfElements, newElement)
			fmt.Printf("we have added the following elements to the temp list %v\nand the list looks like %v\n", tempList, list)
			tempList = append(tempList, endOfElements...)
			added = true
			list = tempList
		}
	}
	if !added {
		list = append(list, newElement)
	}
	fmt.Printf("Here's the list right now %v\n", list)
	return list
}
