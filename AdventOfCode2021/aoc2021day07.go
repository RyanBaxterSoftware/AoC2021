package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CrabSubMovementCalculation() {
	DisplayCrabSubs()
	fmt.Println("The crabs have come to your aid! Help them find out where to line up for your escape route! Are you using standard fuel modeling(1) or variable(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day7Input.txt")
	switch text {
	case "1":
		crabCalc(inputData)
	case "2":
		crabCalcFuelRamp(inputData)
	}
}

func DisplayCrabSubs() {
	fmt.Println("	              _")
	fmt.Println("	.         _____|___")
	fmt.Println("   .          ___/  o o o  \\___")
	fmt.Println("   .         /     _      _    \\")
	fmt.Println("	.   |     (<      >)    |")
	fmt.Println("	.   |      `O,99,O`     |")
	fmt.Println("    .       |     //-\\__/-\\\\    |")
	fmt.Println("	  8-=\\_________________/")
}

func crabCalc(inputData string) {
	crabsString := strings.Split(inputData, ",")
	var crabs []int
	min := 40000
	max := 0
	smallestDiff := math.MaxInt32
	position := -1
	for _, crab := range crabsString {
		crabNum, _ := strconv.Atoi(crab)
		crabs = append(crabs, crabNum)
		if crabNum < min {
			min = crabNum
		}
		if crabNum > max {
			max = crabNum
		}
	}
	for x := min; x <= max; x++ {	
		total := 0
		for _, crab := range crabs {
			total += getDifference(crab, x)
		}
		if smallestDiff > total {
			smallestDiff = total
			position = x
		}
	}
	
	fmt.Printf("The final position is %d and the total fuel consumption is %d\n", position, smallestDiff)
}


func crabCalcFuelRamp(inputData string) {
	crabsString := strings.Split(inputData, ",")
	var crabs []int
	min := 40000
	max := 0
	smallestDiff := math.MaxInt32
	position := -1
	for _, crab := range crabsString {
		crabNum, _ := strconv.Atoi(crab)
		crabs = append(crabs, crabNum)
		if crabNum < min {
			min = crabNum
		}
		if crabNum > max {
			max = crabNum
		}
	}
	for x := min; x <= max; x++ {	
		total := 0
		for _, crab := range crabs {
			total += getModifiedDifference(crab, x)
		}
		if smallestDiff > total {
			smallestDiff = total
			position = x
		}
	}
	
	fmt.Printf("The final position is %d and the total fuel consumption is %d\n", position, smallestDiff)
}

func getDifference(first int, second int) int {
	numberRaw := first - second
	if numberRaw < 0 {
		numberRaw = 0 - numberRaw
	}
	return numberRaw
}

func getModifiedDifference(first int, second int) int {
	newNumber := 0
	difference := getDifference(first, second)
	for x := 1; x <= difference; x++ {
		newNumber += x
	}
	return newNumber
}