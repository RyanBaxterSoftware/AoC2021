package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CalculateOctopusEvents() {
	DisplayOceanVents()
	fmt.Println("we're looking at navigating through the octopus group using their lights. Do you want to do basic calculations?(1) or include incomplete pairs?(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day11Input.txt")
	switch text {
	case "1":
		processOctopus(inputData)
	case "2":
		findUnifiedFlash(inputData)
	}
}

type position struct {
	X int
	Y int
}

func processOctopus(inputData string) {
	lines := strings.Split(inputData, "\n")
	allOctopus := make(map[int][]int)
	for lineindex, line := range lines {
		for _, number := range line {
			allOctopus[lineindex] = append(allOctopus[lineindex], int(number-'0'))
		}
	}
	totalFlash := 0
	for x := 0; x < 100; x++ {
		allOctopusFlashingThisTurn := make([]position, 0)
		for lineindex, line := range allOctopus {
			for elementindex, _ := range line {
				allOctopus[lineindex][elementindex] = allOctopus[lineindex][elementindex] + 1
			}
		}
		noFlashes := false
		for !noFlashes {
			noFlashes = true
			allFlashingOctos := make([]position, 0)
			for lineindex, line := range allOctopus {
				for elementindex, element := range line {
					if element >= 10 {
						totalFlash++
						allFlashingOctos = append(allFlashingOctos, position{lineindex, elementindex})
						noFlashes = false
					}
				}
			}
			for _, flash := range allFlashingOctos {
				allOctopusFlashingThisTurn = append(allOctopusFlashingThisTurn, flash)
				allOctopus = flashAtGivenPoint(flash.X, flash.Y, allOctopus)
			}
			for _, flashingOcto := range allOctopusFlashingThisTurn {
				allOctopus[flashingOcto.X][flashingOcto.Y] = 0
			}
		}
		fmt.Printf("Here's the octopuses %v\n", allOctopus)
		//time.Sleep(10 * time.Second)
	}
	fmt.Printf("We had %d flashes total", totalFlash)
}

func findUnifiedFlash(inputData string) {
	lines := strings.Split(inputData, "\n")
	allOctopus := make(map[int][]int)
	for lineindex, line := range lines {
		for _, number := range line {
			allOctopus[lineindex] = append(allOctopus[lineindex], int(number-'0'))
		}
	}
	allFlashed := false
	x := 0
	for ; !allFlashed; x++ {
		allOctopusFlashingThisTurn := make([]position, 0)
		for lineindex, line := range allOctopus {
			for elementindex, _ := range line {
				allOctopus[lineindex][elementindex] = allOctopus[lineindex][elementindex] + 1
			}
		}
		noFlashes := false
		for !noFlashes {
			noFlashes = true
			allFlashingOctos := make([]position, 0)
			for lineindex, line := range allOctopus {
				for elementindex, element := range line {
					if element >= 10 {
						allFlashingOctos = append(allFlashingOctos, position{lineindex, elementindex})
						noFlashes = false
					}
				}
			}
			for _, flash := range allFlashingOctos {
				allOctopusFlashingThisTurn = append(allOctopusFlashingThisTurn, flash)
				allOctopus = flashAtGivenPoint(flash.X, flash.Y, allOctopus)
			}
			for _, flashingOcto := range allOctopusFlashingThisTurn {
				allOctopus[flashingOcto.X][flashingOcto.Y] = 0
			}
		}
		fmt.Printf("Here's the octopuses %v\n", allOctopus)
		//time.Sleep(10 * time.Second)
		if len(allOctopusFlashingThisTurn) == 100 {
			allFlashed = true
		}
	}
	fmt.Printf("We had a unified flash on turn %d", x)
}

func flashAtGivenPoint(line int, column int, octos map[int][]int) map[int][]int {
	fmt.Printf("Octos in flash method: %v\nAbout to process %d, %d", octos, line, column)
	if octos[line][column] < 10 {
		fmt.Println("THIS VALUE IS NOT 10 SO IDK WHY IT'S FLASHING")
	}
	octos[line][column] = 0

	if line > 0 && column > 0 {
		octos[line-1][column-1] = octos[line-1][column-1] + 1
	}

	if line > 0 {
		octos[line-1][column] = octos[line-1][column] + 1
	}

	if line > 0 && column < len(octos[line-1])-1 {
		octos[line-1][column+1] = octos[line-1][column+1] + 1
	}

	if column > 0 {
		octos[line][column-1] = octos[line][column-1] + 1
	}

	if column < len(octos[line])-1 {
		octos[line][column+1] = octos[line][column+1] + 1
	}

	if line < len(octos)-1 && column > 0 {
		octos[line+1][column-1] = octos[line+1][column-1] + 1
	}

	if line < len(octos)-1 {
		octos[line+1][column] = octos[line+1][column] + 1
	}

	if line < len(octos)-1 && column < len(octos[line+1])-1 {
		octos[line+1][column+1] = octos[line+1][column+1] + 1
	}

	return octos
}
