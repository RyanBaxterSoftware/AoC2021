package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NumpadDisplayMixupFixup() {
	DisplayCrabSubs()
	fmt.Println("Oh no! Our display pad! It's broken! Do you want to check how many numbers exist?(1) or(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day8Input.txt")
	switch text {
	case "1":
		countNumOutputs(inputData)
		//crabCalc(inputData)
	case "2":
		fixNumOutputs(inputData)
	}
}

func DisplayNumOuput() {
	fmt.Println("	              _")
	fmt.Println("	.         _____|___")
	fmt.Println("   .          ___/  o o o  \\___")
	fmt.Println("   .         /     _      _    \\")
	fmt.Println("	.   |     (<      >)    |")
	fmt.Println("	.   |      `O,99,O`     |")
	fmt.Println("    .       |     //-\\__/-\\\\    |")
	fmt.Println("	  8-=\\_________________/")
}

func countNumOutputs(inputData string) {
	valuesWeCareAbout := 0
	lines := strings.Split(inputData, "\n")
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
		halves := strings.Split(line, " | ")
		outputelements := strings.Split(halves[1], " ")
		for _, element := range outputelements {
			if len(element) == 2 || len(element) == 3 || len(element) == 4 || len(element) == 7 {
				valuesWeCareAbout++
				fmt.Printf("%v, ", element)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("We have found %d instances of 1, 4, 7, and 8", valuesWeCareAbout)
}

func fixNumOutputs(inputData string) {
	valuesWeCareAbout := 0
	lines := strings.Split(inputData, "\n")
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
		//halves := strings.Split(line, " | ")
		allElements := strings.Split(line, " ")
		//outputelements := strings.Split(halves[1], " ")
		possibilities := make(map[rune][]int)
		characterOptions := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
		for _, character := range characterOptions {
			possibilities[character] = []int{1, 2, 3, 4, 5, 6, 7}
		}
		/*
			 111111
			2      3
			2      3
			2      3
			 444444
			5      6
			5      6
			5      6
			 777777

		*/
		for _, element := range allElements {

			if len(element) == 2 {
				for _, input := range element {
					onePossibilities := make([]int, 0)
					onePossibilities = append(onePossibilities, 3)
					onePossibilities = append(onePossibilities, 6)
					possibilities = setOptions(input, onePossibilities, possibilities)
				}
				// 1, two right lights
			} else if len(element) == 3 {
				for _, input := range element {
					onePossibilities := make([]int, 0)
					onePossibilities = append(onePossibilities, 1)
					onePossibilities = append(onePossibilities, 3)
					onePossibilities = append(onePossibilities, 6)
					possibilities = setOptions(input, onePossibilities, possibilities)
				}
				// 7, two right and top
			} else if len(element) == 4 {
				for _, input := range element {
					onePossibilities := make([]int, 0)
					onePossibilities = append(onePossibilities, 2)
					onePossibilities = append(onePossibilities, 3)
					onePossibilities = append(onePossibilities, 4)
					onePossibilities = append(onePossibilities, 6)
					possibilities = setOptions(input, onePossibilities, possibilities)
				}
				// 4 top left, middle, and two right
			} else if len(element) == 7 {
				// 8, all lights
			}
		}
		fmt.Printf("Here are the possibilities: %v \n", possibilities)
	}
	fmt.Printf("We have found %d instances of 1, 4, 7, and 8", valuesWeCareAbout)
}

func setOptions(character rune, options []int, possibilities map[rune][]int) map[rune][]int {
	characterPossiblities := possibilities[character]
	newPossibilities := make([]int, 0)
	for _, possibility := range characterPossiblities {
		if containsInt(options, possibility) {
			newPossibilities = append(newPossibilities, possibility)
		}
	}
	possibilities[character] = newPossibilities
	anElementWasChanged := true
	for anElementWasChanged {
		anElementWasChanged = false
		for character, options := range possibilities {
			for checkcharacter, checkoptions := range possibilities {
				if checkcharacter != character && compareIntArray(options, checkoptions) && len(options) == 2 {

				}
			}
		}
	}

	return possibilities
}

func containsInt(array []int, element int) bool {
	exists := false
	for _, arrayelement := range array {
		if element == arrayelement {
			exists = true
		}
	}
	return exists
}

func compareIntArray(array1 []int, array2 []int) bool {
	equal := true
	if len(array1) != len(array2) {
		equal = false
	}
	for _, element := range array1 {
		if !containsInt(array2, element) {
			equal = false
		}
	}
	return equal
}
