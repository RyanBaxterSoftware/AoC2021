package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		findUnifiedFlash(inputData)
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
	fmt.Printf("All connects are: %v\n", exits)
}
