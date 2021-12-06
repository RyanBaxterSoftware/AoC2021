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
	//checkForColumnBingo(card, numbers)
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
