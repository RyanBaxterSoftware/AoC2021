package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func VentNavigation() {
	DisplayOceanVents()
	fmt.Println("Oh no!!!! We found some thermal vents we'll have to avoid! Would you like to run them with simple analysis(1) or complete analysis(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day5Input.txt")
	switch text {
	case "1":
		simpleVentAnalysis(inputData)
	case "2":
		fullVentAnalysis(inputData)
	}
}

func DisplayOceanVents() {
	fmt.Println("                 ooO")
	fmt.Println("	         ooOO")
	fmt.Println("               oOOOO")
	fmt.Println("              ooOO")
	fmt.Println("             /vvv\\")
	fmt.Println("            /V V V\\") 
	fmt.Println("           /V  V  V\\")
	fmt.Println("          /         \\")
	fmt.Println("         /           \\") 
	fmt.Println("        /             \\")
	fmt.Println("       /               \\")
	fmt.Println("      /                 \\")
}

func simpleVentAnalysis(inputVents string) {
	allVentStrings := strings.Split(inputVents, "\n")

	allVents := make([]vent, 0)
	for _, ventstring := range allVentStrings {
		allVents = append(allVents, createVent(ventstring))
	}

	allVents = simplifyVents(allVents)
	hitSpots := make(map[int][]int)
	var danger []string
	for _, nextvent := range allVents {
		var spots []string
		hitSpots, spots = addVentToMapSimple(nextvent, hitSpots)
		for _, spot := range spots {
			if !contains(danger, spot) {
				danger = append(danger, spot)
			}
		}
	}
	fmt.Printf("Here's all the danger spots\n%v\n", danger)
	fmt.Printf("If it worked right the length is %d\n", len(danger))
}

func fullVentAnalysis(inputVents string) {
	allVentStrings := strings.Split(inputVents, "\n")

	allVents := make([]vent, 0)
	for _, ventstring := range allVentStrings {
		allVents = append(allVents, createVent(ventstring))
	}

	hitSpots := make(map[int][]int)
	var danger []string
	for _, nextvent := range allVents {
		var spots []string
		hitSpots, spots = addVentToMapSimple(nextvent, hitSpots)
		for _, spot := range spots {
			if !contains(danger, spot) {
				danger = append(danger, spot)
			}
		}
	}
	fmt.Printf("Here's all the danger spots\n%v\n", danger)
	fmt.Printf("If it worked right the length is %d\n", len(danger))
}

type vent struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func createVent(inputVentData string) vent {
	var newVent vent
	points := strings.Split(inputVentData, " -> ")
	startingPoints := strings.Split(points[0], ",")
	endingPoints := strings.Split(points[1], ",")
	newVent.X1, _ = strconv.Atoi(startingPoints[0])
	newVent.Y1, _ = strconv.Atoi(startingPoints[1])
	newVent.X2, _ = strconv.Atoi(endingPoints[0])
	newVent.Y2, _ = strconv.Atoi(endingPoints[1])
	return newVent
}

func simplifyVents(vents []vent) []vent {
	var remainingVents []vent
	for _, vent := range vents {
		if vent.X1 == vent.X2 || vent.Y1 == vent.Y2 {
			remainingVents = append(remainingVents, vent)
		}
	}

	return remainingVents
}

func displayVent(displayVent vent) {
	fmt.Printf("[%d, %d] to [%d, %d]\n", displayVent.X1, displayVent.Y1, displayVent.X2, displayVent.Y2)
}

func addVentToMapSimple(addedVent vent, existingMap map[int][]int) (map[int][]int, []string) {
	var danger []string
	var max int
	var min int
	if addedVent.X1 == addedVent.X2 {
		if addedVent.Y1 > addedVent.Y2 {
			max = addedVent.Y1
			min = addedVent.Y2
		} else {
			max = addedVent.Y2
			min = addedVent.Y1
		}
		for y := min; y <= max; y++ {
			if !alreadyVisitedCheck(addedVent.X1, y, existingMap) {
				if existingMap[addedVent.X1] == nil {
					existingMap[addedVent.X1] = make([]int, 0)
				}
				existingMap[addedVent.X1] = append(existingMap[addedVent.X1], y)
			} else {
				danger = append(danger, fmt.Sprintf("[%d, %d]", addedVent.X1, y))
			}
		}
	} else if addedVent.Y1 == addedVent.Y2 {
		if addedVent.X1 > addedVent.X2 {
			max = addedVent.X1
			min = addedVent.X2
		} else {
			max = addedVent.X2
			min = addedVent.X1
		}
		for x := min; x <= max; x++ {
			if !alreadyVisitedCheck(x, addedVent.Y1, existingMap) {
				if existingMap[x] == nil {
					existingMap[x] = make([]int, 0)
				}
				existingMap[x] = append(existingMap[x], addedVent.Y1)
			} else {
				danger = append(danger, fmt.Sprintf("[%d, %d]", x, addedVent.Y1))
			}
		} 
	} else {
			var xmax int
			var xmin int
			var ystart int
			var yend int
			if addedVent.X1 > addedVent.X2 {
				xmax = addedVent.X1
				xmin = addedVent.X2
				ystart = addedVent.Y2
				yend = addedVent.Y1
			} else {
				xmax = addedVent.X2
				xmin = addedVent.X1
				ystart = addedVent.Y1
				yend = addedVent.Y2
			}
			y := ystart
			for x := xmin; x <= xmax; x++ {
				if !alreadyVisitedCheck(x, y, existingMap) {
					if existingMap[x] == nil {
						existingMap[x] = make([]int, 0)
					}
					existingMap[x] = append(existingMap[x], y)
				} else {
					danger = append(danger, fmt.Sprintf("[%d, %d]", x, y))
				}
				if yend < ystart {
					y--
				} else {
					y++
				}
			}
	}
	return existingMap, danger
}

func alreadyVisitedCheck(addedX int, addedY int, existingMap map[int][]int) bool {
	pointFound := false
	yPoints := existingMap[addedX]
	for _, point := range yPoints {
		if point == addedY {
			pointFound = true
		}
	}
	return pointFound
}