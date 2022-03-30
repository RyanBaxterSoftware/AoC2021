package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OrigamiCode() {
	DisplayOrigami()
	fmt.Println("You're telling me there's a code in this paper? Want to do the first fold?(1) or All of them?(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day13Input.txt")
	switch text {
	case "1":
		foldOnce(inputData)
	case "2":
		foldAll(inputData)
	}
}

func DisplayOrigami() {
	fmt.Println("                      .")
	fmt.Println("                     /|\\")
	fmt.Println("                    / |)\\")
	fmt.Println("                   /  I( \\")
	fmt.Println("   ,._            /   I`) \\")
	fmt.Println("  /'\\  `~.       /    | (  \\")
	fmt.Println(" /  _\\    '.    /     | `)  \\")
	fmt.Println("/,-'  \\     `. /      I  )   \\/`-..")
	fmt.Println("       \\      /       |  )    \\    `;-,.._")
	fmt.Println("        \\    /        I  )     \\           ``-.._")
	fmt.Println("         \\  /         I.)'      \\                ``\"..")
	fmt.Println("          \\/          |J_,,..__.,\\.,.__..,,._,,,._,,._`;-,..")
	fmt.Println("           \\      _,.;'")
	fmt.Println("            \\_,-'")
}

func foldOnce(data string) {
	lines := strings.Split(data, "\n")
	instructionsNow := false
	points := make([][]int, 0)
	foldsIndexes := make([]string, 0)
	foldsPositions := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			instructionsNow = true
		} else if !instructionsNow {
			thisPointStr := strings.Split(line, ",")
			thisPointInts := make([]int, 0)
			x, _ := strconv.Atoi(thisPointStr[0])
			y, _ := strconv.Atoi(thisPointStr[1])
			thisPointInts = append(thisPointInts, x, y)
			points = append(points, thisPointInts)
		} else {
			fmt.Printf("%v\n", line)
			lastElement := strings.Split(line, " ")[2]
			splitElement := strings.Split(lastElement, "=")
			foldsIndexes = append(foldsIndexes, splitElement[0])
			position, _ := strconv.Atoi(splitElement[1])
			foldsPositions = append(foldsPositions, position)
		}
	}
	fmt.Printf("The number of points remaining are %d\n", len(accomplishFold(points, foldsIndexes[0], foldsPositions[0])))

	fmt.Printf("The points are %v\nThe folds are %v", points, foldsPositions)

}

func foldAll(data string) {
	lines := strings.Split(data, "\n")
	instructionsNow := false
	points := make([][]int, 0)
	foldsIndexes := make([]string, 0)
	foldsPositions := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			instructionsNow = true
		} else if !instructionsNow {
			thisPointStr := strings.Split(line, ",")
			thisPointInts := make([]int, 0)
			x, _ := strconv.Atoi(thisPointStr[0])
			y, _ := strconv.Atoi(thisPointStr[1])
			thisPointInts = append(thisPointInts, x, y)
			points = append(points, thisPointInts)
		} else {
			fmt.Printf("%v\n", line)
			lastElement := strings.Split(line, " ")[2]
			splitElement := strings.Split(lastElement, "=")
			foldsIndexes = append(foldsIndexes, splitElement[0])
			position, _ := strconv.Atoi(splitElement[1])
			foldsPositions = append(foldsPositions, position)
		}
	}

	remainingPoints := points
	for index := range foldsIndexes {
		remainingPoints = accomplishFold(remainingPoints, foldsIndexes[index], foldsPositions[index])
	}

	fmt.Printf("The points are as follows: %v\n", remainingPoints)

	displayPointsGraphic(remainingPoints)
}

func accomplishFold(allPoints [][]int, foldIndex string, foldPosition int) [][]int {
	pointIndex := 0
	if foldIndex == "y" {
		pointIndex++
	}
	newPoints := make([][]int, 0)
	for _, point := range allPoints {
		foundPoint := make([]int, 0)
		if point[pointIndex] > foldPosition {
			newPosition := foldPosition - (point[pointIndex] - foldPosition)
			tempPoint := make([]int, 2)
			tempPoint[0] = point[0]
			tempPoint[1] = point[1]
			tempPoint[pointIndex] = newPosition
			foundPoint = tempPoint
		} else {
			foundPoint = point
		}
		if !pointExists(newPoints, foundPoint) {
			newPoints = append(newPoints, foundPoint)
		}
	}

	return newPoints
}

func pointExists(existingPoints [][]int, testedPoint []int) bool {
	foundPoint := false
	for _, listPoint := range existingPoints {
		if testedPoint[0] == listPoint[0] && testedPoint[1] == listPoint[1] {
			foundPoint = true
		}
	}
	return foundPoint
}

func displayPointsGraphic(existingPoints [][]int) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 50; x++ {
			symbol := "."
			pointHere := []int{x, y}
			if pointExists(existingPoints, pointHere) {
				symbol = "#"
			}
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}
