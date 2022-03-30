package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func CavePathNavigation() {
	DisplayOceanVents()
	fmt.Println("we're looking at navigating through the octopus group using their lights. Do you want to do basic calculations?(1) or include incomplete pairs?(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day12Input.txt")
	switch text {
	case "1":
		navigatePathSimple(inputData)
	case "2":
		navigatePathAdjusted(inputData)
	}
}

func navigatePathSimple(inputData string) {
	exits := make(map[string][]string)
	for _, line := range strings.Split(inputData, "\n") {
		connection := strings.Split(line, "-")
		if !contains(exits[connection[0]], connection[1]) {
			exits[connection[0]] = append(exits[connection[0]], connection[1])
		}
		if !contains(exits[connection[1]], connection[0]) {
			exits[connection[1]] = append(exits[connection[1]], connection[0])
		}
	}
	startPath := make([]string, 0)
	startPath = append(startPath, "start")
	fmt.Printf("The options are %v\n", exits)
	allEndOptions := getAllOptions(startPath, exits)
	fmt.Printf("The total number of options we found is %d\n", len(allEndOptions))
}

func getAllOptions(currentPath []string, options map[string][]string) [][]string {
	// for each next step option
	// add option to path
	// get all the possible paths from there
	// return the end

	allOptions := make([][]string, 0)
	for _, option := range options[currentPath[len(currentPath)-1]] {
		tempPath := append(currentPath, option)
		if option != "end" && (!contains(currentPath, option) || unicode.IsUpper(rune(option[0]))) {
			allOptions = append(allOptions, getAllOptions(tempPath, options)...)
		} else if option == "end" {
			newTempPath := make([]string, 0)
			// this is done to work around the issue with golang holding slice values by reference instead of by value.
			for _, cave := range tempPath {
				newTempPath = append(newTempPath, cave)
			}
			allOptions = append(allOptions, newTempPath)
		}
	}

	return allOptions
}

func navigatePathAdjusted(inputData string) {
	exits := make(map[string][]string)
	for _, line := range strings.Split(inputData, "\n") {
		connection := strings.Split(line, "-")
		if !contains(exits[connection[0]], connection[1]) {
			exits[connection[0]] = append(exits[connection[0]], connection[1])
		}
		if !contains(exits[connection[1]], connection[0]) {
			exits[connection[1]] = append(exits[connection[1]], connection[0])
		}
	}
	startPath := make([]string, 0)
	startPath = append(startPath, "start")
	fmt.Printf("The options are %v\n", exits)
	allEndOptions := getAllOptionsWithSmallCaveAllowance(startPath, exits, false)
	fmt.Printf("The total number of options we found is %d\n", len(allEndOptions))
}

func getAllOptionsWithSmallCaveAllowance(currentPath []string, options map[string][]string, smallCaveRecurred bool) [][]string {
	// for each next step option
	// add option to path
	// get all the possible paths from there
	// return the end

	allOptions := make([][]string, 0)
	for _, option := range options[currentPath[len(currentPath)-1]] {
		tempPath := append(currentPath, option)
		if option != "end" && option != "start" && (unicode.IsUpper(rune(option[0])) || (!(contains(currentPath, option) && smallCaveRecurred))) {
			futureSmallCaveRecurred := smallCaveRecurred
			if contains(currentPath, option) && unicode.IsLower(rune(option[0])) {
				futureSmallCaveRecurred = true
			}
			allOptions = append(allOptions, getAllOptionsWithSmallCaveAllowance(tempPath, options, futureSmallCaveRecurred)...)
		} else if option == "end" {
			newTempPath := make([]string, 0)
			// this is done to work around the issue with golang holding slice values by reference instead of by value.
			for _, cave := range tempPath {
				newTempPath = append(newTempPath, cave)
			}
			allOptions = append(allOptions, newTempPath)
		}
	}

	return allOptions
}
