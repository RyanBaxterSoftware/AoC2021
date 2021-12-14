package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bingoBaybee() {
	DisplayBingoSquid()
	// TODO: Update display to print out nicer cards and less gross numbers
	fmt.Println("This squid is looking to play some bingo!!!! (bingo squid) Would you like to use traditional rows and columns(1) or (2)?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	inputData, _ := ReadFile("./Day4Input.txt")
	switch text {
	case "1":
		standardRowsAndColumns(inputData)
	case "2":
		lastWinnerRowsAndColumns(inputData)
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
			currentCard = append(currentCard, strings.TrimSpace(strings.Replace(line, "  ", " ", -1)))
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
				winners = append(winners, cards[cardIndex])
			}
		}
	}
	fmt.Printf("The calling numbers are as follows: \n%v\n", callingNumbers)
	fmt.Printf("The winning cards are as follows: \n%v", winners)

	fmt.Printf("The reportable value for the winning card is %d", calcCardSum(winners[0]))
}

func lastWinnerRowsAndColumns(inputData string) {
	bingolines := strings.Split(inputData, "\n")

	callingNumbers := bingolines[0]

	currentCard := make([]string, 0)
	cards := make([]bingoCard, 0)
	for _, line := range bingolines[1:] {
		if line == "" && len(currentCard) > 0 {
			cards = append(cards, createBingoCard(currentCard))
			currentCard = make([]string, 0)
		} else {
			currentCard = append(currentCard, strings.TrimSpace(strings.Replace(line, "  ", " ", -1)))
			fmt.Printf("The current card we have is %v\nThe length is %v", currentCard, len(currentCard))
		}
	}
	if len(currentCard) > 0 {
		cards = append(cards, createBingoCard(currentCard))
	}

	var lastWinner bingoCard
	for _, number := range strings.Split(callingNumbers, ",") {
		tempCards := make([]bingoCard, 0)
		for cardIndex, card := range cards {
			var isBingo bool
			isBingo, cards[cardIndex] = AddNumber(card, number)
			if isBingo {
				lastWinner = cards[cardIndex]
			} else {
				tempCards = append(tempCards, cards[cardIndex])
			}
		}
		cards = tempCards
	}
	fmt.Printf("The calling numbers are as follows: \n%v\n", callingNumbers)
	fmt.Printf("The winning cards are as follows: \n%v", lastWinner)

	fmt.Printf("The reportable value for the winning card is %d", calcCardSum(lastWinner))
}

func calcCardSum(card bingoCard) int {
	sumOfNumbers := 0
	for _, row := range card.Rows {
		for _, number := range strings.Split(row, " ") {
			if !contains(card.CalledNumbers, number) {
				numInt, _ := strconv.Atoi(number)
				sumOfNumbers += numInt
			}
		}
	}
	finalNum, _ := strconv.Atoi(card.CalledNumbers[len(card.CalledNumbers)-1])
	finalValue := sumOfNumbers * finalNum
	return finalValue
}
