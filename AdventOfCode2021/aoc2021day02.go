package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func CalculateMovement() {
	displaySub()
	fmt.Println("Our submarine is navigating. Would you like to track this navigation through traditional definitions (1) or through aim guidance(2)?")
	fmt.Printf("Here's a test of your program: %d\n", absVal(-1))
	
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day2Input.txt")
	switch text {
		case "1":
			countDistance(inputData)
		case "2":
			countDistanceByAim(inputData)
	}
}

func countDistance(inputData string) {
	actions := strings.Split(inputData, "\n")
	height := 0
	horizontal := 0
	for _, action := range actions {
		actionSplit := strings.Split(action, " ")
		numberValue, _ := strconv.Atoi(actionSplit[1])
		switch actionSplit[0] {
			case "forward":
				horizontal += numberValue
			case "up":
				height -= numberValue
			case "down":
				height += numberValue
		}
		fmt.Printf("Current position:\nHeight: %d Horizontal: %d\n", height, horizontal)
	}
	fmt.Printf("Our combined value is %d\n", (absVal(height * horizontal)))
}

func countDistanceByAim(inputData string) {
	actions := strings.Split(inputData, "\n")
	height := 0
	horizontal := 0
	aim := 0
	for _, action := range actions {
		actionSplit := strings.Split(action, " ")
		numberValue, _ := strconv.Atoi(actionSplit[1])
		switch actionSplit[0] {
			case "forward":
				horizontal += numberValue
				height += aim*numberValue
			case "up":
				aim -= numberValue
			case "down":
				aim += numberValue
		}
		fmt.Printf("Current position:\nHeight: %d Horizontal: %d\n", height, horizontal)
	}
	fmt.Printf("Our combined value is %d\n", (absVal(height * horizontal)))
}

func absVal(value int) int {
	if(value < 0) {
		return 0 - value
	} else {
		return value
	}
}