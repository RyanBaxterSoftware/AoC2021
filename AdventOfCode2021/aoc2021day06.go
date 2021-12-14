package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FishSpawnCalculations() {
	DisplayLanternFish()
	fmt.Println("Look at all these fishes! Would you like to calculate their spawning in the small scale(1) or full scale(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day6Input.txt")
	switch text {
	case "1":
		simpleFishCalc(inputData)
	case "2":
		bigFishCalc(inputData)
	}
}

func DisplayLanternFish() {
	fmt.Println("    _       __ ")
	fmt.Println("   / \\     /  ;  ")
	fmt.Println("  O   \\_.--\"\"\"-..   _.")
	fmt.Println("      /F         `-'  [")
	fmt.Println("     ]  ,    ,    ,    ;")
	fmt.Println("      '--L__J_.-\"\" ',_;")
	fmt.Println("          '-._J")
}

func simpleFishCalc(inputData string) {
	mapOfFish := make(map[int]int)

	eachFish := strings.Split(inputData, ",")
	for _, fish := range eachFish {
		fishRemainingTime, _ := strconv.Atoi(fish)
		mapOfFish[fishRemainingTime] = mapOfFish[fishRemainingTime] + 1
		fmt.Printf("Adding in a fish with %v days", fish)
	}
	for x := 0; x < 80; x++ {
		numOfBreeders := mapOfFish[0]
		mapOfFish[9] = numOfBreeders
		mapOfFish[7] = mapOfFish[7] + numOfBreeders

		mapOfFish[0] = mapOfFish[1]
		mapOfFish[1] = mapOfFish[2]
		mapOfFish[2] = mapOfFish[3]
		mapOfFish[3] = mapOfFish[4]
		mapOfFish[4] = mapOfFish[5]
		mapOfFish[5] = mapOfFish[6]
		mapOfFish[6] = mapOfFish[7]
		mapOfFish[7] = mapOfFish[8]
		mapOfFish[8] = mapOfFish[9]
		mapOfFish[9] = 0
		fmt.Printf("After iteration %d, we have %d fish\n", x, countFish(mapOfFish))
	}

	fmt.Printf("The total number of fish is %d\n", countFish(mapOfFish))
}

func bigFishCalc(inputData string) {
	mapOfFish := make(map[int]int)

	eachFish := strings.Split(inputData, ",")
	for _, fish := range eachFish {
		fishRemainingTime, _ := strconv.Atoi(fish)
		mapOfFish[fishRemainingTime] = mapOfFish[fishRemainingTime] + 1
		fmt.Printf("Adding in a fish with %v days", fish)
	}
	for x := 0; x < 256; x++ {
		numOfBreeders := mapOfFish[0]
		mapOfFish[9] = numOfBreeders
		mapOfFish[7] = mapOfFish[7] + numOfBreeders

		mapOfFish[0] = mapOfFish[1]
		mapOfFish[1] = mapOfFish[2]
		mapOfFish[2] = mapOfFish[3]
		mapOfFish[3] = mapOfFish[4]
		mapOfFish[4] = mapOfFish[5]
		mapOfFish[5] = mapOfFish[6]
		mapOfFish[6] = mapOfFish[7]
		mapOfFish[7] = mapOfFish[8]
		mapOfFish[8] = mapOfFish[9]
		mapOfFish[9] = 0
		fmt.Printf("After iteration %d, we have %d fish\n", x, countFish(mapOfFish))
	}

	fmt.Printf("The total number of fish is %d\n", countFish(mapOfFish))
}

func countFish(mapOfFish map[int]int) int {
	numberOfAllFish := mapOfFish[0] + mapOfFish[1] + mapOfFish[2] + mapOfFish[3] + mapOfFish[4] + mapOfFish[5] + mapOfFish[6] + mapOfFish[7] + mapOfFish[8]
	return numberOfAllFish
}