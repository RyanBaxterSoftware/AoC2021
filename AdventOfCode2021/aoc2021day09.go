package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func VolcanicVentAnalysis() {
	DisplayOceanVents()
	fmt.Println("Navigate the volanic vents! Want to check the low points?(1) or(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day9Input.txt")
	switch text {
	case "1":
		findLowestPoints(inputData)
		//crabCalc(inputData)
	case "2":
		findLowestBasins(inputData)
	}
}

func findLowestPoints(inputData string) {
	lines := strings.Split(inputData, "\n")
	collectionOfLowestPoints := make([]int, 0)
	for acrossdepth, line := range lines {
		for downdepth := 0; downdepth < len(line); downdepth++ {
			otherPoints := make([]int, 0)
			currentElement := int(lines[acrossdepth][downdepth] - '0')
			if acrossdepth > 0 {
				var element int
				lineForElement := lines[acrossdepth-1]
				element = int(lineForElement[downdepth] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
			}
			if acrossdepth < len(lines)-1 {
				var element int
				lineForElement := lines[acrossdepth+1]
				element = int(lineForElement[downdepth] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
				//element := lines[acrossdepth+1][downdepth]

			}
			if downdepth < len(line)-1 {
				var element int
				lineForElement := lines[acrossdepth]
				element = int(lineForElement[downdepth+1] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
				//element := lines[acrossdepth][downdepth+1]

			}
			if downdepth > 0 {
				var element int
				lineForElement := lines[acrossdepth]
				element = int(lineForElement[downdepth-1] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
				//element := lines[acrossdepth][downdepth-1]

			}

			lowestPoint := true
			for _, otherpoint := range otherPoints {
				if otherpoint <= currentElement {
					lowestPoint = false
				}
			}
			if lowestPoint {
				fmt.Printf("We found a new lowest point. It is %d at %d and %d\n", currentElement, acrossdepth, downdepth)
				collectionOfLowestPoints = append(collectionOfLowestPoints, currentElement)

			}

			//element := lines[acrossdepth][downdepth]
			//fmt.Printf("This is the element%c\n", element)
			//fmt.Printf("The here's the current line: %v\nhere is it's element %c\n", lines[acrossdepth], lines[acrossdepth][downdepth])
		}

	}

	sumOfHeightValues := 0
	for _, point := range collectionOfLowestPoints {
		sumOfHeightValues += point + 1
	}
	fmt.Printf("The sum of depths is %d \n", sumOfHeightValues)
}

func findLowestBasins(inputData string) {
	lines := strings.Split(inputData, "\n")
	collectionOfLowestPoints := make([]int, 0)
	threeLargestBasins := make([]int, 3)
	threeLargestBasins[0] = 0
	threeLargestBasins[1] = 0
	threeLargestBasins[2] = 0
	for acrossdepth, line := range lines {
		for downdepth := 0; downdepth < len(line); downdepth++ {
			otherPoints := make([]int, 0)
			currentElement := int(lines[acrossdepth][downdepth] - '0')
			if acrossdepth > 0 {
				var element int
				lineForElement := lines[acrossdepth-1]
				element = int(lineForElement[downdepth] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
			}
			if acrossdepth < len(lines)-1 {
				var element int
				lineForElement := lines[acrossdepth+1]
				element = int(lineForElement[downdepth] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
				//element := lines[acrossdepth+1][downdepth]

			}
			if downdepth < len(line)-1 {
				var element int
				lineForElement := lines[acrossdepth]
				element = int(lineForElement[downdepth+1] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
				//element := lines[acrossdepth][downdepth+1]

			}
			if downdepth > 0 {
				var element int
				lineForElement := lines[acrossdepth]
				element = int(lineForElement[downdepth-1] - '0')
				otherPoints = append(otherPoints, element)
				//fmt.Printf("The number we got was %d\n", element)
				//element := lines[acrossdepth][downdepth-1]

			}

			lowestPoint := true
			for _, otherpoint := range otherPoints {
				if otherpoint <= currentElement {
					lowestPoint = false
				}
			}
			if lowestPoint {
				fmt.Printf("We found a new lowest point. It is %d at %d and %d\n", currentElement, acrossdepth, downdepth)
				collectionOfLowestPoints = append(collectionOfLowestPoints, currentElement)
				thisBasinSize := findSizeOfBasin(fmt.Sprintf("%d,%d", acrossdepth, downdepth), lines)
				if thisBasinSize > threeLargestBasins[0] {
					threeLargestBasins[2] = threeLargestBasins[1]
					threeLargestBasins[1] = threeLargestBasins[0]
					threeLargestBasins[0] = thisBasinSize
				} else if thisBasinSize > threeLargestBasins[1] {
					threeLargestBasins[2] = threeLargestBasins[1]
					threeLargestBasins[1] = thisBasinSize
				} else if thisBasinSize > threeLargestBasins[2] {
					threeLargestBasins[2] = thisBasinSize
				}
			}

			//element := lines[acrossdepth][downdepth]
			//fmt.Printf("This is the element%c\n", element)
			//fmt.Printf("The here's the current line: %v\nhere is it's element %c\n", lines[acrossdepth], lines[acrossdepth][downdepth])
		}

	}

	sumOfHeightValues := 0
	for _, point := range collectionOfLowestPoints {
		sumOfHeightValues += point + 1
	}
	fmt.Printf("The sum of depths is %d \n", sumOfHeightValues)
	fmt.Printf("The collection of the largest basins is %d", threeLargestBasins[0]*threeLargestBasins[1]*threeLargestBasins[2])
}

func findSizeOfBasin(spotToCheck string, allPoints []string) int {
	allHigherPoints := getAllHigherPoints(spotToCheck, allPoints)
	allHigherPoints = append(allHigherPoints, spotToCheck)
	uniqueHigherPoints := make([]string, 0)
	for _, newPoint := range allHigherPoints {
		if !contains(uniqueHigherPoints, newPoint) {
			uniqueHigherPoints = append(uniqueHigherPoints, newPoint)
		}
	}
	sizeOfBasin := len(uniqueHigherPoints)
	fmt.Printf("The size of the basin is %d\n", sizeOfBasin)
	return sizeOfBasin
}

func getAllHigherPoints(spotToCheck string, allPoints []string) []string {
	higherPoints := make([]string, 0)
	pointsVerts := strings.Split(spotToCheck, ",")
	acrossdepth, _ := strconv.Atoi(pointsVerts[0])
	downdepth, _ := strconv.Atoi(pointsVerts[1])
	currentHeight := int(allPoints[acrossdepth][downdepth] - '0')
	if acrossdepth > 0 {
		var element int
		lineForElement := allPoints[acrossdepth-1]
		element = int(lineForElement[downdepth] - '0')
		if element > currentHeight && element != 9 {
			higherPoints = append(higherPoints, fmt.Sprintf("%d,%d", acrossdepth-1, downdepth))
		}
		//fmt.Printf("The number we got was %d\n", element)
	}
	if acrossdepth < len(allPoints)-1 {
		var element int
		lineForElement := allPoints[acrossdepth+1]
		element = int(lineForElement[downdepth] - '0')
		if element > currentHeight && element != 9 {
			higherPoints = append(higherPoints, fmt.Sprintf("%d,%d", acrossdepth+1, downdepth))
		}
		//fmt.Printf("The number we got was %d\n", element)
		//element := lines[acrossdepth+1][downdepth]

	}
	if downdepth < len(allPoints[0])-1 {
		var element int
		lineForElement := allPoints[acrossdepth]
		element = int(lineForElement[downdepth+1] - '0')
		if element > currentHeight && element != 9 {
			higherPoints = append(higherPoints, fmt.Sprintf("%d,%d", acrossdepth, downdepth+1))
		}
		//fmt.Printf("The number we got was %d\n", element)
		//element := lines[acrossdepth][downdepth+1]

	}
	if downdepth > 0 {
		var element int
		lineForElement := allPoints[acrossdepth]
		element = int(lineForElement[downdepth-1] - '0')
		if element > currentHeight && element != 9 {
			higherPoints = append(higherPoints, fmt.Sprintf("%d,%d", acrossdepth, downdepth-1))
		}
		//fmt.Printf("The number we got was %d\n", element)
		//element := lines[acrossdepth][downdepth-1]

	}
	allHigherPoints := make([]string, 0)
	for _, higherpoint := range higherPoints {
		allHigherPoints = append(allHigherPoints, higherpoint)
		iterateHigherPoints := getAllHigherPoints(higherpoint, allPoints)
		allHigherPoints = append(allHigherPoints, iterateHigherPoints...)
	}
	fmt.Printf("The number of elements we got is %d and the elements are %v\n", len(allHigherPoints), allHigherPoints)

	return allHigherPoints
}
