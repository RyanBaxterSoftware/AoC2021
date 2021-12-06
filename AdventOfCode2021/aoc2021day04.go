package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bingoBaybee() {
	DisplayBingoSquid()
	fmt.Println("This squid is looking to play some bingo!!!! (bingo squid) Would you like to use traditional rows and columns(1) or (2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day4Input.txt")
	switch text {
	case "1":
		standardRowsAndColumns(inputData)
	case "2":
		CalculateOxygenScrubbing(inputData)
	}
}

func DisplayBingoSquid() {
	fmt.Println("          ^")
	fmt.Println("	/   \\            _     _   ")
	fmt.Println("	\\   /         	| |   (_)   ")
	fmt.Println("	|   |           | |__  _ _ __   __ _  ___")
	fmt.Println("	|   |           | '_ \\| | '_ \\ / _` |/ _ \\")
	fmt.Println("	| 0 |           | |_) | | | | | (_| | (_) |")
	fmt.Println("       // ||\\\\          |_.__/|_|_| |_|\\__, |\\___/")
	fmt.Println("      (( // || 	                        __/ |      ")
	fmt.Println("       \\\\))  \\\\                        |___/   ")
	fmt.Println("       //||   ))")
	fmt.Println("       ( ))  //")
	fmt.Println("       //   ((")
}

func standardRowsAndColumns(inputData string) {
	bingolines := strings.Split(inputData, "\n")

	callingNumbers := bingolines[0]

	currentCard := make([]string, 0)
	cards := make([]bingoCard, 0)
	for _, line := range bingolines[1:] {
		if line == "" && len(currentCard) > 0 {
			cards = append(cards, createBingoCard(currentCard))
			currentCard = make([]string, 0)
		} else {
			currentCard = append(currentCard, line)
			fmt.Printf("The current card we have is %v\nThe length is %v", currentCard, len(currentCard))
		}
	}
	if len(currentCard) > 0 {
		cards = append(cards, createBingoCard(currentCard))
	}

	winners := make([]bingoCard, 0)
	for _, number := range strings.Split(callingNumbers, ",") {
		for cardIndex, card := range cards {
			var isBingo bool
			isBingo, cards[cardIndex] = AddNumber(card, number)
			if isBingo && len(winners) == 0 {
				winners = append(winners, card)
			}
		}
	}
	fmt.Printf("The calling numbers are as follows: \n%v\nand the cards we have are as follows %v\n", callingNumbers, cards)
	fmt.Printf("The winning cards are as follows: \n%v", winners)
}
