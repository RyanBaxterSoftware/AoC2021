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
		//crabCalcFuelRamp(inputData)
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
		characterOptions := []rune{ 'a', 'b', 'c', 'd', 'e', 'f', 'g'}
		for _, character := range characterOptions {
			possibilities[character] = []int{ 1,2,3,4,5,6,7 }
		}
		for _, element := range allElements {
			if len(element) == 2 {
			// 1, two right lights
			} else if len(element) == 3 {
			// 7, two right and top
			} else if len(element) == 4 {
			// 4 top left, middle, and two right
			} else if len(element) == 7 {
			// 8, all lights
				fmt.Printf("%v, ", element)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("We have found %d instances of 1, 4, 7, and 8", valuesWeCareAbout)
}