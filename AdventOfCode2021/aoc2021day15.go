package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NavigateChitons() {
	DisplayChemicals()
	fmt.Println("Look out for those chitons! Want to navigate the lowest total risk?(1) or All of them?(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day14Input.txt")
	switch text {
	case "1":
		originalRisk(inputData)
	case "2":
		fullExpanse(inputData)
	}
}

/*
	x---->    y
			  |
			  |
			  |
			  v
*/

func originalRisk(input string) {
	// start at the top left corner
	// go through every option that doesn't backtrack until it hits the previous minimum
	// if we get to the end, it's the new minimum

}

func navigateNextStep(layout [][]int, position []int, previousPositions [][]int, currentRisk int, previousMinimum int) int {
	newRisk = currentRisk + layout[position[0]][position[1]]

	// perform safety check before making each of these positions
	positions := make([][]int, 0)
	north := make([]int, 0)
	north = append(north, position[0]+1, position[1])
	if !pointExists(positions, north) {
		positions = append(positions, north)
	}
	south := make([]int, 0)
	south = append(south, position[0]-1, position[1])
	positions = append(positions, south)
	east := make([]int, 0)
	east = append(east, position[0], position[1]+1)
	positions = append(positions, east)
	west := make([]int, 0)
	west = append(west, position[0], position[1]-1)
	positions = append(positions, west)

	for _, position := range positions {

	}
	return -1
}
