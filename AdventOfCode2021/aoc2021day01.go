package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func CalculateDescent() {
	displaySub()
	fmt.Println("Our submarine is in the process of descending. Would you like to verify this descent through data points (1) or through windows(2)?")
	
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day1Input.txt")
	switch text {
		case "1":
			numIncreases, _ := countDepth(inputData)
			fmt.Printf("There were %d increase occurences\n", numIncreases)
		case "2":
			numIncreases, _ := countDepthWindows(inputData)
			fmt.Printf("There were %d increases between windows\n", numIncreases)
	}
}

// TODO: add in err values and checks
func countDepth(inputData string) (int, error) {
	depthPoints := strings.Split(inputData, "\n")
	previousPoint, _ := strconv.Atoi(depthPoints[0])
	numberOfIncreases := 0
	for _, point := range depthPoints[1:] {
		pointNum, _ := strconv.Atoi(point)
		if(previousPoint < pointNum) {
			numberOfIncreases++
			fmt.Println(point + ", increase")
		} else {
			fmt.Println(point + ", decrease")
		}
		previousPoint = pointNum
	}
	return numberOfIncreases, nil
}

func countDepthWindows(inputData string) (int, error) {
	depthPoints := strings.Split(inputData, "\n")
	firstPoint, _ := strconv.Atoi(depthPoints[2])
	secondPoint, _ := strconv.Atoi(depthPoints[1])
	thirdPoint, _ := strconv.Atoi(depthPoints[0])
	previousWindow := firstPoint + secondPoint + thirdPoint
	numberOfIncreases := 0
	windowNum := 0
	for _, point := range depthPoints[3:] {
		thirdPoint = secondPoint
		secondPoint = firstPoint
		firstPoint, _ = strconv.Atoi(point)
		windowNum = firstPoint + secondPoint + thirdPoint
		if(previousWindow < windowNum) {
			numberOfIncreases++
			fmt.Printf("%d, increase\n", windowNum)
		} else {
			fmt.Printf("%d, decrease\n", windowNum)
		}
		previousWindow = windowNum
	}
	return numberOfIncreases, nil
}

func displaySub() {
	fmt.Println("                    _")
	fmt.Println("		   | \\")
	fmt.Println("		    '.|")
	fmt.Println("    _-   _-    _-  _-||    _-    _-  _-   _-    _-    _-")
	fmt.Println("      _-    _-   - __||___    _-       _-    _-    _-")
	fmt.Println("   _-   _-    _-  |   _   |       _-   _-    _-")
	fmt.Println("     _-    _-    /_) (_) (_\\        _-    _-       _-")
	fmt.Println("             _.-'           `-._      ________       _-")
	fmt.Println("        _..--`                   `-..'       .'")
	fmt.Println("    _.-'  o/o                     o/o`-..__.'        ~  ~")
	fmt.Println(" .-'      o|o                     o|o      `.._.  // ~  ~")
	fmt.Println(" `-._     o|o                     o|o        |||<|||~  ~")
	fmt.Println("     `-.__o\\o                     o|o       .'-'  \\ ~  ~")
	fmt.Println("LGB       `-.______________________\\_...-``'.       ~  ~")
	fmt.Println("			             `._______`.")
}