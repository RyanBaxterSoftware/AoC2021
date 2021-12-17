package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Advent of Code 2021")
	fmt.Println("Select from the following days challenges:")
	displayMenu()
	fmt.Println("Make your selection:")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	fmt.Println()
	processText(text)
}

func displayMenu() {
	fmt.Println()
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	dayFiles := make([]fs.FileInfo, 0)
	for _, file := range files {
		if strings.Contains(file.Name(), "aoc2021day") {
			dayFiles = append(dayFiles, file)
		}
	}
	for num, file := range dayFiles {
		fmt.Println(num, file.Name())
	}
	fmt.Println()
}

func processText(text string) {
	switch text {
	case "0":
		helloWorld()
	case "1":
		CalculateDescent()
	case "2":
		CalculateMovement()
	case "3":
		CalculatePowerUsage()
	case "4":
		bingoBaybee()
	case "5":
		VentNavigation()
	case "6":
		FishSpawnCalculations()
	case "7":
		CrabSubMovementCalculation()
	case "8":
		NumpadDisplayMixupFixup()
	default:
		fmt.Println("I didn't recognize " + text + " as a valid input.")
	}
}
