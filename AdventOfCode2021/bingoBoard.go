package main

import (
	"fmt"
	"strings"
)

type bingoCard struct {
	Rows          []string
	CalledNumbers []string
}

func createBingoCard(rows []string) bingoCard {
	fmt.Printf("%v", rows)
	var card bingoCard
	card.Rows = rows
	return card
}

// Returns a boolean indicating if the passed in card has a bingo
func checkForBingoSimple(card bingoCard) bool {
	bingoFound := false
	if checkForRowBingo(card) > 0 {
		bingoFound = true
	}
	if checkForColumnBingo(card) > 0 {
		bingoFound = true
	}

	return bingoFound
}

// Returns a int value of the row the bingo was in
func checkForRowBingo(card bingoCard) int {
	bingos := 0
	for which, row := range card.Rows {
		bingo := true
		for _, number := range strings.Split(row, " ") {
			if !contains(card.CalledNumbers, number) {
				bingo = false
			}
		}
		if bingo {
			bingos = which + 1
		}
	}
	if bingos != 0 {
		fmt.Printf("We have a bingo from the row of the bingo card. Here's the card: \n %v \n and here's the winning numbers\n%v\n", card.Rows, card.CalledNumbers)
	}
	return bingos
}

func checkForColumnBingo(card bingoCard) int {
	bingos := 0
	var realWin []string
	for column := 0; column < len(strings.Split(card.Rows[0], " ")); column++ {
		bingo := true
		winningNumbers := make([]string, 0)
		for _, row := range card.Rows {
			winningNumbers = append(winningNumbers, strings.Split(row, " ")[column])
			if !contains(card.CalledNumbers, strings.Split(row, " ")[column]) {
				bingo = false
			}
		}
		if bingo {
			bingos = column + 1
			realWin = winningNumbers
		}
	}
	if(bingos != 0) {
		fmt.Printf("We got a bingo in column %v and the called numbers are %v which fit into %v and hte winning numbers there are %v\n", bingos, card.CalledNumbers, card.Rows, realWin)
	}
	return bingos
}

func contains(slice []string, element string) bool {
	foundIt := false
	for _, checkItems := range slice {
		if checkItems == element {
			foundIt = true
		}
	}
	return foundIt
}

// returns whether the new card gives a bingo
func AddNumber(card bingoCard, number string) (bool, bingoCard) {
	card.CalledNumbers = append(card.CalledNumbers, number)
	fmt.Printf("We have added %v to the list of numbers, giving us %v\n", number, card.CalledNumbers)
	return checkForBingoSimple(card), card
}
