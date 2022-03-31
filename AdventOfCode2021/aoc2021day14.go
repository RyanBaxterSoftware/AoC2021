package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func ChemicalExpansion() {
	DisplayChemicals()
	fmt.Println("Wow, that's some chemical expansion! Want to do the first 10 expansions?(1) or All of them?(2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day14Input.txt")
	switch text {
	case "1":
		initialExpansion(inputData)
	case "2":
		fullExpanse(inputData)
	}
}

func DisplayChemicals() {
	fmt.Println("               .")
	fmt.Println("           	.")
	fmt.Println("           	):")
	fmt.Println("              /.:)")
	fmt.Println("             (:./")
	fmt.Println("              )(")
	fmt.Println("             (.;)")
	fmt.Println("             _)/")
	fmt.Println("             |:|")
	fmt.Println("             |;|")
	fmt.Println("             |:|")
	fmt.Println("            .';'.")
	fmt.Println("           /` .: \\")
	fmt.Println("          :~:~8~o~:")
	fmt.Println("       .___'._____.'___.")
	fmt.Println("       |___________|%%%|")
	fmt.Println("       .:  ._\\/_.   :.")
	fmt.Println("       :'  |___%|.  ':")
	fmt.Println("      .:    | %| :   :.")
	fmt.Println("      :'    | %| :   ':")
	fmt.Println("     .:     | %| :.   :.")
	fmt.Println("     :'    .|_%|.':   ':")
	fmt.Println("    .:     | ()%| :    :.")
	fmt.Println("    :'  .__|___%|_:.   ':")
	fmt.Println("   :'   |      ___%|_   ':")
	fmt.Println("  //    |_____(____  )   \\\\")
	fmt.Println("           	   ( (")
	fmt.Println("           		\\ '._____.-")
	fmt.Println("           		 '.___grp_.-\")")
	fmt.Println("           ")
}

func initialExpansion(input string) {
	lines := strings.Split(input, "\n")
	currentChemical := ""
	expansions := make(map[string]string, 0)
	for _, line := range lines {
		if currentChemical == "" {
			currentChemical = line
		} else if line != "" {
			rule := strings.Split(line, " -> ")
			expansions[rule[0]] = rule[1]
		}
	}

	fmt.Printf("The results are %v and %v\n", currentChemical, expansions)

	newChemical := ""
	for x := 0; x < 10; x++ {
		newChemical = ""
		for pos := 0; pos < len(currentChemical)-1; pos++ {
			testingValue := (currentChemical[pos : pos+1]) + (currentChemical[pos+1 : pos+2])
			newChemical += currentChemical[pos : pos+1]
			newChemical += expansions[testingValue]
		}
		newChemical += currentChemical[len(currentChemical)-1:]
		currentChemical = newChemical
		fmt.Printf("The chemical is now %v\n", newChemical)
	}

	chemCount := make(map[rune]int, 0)
	for _, element := range newChemical {
		if val, ok := chemCount[element]; ok {
			chemCount[element] = val + 1
		} else {
			chemCount[element] = 1
		}
	}

	fmt.Printf("The count of elements is %v\n", chemCount)

	max := 0
	min := (1<<bits.UintSize)/2 - 1

	for _, value := range chemCount {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	fmt.Printf("The difference between the occurences that happen most and least is %d\n", max-min)
}

func fullExpanse(input string) {
	// same process, but use a dictionary to follow the rules instead. Keep a count of each 2 elements, and add the new 2 elements pair to iterate on next
	lines := strings.Split(input, "\n")
	chemicalLine := false
	existingChemicals := make(map[string]int, 0)
	expansions := make(map[string]string, 0)
	firstAndLast := make([]string, 0)
	for _, line := range lines {
		if !chemicalLine {
			firstAndLast = append(firstAndLast, line[0:1], line[len(line)-1:])
			fmt.Printf("first and last: %v\n", firstAndLast)
			for x := 0; x < len(line)-1; x++ {
				elementCombo := line[x : x+2]
				if val, ok := existingChemicals[elementCombo]; ok {
					existingChemicals[elementCombo] = val + 1
				} else {
					existingChemicals[elementCombo] = 1
				}
			}
			chemicalLine = true
		} else if line != "" {
			rule := strings.Split(line, " -> ")
			expansions[rule[0]] = rule[1]
		}
	}

	fmt.Printf("The current elements are %v\n", existingChemicals)

	for x := 0; x < 40; x++ {
		newExistingChemicals := make(map[string]int, 0)
		for key, count := range existingChemicals {
			results := make([]string, 0)
			extraLetter := expansions[key]
			results = append(results, key[0:1]+extraLetter)
			results = append(results, extraLetter+key[1:])
			for _, result := range results {
				if val, ok := newExistingChemicals[result]; ok {
					newExistingChemicals[result] = val + count
				} else {
					newExistingChemicals[result] = count
				}
			}
		}
		existingChemicals = newExistingChemicals
		fmt.Printf("Existing chemical combos now %v\n", existingChemicals)
	}
	fmt.Printf("Existing chemical combos now %v\n", existingChemicals)

	// find out how these combos account for element existence.
	elementOccurences := make(map[string]int, 0)
	for pair, occurence := range existingChemicals {
		first := pair[0:1]
		second := pair[1:]
		if val, ok := elementOccurences[first]; ok {
			elementOccurences[first] = val + occurence
		} else {
			elementOccurences[first] = occurence
		}
		if val, ok := elementOccurences[second]; ok {
			elementOccurences[second] = val + occurence
		} else {
			elementOccurences[second] = occurence
		}
	}

	for element, occurence := range elementOccurences {
		elementOccurences[element] = occurence / 2
	}
	// account for the fact that each of these elements have an occurence that isn't double counted like the rest
	elementOccurences[firstAndLast[0]] = elementOccurences[firstAndLast[0]] + 1
	elementOccurences[firstAndLast[0]] = elementOccurences[firstAndLast[1]] + 1

	max := 0
	min := (1<<bits.UintSize)/2 - 1

	for _, value := range elementOccurences {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	fmt.Printf("The difference between the occurences that happen most and least is %d\n", max-min)
}
