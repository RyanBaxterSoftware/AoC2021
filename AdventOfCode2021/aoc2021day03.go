package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculatePowerUsage() {
	DisplayPowerSub()
	fmt.Println("Our submarine is outputting power data. Would you like to process this data using standard processes (1) or ????(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day3Input.txt")
	switch text {
	case "1":
		calculatePowerStandard(inputData)
	case "2":
		CalculateOxygenScrubbing(inputData)
	}
}

func DisplayPowerSub() {
	fmt.Println("        ⚡      |_      ⚡")
	fmt.Println("    ⚡    _____|~ |____        ⚡")
	fmt.Println("        (  --         ~~~~--_,")
	fmt.Println("        ~~~~~~~~~~~~~~~~~~~'`   dr")
}

func calculatePowerStandard(inputData string) {
	bytes := strings.Split(inputData, "\n")
	//we're analyzing binary so we can make it expressly defined
	zeroOccurences := make([]int, len(bytes[0]))
	oneOccurences := make([]int, len(bytes[0]))
	for _, byte := range bytes {
		for i, bit := range byte {
			//fmt.Printf("This bit is %c in spot %v\n", bit, i)
			switch bit {
			case '1':
				oneOccurences[i]++
			case '0':
				zeroOccurences[i]++
			}
		}
		fmt.Printf("We have processed the byte of %v\n", byte)
	}
	fmt.Println(zeroOccurences)
	fmt.Println(oneOccurences)
	gammaNumber := ""
	epsilonNumber := ""
	for i := 0; i < len(oneOccurences); i++ {
		if zeroOccurences[i] < oneOccurences[i] {
			gammaNumber += "1"
			epsilonNumber += "0"
		} else {
			gammaNumber += "0"
			epsilonNumber += "1"
		}
	}
	fmt.Printf("Gamma number is %v and the epsilon number is %v\n", gammaNumber, epsilonNumber)
	gammaNumberInt, _ := strconv.ParseInt(gammaNumber, 2, 64)
	epsilonNumberInt, _ := strconv.ParseInt(epsilonNumber, 2, 64)
	fmt.Printf("We are multiplying %d and %d to get %d\n", gammaNumberInt, epsilonNumberInt, (gammaNumberInt * epsilonNumberInt))
}

// TODO: add a lot of output to show what change is being made.
func CalculateOxygenScrubbing(inputData string) {
	bytes := strings.Split(inputData, "\n")
	oxygenPossibles := bytes
	cotwoPossibles := bytes
	for i := 0; i < len(bytes[0]); i++ {
		countZero := 0
		countOne := 0

		var soughtValue byte
		var tempSlice []string
		if len(oxygenPossibles) > 1 {
			for _, byte := range oxygenPossibles {
				switch byte[i] {
				case '0':
					countZero++
				case '1':
					countOne++
				}
			}
			// process out all values without that number in that position
			soughtValue = '1'
			if countZero > countOne {
				soughtValue = '0'
			}
			tempSlice = make([]string, 0)
			for _, thisValue := range oxygenPossibles {
				if thisValue[i] == soughtValue {
					tempSlice = append(tempSlice, thisValue)
				}
			}
			fmt.Println(tempSlice)
			oxygenPossibles = tempSlice
		}

		if len(cotwoPossibles) > 1 {
			countZero = 0
			countOne = 0
			for _, byte := range cotwoPossibles {
				switch byte[i] {
				case '0':
					countZero++
				case '1':
					countOne++
				}
			} // process out all values without that number in that position
			soughtValue = '0'
			if countZero > countOne {
				soughtValue = '1'
			}
			tempSlice = make([]string, 0)
			for _, thisValue := range cotwoPossibles {
				if thisValue[i] == soughtValue {
					tempSlice = append(tempSlice, thisValue)
				}
			}
			fmt.Println(tempSlice)
			cotwoPossibles = tempSlice
		}
		// process out all values without that number in that position
	}

	oxygenNumberInt, _ := strconv.ParseInt(oxygenPossibles[0], 2, 64)
	cotwoNumberInt, _ := strconv.ParseInt(cotwoPossibles[0], 2, 64)
	fmt.Printf("We are multiplying %d and %d to get %d\n", oxygenNumberInt, cotwoNumberInt, (oxygenNumberInt * cotwoNumberInt))
}
