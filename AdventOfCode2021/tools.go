package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if(err != nil) {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	output := ""

	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}
	output = output[:len(output)-1]

	if err := scanner.Err(); err != nil {
		
		fmt.Println("Fatal error with scanner.")
	}
	return output, nil
}